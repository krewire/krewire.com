---
title: "Validasi Struct Deklaratif"
description: "Validasi struct dan skalar berbasis refleksi menggunakan tag validate murni di Go dalam pkg/validation."
date: "2026-10-25"
---

# Validasi Struct Deklaratif (`pkg/validation`)

Jalur impor: `github.com/krewire/krewire/pkg/validation`

Paket `pkg/validation` memvalidasi struct dan field skalar Go menggunakan tag struct `validate`. Pustaka ini beroperasi bebas framework dan 100% menggunakan pustaka standar Go, memungkinkan endpoint HTTP, perkakas CLI, dan entitas domain berbagi satu model aturan validasi terpadu.

---

## 1. Aturan yang Didukung

- **`required`**: Nilai tidak boleh nol, kosong, atau nil.
- **`min=N`**: Nilai numerik minimum atau panjang string/slice minimum.
- **`max=N`**: Nilai numerik maksimum atau panjang string/slice maksimum.
- **`len=N`**: Panjang eksak untuk string atau slice.
- **`email`**: Format alamat email yang valid sesuai RFC 5322.
- **`oneof=a b c`**: Nilai harus cocok dengan salah satu pilihan yang dipisahkan spasi.
- **`pattern=regex`**: Nilai harus cocok dengan ekspresi reguler (*regex*).

---

## 2. Referensi API Utama

- **`validation.Struct(v any) error`**: Memvalidasi struct atau pointer ke struct. Mengembalikan `nil` jika semua aturan lolos, atau `*validation.ValidationError` yang merinci setiap field yang gagal.
- **`validation.Field(value any, tag string) error`**: Memvalidasi satu nilai skalar terhadap string aturan.

---

## 3. Contoh Penggunaan

```go
package main

import (
	"errors"
	"fmt"

	"github.com/krewire/krewire/pkg/validation"
)

type UserRegistration struct {
	Username string `json:"username" validate:"required,min=3,max=32,pattern=^[a-zA-Z0-9_-]+$"`
	Email    string `json:"email"    validate:"required,email"`
	Role     string `json:"role"     validate:"required,oneof=developer admin observer"`
	Age      int    `json:"age"      validate:"min=13,max=120"`
}

func main() {
	user := UserRegistration{
		Username: "al",
		Email:    "invalid-email",
		Role:     "superuser",
		Age:      10,
	}

	err := validation.Struct(user)
	if err != nil {
		var valErr *validation.ValidationError
		if errors.As(err, &valErr) {
			for _, f := range valErr.Fields {
				fmt.Printf("Field %q gagal aturan %q (didapat: %v, parameter: %q)\n",
					f.Field, f.Rule, f.Value, f.Param)
			}
		}
	}
}
```

---

## 4. Keselarasan Tiga Pilar (3S)

- **Secure (S1)**: Invarian divalidasi langsung di batas terluar sebelum input pengguna yang tidak dipercaya menyentuh query basis data, eksekutor shell, atau penyimpanan berkas.
- **Sustainable (S2)**: Refleksi pustaka standar Go murni; tanpa dependensi pihak ketiga yang membengkak, bebas CGO.
- **Scalable (S3)**: Skema validasi terpadu yang dapat digunakan berulang kali pada handler web, argumen CLI, dan muatan pekerjaan worker.
