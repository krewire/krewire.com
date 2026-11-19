---
title: "Security Controls & Hardening"
description: "Hardened security defaults, SafePath sandbox, SafeHTTPClient SSRF defense, PBKDF2 hashing, OWASP middleware, and PII redaction in pkg/sec."
date: "2026-10-25"
---

# Security Controls & Hardening (`pkg/sec`)

Import path: `github.com/krewire/krewire/pkg/sec`

Security is Krewire's primary pillar (**Pillar 1: Secure / S1: Safety by Default**). Rather than treating hardening as an afterthought or an enterprise add-on, `pkg/sec` delivers battle-tested, secure-by-default primitives aligned with OWASP Top 10 and CWE benchmarks.

---

## 1. Directory Traversal Sandbox (`sec.SafePath`)

Direct filesystem operations with user input risk Directory Traversal vulnerabilities (CWE-22). `SafePath` enforces strict boundary sandboxing:

- Rejects paths containing null bytes (`\x00`).
- Normalizes Windows backslashes to standard forward slashes.
- Canonicalizes relative sequences (`..`, `.`).
- Confirms the resolved target path resides strictly inside the designated root directory.

```go
package main

import (
	"fmt"
	"os"

	"github.com/krewire/krewire/pkg/sec"
)

func ServeUserUpload(filename string) ([]byte, error) {
	baseDir := "/var/data/uploads"

	// Returns sec.ErrPathTraversal if filename escapes baseDir
	safePath, err := sec.SafePath(baseDir, filename)
	if err != nil {
		return nil, fmt.Errorf("security violation: %w", err)
	}

	return os.ReadFile(safePath)
}
```

---

## 2. SSRF Protection (`sec.SafeHTTPClient`)

Outbound HTTP requests (webhooks, proxying, external API integrations) are vulnerable to Server-Side Request Forgery (SSRF / CWE-918). `sec.SafeHTTPClient` intercepts connections at the network dialer level:

- **Blocks Private IP Ranges (RFC 1918)**: `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`.
- **Blocks Loopback Addresses**: `127.0.0.0/8`, `::1`.
- **Blocks Cloud Metadata Endpoints**: `169.254.169.254` (AWS, GCP, Azure, DigitalOcean).
- **Prevents DNS Rebinding**: Resolves and audits destination IPs immediately before socket establishment.

```go
package main

import (
	"context"
	"time"

	"github.com/krewire/krewire/pkg/sec"
)

func DeliverWebhook(ctx context.Context, webhookURL string) error {
	client := sec.NewSafeHTTPClient(sec.SafeClientConfig{
		Timeout:         5 * time.Second,
		AllowPrivateIPs: false, // Default: false
	})

	resp, err := client.Get(webhookURL)
	if err != nil {
		return err // Blocked if webhookURL resolves to private or metadata IP
	}
	defer resp.Body.Close()
	return nil
}
```

---

## 3. Cryptographic Password Hashing & Constant-Time Checks

Implements pure-Go PBKDF2-HMAC-SHA256 password hashing with zero CGO dependencies:

```go
package main

import (
	"fmt"

	"github.com/krewire/krewire/pkg/sec"
)

func main() {
	rawPassword := "CorrectHorseBatteryStaple!"

	// 1. Hash with secure 16-byte cryptographic salt
	hash, err := sec.HashPasswordPBKDF2(rawPassword)
	if err != nil {
		panic(err)
	}

	// 2. Timing-safe verification using crypto/subtle
	valid := sec.VerifyPasswordPBKDF2(rawPassword, hash)
	fmt.Printf("Password valid: %t\n", valid)
}
```

---

## 4. OWASP Defense Headers Middleware

Automatically injects baseline defense headers into all HTTP responses:

- `X-Frame-Options: DENY` (Clickjacking prevention)
- `X-Content-Type-Options: nosniff` (MIME sniffing prevention)
- `Referrer-Policy: strict-origin-when-cross-origin`
- `Content-Security-Policy (CSP)`
- `Strict-Transport-Security (HSTS)`

```go
mux := http.NewServeMux()
handler := sec.SecurityHeadersMiddleware(sec.SecurityHeadersConfig{
	ContentSecurityPolicy: "default-src 'self';",
	FrameOptions:          "DENY",
	ContentTypeOptions:    "nosniff",
	ReferrerPolicy:        "strict-origin-when-cross-origin",
	HSTSMaxAgeSeconds:     31536000,
})(mux)
```

---

## 5. Double-Submit CSRF Protection

Provides cookie-to-header token verification for state-mutating requests (`POST`, `PUT`, `PATCH`, `DELETE`):

```go
csrfMiddleware := sec.CSRF(sec.CSRFConfig{
	Secret:   os.Getenv("SESSION_SECRET"),
	Secure:   true,
	SameSite: http.SameSiteStrictMode,
})
```

---

## 6. PII Masking & Log Redaction

Structured redaction helper for user inputs, error traces, and logs:

```go
// Masks credit cards, API keys, JWT tokens, and passwords
safeLog := sec.MaskPII(rawUserData)
```

---

## 7. Alignment with the Three Pillars (3S)

- **Secure (S1)**: Active prevention against OWASP Top 10 risks: Path Traversal (CWE-22), SSRF (CWE-918), XSS, and Clickjacking.
- **Sustainable (S2)**: 100% pure Go using standard library primitives (`crypto/sha256`, `crypto/subtle`, `crypto/rand`, `net/http`) with zero native linkages.
- **Scalable (S3)**: Stateless and reusable across edge gateways, API handlers, background workers, and CLI devtools.
