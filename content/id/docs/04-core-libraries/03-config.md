---
title: "Konfigurasi Deklaratif"
description: "Pemuatan konfigurasi YAML tipe-aman dengan overlay variabel lingkungan dan pengikatan field struct di pkg/config."
date: "2026-10-25"
---

# Konfigurasi Deklaratif (`pkg/config`)

Jalur impor: `github.com/krewire/krewire/pkg/config`

Paket `pkg/config` menyediakan pemuatan konfigurasi bertipe kuat dari berkas YAML dengan *overlay* variabel lingkungan untuk seluruh ekosistem Krewire.

Urutan prioritas berlaku mutlak:
$$\text{Bawaan Struct} < \text{Berkas Konfigurasi YAML} < \text{Variabel Lingkungan}$$

---

## 1. Referensi API Utama

- **`config.Load(path string, dst any) error`**: Melakukan unmarshal berkas YAML pada `path` ke `dst` (harus berupa pointer non-nil ke struct). Jika berkas tidak ditemukan, membiarkan nilai `dst` tetap dan mengembalikan `nil`.
- **`config.LoadOrDefault(path string, dst any) error`**: Memuat YAML jika ada; mentoleransi ketiadaan berkas konfigurasi tanpa memicu error.
- **`config.Override(dst any, lookup func(string) (string, bool), opts ...Option) error`**: Menerapkan variabel lingkungan di atas struct yang telah dimuat sebelumnya.
- **`config.WithPrefix(prefix string) Option`**: Menyesuaikan awalan untuk kunci variabel lingkungan implisit (default: `APP`).

---

## 2. Pemetaan Tag Struct & Lingkungan

- Field dengan tag `yaml:"key"` terikat dari berkas konfigurasi YAML.
- Field dengan tag `env:"EXACT_KEY"` terikat langsung dari nama variabel lingkungan yang sama persis.
- Field tanpa tag menghasilkan kunci lingkungan implisit dari nama atau tag YAML yang diubah ke huruf kapital dengan awalan (misalnya `Server.Port` menjadi `APP_SERVER_PORT`).

---

## 3. Contoh Penggunaan

```go
package main

import (
	"fmt"
	"os"

	"github.com/krewire/krewire/pkg/config"
	"github.com/krewire/krewire/pkg/validation"
)

type ServerConfig struct {
	Host string `yaml:"host" env:"HOST"`
	Port int    `yaml:"port" env:"PORT" validate:"required,min=1,max=65535"`
}

type AppConfig struct {
	AppName string       `yaml:"app_name"`
	Server  ServerConfig `yaml:"server"`
	DBURL   string       `yaml:"db_url" env:"DATABASE_URL"`
}

func main() {
	// 1. Nilai bawaan yang aman
	cfg := AppConfig{
		AppName: "Layanan Krewire",
		Server: ServerConfig{
			Host: "127.0.0.1",
			Port: 8080,
		},
	}

	// 2. Muat berkas YAML
	if err := config.Load("krewire.yaml", &cfg); err != nil {
		panic(err)
	}

	// 3. Timpa dengan variabel lingkungan
	if err := config.Override(&cfg, os.LookupEnv, config.WithPrefix("APP")); err != nil {
		panic(err)
	}

	// 4. Validasi invarian
	if err := validation.Struct(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("%s berjalan di %s:%d\n", cfg.AppName, cfg.Server.Host, cfg.Server.Port)
}
```

---

## 4. Keselarasan Tiga Pilar (3S)

- **Secure (S1)**: Kunci rahasia, URL basis data, dan token diinjeksikan secara dinamis via variabel lingkungan (`${env:DATABASE_URL}`) tanpa pernah di-*commit* mentah ke repositori.
- **Sustainable (S2)**: Refleksi standar pustaka Go dan parser YAML murni; bebas dari framework konfigurasi eksternal yang membengkak.
- **Scalable (S3)**: Model konfigurasi tunggal terpadu yang dipakai bersama lintas server HTTP, perkakas CLI, worker, dan daemon latar belakang.
