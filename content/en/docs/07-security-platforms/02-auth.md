---
title: "Authentication Primitives"
description: "Cryptographic JWT signing and verification, HTTP Basic Authentication with constant-time password comparison in pkg/auth."
date: "2026-10-25"
---

# Authentication Primitives (`pkg/auth`)

Import path: `github.com/krewire/krewire/pkg/auth`

`pkg/auth` provides core cryptographic authentication primitives and token verification routines for APIs and web services.

---

## JWT Authentication

Sign, parse, and verify cryptographically signed JSON Web Tokens using HMAC-SHA256:

```go
package main

import (
	"time"

	"github.com/krewire/krewire/pkg/auth"
	"github.com/krewire/krewire/pkg/web"
)

func main() {
	secretKey := "my-secret-key"

	// Sign token
	token, err := auth.SignJWT(auth.Claims{
		Subject:   "user_123",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}, []byte(secretKey))

	// Middleware verification
	router := web.NewRouter()
	router.Use(auth.JWTAuth(auth.JWTConfig{
		Secret: []byte(secretKey),
	}))
}
```

---

## Basic Authentication

RFC 7617 HTTP Basic Authentication with constant-time password comparison to prevent timing attacks:

```go
router.Use(auth.BasicAuth("AdminRealm", map[string]string{
    "admin": os.Getenv("ADMIN_PASSWORD"),
}))
```\n