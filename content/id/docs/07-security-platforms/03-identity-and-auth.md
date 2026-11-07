---
title: "Platform Identitas & Autentikasi Enterprise"
description: "Arsitektur, hashing PBKDF2, TOTP 2FA RFC 6238, SSO enterprise (OIDC/SAML2), RBAC, session store, proteksi CSRF, dan integrasi di Krewire Auth."
date: "2026-10-31"
---

# Platform Identitas & Autentikasi Enterprise (Krewire Auth)

Krewire Auth adalah platform identitas enterprise terdedikasi yang menyediakan spektrum penuh manajemen identitas, Single Sign-On (SSO), autentikasi multifaktor (MFA), tata kelola hak akses berbasis peran (RBAC), dan batas keamanan agen otonom di seluruh ekosistem Krewire.

---

## Subsistem & Arsitektur

### 1. `pkg/identity`
- **Hashing Kata Sandi:** PBKDF2-HMAC-SHA256 sesuai standar OWASP (600.000 iterasi).
- **Autentikasi Multifaktor:** Time-based One-Time Passwords (TOTP RFC 6238) dengan kode QR dan URI provisi `otpauth://`.
- **Kode Pemulihan Cadangan:** 8 kode pemulihan darurat sekali pakai yang divalidasi hash SHA-256.
- **Penyimpanan Persisten:** SQL user store dan session store persisten dengan dukungan PostgreSQL dan SQLite.
- **Redis Session Store:** `RedisSessionStore` throughput tinggi yang berkomunikasi via protokol asli RESP dengan pembersihan TTL otomatis.

### 2. `pkg/sso`
- **OIDC / OAuth2:** OIDC/OAuth2 generik dengan PKCE (RFC 7636) dan verifikasi status (state).
- **Preset Penyedia:** Driver siap pakai untuk Google Workspace, GitHub, dan Microsoft Azure AD / Entra ID.
- **SAML 2.0:** Pembuatan metadata SP, validasi cryptographic assertion toleran clock-skew, dan verifikasi tanda tangan digital.
- **OIDC Discovery:** Endpoint metadata publik (`/.well-known/openid-configuration`) dan manajemen kunci JWKS.

### 3. `pkg/rbac`
- **Peran Pengguna:** Hierarki pengguna (`superadmin` > `org_admin` > `developer` > `viewer`).
- **Izin Granular:** Aturan resource-action dengan pencocokan wildcard (`workloads:*`, `projects:deploy`, `secrets:read`).
- **RBAC Agen & Tool:** Hak akses terdedikasi untuk AI agent otonom (`AgentPermission`, `ToolPermission`) dengan evaluator kebijakan untuk mencegah eksekusi tool tanpa wewenang.
- **Audit Trail Keamanan:** `AuditLogger` terstruktur yang mencatat seluruh aksi autentikasi, otorisasi, dan administratif beserta konteks IP dan pengguna.

### 4. `pkg/middleware`
- **Double-Submit CSRF Cookie:** Verifikasi token kriptografis cookie-ke-header sesuai standar keamanan berlapis (defense-in-depth).
- **Konteks Sesi:** Helper ekstraksi sesi yang aman (`WithSession`, `SessionFromContext`) terintegrasi langsung dengan siklus permintaan HTTP.

### 5. `pkg/webauth`
Handler HTTP siap pasang dengan antarmuka web responsif dark/light:
- `GET /login`, `POST /login`
- `GET /register`, `POST /register`
- `GET /forgot-password`, `POST /forgot-password`
- `GET /reset-password`, `POST /reset-password`
- `GET /2fa/challenge`, `POST /2fa/challenge`
- `GET /2fa/setup`, `POST /2fa/setup`
- `GET /logout`, `POST /logout`

---

## Akses Platform & Integrasi

Krewire Auth tersedia sebagai platform identitas terkelola di [auth.krewire.com](https://auth.krewire.com), dengan edisi on-premise privat yang tersedia di bawah lisensi enterprise.

### Integrasi Standar OIDC & OAuth2

Setiap aplikasi, API gateway, atau layanan Krewire dapat terhubung secara mulus ke Krewire Auth melalui protokol standar industri:
- **Discovery Endpoint:** `https://auth.krewire.com/.well-known/openid-configuration`
- **JWKS Key Set:** `https://auth.krewire.com/.well-known/jwks.json`
- **Authorization Endpoint:** `https://auth.krewire.com/oauth/authorize`
- **Token Endpoint:** `https://auth.krewire.com/oauth/token`

### Keamanan Lokal untuk Workload Open-Source

Bagi proyek open-source yang berjalan mandiri tanpa layanan cloud, Krewire menyediakan library autentikasi dan keamanan native tanpa dependensi di monorepo inti:
- `pkg/sec`: Perbandingan constant-time, header keamanan OWASP, validasi token CSRF, dan pencegahan SSRF.
- `pkg/auth`: Token JWT stateless, hashing password PBKDF2, dan middleware HTTP Basic Auth.

```go
import (
    "github.com/krewire/krewire/pkg/auth"
    "github.com/krewire/krewire/pkg/sec"
)
```

### Autentikasi Enterprise On-Premise

Organisasi yang membutuhkan isolasi identitas di server lokal, federasi SAML 2.0 dengan Active Directory internal, atau audit kepatuhan khusus dapat melisensikan server Krewire Auth mandiri melalui `enterprise@krewire.com`.
