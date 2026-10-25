---
title: "Web & Aplikasi (web, app)"
description: "HTTP routing inti, pipeline middleware, format problem details RFC 7807, container dependency injection, dan supervisi siklus hidup server."
date: "2026-10-08"
---

# Web & Aplikasi (`web`, `app`)

Jalur impor (*import paths*):
- `github.com/krewire/krewire/pkg/web`
- `github.com/krewire/krewire/pkg/app`

Paket **`web`** dan **`app`** menyediakan engine web HTTP berkinerja tinggi dan container dependency injection yang mentenagai workload `app`, `service`, dan `site`.

---

## 1. Engine Web (`pkg/web`)

Paket `web` membungkus pustaka `net/http` bawaan Go dengan routing ekspresif, ekstraksi parameter, rantai middleware, dan serialisasi JSON berkecepatan tinggi.

### Fitur Utama
- **Routing Bebas Alokasi (Zero Allocations):** Router radix tree berperforma tinggi yang mendukung grup rute, parameter path (`/users/:id`), dan wildcard (`/static/*`).
- **Kompatibilitas Pustaka Standar:** Seluruh middleware dan handler menggunakan tipe standar `http.Handler` dan `http.HandlerFunc`.
- **Standar RFC 7807 Problem Details:** Respon error JSON terstandarisasi untuk kenyamanan konsumsi API.
- **Helper Rendering:** Fungsi pembantu kelas satu seperti `web.JSON()`, `web.HTML()`, dan `web.Redirect()`.

### Contoh Penggunaan

```go
package main

import (
	"net/http"
	"github.com/krewire/krewire/pkg/web"
)

func main() {
	r := web.NewRouter()

	// Middleware Global
	r.Use(web.Recovery())
	r.Use(web.Logger())

	// Grup Rute
	api := r.Group("/api/v1")
	{
		api.Get("/items", func(w http.ResponseWriter, r *http.Request) {
			items := []string{"buku", "keyboard", "monitor"}
			web.JSON(w, http.StatusOK, items)
		})

		api.Get("/items/:id", func(w http.ResponseWriter, r *http.Request) {
			id := web.Param(r, "id")
			web.JSON(w, http.StatusOK, map[string]string{"id": id})
		})
	}

	http.ListenAndServe(":8080", r)
}
```

---

## 2. Wadah Aplikasi (`pkg/app`)

Paket `app` mengelola siklus hidup aplikasi, injeksi dependensi, dan supervisi proses dengan graceful shutdown.

### Fitur Utama
- **Perakitan Wadah (Container Assembly):** Pendaftaran layanan singleton dan transient dengan jaminan type-safety.
- **Graceful Shutdown:** Menangkap sinyal OS (`SIGINT`, `SIGTERM`), menyelesaikan koneksi HTTP aktif dengan batas waktu timeout yang dapat dikonfigurasi.
- **Probe Kesehatan (Health Probes):** Hook bawaan untuk ping koneksi database dan verifikasi kesiapan layanan (readiness check).

### Contoh Penggunaan

```go
application := app.New(app.Config{
    Name:    "storefront",
    Version: "0.1.0",
})

// Mendaftarkan koneksi database bersama
application.Provide(func() (*sql.DB, error) {
    return sql.Open("postgres", os.Getenv("DATABASE_URL"))
})

// Menjalankan server dengan penanganan sinyal OS
if err := application.RunServer(":8080", router); err != nil {
    log.Fatal(err)
}
```
