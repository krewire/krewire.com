---
title: "Workspace Virtual Native AI"
description: "Arsitektur, isolasi proses sandbox, terminal WebSocket PTY, asisten pair programmer, hub kolaborasi sesi, dan daemon di Krewire Dash."
date: "2026-11-05"
---

# Workspace Virtual Native AI (Krewire Dash)

Krewire Dash menyediakan mesin workspace virtual berbasis AI dan platform cloud IDE. Paket ini menggabungkan isolasi proses yang hermetis, streaming terminal dua arah, editor web berbasis Monaco, dan asisten pair programmer `krew` tersemat ke dalam satu daemon server tunggal.

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

## Akses Workspace & Integrasi Platform

Krewire Workspace di-host sebagai layanan cloud kolaboratif di [dash.krewire.com](https://dash.krewire.com).

### Portal Workspace Cloud

Developer dapat menyalakan lingkungan pengembangan instan langsung dari browser:
- **Portal Web:** `https://dash.krewire.com`
- **Fitur:** Editor web berbasis VS Code / Monaco, streaming terminal dua arah, dan asisten AI `krew` tersemat.
- **Single Sign-On:** Terintegrasi penuh dengan Krewire Auth (`auth.krewire.com`).

### Alternatif Pengembangan Lokal

Untuk alur kerja pengembangan lokal tanpa dependensi cloud, developer dapat memanfaatkan perkakas native CLI `kiw`:
- `kiw dev`: Menjalankan aplikasi secara lokal dengan auto-restart dan hot reload.
- `kiw dashboard`: Menjalankan dashboard observabilitas lokal di `http://127.0.0.1:4000`.
- `kiw mcp --stdio`: Menghubungkan asisten coding AI (seperti Antigravity atau Cursor) langsung ke proyek melalui Model Context Protocol (MCP).

### Edisi Enterprise On-Premise

Workspace enterprise dengan VPC peering privat, sandbox air-gapped, dan audit log terpusat tersedia melalui kontrak enterprise. Hubungi `enterprise@krewire.com`.
