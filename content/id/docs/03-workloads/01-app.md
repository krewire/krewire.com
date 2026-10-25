---
title: "Monolit Fullstack (app)"
description: "Panduan arsitektur dan implementasi untuk workload app: fondasi Go net/http, embedded asset, template SSR, dependency injection, dan operasional produksi."
date: "2026-10-08"
---

# Monolit Fullstack (`app`)

Workload **`app`** adalah fondasi Krewire untuk membangun monolit web tingkat produksi. Beban kerja ini mengompilasi seluruh permukaan web—perutean HTTP, tampilan yang dirender di sisi server (SSR), aset statis, akses basis data, tugas latar belakang, dan endpoint API—ke dalam **satu binary Go mandiri**.

Dibangun di atas primitif standar Go (`net/http`, `embed.FS`, `context`, `slog`), workload `app` tidak memerlukan server web eksternal seperti Nginx atau runtime Node.js untuk pengembangan lokal maupun deployment dasar.

---

## Karakteristik Arsitektur

1. **Nol Dependensi Runtime:** Semua yang dibutuhkan untuk menjalankan aplikasi dikompilasi ke dalam satu file executable statis.
2. **Utamakan Pustaka Standar:** Sistem perutean dan middleware mengimplementasikan antarmuka standar Go `http.Handler` dan `http.HandlerFunc`.
3. **Embedded Static Assets:** File CSS, JavaScript, ikon, dan template dikompilasi langsung ke dalam binary menggunakan package `embed` bawaan Go, mencegah galat aset hilang (missing assets) di lingkungan produksi.
4. **Batas Domain yang Rapi:** Dibangun di sekitar pola modular monolit (`KWF-5ZHQV`), memungkinkan ekstraksi masa depan ke dalam microservice atau worker latar belakang tanpa perlu menulis ulang kode.

---

## Scaffolding Proyek `app`

Gunakan perintah `kiw new` dengan flag `--app`:

```bash
kiw new storefront --app
cd storefront
```

Struktur direktori yang dihasilkan mengikuti tata letak kanonikal Go:

```text
storefront/
├── krewire.yaml          # Konfigurasi proyek
├── go.mod                # Definisi modul Go
├── cmd/
│   └── server/
│       └── main.go       # Titik masuk aplikasi (entry point)
├── internal/
│   ├── handlers/         # Handler permintaan HTTP
│   ├── models/           # Model & entitas domain
│   ├── services/         # Logika bisnis & operasi
│   └── storage/          # Migrasi basis data & repositori
├── views/                # Template HTML yang dirender server (SSR)
│   ├── layouts/
│   └── pages/
└── public/               # Aset statis (CSS, JS, gambar)
```

---

## Titik Masuk Aplikasi (`main.go`)

Workload `app` minimal berstandar produksi menggunakan `pkg/app` untuk supervisi daur hidup (lifecycle) dan `pkg/web` untuk perutean HTTP:

```go
package main

import (
	"context"
	"embed"
	"net/http"
	"os"

	"github.com/krewire/krewire/pkg/app"
	"github.com/krewire/krewire/pkg/web"
	"github.com/krewire/krewire/pkg/sec"
)

//go:embed public/* views/*
var embeddedAssets embed.FS

func main() {
	// Inisialisasi kontainer aplikasi Krewire
	application := app.New(app.Config{
		Name:    "storefront",
		Version: "0.1.0",
	})

	// Buat web engine
	router := web.NewRouter()

	// Terapkan header keamanan default dan middleware
	router.Use(sec.SecurityHeaders())
	router.Use(web.Recovery())
	router.Use(web.Logger())

	// Daftarkan route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Welcome to Storefront</h1>"))
	})

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		web.JSON(w, http.StatusOK, map[string]string{
			"status": "healthy",
			"version": "0.1.0",
		})
	})

	// Mount aset statis yang di-embed
	router.ServeFiles("/static/*", http.FS(embeddedAssets))

	// Jalankan dengan graceful shutdown
	if err := application.RunServer(":8080", router); err != nil {
		os.Exit(1)
	}
}
```

---

## Konfigurasi (`krewire.yaml`)

Konfigurasikan binding port server, output build, dan flag runtime di `krewire.yaml`:

```yaml
project:
  name: "storefront"
  kind: app
  version: "0.1.0"

dev:
  port: 8080

build:
  output: ".krewire/build"      # binary ditempatkan di .krewire/build/bin/storefront
  depends_on:
    - "go generate ./..."
```

Flag `--addr` menimpa nilai `dev.port` untuk satu pemanggilan (`kiw run --addr :3000`).

---

## Pengembangan & Operasional

### Pengembangan Lokal (`kiw dev`)

Jalankan aplikasi dalam mode pengembangan dengan pemantauan file otomatis (file watching) dan kompilasi ulang instan:

```bash
kiw dev
```

Setiap kali file `.go`, `.html`, atau aset berubah, `kiw` secara otomatis mengompilasi ulang dan me-restart proses server dalam hitungan milidetik.

### Melakukan Build untuk Produksi (`kiw build`)

Kompilasi binary siap produksi:

```bash
kiw build
```

Untuk cross-compile ke server Linux:

```bash
GOOS=linux GOARCH=amd64 kiw build
```

Binary yang dihasilkan di `bin/server` bersifat mandiri (standalone) dan siap di-deploy ke VPS Linux atau lingkungan kontainer apa pun.
