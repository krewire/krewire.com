---
title: "Monolit Fullstack (app)"
description: "Panduan arsitektur dan implementasi untuk workload app: fondasi Go net/http, embedded asset, template SSR, dependency injection, dan operasional produksi."
date: "2026-10-31"
---

# Monolit Fullstack (`app`)

Workload **`app`** adalah fondasi Krewire untuk membangun monolit web tingkat produksi. Beban kerja ini mengompilasi seluruh permukaan web—perutean HTTP, tampilan yang dirender di sisi server (SSR), embedded static assets, typed dependency injection, konfigurasi, dan endpoint API JSON—ke dalam **satu binary Go mandiri**.

Dibangun di atas primitif standar Go (`net/http`, `embed.FS`, `context`, `slog`), workload `app` tidak memerlukan server web eksternal seperti Nginx atau runtime Node.js untuk pengembangan lokal maupun deployment dasar.

---

## Karakteristik Arsitektur

1. **Nol Dependensi Runtime:** Semua yang dibutuhkan untuk menjalankan aplikasi dikompilasi ke dalam satu file executable statis.
2. **Utamakan Pustaka Standar:** Sistem perutean, middleware, dan handler terintegrasi langsung dengan ekosistem antarmuka standar Go `net/http`.
3. **Embedded Static Assets:** File CSS, JavaScript, dan template dikompilasi langsung ke dalam binary menggunakan package `embed` bawaan Go, mencegah galat aset hilang (missing assets) di lingkungan produksi.
4. **Perakitan Dependensi Deterministik:** Memanfaatkan `pkg/app` untuk typed dependency injection, registrasi penyedia (`Register` dan `Boot`), serta siklus hidup yang terpantau (observable).
5. **Batas Domain yang Rapi:** Dibangun di sekitar pola modular monolit (`KWF-5ZHQV`), memisahkan konfigurasi, handler HTTP, presentasi SSR, dan logika domain secara bersih.

---

## Scaffolding Proyek `app`

Gunakan perintah `kiw new` (yang memilih varian `app` secara default atau dengan flag `--app`), atau lengkapi direktori yang ada dengan `kiw init --app`:

```bash
kiw new storefront --app
cd storefront
```

Struktur direktori yang dihasilkan mengikuti arsitektur modular monolit kanonikal Krewire:

```text
storefront/
├── krewire.yaml          # Konfigurasi proyek (kind: app)
├── go.mod                # Definisi modul Go
├── main.go               # Titik masuk aplikasi yang ringkas (entry point)
├── internal/
│   ├── app/
│   │   └── app.go        # Komposisi root: merakit provider ke dalam *web.App
│   ├── config/
│   │   └── config.go     # Konfigurasi bertipe dimuat dari krewire.yaml & env
│   └── http/
│       └── http.go       # Handler HTTP API JSON & registrasi rute
├── web/
│   ├── layouts/
│   │   └── shell.go      # Shell layout HTML dengan theme toggle
│   ├── pages/
│   │   └── pages.go      # Registrasi halaman SSR (web.PageSpec)
│   ├── components/
│   │   └── components.go # Komponen UI SSR (misalnya Hero)
│   └── theme/
│       └── theme.go      # Palet tema aplikasi (light / dark)
├── assets/
│   ├── embed.go          # Filesystem aset statis embedded (embed.FS)
│   └── public/
│       ├── app.css       # Style CSS klien
│       └── app.js        # Script JS klien
├── README.md             # Dokumentasi proyek
├── .gitignore            # Aturan eksklusi Git
├── .env.example          # Template variabel lingkungan
└── .env                  # File lingkungan lokal runtime
```

---

## Titik Masuk Aplikasi (`main.go`)

Dalam arsitektur modular Krewire, `main.go` dirancang sebagai titik masuk yang ramping: memuat konfigurasi bertipe, menginisialisasi kontainer aplikasi, dan menjalankan HTTP runner:

```go
// Titik masuk ringkas: muat konfigurasi, bangun app, jalankan.
package main

import (
	"fmt"
	"os"

	"storefront/internal/app"
	"storefront/internal/config"
)

// Signature adalah watermark ekosistem yang disematkan saat kompilasi via -ldflags.
var Signature = "build with krewire"

func main() {
	meta, err := config.LoadMetadata("krewire.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	a, err := app.New(meta)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := a.Run(meta.Addr()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

---

## Komponen Utama & Arsitektur

### 1. Komposisi Root (`internal/app/app.go`)

Komposisi root merakit semua subsistem provider (`configProvider`, `webProvider`, `http.Provider`) melalui `pkg/app`:

```go
package app

