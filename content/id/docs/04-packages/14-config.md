---
title: "Konfigurasi Hierarkis"
description: "Pemuatan konfigurasi aplikasi hierarkis dengan flag CLI, variabel lingkungan, dan dukungan dotenv di pkg/config."
date: "2026-10-25"
---

# Konfigurasi Hierarkis (`pkg/config`)

Jalur impor: `github.com/krewire/krewire/pkg/config`

Paket `config` memuat pengaturan aplikasi bertipe dengan urutan prioritas deterministik:
1. Flag CLI eksplisit
2. Variabel lingkungan (awalan `APP_*`)
3. Berkas lingkungan lokal (`.env.local`)
4. Berkas lingkungan dasar (`.env`)
5. Berkas konfigurasi (`krewire.yaml` / JSON)
6. Tag default pada field struct

---

## Contoh Penggunaan

```go
package main

import (
	"fmt"
	"github.com/krewire/krewire/pkg/config"
)

type AppConfig struct {
	Port     int    `env:"PORT" default:"8080"`
	DBHost   string `env:"DB_HOST" default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" default:"info"`
}

func main() {
	var cfg AppConfig
	// Memuat dari .env, env var, dan default secara otomatis
	if err := config.Load(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("Server mendengarkan di port %d dengan log level %s\n", cfg.Port, cfg.LogLevel)
}
```\n