---
title: "Primitif Autentikasi"
description: "Penandatanganan dan verifikasi kriptografis JWT, HTTP Basic Authentication dengan perbandingan waktu konstan di pkg/auth."
date: "2026-10-25"
---

# Primitif Autentikasi (`pkg/auth`)

Jalur impor: `github.com/krewire/krewire/pkg/auth`

`pkg/auth` menyediakan primitif autentikasi kriptografis inti dan rutinitas verifikasi token untuk API dan layanan web.

---

## Autentikasi JWT

Membuat tanda tangan, mengurai, dan memverifikasi JSON Web Token secara kriptografis menggunakan HMAC-SHA256:

```go
package main

import (
	"time"

	"github.com/krewire/krewire/pkg/auth"
	"github.com/krewire/krewire/pkg/web"
)

func main() {
	secretKey := "my-secret-key"

	// Menandatangani token
	token, err := auth.SignJWT(auth.Claims{
		Subject:   "user_123",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}, []byte(secretKey))

	// Verifikasi middleware
	router := web.NewRouter()
	router.Use(auth.JWTAuth(auth.JWTConfig{
		Secret: []byte(secretKey),
	}))
}
```

---

## Basic Authentication

Autentikasi dasar RFC 7617 HTTP Basic Authentication dengan perbandingan waktu konstan (*constant-time comparison*) guna mencegah serangan timing:

```go
router.Use(auth.BasicAuth("AdminRealm", map[string]string{
    "admin": os.Getenv("ADMIN_PASSWORD"),
}))
```\n