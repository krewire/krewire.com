---
title: "Platform PaaS Hibrida"
description: "Arsitektur, runner, supervisor proses, dan router edge reverse proxy dinamis di krewire/cloud."
date: "2026-10-24"
---

# Platform PaaS Hibrida (`github.com/krewire/cloud`)

`krewire/cloud` mentenagai mesin PaaS open-core yang di-host mandiri (self-hosted) sekaligus platform SaaS multi-tenant terkelola (`cloud.krewire.com`), mengimplementasikan [ADR-C9L2P](file:///home/reasvyn/workspace/krewire/internal/docs/adr/ADR-C9L2P-hybrid-paas-platform-architecture.md).

## Subsistem

1. **`pkg/controlplane`**:
   - Mengelola akun tenant, proyek, riwayat deployment, dan secret variabel lingkungan.
   - Menegakkan batasan paket dan kuota (`free`, `pro`, `team`, `enterprise`).
2. **`pkg/runner`**:
   - Menjalankan kompilasi Go hermetis (`CGO_ENABLED=0`, `-trimpath`, `-s -w`).
   - Menghasilkan binary mandiri yang ringkas dan berjalan dengan memori ~15MB RAM.
3. **`pkg/supervisor`**:
   - Supervisor pool proses dinamis yang mengalokasikan port localhost terisolasi.
   - Menjalankan probe health check otomatis sebelum mengalirkan trafik.
   - Rolling deploy berkesinambungan tanpa downtime (zero-downtime).
4. **`pkg/gateway`**:
   - Router edge reverse proxy dinamis.
   - Menyelesaikan subdomain (`<app-slug>.krewire.dev`) dan domain kustom.
   - Hook provisi sertifikat TLS otomatis melalui ACME.
5. **`cmd/cloud-server`**:
   - Daemon binary tunggal yang dapat dideploy di VPS Linux mana pun hanya dengan satu perintah:
     ```bash
     go run ./cmd/cloud-server
     ```
