---
title: "Sistem Berkas Virtual & I/O Atomik"
description: "Batas sistem berkas virtual yang dapat diuji, penulisan berkas atomik bebas race condition, dan perlindungan symlink di fs dan file."
date: "2026-10-25"
---

# Sistem Berkas Virtual & I/O Atomik (`fs`, `file`)

Jalur impor:
- `github.com/krewire/krewire/fs`
- `github.com/krewire/krewire/file`

Krewire menyediakan abstraksi sistem berkas yang bersih dan dapat diuji, serta primitif penulisan berkas atomik yang dirancang untuk mencegah kerusakan akibat penulisan parsial, race condition symlink, serta serangan traversal berkas.

---

## 1. Antarmuka Sistem Berkas yang Dapat Diuji (`fs`)

Sementara pustaka standar Go `io/fs` hanya menyediakan abstraksi baca-saja (read-only), antarmuka `fs.FileSystem` menyediakan abstraksi baca-tulis penuh yang dapat diuji:

```go
package fs

import (
	"io"
	"io/fs"
)

type FileSystem interface {
	Open(name string) (io.ReadCloser, error)
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, data []byte, perm fs.FileMode) error
	Stat(name string) (fs.FileInfo, error)
	MkdirAll(path string, perm fs.FileMode) error
	Rename(oldPath, newPath string) error
	Remove(name string) error
}
```

### Implementasi Tersedia

- **`fs.OS{}`**: Berinteraksi langsung dengan disk sistem operasi host yang sebenarnya.
- **`fs.NewMemory()`**: Sistem berkas in-memory berkinerja tinggi untuk pengujian unit, mencegah polusi direktori sementara selama eksekusi CI otomatis.

---

## 2. Penulisan Atomik Bebas Race Condition (`file`)

Menulis berkas langsung ke jalur tujuan dapat menghasilkan berkas yang rusak jika proses berhenti di tengah jalan. Fungsi `file.WriteAtomic` menjamin atomisitas POSIX dan menegakkan keamanan defensif:

### Protokol Keamanan Defensif pada `file.WriteAtomic`

- **Perlindungan Symlink (CWE-59 / CWE-61)**: Menolak penulisan jika target adalah symlink yang sudah ada (`ErrSymlinkTarget`).
- **Sufiks Acak Kriptografis (CWE-377)**: Menghasilkan nama berkas sementara yang tidak dapat diprediksi menggunakan 16 byte acak dari `crypto/rand`.
- **Izin Berkas Privat Bawaan**: Membuat direktori induk dengan izin `0700` (`ModeDirPrivate`) dan berkas dengan izin `0600` (`ModeFilePrivate`).
- **Rename Atomik POSIX**: Menggunakan `os.Rename` sehingga pembaca selalu melihat berkas lama yang utuh atau berkas baru yang telah selesai sepenuhnya, tanpa kondisi setengah jadi.

### Contoh Kode

```go
package main

import (
	"fmt"
	"log"

	"github.com/krewire/krewire/file"
)

func SaveConfiguration(configJSON []byte) error {
	target := "/opt/app/config.json"

	// Berikan nil untuk filesystem OS default atau sematkan mock fs.FileSystem
	err := file.WriteAtomic(nil, target, configJSON, 0o600)
	if err != nil {
		return fmt.Errorf("gagal menulis berkas secara atomik: %w", err)
	}

	return nil
}
```

---

## 3. Keselarasan dengan Tiga Pilar (3S)

- **Aman (S1 - Secure)**: Mencegah pembajakan symlink (CWE-59) dan tabrakan nama berkas sementara (CWE-377) secara bawaan.
- **Berkelanjutan (S2 - Sustainable)**: Murni Go menggunakan stdlib `io/fs`, `os`, dan `crypto/rand` tanpa CGO atau pustaka pihak ketiga.
- **Skalabel (S3 - Scalable)**: Batas sistem berkas virtual yang decoupled memungkinkan pengujian berbasis memori yang selesai dalam hitungan milidetik.
