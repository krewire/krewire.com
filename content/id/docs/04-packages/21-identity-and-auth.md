---
title: "Platform Identitas & Autentikasi Enterprise"
description: "Arsitektur, hashing PBKDF2, TOTP 2FA RFC 6238, SSO enterprise (OIDC/SAML2), RBAC, session store, proteksi CSRF, dan deployment server di krewire/auth."
date: "2026-10-31"
---

# Platform Identitas & Autentikasi Enterprise (`github.com/krewire/auth`)

`krewire/auth` adalah monorepo enterprise terdedikasi yang menyediakan spektrum penuh manajemen identitas, Single Sign-On (SSO), autentikasi multifaktor (MFA), tata kelola hak akses berbasis peran (RBAC), dan batas keamanan agen otonom di seluruh ekosistem Krewire.

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
- **Double-Submit CSRF Cookie:** Verifikasi token kriptografis cookie-ke-header sesuai ADR-D1F3N.
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

## Daemon & Deployment (`cmd/krewire-auth-server`)

Daemon mandiri berjalan sebagai server autentikasi terisolasi:

```bash
cd auth
go run ./cmd/krewire-auth-server
```

Atau lakukan build dan deployment menggunakan Dockerfile multi-stage produksi:

```bash
docker build -t krewire-auth-server ./auth
```
