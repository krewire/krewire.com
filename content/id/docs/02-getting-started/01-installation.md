---
title: "Instalasi"
description: "Panduan instalasi komprehensif untuk alat pengembang CLI kiw di Linux, macOS, dan Windows."
date: "2026-09-29"
---

# Instalasi

CLI **`kiw`** adalah multi-tool terpadu bagi pengembang di seluruh ekosistem Krewire. Satu binary mandiri mendukung scaffolding proyek, pembuatan kode (code generation), server pengembangan lokal, pengujian, kompilasi aset, dan deployment produksi di seluruh delapan jenis workload.

---

## 1. Persyaratan Sistem

Sebelum menginstal `kiw`, pastikan lingkungan pengembangan Anda memenuhi persyaratan dasar berikut:

| Persyaratan | Spesifikasi yang Didukung | Catatan |
| :--- | :--- | :--- |
| **Sistem Operasi** | Linux (Ubuntu, Debian, Fedora, Arch, Alpine), macOS (12+), Windows (10/11) | Binary bawaan (native) tersedia untuk semua platform. |
| **Arsitektur** | `amd64` (x86_64), `arm64` (Apple Silicon, server ARM64) | Build statis multi-arsitektur terverifikasi. |
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

Pilih metode penginstalan yang paling sesuai dengan platform dan alur kerja Anda:

### Metode A: Skrip Penginstal Otomatis (Linux & macOS)

Metode instalasi yang direkomendasikan untuk sistem Linux dan macOS menggunakan skrip penginstalan otomatis:

```bash
curl -fsSL https://krewire.com/scripts/install.sh | sh
```

#### Cara Kerja Penginstal:
1. **Deteksi Platform:** Memeriksa kernel Anda (`Linux`/`Darwin`) dan arsitektur CPU (`x86_64`/`arm64`).
2. **Resolusi Tag:** Menemukan rilis stabil resmi terbaru (misalnya `v0.1.0`) dari registry rilis GitHub.
3. **Verifikasi Checksum:** Mengunduh manifes resmi `kiw_checksums.txt` dan memvalidasi hash SHA256 dari arsip yang diunduh secara kriptografis sebelum ekstraksi.
4. **Penempatan Atomik:** Menempatkan binary ke `/usr/local/bin` menggunakan `install -m 755`. Jika dijalankan tanpa izin root, secara otomatis beralih ke `~/.local/bin`.
5. **Pembersihan Aman:** Menangkap sinyal keluar (`INT`, `TERM`, `EXIT`) untuk menjamin direktori ekstraksi sementara dihapus secara aman.

Untuk menyesuaikan direktori instalasi, gunakan variabel lingkungan `INSTALL_DIR`:

```bash
curl -fsSL https://krewire.com/scripts/install.sh | INSTALL_DIR="$HOME/bin" sh
```

---

### Metode B: Instal via Go Toolchain (Lintas Platform / Windows)

Jika Anda sudah menginstal Go dan mengonfigurasi `$GOPATH/bin` dalam `PATH` sistem Anda, instal `kiw` langsung dari source code:

```bash
go install github.com/krewire/krewire/tools/kiw/cmd/kiw@latest
```

Untuk mengunci versi tertentu (misalnya `v0.1.0`):

```bash
go install github.com/krewire/krewire/tools/kiw/cmd/kiw@v0.1.0
```

#### Mengonfigurasi PATH Anda:
Pastikan direktori binary Go Anda disertakan dalam profil shell (`~/.bashrc`, `~/.zshrc`, atau Environment Variables Windows):

```bash
# Tambahkan ke ~/.bashrc atau ~/.zshrc jika belum ada:
export PATH="$PATH:$(go env GOPATH)/bin"
```

Pada Windows PowerShell:

```powershell
$env:Path += ";$((go env GOPATH))\bin"
```

---

## 3. Memverifikasi Instalasi

Setelah instalasi selesai, pastikan bahwa binary dapat diakses dan berfungsi normal:

### 1. Pemeriksaan Versi

```bash
kiw version
```

Output yang diharapkan (baris `Commit` hanya dicetak untuk build dengan informasi provenance modul):

```text
kiw Krewire Devtool
  Krewire Ecosystem  v0.1.0
  Signature          build with krewire
  Go                 go1.27.1 (linux/amd64)
  Commit             a1b2c3d4
```

### 2. Diagnostik Lingkungan

Jalankan `kiw info` untuk memeriksa toolchain aktif, sistem operasi, dan lingkungan Go Anda:

```bash
kiw info
```

### 3. Pemeriksaan Kompatibilitas

Jalankan `kiw compat` di dalam direktori proyek Krewire mana pun untuk memverifikasi bahwa CLI lokal Anda sesuai dengan dependensi paket di `go.mod`:

```bash
kiw compat
```

---

## 4. Memperbarui `kiw`

Untuk memperbarui ke rilis stabil terbaru:

- **Via Skrip Penginstal Shell:**
  ```bash
  curl -fsSL https://krewire.com/scripts/install.sh | sh
  ```
- **Via Go Toolchain:**
  ```bash
  go install github.com/krewire/krewire/tools/kiw/cmd/kiw@latest
  ```

---

## 5. Menghapus (Uninstall) `kiw`

Karena `kiw` didistribusikan sebagai satu binary statis mandiri tanpa daemon latar belakang atau entri registry, penghapusannya sangat mudah:

```bash
# Untuk instalasi standar di seluruh sistem:
sudo rm -f /usr/local/bin/kiw

# Untuk instalasi lokal pengguna:
rm -f ~/.local/bin/kiw

# Untuk instalasi via Go bin:
rm -f "$(go env GOPATH)/bin/kiw"
```

---

## Langkah Berikutnya

Sekarang toolchain Anda telah terinstal, lanjutkan ke [**2.2 Konfigurasi →**](/getting-started/configuration) untuk mempelajari cara mengonfigurasi proyek menggunakan `krewire.yaml`.
