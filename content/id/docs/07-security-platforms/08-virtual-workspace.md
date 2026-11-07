---
title: "Workspace Virtual Native AI"
description: "Arsitektur, isolasi proses sandbox, terminal WebSocket PTY, asisten pair programmer, hub kolaborasi sesi, dan daemon di krewire/dash."
date: "2026-11-05"
---

# Workspace Virtual Native AI (`github.com/krewire/dash`)

`krewire/dash` menyediakan mesin workspace virtual berbasis AI dan platform cloud IDE. Paket ini menggabungkan isolasi proses yang hermetis, streaming terminal dua arah, editor web berbasis Monaco, dan asisten pair programmer `krew` tersemat ke dalam satu daemon server tunggal.

---

## Subsistem & Arsitektur

### 1. `pkg/sandbox` — Mesin Isolasi Hermetis
- **Process Jail:** Sandboxing kontainer ringan menggunakan Linux Bubblewrap (`bwrap`) dengan user namespace unprivileged.
- **Batasan Sumber Daya:** Kuota memori ketat, batasan berkas sementara, dan isolasi pohon proses untuk setiap workspace.
- **Sweeper Idle Otomatis:** Worker daemon background yang memantau keaktifan workspace dan menjeda sandbox secara aman setelah 30 menit tidak aktif.
- **Fallback Lintas Platform:** Penanganan proses yang aman dengan pemisahan atribut proses OS secara modular antara Linux dan macOS.

### 2. `pkg/terminal` — PTY WebSocket Dua Arah
- **Alokasi PTY Riil:** Alokasi pseudo-terminal kernel (`pty`) yang menjalankan shell bash/sh interaktif nyata.
- **Streaming WebSocket:** Latensi sub-10ms antara xterm.js di peramban web dan runner host PTY.
- **Sinkronisasi Ukuran Layar:** Penyesuaian dinamis jumlah kolom dan baris jendela terminal secara otomatis.

### 3. `pkg/session` — Hub Kolaborasi Real-Time
- **Multi-Tenant Hub:** Mengelola sesi workspace aktif dan koneksi klien yang terhubung secara bersamaan.
- **Kehadiran & Broadcast:** Menyelaraskan posisi kursor, pemilihan berkas, dan langkah penalaran agen secara langsung antar kolaborator.

### 4. `pkg/ai` — Kumpulan Asisten AI Tersemat
- **Konteks Sadar-Workspace:** Menghubungkan runtime agen `krew` langsung ke sistem berkas sandbox workspace aktif.
- **Pencarian Kode Semantik:** Navigasi berbasis AST kode untuk menemukan deklarasi simbol, interface, dan fungsi lintas package.
- **Review Diff Interaktif:** Menghasilkan unified diff visual terhadap usulan perubahan kode sebelum diterapkan ke berkas kerja.

### 5. `pkg/api` — API REST & Streaming Terpadu
- `POST /api/v1/workspaces`: Mengalokasikan workspace sandbox baru.
- `GET /api/v1/workspaces`: Menampilkan daftar workspace tenant yang sedang berjalan atau dijeda.
- `DELETE /api/v1/workspaces/{id}`: Menghapus dan membersihkan sumber daya workspace.
- `GET /api/v1/workspaces/{id}/files`: Menelusuri struktur direktori workspace.
- `GET /api/v1/workspaces/{id}/terminal/ws`: Terhubung ke terminal interaktif melalui WebSocket.
- `GET /api/v1/workspaces/{id}/session/ws`: Terhubung ke hub kolaborasi real-time.
- `GET /metrics`: Endpoint Prometheus melaporkan sandbox aktif, penggunaan CPU/RAM, dan sesi terminal.

---

## Akses Cloud Workspace & Deployment Enterprise

Krewire Workspace disediakan sebagai layanan cloud terkelola di `https://dash.krewire.com` dan dapat di-host mandiri pada lingkungan enterprise melalui kontainer:

### Akses Web Workspace Publik

Buka dan gunakan ruang kerja virtual Anda langsung melalui browser:
- **Portal Cloud:** `https://dash.krewire.com`
- **Autentikasi:** Single Sign-On terintegrasi via Krewire Auth (`auth.krewire.com`)

### Menjalankan Kontainer Enterprise Self-Hosted

```bash
docker run -d -p 9000:9000 \
  -e PORT=9000 \
  -e DASH_DATA_DIR="/data/workspaces" \
  -e AUTH_URL="https://auth.krewire.com" \
  -v /var/data/dash:/data/workspaces \
  ghcr.io/krewire/dash-server:latest
```
