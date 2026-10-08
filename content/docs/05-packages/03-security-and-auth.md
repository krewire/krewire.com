---
title: "Security & Auth (sec, auth)"
description: "Hardened security defaults, OWASP Top 10 middleware, CSRF prevention, SSRF protection, PII redaction, JWT tokens, and basic authentication."
date: "2026-10-08"
---

# Security & Auth (`sec`, `auth`)

Import paths:
- `github.com/krewire/krewire/packages/sec`
- `github.com/krewire/krewire/packages/auth`

Security is Krewire's primary pillar. Rather than leaving hardening as an afterthought, `packages/sec` and `packages/auth` provide battle-tested, secure-by-default primitives aligned with OWASP Top 10 and CWE benchmarks.

---

## 1. Security Controls (`packages/sec`)

### Security Headers Middleware
Automatically applies modern HTTP defense headers:
- `X-Frame-Options: DENY`
- `X-Content-Type-Options: nosniff`
- `Referrer-Policy: strict-origin-when-cross-origin`
- `Content-Security-Policy (CSP)`
- `Strict-Transport-Security (HSTS)`

```go
import "github.com/krewire/krewire/packages/sec"

router.Use(sec.SecurityHeaders(sec.HeadersConfig{
    HSTSMaxAge: 31536000,
    FrameDeny:  true,
}))
```

### CSRF Protection
Double-submit cookie and synchronization token protection for forms and state-mutating requests (`POST`, `PUT`, `DELETE`).

```go
router.Use(sec.CSRF(sec.CSRFConfig{
    Secret:   os.Getenv("APP_KEY"),
    Secure:   true,
    SameSite: http.SameSiteStrictMode,
}))
```

### SSRF Protection
Validates outbound URLs before HTTP fetch, blocking private loopback and metadata addresses (`127.0.0.1`, `169.254.169.254`, `::1`):

```go
if err := sec.ValidateOutboundURL("https://example.com/webhook"); err != nil {
    // Blocked malicious SSRF attempt
}
```

### PII Redaction
Redacts sensitive identifiers (credit cards, passwords, tokens) in logs and error traces:

```go
safeLog := sec.MaskPII(rawUserData)
```

---

## 2. Authentication Primitives (`packages/auth`)

### JWT Authentication
Sign, parse, and verify cryptographically signed JSON Web Tokens using HMAC-SHA256:

```go
import "github.com/krewire/krewire/packages/auth"

// Sign token
token, err := auth.SignJWT(auth.Claims{
    Subject:   "user_123",
    ExpiresAt: time.Now().Add(24 * time.Hour),
}, []byte(secretKey))

// Middleware verification
router.Use(auth.JWTAuth(auth.JWTConfig{
    Secret: []byte(secretKey),
}))
```

### Basic Authentication
RFC 7617 HTTP Basic Authentication with constant-time password comparison to prevent timing attacks:

```go
router.Use(auth.BasicAuth("AdminRealm", map[string]string{
    "admin": os.Getenv("ADMIN_PASSWORD"),
}))
```
