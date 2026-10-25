---
title: "Wadah Aplikasi & Monolit"
description: "Wadah dependency injection, manajemen siklus hidup aplikasi, dan supervisi proses graceful shutdown di pkg/app."
date: "2026-10-25"
---

# Wadah Aplikasi & Monolit (`pkg/app`)

Jalur impor: `github.com/krewire/krewire/pkg/app`

Paket `app` mengelola siklus hidup aplikasi, injeksi dependensi (DI), dan supervisi proses yang aman untuk monolit web tingkat produksi.

---

## Fitur Utama

- **Penyusunan Kontainer:** Mendaftarkan dependensi singleton dan transient dengan keamanan tipe statis Go.
- **Graceful Shutdown:** Mengintersepsi sinyal OS (`SIGINT`, `SIGTERM`), mengosongkan koneksi HTTP aktif dalam batas waktu yang dapat dikonfigurasi.
- **Health Probe:** Hook bawaan untuk ping database, kesiapan sistem (readiness), dan kelangsungan proses (liveliness).

---

## Contoh Penggunaan

```go
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/krewire/krewire/pkg/app"
)

func main() {
	application := app.New(app.Config{
		Name:    "storefront",
		Version: "0.1.0",
	})

	// Mendaftarkan koneksi database bersama
	application.Provide(func() (*sql.DB, error) {
		return sql.Open("postgres", os.Getenv("DATABASE_URL"))
	})

	// Menjalankan server dengan signal draining
	if err := application.RunServer(":8080", router); err != nil {
		log.Fatal(err)
	}
}
```\n