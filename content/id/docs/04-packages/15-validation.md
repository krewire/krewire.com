---
title: "Validasi Input Komposabel"
description: "Validasi struktur data komposabel tanpa overhead refleksi, menyediakan error terstruktur untuk API dan form di pkg/validation."
date: "2026-10-25"
---

# Validasi Input Komposabel (`pkg/validation`)

Jalur impor: `github.com/krewire/krewire/pkg/validation`

Paket `validation` memvalidasi struktur data tanpa overhead refleksi runtime yang berat, menyediakan struktur error yang jelas untuk API dan form web.

---

## Aturan Umum

- **`rules.Required`**: Nilai tidak boleh kosong atau nol.
- **`rules.Email`**: Format email valid sesuai standar RFC 5322.
- **`rules.Min` / `rules.Max`**: Batas numerik atau panjang karakter string minimum/maksimum.
- **`rules.OneOf`**: Nilai harus cocok dengan salah satu opsi yang diizinkan.
- **`rules.Pattern`**: Validasi ekspresi reguler (regex).

---

## Contoh Penggunaan

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
```\n