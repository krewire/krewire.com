---
title: "Microservice (service)"
description: "Microservice cloud-native dengan penemuan layanan (service discovery), perutean API gateway, pembatasan laju (rate limiting), circuit breaking, dan tracing terdistribusi OpenTelemetry."
date: "2026-10-08"
---

# Microservice (`service`)

Workload **`service`** memperlengkapi para pengembang untuk membangun microservice cloud-native ber-throughput tinggi dan API terdistribusi.

Ditenagai oleh `pkg/cloud/service`, `pkg/sec`, dan `pkg/resilience`, workload `service` sudah mencakup penemuan layanan (service discovery) kelas produksi, tracing terdistribusi, perutean gateway, dan telemetri langsung sejak awal.

---

## Karakteristik Arsitektur

1. **Service Discovery:** Driver bawaan untuk Consul, etcd, Kubernetes DNS, dan AWS Cloud Map.
2. **Telemetri & Observabilitas:** Instrumentasi bawaan OpenTelemetry (OTel) untuk tracing terdistribusi, metrik, dan structured logging dengan propagasi W3C `traceparent`.
3. **Ketahanan Lalu Lintas (Traffic Resilience):** Rate limiting token-bucket, circuit breaker, propagasi batas waktu (timeout), dan bulkhead konkurensi.
4. **Pemeriksaan Kesehatan & Kesiapan:** Endpoint `/healthz` dan `/readyz` otomatis yang mematuhi konvensi kontainer cloud.

---

## Scaffolding Proyek `service`

```bash
kiw new auth-service --service
cd auth-service
```

Tata letak direktori:

```text
auth-service/
├── krewire.yaml          # Konfigurasi proyek
├── go.mod
├── cmd/
│   └── service/
│       └── main.go       # Titik masuk service
└── internal/
    ├── api/              # Handler endpoint HTTP / gRPC
    ├── discovery/        # Logika registrasi service
    └── domain/           # Logika bisnis domain inti
```

---

## Implementasi Service

```go
package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/krewire/krewire/pkg/cloud/service"
	"github.com/krewire/krewire/pkg/web"
)

func main() {
	svc := service.New(service.Config{
		Name:    "auth-service",
		Version: "0.1.0",
		Port:    ":8080",
	})

	router := web.NewRouter()

	// Endpoint kesehatan dan kesiapan
	router.Get("/healthz", service.HealthHandler())
	router.Get("/readyz", service.ReadinessHandler())

	// Route bisnis yang dilindungi
	router.Post("/v1/tokens/verify", func(w http.ResponseWriter, r *http.Request) {
		web.JSON(w, http.StatusOK, map[string]any{
			"valid":     true,
			"verified": time.Now().UTC(),
		})
	})

	// Jalankan service dengan registrasi penemuan otomatis
	if err := svc.Start(context.Background(), router); err != nil {
		os.Exit(1)
	}
}
```

---

## Konfigurasi (`krewire.yaml`)

```yaml
project:
  name: "auth-service"
  kind: service
  version: "0.1.0"

service:
  name: "auth-service"
  port: 8080
```

`name` dan `port` adalah satu-satunya kunci di bawah `service:`. Penemuan layanan (`pkg/cloud/service/registry`), rate limiting dan circuit breaking (`pkg/cloud/service/resilience`), serta tracing (`pkg/cloud/service/tracing`) disusun langsung dalam kode Go tempat registry dan transport dihubungkan.

---

## Menjalankan dan Memverifikasi

Jalankan service secara lokal:

```bash
kiw run
```

Periksa endpoint yang terdaftar dan status kesehatan service:

```bash
curl http://localhost:8080/healthz
