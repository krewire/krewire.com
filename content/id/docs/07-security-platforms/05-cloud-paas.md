---
title: "Platform Hybrid PaaS"
description: "Arsitektur, runner, supervisi proses, WebAuth console UI, webhook Git push, dan reverse proxy edge router di Krewire Cloud."
date: "2026-10-31"
---

# Platform Hybrid PaaS (Krewire Cloud)

Krewire Cloud menenagai mesin PaaS hibrida mandiri (self-hosted) sekaligus platform SaaS cloud multi-tenant terkelola (`cloud.krewire.com`).

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

## Akses Platform & Deployment

Krewire Cloud disediakan sebagai control plane PaaS multi-tenant terkelola di [cloud.krewire.com](https://cloud.krewire.com), dengan edisi on-premise khusus yang tersedia untuk pelanggan enterprise.

### Men-deploy Aplikasi Anda

Anda dapat men-deploy proyek Krewire apa pun secara langsung menggunakan CLI `kiw`:

```bash
# Menyiapkan binary produksi dan artefak siap rilis
kiw deploy --target binary

# Membuat definisi layanan systemd untuk server Linux
kiw deploy --target systemd

# Untuk situs statis dan buku dokumentasi (GitHub Pages)
kiw deploy --target gh-pages
```

### Kontainer Docker Produksi

Untuk mengemas aplikasi Krewire ke dalam kontainer mandiri yang sangat ringan (<15MB) untuk dideploy ke VPS, host Docker, atau cluster Kubernetes:

```dockerfile
# Build multi-stage untuk workload Krewire
FROM golang:1.27-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/server .

FROM alpine:3.20
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/server /app/server
EXPOSE 8080
ENTRYPOINT ["/app/server"]
```

### PaaS Enterprise On-Premise

Bagi organisasi yang membutuhkan hosting privat air-gapped atau control plane mandiri di infrastruktur internal, Krewire Cloud disediakan melalui langganan enterprise. Hubungi tim enterprise melalui `enterprise@krewire.com` untuk akses registri privat dan panduan provisi.
