---
title: "API Gateway (gateway)"
description: "Posisikan di depan satu atau beberapa layanan upstream dengan tabel rute deklaratif, middleware per-route, pembatasan laju (rate limiting), batas waktu (timeouts), dan penyeimbangan beban service-discovery."
date: "2026-10-25"
---

# API Gateway (`gateway`)

Workload **`gateway`** adalah reverse proxy mandiri yang menjadi gerbang utama di depan satu atau beberapa layanan upstream. Beban kerja ini memiliki tabel rute deklaratif, menerapkan middleware (CORS, tracing, timeout, rate limiting, circuit breaking), dan menguraikan alamat upstream melalui service registry.

Ditenagai oleh `pkg/web/gateway`, tabel rute yang sama dapat dijalankan dari binary proyek maupun langsung dari CLI melalui `kiw gateway`.

---

## Karakteristik Arsitektur

1. **Tabel Rute Deklaratif:** Rute eksak dan prefiks, pencocokan method per-route, serta penentuan alamat baik dengan `target_url` tetap maupun `service` yang diuraikan oleh registry.
2. **Middleware Per-Route:** CORS, tracing, timeout, rate limiting, dan circuit breaking didaftarkan berdasarkan nama dan disusun per-route alih-alih secara global.
3. **Penemuan Layanan (Service Discovery):** Penyeimbangan beban round-robin di atas endpoint yang terdaftar; upstream yang tidak tersedia akan mengembalikan `502` dengan payload bodi Problem JSON terstruktur.
4. **Satu Tabel Rute, Dua Cara Eksekusi:** `kiw dev` / `kiw run` mengeksekusi binary proyek; `kiw gateway` menjalankan blok `gateway:` yang identik secara langsung tanpa perlu build ulang.

---

## Scaffolding Proyek `gateway`

```bash
kiw new my-gateway --gateway
cd my-gateway
```

Tata letak direktori:

```text
my-gateway/
├── krewire.yaml          # Konfigurasi proyek dan tabel rute
├── go.mod
├── main.go               # Memuat krewire.yaml dan menyajikan gateway
├── internal/
│   └── config/           # Loader bertipe data untuk blok gateway:
└── README.md
```

---

## Implementasi Tabel Rute (`main.go`)

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krewire/krewire/pkg/cloud/service/registry"
	"github.com/krewire/krewire/pkg/web/gateway"
	"my-gateway/internal/config"
)

func main() {
	cfg, err := config.Load("krewire.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "gateway: load config: %v\n", err)
		os.Exit(1)
	}

	routes, err := cfg.Routes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gateway: %v\n", err)
		os.Exit(1)
	}

	gw := gateway.New(
		gateway.WithRegistry(registry.NewMemoryRegistry()),
		gateway.WithMiddleware("cors", gateway.CORSMiddleware("*")),
		gateway.WithMiddleware("tracing", gateway.TracingMiddleware()),
	)
	for _, route := range routes {
		gw.AddRoute(route)
	}

	server := &http.Server{Addr: cfg.Addr(), Handler: gw}
	slog.Info("gateway listening", "addr", server.Addr, "routes", len(routes))

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintf(os.Stderr, "gateway error: %v\n", err)
			os.Exit(1)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	slog.Info("shutting down gateway...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
}
```

---

## Konfigurasi (`krewire.yaml`)

```yaml
project:
  name: "my-gateway"
  kind: gateway
  version: "0.1.0"

gateway:
  host: ""
  port: 8080
  routes:
    - path: /api
      prefix: true
      strip_prefix: true
      target_url: http://localhost:8081
      timeout: 10s
    - path: /users
      service: user-service
      middleware: [cors, tracing]
      rate_limit: 100
```

Field rute dipetakan langsung ke `gateway.Route`: `path`, `method`, `service`, `target_url`, `prefix`, `strip_prefix`, `middleware`, `rate_limit`, dan `timeout`.

`rate_limit` (permintaan per detik) dan `timeout` ditegakkan langsung oleh gateway itu sendiri—rute akan dibatasi lajunya dan diberi tenggat waktu bahkan jika tidak ada middleware yang dipasang. Permintaan yang melebihi kuota akan menerima status `429` beserta header `Retry-After`; upstream yang lambat akan memicu respons `504 Gateway Timeout`. Variabel lingkungan `GATEWAY_ADDR` menimpa nilai `gateway.host` + `gateway.port`.

---

## Middleware

Middleware didaftarkan pada gateway di bawah sebuah nama dan dipilih per-route melalui daftar `middleware:`. Penamaan bebas ditentukan sendiri; package menyediakan fungsi konstruktor:

- `gateway.CORSMiddleware(allowedOrigins ...string)`
- `gateway.TracingMiddleware()` — mempropagasi `X-Request-ID`
- `gateway.TimeoutMiddleware(d)` — menetapkan batas waktu permintaan (deadline)
- `gateway.RateLimitMiddleware(rps)` — mengatur jarak permintaan sebesar 1/rps
- `gateway.CircuitBreakerMiddleware(breaker)` — fail fast ketika layanan upstream mengalami degradasi

Rute yang menyebutkan nama middleware yang belum didaftarkan pada gateway akan mengabaikannya, sehingga pemasangan alur tetap eksplisit di dalam `main.go`.

---

## Menjalankan Gateway

```bash
# Pengembangan: compile ulang dan restart saat ada perubahan file
kiw dev

# Binary produksi di .krewire/build/bin/<name>
kiw build

# Jalankan tabel rute bersama tanpa perlu melakukan build binary proyek
kiw gateway --port 8080 --routes "/api=user-service,/pay=payment-service"
