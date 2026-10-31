---
title: "Wadah Aplikasi & Assembly"
description: "Wadah dependency injection bertipe, perakitan provider, manajemen siklus hidup, dan supervisi proses graceful shutdown di pkg/app."
date: "2026-10-31"
---

# Wadah Aplikasi & Assembly (`pkg/app`)

Jalur impor: `github.com/krewire/krewire/pkg/app`

Paket `pkg/app` menyediakan kontainer dependensi bertipe (typed dependency container), sistem perakitan penyedia (provider assembly), dan supervisi proses dengan graceful shutdown (`KWF-D1CNT`). Paket ini berfungsi sebagai akar komposisi tunggal (single composition root) aplikasi—setiap dependensi dideklarasikan dengan keamanan tipe statis Go, diselesaikan secara deterministik, dan dapat diobservasi melalui log terstruktur serta hook tracer.

---

## Fitur Utama

- **Kontainer Dependensi Bertipe:** API `Provide[T]`, `Singleton[T]`, dan `Resolve[T]` berbasis generic yang mengeliminasi kesalahan refleksi saat runtime.
- **Perakitan Penyedia (Provider Assembly):** Siklus dua fase (`Register` dilanjutkan dengan `Boot`) yang dikelola oleh `NewApp`, mengunci modifikasi kontainer sebelum fase boot berjalan.
- **Override Pengujian (Test-Time Overrides):** `(*App).Test(overrides...)` memungkinkan penggantian dependensi tiruan (mock) dengan cepat tanpa mengotori konfigurasi produksi.
- **Supervisi Daur Hidup & Graceful Draining:** Penanganan sinyal OS bawaan (`SIGINT`, `SIGTERM`) yang mengoordinasikan shutdown aman di antara listener HTTP dan worker latar belakang.
- **Observabilitas & Tracing:** Hook bawaan untuk `slog.Logger`, deteksi siklus dependensi melingkar (`CycleError`), dan tracer performa khusus.

---

## Ringkasan API

### 1. Operasi Kontainer

```go
// Membuat kontainer dengan opsi
c := app.New(app.WithLogger(logger), app.WithTrace())

// Daftarkan factory transient (instansi baru setiap resolusi)
app.Provide[*UserService](c, func() *UserService {
    return &UserService{}
})

// Daftarkan factory singleton (dihitung sekali lalu di-cache)
app.Singleton[*sql.DB](c, func() (*sql.DB, error) {
    return sql.Open("pgx", os.Getenv("DATABASE_URL"))
})

// Selesaikan dependensi bertipe
db, err := app.Resolve[*sql.DB](c)
```

### 2. Perakitan Penyedia (`app.Provider`)

Provider menstrukturkan subsistem modular antar-domain:

```go
type Provider interface {
    Register(c *app.Container) error
    Boot(c *app.Container) error
}
```

- **Fase `Register`:** Daftarkan factory dan singleton ke dalam kontainer.
- **Fase `Boot`:** Dijalankan setelah semua provider selesai mendaftar dan kontainer dikunci; selesaikan penyambungan antar-layanan dan daftarkan route.

```go
a := app.NewApp(
    &DatabaseProvider{},
    &UserServiceProvider{},
    &HttpServerProvider{},
)

container, err := a.Build()
if err != nil {
    log.Fatalf("gagal menyusun aplikasi: %v", err)
}
```

---

## Contoh Lengkap

Contoh berikut menunjukkan pembuatan composition root yang rapi dengan provider kustom, resolusi layanan bertipe, dan eksekusi dengan manajemen daur hidup graceful shutdown:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/krewire/krewire/pkg/app"
)

// Service mendefinisikan antarmuka logika bisnis.
type GreeterService struct {
	Greeting string
}

// GreeterProvider mengikat layanan ke dalam kontainer app.
type GreeterProvider struct{}

func (p *GreeterProvider) Register(c *app.Container) error {
	return app.Singleton[*GreeterService](c, func() *GreeterService {
		return &GreeterService{Greeting: "Halo dari Krewire"}
	})
}

func (p *GreeterProvider) Boot(c *app.Container) error {
	svc, err := app.Resolve[*GreeterService](c)
	if err != nil {
		return err
	}
	fmt.Printf("[Boot] Greeter diinisialisasi: %s\n", svc.Greeting)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Susun provider aplikasi
	application, err := app.Bootstrap(ctx, &GreeterProvider{})
	if err != nil {
		log.Fatalf("kesalahan bootstrap: %v", err)
	}
	defer application.Stop(ctx)

	// 2. Selesaikan dependensi bertipe
	greeter, err := app.Resolve[*GreeterService](application.Container())
	if err != nil {
		log.Fatalf("kesalahan resolve: %v", err)
	}

	// 3. Tentukan handler HTTP
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s pada %s\n", greeter.Greeting, time.Now().Format(time.RFC3339))
	})

	// 4. Jalankan server dengan penanganan interupsi graceful shutdown
	runner := app.HttpRunner(mux, ":8080")
	if err := application.Run(ctx, runner); err != nil {
		log.Fatalf("server terhenti: %v", err)
	}
}
```

---

## Batasan Desain

`pkg/app` bertanggung jawab atas komposisi aplikasi, urutan registrasi, dan supervisi runner saja. Aturan domain bisnis berada di `pkg/model`, perutean HTTP di `pkg/web`, dan mekanisme mesin runtime di `pkg/lifecycle`. `pkg/app` tidak pernah mengimpor `tools/kiw`.