import (
	"io/fs"

	rvapp "github.com/krewire/krewire/pkg/app"
	rvweb "github.com/krewire/krewire/pkg/web"
	"storefront/assets"
	"storefront/internal/config"
	"storefront/internal/http"
	"storefront/web/components"
	"storefront/web/layouts"
	"storefront/web/pages"
	"storefront/web/theme"
)

// New merakit seluruh aplikasi dan mengembalikan instansi *web.App yang siap digunakan
func New(meta *config.Metadata) (*rvweb.App, error) {
	container, err := rvapp.NewApp(
		&configProvider{meta: meta},
		&webProvider{},
		&http.Provider{},
	).Build()
	if err != nil {
		return nil, err
	}
	return rvapp.Resolve[*rvweb.App](container)
}
```

### 2. Perutean HTTP & Handler API (`internal/http/http.go`)

Handler API HTTP mengikat data input, menjalankan validasi dengan `pkg/validation`, dan dipasang ke router aplikasi:

```go
package http

import (
	"net/http"

	rvapp "github.com/krewire/krewire/pkg/app"
	rvweb "github.com/krewire/krewire/pkg/web"
)

type Handler struct{}

func (h *Handler) Status(w http.ResponseWriter, _ *http.Request, _ rvweb.Params) {
	rvweb.JSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"version": "0.1.0",
	})
}

type Provider struct{}

func (p *Provider) Register(c *rvapp.Container) error {
	return rvapp.Singleton[*Handler](c, func() *Handler { return &Handler{} })
}

func (p *Provider) Boot(c *rvapp.Container) error {
	a, err := rvapp.Resolve[*rvweb.App](c)
	if err != nil {
		return err
	}
	h, err := rvapp.Resolve[*Handler](c)
	if err != nil {
		return err
	}
	a.Router().Get("/api/status", h.Status)
	return nil
}
```

### 3. Server-Side Rendering (`web/pages/pages.go`)

Halaman didaftarkan secara deklaratif menggunakan spesifikasi `web.PageSpec` yang mendukung data statis maupun resolver dinamis per-permintaan:

```go
package pages

import (
	"net/http"
	"time"

	"github.com/krewire/krewire/pkg/web"
)

func Register(a *web.App) {
	a.Page(web.PageSpec{
		Path:   "/",
		Title:  "Storefront",
		Layout: "Shell",
		Root:   "Hero",
		Data: map[string]any{
			"badge":    "v0.1.0",
			"title":    "Selamat Datang di Storefront",
			"subtitle": "SSR, API JSON, dan embedded assets dari satu binary Go.",
		},
		EmitProps: true,
		Scripts:   []string{"/static/app.js"},
	})

	a.Page(web.PageSpec{
		Path:   "/greet",
		Title:  "Greet",
		Layout: "Shell",
		Root:   "Hero",
		DataFunc: func(r *http.Request) (any, error) {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "world"
			}
			return map[string]any{
				"badge":    time.Now().Format("15:04:05"),
				"title":    "Halo, " + name,
				"subtitle": "Halaman ini dirender dinamis per permintaan.",
			}, nil
		},
	})
}
```

---

## Konfigurasi (`krewire.yaml`)

Konfigurasikan port server, metadata proyek, dan target opsional di `krewire.yaml`:

```yaml
project:
  name: "storefront"
  kind: app
  version: "0.1.0"

dev:
  port: 8080

build:
  output: ".krewire/build"
```

Alamat runtime juga dapat disesuaikan melalui variabel lingkungan `APP_ADDR` atau flag `--addr` (contoh: `kiw run --addr :3000`).

---

## Pengembangan & Operasional

### Pengembangan Lokal (`kiw dev` atau `kiw run`)

Jalankan aplikasi dengan pemantauan perubahan file dan kompilasi otomatis:

```bash
kiw dev
```

Atau jalankan aplikasi secara langsung:

```bash
kiw run
```

### Menjalankan Pengujian (`kiw test`)

Jalankan test suite unit dan integrasi:

```bash
kiw test ./...
```

### Melakukan Build untuk Produksi (`kiw build`)

Kompilasi binary siap produksi:

```bash
kiw build
```

Binary yang dihasilkan ditempatkan di `.krewire/build/bin/storefront` (atau cross-compiled untuk platform target dengan variabel env standar Go seperti `GOOS=linux GOARCH=amd64`).
