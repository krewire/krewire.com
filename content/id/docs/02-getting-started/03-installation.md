---
title: "Instalasi"
description: "Panduan instalasi komprehensif untuk alat pengembang CLI krew di Linux, macOS, dan Windows."
date: "2026-09-29"
---

# Instalasi

CLI **`krew`** adalah perkakas terpadu (developer toolchain & CLI driver) untuk seluruh ekosistem Krewire. Satu binary mandiri mendukung scaffolding proyek, pembuatan kode, server pengembangan lokal dengan live-reload, pengujian, migrasi database, dan kompilasi produksi di seluruh delapan jenis workload.

---

## 1. Persyaratan Sistem

Sebelum menginstal `krew`, pastikan lingkungan pengembangan Anda memenuhi persyaratan dasar berikut:

| Persyaratan | Spesifikasi yang Didukung | Catatan |
| :--- | :--- | :--- |
| **Sistem Operasi** | Linux (Ubuntu, Debian, Fedora, Arch, Alpine), macOS (12+), Windows (10/11) | Binary statis bawaan tersedia untuk semua platform. |
| **Arsitektur** | `amd64` (x86_64), `arm64` (Apple Silicon, server ARM64) | Build statis multi-arsitektur terverifikasi tanpa dependensi C native (`CGO_ENABLED=0`). |
| **Go Runtime** | **Go 1.27.1 atau lebih tinggi** | Diperlukan saat membangun dari source code atau menjalankan workload Go. |
| **Version Control** | `git` (2.30+) | Diperlukan untuk bootstrap template dan pengambilan modul. |

Periksa versi Go Anda saat ini:

```bash
go version
```

> [!NOTE]
> Jika Anda belum menginstal Go, unduh dan pasang paket resmi dari [golang.org/dl](https://golang.org/dl).

---

## 2. Metode Instalasi

Pilih metode penginstalan yang paling sesuai dengan platform Anda:

### Metode A: Skrip Penginstal Otomatis (Linux & macOS)

Metode instalasi yang direkomendasikan untuk sistem Linux dan macOS menggunakan skrip penginstalan otomatis:

```bash
curl -fsSL https://krewire.com/install.sh | sh
```

*(Atau: `curl -fsSL https://krewire.com/scripts/install.sh | sh`)*

#### Cara Kerja Penginstal:
1. **Deteksi Platform:** Memeriksa kernel Anda (`Linux`/`Darwin`) dan arsitektur CPU (`x86_64`/`arm64`).
2. **Resolusi Rilis:** Memeriksa ketersediaan arsip binary pre-built di GitHub Releases. Jika binary siap pakai belum tersedia untuk platform Anda, installer secara otomatis beralih ke `go install github.com/krewire/krew/cmd/krew@latest`.
3. **Verifikasi Checksum:** Mengunduh manifes resmi `krew_checksums.txt` dan memvalidasi hash SHA256 sebelum ekstraksi.
4. **Penempatan Atomik:** Menempatkan binary `krew` ke `/usr/local/bin` menggunakan `install -m 755` (atau `~/.local/bin` jika tanpa akses root). Skrip juga secara otomatis membuat symlink alias `kiw`.
5. **Pembersihan Aman:** Menangkap sinyal keluar untuk memastikan direktori sementara segera dihapus.

Untuk menentukan direktori kustom:

```bash
curl -fsSL https://krewire.com/install.sh | INSTALL_DIR="$HOME/bin" sh
```

---

### Metode B: Skrip PowerShell Otomatis (Windows)

Pada Windows 10/11, buka PowerShell dan jalankan:

```powershell
irm https://krewire.com/install.ps1 | iex
```

*(Atau: `irm https://krewire.com/scripts/install.ps1 | iex`)*

#### Cara Kerja Penginstal PowerShell:
1. Mendeteksi arsitektur Windows `amd64` atau `arm64`.
2. Mengunduh arsip binary pre-built atau mengompilasi via `go install`.
3. Memasang `krew.exe` dan `kiw.exe` ke `$env:LOCALAPPDATA\Programs\krew`.
4. Menambahkan direktori instalasi ke variabel lingkungan `PATH` pengguna secara otomatis.

---

### Metode C: Instal via Go Toolchain (Lintas Platform)

Jika Anda sudah menginstal Go dan mengonfigurasi `$GOPATH/bin` dalam `PATH` sistem Anda, instal `krew` langsung dari source code:

```bash
go install github.com/krewire/krew/cmd/krew@latest
```

Untuk mengunci versi tertentu (misalnya `v0.2.3`):

```bash
go install github.com/krewire/krew/cmd/krew@v0.2.3
```

#### Mengonfigurasi PATH Anda:

Pastikan direktori binary Go disertakan dalam profil shell:

**Bash / Zsh (`~/.bashrc` atau `~/.zshrc`):**
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

**Fish (`~/.config/fish/config.fish`):**
```fish
set -gx PATH $PATH (go env GOPATH)/bin
```

**Windows PowerShell:**
```powershell
$env:Path += ";$((go env GOPATH))\bin"
```

---

## 3. Memverifikasi Instalasi

Setelah instalasi selesai, pastikan bahwa binary dapat diakses dan berfungsi normal:

### 1. Pemeriksaan Versi

```bash
krew version
```

Output contoh:
```text
krew v0.2.3 (linux/amd64) commit 7a8f9c1 built 2026-11-20
```

### 2. Menu Bantuan Interaktif

```bash
krew help
```

---

## 4. Langkah Berikutnya

Dengan `krew` terinstal, Anda dapat membuat proyek baru dalam 30 detik:

```bash
# Inisialisasi proyek fullstack baru
krew new my-app

# Masuk ke direktori proyek
cd my-app

# Jalankan server pengembangan lokal (hot-reload)
krew dev

# Kompilasi hasil produksi (binary mandiri atau bundle statis)
krew build
```
