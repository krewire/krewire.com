---
title: "Keamanan & Otentikasi (sec, auth)"
description: "Standar keamanan bawaan yang kokoh, middleware OWASP Top 10, pencegahan CSRF, perlindungan SSRF, penyamaran PII, token JWT, dan autentikasi dasar."
date: "2026-10-08"
---

# Keamanan & Otentikasi (`sec`, `auth`)

Jalur impor (*import paths*):
- `github.com/krewire/krewire/pkg/sec`
- `github.com/krewire/krewire/pkg/auth`

Keamanan adalah pilar utama Krewire. Alih-alih memikirkan pengerasan sistem (*hardening*) belakangan, `pkg/sec` dan `pkg/auth` menyediakan primitif teruji yang aman secara default dan selaras dengan standar OWASP Top 10 serta benchmark CWE.

---

## 1. Kontrol Keamanan (`pkg/sec`)

### Middleware Security Headers
Menerapkan header proteksi HTTP modern secara otomatis:
- `X-Frame-Options: DENY`
- `X-Content-Type-Options: nosniff`
- `Referrer-Policy: strict-origin-when-cross-origin`
- `Content-Security-Policy (CSP)`
- `Strict-Transport-Security (HSTS)`

```go
import "github.com/krewire/krewire/pkg/sec"

router.Use(sec.SecurityHeaders(sec.HeadersConfig{
    HSTSMaxAge: 31536000,
    FrameDeny:  true,
}))
```

### Proteksi CSRF
Perlindungan double-submit cookie dan token sinkronisasi untuk formulir dan permintaan pengubah state (`POST`, `PUT`, `DELETE`).

```go
router.Use(sec.CSRF(sec.CSRFConfig{
    Secret:   os.Getenv("APP_KEY"),
    Secure:   true,
    SameSite: http.SameSiteStrictMode,
}))
```

### Proteksi SSRF
Memvalidasi URL keluar sebelum dilakukan HTTP fetch, memblokir alamat loopback privat dan endpoint metadata cloud (`127.0.0.1`, `169.254.169.254`, `::1`):

```go
if err := sec.ValidateOutboundURL("https://example.com/webhook"); err != nil {
    // Memblokir percobaan serangan SSRF berbahaya
}
```

### Penyamaran PII (PII Redaction)
Menyamarkan data sensitif pengguna (kartu kredit, kata sandi, token) di dalam log dan trace error:

```go
safeLog := sec.MaskPII(rawUserData)
```

---

## 2. Primitif Otentikasi (`pkg/auth`)

### Otentikasi JWT
Menandatangani, mem-parsing, dan memvalidasi JSON Web Tokens (JWT) bertanda tangan kriptografis menggunakan HMAC-SHA256:

```go
import "github.com/krewire/krewire/pkg/auth"

// Menandatangani token
token, err := auth.SignJWT(auth.Claims{
    Subject:   "user_123",
    ExpiresAt: time.Now().Add(24 * time.Hour),
}, []byte(secretKey))

// Verifikasi middleware
router.Use(auth.JWTAuth(auth.JWTConfig{
    Secret: []byte(secretKey),
}))
```

### Otentikasi Dasar (Basic Auth)
Otentikasi HTTP Basic RFC 7617 dengan perbandingan kata sandi waktu-konstan (*constant-time comparison*) guna mencegah serangan timing attack:

```go
router.Use(auth.BasicAuth("AdminRealm", map[string]string{
    "admin": os.Getenv("ADMIN_PASSWORD"),
}))
```
