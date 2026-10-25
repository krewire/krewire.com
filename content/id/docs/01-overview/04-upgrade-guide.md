---
title: "Panduan Pembaruan"
description: "Kebijakan versi, instruksi pembaruan CLI kiw, dan arsitektur konfigurasi untuk Krewire."
date: "2026-09-29"
---

# Panduan Pembaruan (Upgrade Guide)

Panduan ini merinci kebijakan penomoran versi, prosedur pembaruan perangkat, dan standar konfigurasi di seluruh ekosistem Krewire.

---

## 1. Strategi Versi & Janji Kompatibilitas

Krewire mematuhi standar **Semantic Versioning (SemVer 2.0.0)** secara ketat di seluruh ekosistem:

$$\text{vMAJOR}.\text{MINOR}.\text{PATCH}$$

- **Rilis `PATCH` (`v0.1.0` → `v0.1.1`):** Perbaikan bug, optimasi performa, dan pembaruan dokumentasi. Dijamin 100% kompatibel ke belakang tanpa perubahan yang merusak (*zero breaking changes*).
- **Rilis `MINOR` (`v0.1.0` → `v0.2.0`):** Kapabilitas workload baru, paket tambahan, dan flag CLI baru. Sesuai spesifikasi `KWF-META-CMBZJ` dan `KWN-DEVTOOL-Z0VFC`:
  > *"Kapabilitas baru diperkenalkan sebagai paket tambahan; perubahan yang merusak tidak diizinkan pada rilis minor maupun patch."*
- **Rilis `MAJOR` (`v1.0.0`):** Evolusi arsitektur berskala besar. Setiap perubahan yang tidak kompatibel ke belakang akan didahului oleh peringatan depresiasi dan skrip migrasi otomatis.

---

## 2. Memperbarui CLI `kiw`

Untuk memperbarui instalasi lokal `kiw` Anda ke versi rilis terbaru:

### Opsi A: Skrip Shell Otomatis (Linux & macOS)

Jalankan kembali skrip installer resmi:

```bash
curl -fsSL https://krewire.com/scripts/install.sh | sh
```

Skrip ini akan secara otomatis mendeteksi arsitektur CPU Anda, mengunduh rilis binary terbaru, dan mengganti file binary di direktori `/usr/local/bin/kiw` (atau `~/.local/bin/kiw`).

### Opsi B: Menggunakan `go install`

```bash
go install github.com/krewire/krewire/tools/kiw/cmd/kiw@latest
```

### Verifikasi Versi Terbaru

```bash
kiw version
```

Pastikan keluaran versi CLI telah sesuai dengan rilis terbaru yang dituju (`v0.1.0`+).

---

## 3. Memperbarui Dependensi Go

Di dalam repositori aplikasi Anda, perbarui paket-paket domain Krewire:

```bash
go get -u github.com/krewire/krewire@latest
go mod tidy
```

---

## 4. Memverifikasi Kompatibilitas Proyek

Setelah melakukan pembaruan, jalankan perintah validator kompatibilitas bawaan:

```bash
kiw compat
```

Perintah `kiw compat` akan memvalidasi keharmonisan antara konfigurasi `krewire.yaml`, versi modul Go, dan dukungan kompiler.

---

## Langkah Selanjutnya

Setelah memahami arsitektur dan strategi pembaruan Krewire, lanjutkan ke langkah onboarding praktis:

Lanjutkan ke [**2. Memulai (Getting Started) →**](/getting-started).
