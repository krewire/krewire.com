---
title: "Kontrol Keamanan & Hardening"
description: "Default keamanan tangguh, sandbox SafePath, proteksi SSRF SafeHTTPClient, hashing PBKDF2, middleware OWASP, dan redaksi PII di pkg/sec."
date: "2026-10-25"
---

# Kontrol Keamanan & Hardening (`pkg/sec`)

Jalur impor: `github.com/krewire/krewire/pkg/sec`

Keamanan adalah pilar utama Krewire (**Pilar 1: Aman / S1: Keamanan Secara Bawaan**). Alih-alih menjadikan pengerasan sistem sebagai pemikiran tambahan atau fitur komersial berbayar, `pkg/sec` menghadirkan primitif teruji yang aman secara default (*secure-by-default*) selaras dengan tolok ukur OWASP Top 10 dan CWE.

---

## 1. Sandbox Traversal Direktori (`sec.SafePath`)

Operasi sistem berkas langsung menggunakan masukan pengguna rentan terhadap celah Directory Traversal (CWE-22). `SafePath` menegakkan isolasi batasan yang ketat:

- Menolak jalur yang mengandung byte null (`\x00`).
- Menormalisasi garis miring terbalik (backslash) Windows ke garis miring standar.
- Mengkanonikalkan urutan relatif (`..`, `.`).
- Memastikan berkas target hasil resolusi berada secara mutlak di dalam direktori root yang diizinkan.

```go
package main

import (
	"fmt"
	"os"

	"github.com/krewire/krewire/pkg/sec"
)

func ServeUserUpload(filename string) ([]byte, error) {
	baseDir := "/var/data/uploads"

	// Mengembalikan sec.ErrPathTraversal jika filename berusaha keluar dari baseDir
	safePath, err := sec.SafePath(baseDir, filename)
	if err != nil {
		return nil, fmt.Errorf("pelanggaran keamanan: %w", err)
	}

	return os.ReadFile(safePath)
}
```

---

## 2. Proteksi SSRF (`sec.SafeHTTPClient`)

Permintaan HTTP keluar (webhook, proxying, integrasi API eksternal) rentan terhadap serangan Server-Side Request Forgery (SSRF / CWE-918). `sec.SafeHTTPClient` mengintersepsi koneksi pada tingkat dialer soket jaringan:

- **Memblokir Rentang IP Privat (RFC 1918)**: `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`.
- **Memblokir Alamat Loopback**: `127.0.0.0/8`, `::1`.
- **Memblokir Endpoint Metadata Cloud**: `169.254.169.254` (AWS, GCP, Azure, DigitalOcean).
- **Mencegah DNS Rebinding**: Mengaudit dan memvalidasi IP tujuan tepat sebelum koneksi soket terbentuk.

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
		AllowPrivateIPs: false, // Bawaan: false
	})

	resp, err := client.Get(webhookURL)
	if err != nil {
		return err // Diblokir jika webhookURL merujuk ke IP privat atau metadata
	}
	defer resp.Body.Close()
	return nil
}
```

---

## 3. Hashing Kata Sandi Kriptografis & Pemeriksaan Constant-Time

Mengimplementasikan hashing password PBKDF2-HMAC-SHA256 dalam Go murni tanpa dependensi CGO:

```go
package main

import (
	"fmt"

	"github.com/krewire/krewire/pkg/sec"
)

func main() {
	rawPassword := "CorrectHorseBatteryStaple!"

	// 1. Hash dengan salt kriptografis 16-byte yang aman
	hash, err := sec.HashPasswordPBKDF2(rawPassword)
	if err != nil {
		panic(err)
	}

	// 2. Verifikasi timing-safe menggunakan crypto/subtle
	valid := sec.VerifyPasswordPBKDF2(rawPassword, hash)
	fmt.Printf("Kata sandi valid: %t\n", valid)
}
```

---

## 4. Middleware Header Pertahanan OWASP

Menyematkan header pertahanan dasar secara otomatis ke setiap respons HTTP:

- `X-Frame-Options: DENY` (Pencegahan Clickjacking)
- `X-Content-Type-Options: nosniff` (Pencegahan MIME sniffing)
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

## 5. Proteksi CSRF Double-Submit

Menyediakan verifikasi token cookie-ke-header untuk permintaan perubahan status (`POST`, `PUT`, `PATCH`, `DELETE`):

```go
csrfMiddleware := sec.CSRF(sec.CSRFConfig{
	Secret:   os.Getenv("SESSION_SECRET"),
	Secure:   true,
	SameSite: http.SameSiteStrictMode,
})
```

---

## 6. Penyamaran PII & Redaksi Log

Helper redaksi terstruktur untuk masukan pengguna, pelacakan galat, dan pencatatan log:

```go
// Menyamarkan kartu kredit, token API, token JWT, dan kata sandi
safeLog := sec.MaskPII(rawUserData)
```

---

## 7. Keselarasan dengan Tiga Pilar (3S)

- **Aman (S1 - Secure)**: Pencegahan aktif terhadap risiko OWASP Top 10: Path Traversal (CWE-22), SSRF (CWE-918), XSS, dan Clickjacking.
- **Berkelanjutan (S2 - Sustainable)**: 100% pure Go menggunakan primitif pustaka standar (`crypto/sha256`, `crypto/subtle`, `crypto/rand`, `net/http`) tanpa tautan binary native eksternal.
- **Skalabel (S3 - Scalable)**: Bersifat stateless dan dapat digunakan ulang di edge gateway, handler API, background worker, maupun perkakas CLI.
