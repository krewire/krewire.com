---
title: "Platform Hybrid PaaS"
description: "Arsitektur, runner, supervisi proses, WebAuth console UI, webhook Git push, dan reverse proxy edge router di krewire/cloud."
date: "2026-10-31"
---

# Platform Hybrid PaaS (`github.com/krewire/cloud`)

`krewire/cloud` menenagai mesin PaaS open-core mandiri (self-hosted) sekaligus platform SaaS cloud multi-tenant terkelola (`cloud.krewire.com`), mengimplementasikan `ADR-C9L2P`.

---

## Subsistem & Arsitektur

### 1. `pkg/controlplane`
- Mengelola akun tenant, proyek, catatan deployment, dan secret variabel lingkungan.
- Menegakkan batas paket dan kuota akun (`free`, `pro`, `team`, `enterprise`).
- Kepatuhan isolasi data dan audit hukum Perlindungan Data Pribadi (UU PDP) Indonesia.

### 2. `pkg/runner`
- Menjalankan kompilasi Go hermetis (`CGO_ENABLED=0`, `-trimpath`, `-s -w`).
- Menghasilkan binary mandiri yang hemat sumber daya (~15MB RAM) tanpa dependensi eksternal.

### 3. `pkg/supervisor`
- Supervisor kolam proses dinamis yang mengalokasikan port localhost terisolasi.
- Probe health check otomatis sebelum pengalihan lalu lintas jaringan (traffic cutover).
- Rolling deployment tanpa jeda (zero-downtime rolling deploys).

### 4. `pkg/gateway`
- Edge router reverse proxy dinamis.
- Provisi sertifikat TLS otomatis Let's Encrypt ACME untuk subdomain dan custom domain.
- Perutean tepi HTTP/WebSocket berbasis host serta connection pooling.

### 5. `pkg/api` & Konsol Web
- **WebAuth Console UI:** Antarmuka web terintegrasi untuk login, SSO, dan 2FA TOTP (`/webauth/login`, `/webauth/sso`, `/webauth/mfa`).
- **Webhook Continuous Deployment:** Listener webhook Git push otomatis (`/api/v1/webhooks/git`) yang memicu pipeline kompilasi dan rollout instan.
- **Streaming Log Real-Time:** Pengiriman log build dan eksekusi kontainer langsung melalui Server-Sent Events (SSE) di `/api/v1/apps/{id}/logs`.

---

## Daemon & Deployment (`cmd/krewire-cloud-server`)

Daemon binary tunggal dapat dijalankan di VPS Linux apa pun (Ubuntu, Debian, Alpine):

```bash
cd cloud
go run ./cmd/krewire-cloud-server
```

Atau lakukan deployment kontainer menggunakan Dockerfile multi-stage produksi:

```bash
docker build -t krewire-cloud-server ./cloud
```
