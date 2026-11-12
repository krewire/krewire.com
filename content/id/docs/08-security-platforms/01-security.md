---
title: "Kontrol Keamanan & Hardening"
description: "Default keamanan tangguh, middleware OWASP Top 10, pencegahan CSRF, proteksi SSRF, dan redaksi PII di pkg/sec."
date: "2026-10-25"
---

# Kontrol Keamanan & Hardening (`pkg/sec`)

Jalur impor: `github.com/krewire/krewire/pkg/sec`

Keamanan adalah pilar utama Krewire. Alih-alih menjadikan pengamanan sebagai pemikiran sekunder, `pkg/sec` menyediakan primitif teruji yang aman secara default (*secure-by-default*) selaras dengan tolok ukur OWASP Top 10 dan CWE.

---

## Middleware Header Keamanan

Secara otomatis menerapkan header pertahanan HTTP modern:
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

---

## Proteksi CSRF

Proteksi double-submit cookie dan token sinkronisasi untuk formulir HTML dan permintaan perubahan status (`POST`, `PUT`, `DELETE`).

```go
router.Use(sec.CSRF(sec.CSRFConfig{
    Secret:   os.Getenv("APP_KEY"),
    Secure:   true,
    SameSite: http.SameSiteStrictMode,
}))
```

---

## Proteksi SSRF

Memvalidasi URL keluar sebelum panggilan HTTP dilakukan, memblokir alamat loopback privat dan IP metadata cloud (`127.0.0.1`, `169.254.169.254`, `::1`):

```go
if err := sec.ValidateOutboundURL("https://example.com/webhook"); err != nil {
    // Upaya SSRF berbahaya berhasil digagalkan
}
```

---

## Redaksi Data Sensitif (PII)

Menyamarkan pengenal sensitif (nomor kartu kredit, kata sandi, token API) di dalam log dan error trace:

```go
safeLog := sec.MaskPII(rawUserData)
```\n