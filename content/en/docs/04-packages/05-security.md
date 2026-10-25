---
title: "Security Controls & Hardening"
description: "Hardened security defaults, OWASP Top 10 middleware, CSRF prevention, SSRF protection, and PII redaction in pkg/sec."
date: "2026-10-25"
---

# Security Controls & Hardening (`pkg/sec`)

Import path: `github.com/krewire/krewire/pkg/sec`

Security is Krewire's primary pillar. Rather than leaving hardening as an afterthought, `pkg/sec` provides battle-tested, secure-by-default primitives aligned with OWASP Top 10 and CWE benchmarks.

---

## Security Headers Middleware

Automatically applies modern HTTP defense headers:
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

## CSRF Protection

Double-submit cookie and synchronization token protection for forms and state-mutating requests (`POST`, `PUT`, `DELETE`).

```go
router.Use(sec.CSRF(sec.CSRFConfig{
    Secret:   os.Getenv("APP_KEY"),
    Secure:   true,
    SameSite: http.SameSiteStrictMode,
}))
```

---

## SSRF Protection

Validates outbound URLs before HTTP fetch, blocking private loopback and metadata addresses (`127.0.0.1`, `169.254.169.254`, `::1`):

```go
if err := sec.ValidateOutboundURL("https://example.com/webhook"); err != nil {
    // Blocked malicious SSRF attempt
}
```

---

## PII Redaction

Redacts sensitive identifiers (credit cards, passwords, API tokens) in logs and error traces:

```go
safeLog := sec.MaskPII(rawUserData)
```\n