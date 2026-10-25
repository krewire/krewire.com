---
title: "Platform Identitas & Autentikasi Enterprise"
description: "Arsitektur, hashing PBKDF2, TOTP 2FA RFC 6238, SSO enterprise (OIDC/SAML2), RBAC, dan web UI handler di krewire/auth."
date: "2026-10-24"
---

# Platform Identitas & Autentikasi Enterprise (`github.com/krewire/auth`)

`krewire/auth` adalah monorepo enterprise terdedikasi yang menyediakan spektrum penuh manajemen identitas, Single Sign-On (SSO), autentikasi multifaktor, dan tata kelola akses untuk ekosistem Krewire.

## Subsistem

1. **`pkg/identity`**:
   - Hashing kata sandi PBKDF2-HMAC-SHA256 sesuai rekomendasi OWASP (600.000 iterasi).
   - Time-based One-Time Passwords (TOTP RFC 6238) dengan URI provisi `otpauth://`.
   - Kode pemulihan darurat cadangan (8 kode) dengan validasi hash SHA-256.
   - Toko token reset kata sandi sekali pakai yang dibatasi waktu.
2. **`pkg/sso`**:
   - OIDC/OAuth2 generik dengan PKCE (RFC 7636).
   - Preset penyedia identitas bawaan: Google Workspace, GitHub, Microsoft Azure AD / Entra ID.
   - Pembuatan metadata SP SAML 2.0 dan validasi cryptographic assertion.
3. **`pkg/rbac`**:
   - Peran hierarkis (`superadmin` > `org_admin` > `developer` > `viewer`).
   - Izin granular `resource:action` dengan dukungan wildcard (`workloads:*`, `projects:deploy`).
   - Pemeriksaan isolasi batas multi-tenant (`evaluator.Can(ctx, user, action, tenant)`).
4. **`pkg/webauth`**:
   - Handler HTTP siap pasang dengan antarmuka web responsif dark/light:
     - `GET /login`, `POST /login`
     - `GET /register`, `POST /register`
     - `GET /forgot-password`, `POST /forgot-password`
     - `GET /reset-password`, `POST /reset-password`
     - `GET /2fa/challenge`, `POST /2fa/challenge`
     - `GET /2fa/setup`, `POST /2fa/setup`
     - `GET /logout`, `POST /logout`
