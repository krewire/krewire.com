---
title: "PaaS Hibrida & Cloud Terkelola"
description: "Mengoperasikan Krewire Cloud: mesin PaaS binary tunggal self-hosted di VPS privat vs platform SaaS cloud terkelola."
date: "2026-10-24"
---

# PaaS Hibrida & Cloud Terkelola (Krewire Cloud)

Krewire Cloud dirancang sebagai mesin PaaS hibrida mandiri untuk hosting VPS privat sekaligus platform cloud multi-tenant terkelola.

> **Portal Live:** Krewire Cloud tersedia di [cloud.krewire.com](https://cloud.krewire.com). Perintah `kiw deploy` di bagian 1 adalah bagian dari inti open-source dan sudah berfungsi mandiri di infrastruktur Anda sendiri tanpa Krewire Cloud. Lihat [Inti & Layanan Platform](/project/core-and-platform-services).

## 1. Mesin PaaS Hibrida Self-Hosted

Rilis workload native-Go privat Anda di Linux VPS mana pun (Ubuntu, Debian, Alpine) menggunakan perkakas standar:

```bash
# Kemas binary produksi dengan aset tersemat
kiw deploy --target binary

# Hasilkan unit supervisor systemd untuk pengelolaan layanan tanpa downtime
kiw deploy --target systemd
```

- **Efisiensi RAM Ekstrem**: Workload Go berjalan hanya dalam 15–30MB RAM. Satu VPS seharga $5/bulan dapat menampung puluhan aplikasi aktif.
- **Rolling Deployment Tanpa Downtime (Zero-Downtime)**: Pool proses tersupervisi menjalankan instance baru di sandbox eksekusi terisolasi, memvalidasi health check, dan secara dinamis memperbarui reverse proxy.
- **TLS Otomatis Let's Encrypt**: Sertifikat ACME otomatis untuk wildcard subdomain dan domain kustom.

## 2. Platform Cloud Multi-Tenant Terkelola

Platform terkelola ini melayani tim yang menginginkan deployment tanpa repot mengelola VPS:

- `kiw deploy` + webhook Git push-to-deploy: membundel dan menyiapkan proyek secara lokal, lalu continuous deployment merilisnya ke produksi.
- Streaming build real-time dan agregasi log terpusat.
- Pembayaran lokal Indonesia melalui QRIS dan Virtual Account Bank.
- Kepatuhan kedaulatan data penuh di bawah Undang-Undang Perlindungan Data Pribadi (UU PDP) Indonesia.
