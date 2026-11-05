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

## Daemon & Deployment (`cmd/krewire-dash-server`)

Jalankan daemon workspace virtual secara lokal maupun di server produksi:

```bash
cd dash
go run ./cmd/krewire-dash-server
```

Konfigurasi via environment variables:
- `PORT`: Port listen HTTP (default: `9000`).
- `DASH_DATA_DIR`: Direktori penyimpanan data workspace tenant (default: `$TMPDIR/krewire-dash`).
- `AUTH_URL`: URL server Krewire Auth eksternal untuk verifikasi SSO/RBAC di lingkungan produksi.
