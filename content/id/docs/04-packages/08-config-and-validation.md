---
title: "Konfigurasi & Validasi (config, validation)"
description: "Panduan referensi pemuatan konfigurasi hierarkis, override environment, dukungan dotenv, dan aturan validasi yang dapat dikomposisikan."
date: "2026-10-10"
---

# Konfigurasi & Validasi (`config`, `validation`)

Jalur impor (*import paths*):
- `github.com/krewire/krewire/pkg/config`
- `github.com/krewire/krewire/pkg/validation`

Pemuatan konfigurasi yang tangguh dan validasi input yang ketat mencegah error saat runtime serta celah keamanan sebelum permintaan masuk ke logika domain inti.

---

## 1. Konfigurasi Hierarkis (`pkg/config`)

Paket `config` memuat pengaturan aplikasi bertipe (*typed settings*) dengan prioritas deterministik:
1. Flag eksplisit CLI
2. Variabel environment (awalan prefix `APP_*`)
3. Berkas environment lokal (`.env.local`)
4. Berkas environment dasar (`.env`)
5. Berkas konfigurasi (`krewire.yaml` / JSON)
6. Tag nilai default pada field struct

### Contoh Penggunaan

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
	// Memuat dari .env, variabel environment, dan nilai default secara otomatis
	if err := config.Load(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("Server berjalan di port %d dengan level log %s\n", cfg.Port, cfg.LogLevel)
}
```

---

## 2. Validasi Input yang Dapat Dikomposisikan (`pkg/validation`)

Paket `validation` memvalidasi struktur data tanpa beban performa reflection berlebih, menyajikan struktur error yang terang untuk API dan formulir web.

### Aturan Umum (*Common Rules*):
- **`rules.Required`**: Nilai tidak boleh kosong atau bernilai nol.
- **`rules.Email`**: Format email yang mematuhi standar RFC 5322.
- **`rules.Min` / `rules.Max`**: Batas numerik atau batas minimum/maksimum panjang string.
- **`rules.OneOf`**: Nilai harus cocok dengan salah satu opsi yang diizinkan.
- **`rules.Pattern`**: Validasi berbasis ekspresi reguler (regex).

### Contoh Penggunaan

```go
package main

import (
	"fmt"
	"github.com/krewire/krewire/pkg/validation"
	"github.com/krewire/krewire/pkg/validation/rules"
)

type RegisterRequest struct {
	Username string
	Email    string
	Age      int
}

func (r RegisterRequest) Validate() validation.Errors {
	v := validation.New()

	v.Check("username", r.Username, rules.Required(), rules.Min(3), rules.Max(32))
	v.Check("email", r.Email, rules.Required(), rules.Email())
	v.Check("age", r.Age, rules.Min(18))

	return v.Errors()
}

func main() {
	req := RegisterRequest{Username: "al", Email: "invalid-email", Age: 16}
	errs := req.Validate()

	if errs.HasErrors() {
		for field, msg := range errs.All() {
			fmt.Printf("- %s: %s\n", field, msg)
		}
	}
}
```
