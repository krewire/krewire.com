---
title: "Penyimpanan & Sistem Berkas (storage, fs, file)"
description: "Penyimpanan objek independen dari backend, operasi berkas atomik, dan batasan virtual filesystem yang mudah diuji."
date: "2026-10-08"
---

# Penyimpanan & Sistem Berkas (`storage`, `fs`, `file`)

Jalur impor (*import paths*):
- `github.com/krewire/krewire/pkg/storage`
- `github.com/krewire/krewire/pkg/file`
- `github.com/krewire/krewire/pkg/fs`

Akses filesystem dan penyimpanan objek persisten adalah kebutuhan esensial dalam aplikasi web modern. Krewire menyediakan batasan bersih yang memisahkan logika bisnis dari disk lokal maupun penyedia penyimpanan cloud.

---

## 1. Abstraksi Penyimpanan Objek (`pkg/storage`)

`pkg/storage` mendefinisikan antarmuka seragam untuk menyimpan, membaca, dan menghapus objek binary, dengan driver yang dapat ditukar (*swappable*):
- **Driver In-Memory:** Eksekusi pengujian instan dan deterministik tanpa efek samping ke disk.
- **Driver Disk Lokal:** Menyimpan berkas unggahan dan artefak terkompilasi pada filesystem server.
- **Driver Cloud S3:** Menyimpan aset di Amazon S3, MinIO, atau Cloudflare R2.

```go
package main

import (
	"context"
	"strings"
	"github.com/krewire/krewire/pkg/storage"
)

func main() {
	// Inisialisasi driver disk lokal atau cloud
	store := storage.NewLocalDriver("/var/data/uploads")

	ctx := context.Background()

	// Menulis objek
	err := store.Put(ctx, "avatar_12.png", strings.NewReader("image-bytes"), "image/png")

	// Membaca objek
	reader, err := store.Get(ctx, "avatar_12.png")

	// Menghapus objek
	err = store.Delete(ctx, "avatar_12.png")
}
```

---

## 2. Penulisan Berkas Atomik (`pkg/file`)

Menulis file langsung ke disk berisiko menghasilkan file rusak atau parsial jika terjadi crash atau listrik padam saat penulisan berlangsung. `pkg/file` menyediakan operasi penulisan atomik yang aman dari kegagalan:

```go
import "github.com/krewire/krewire/pkg/file"

// Menulis ke file sementara terlebih dahulu, kemudian me-rename secara atomik
err := file.WriteAtomic("config.json", data, 0644)
```

---

## 3. Filesystem Virtual yang Mudah Diuji (`pkg/fs`)

`pkg/fs` mengekspos antarmuka bersih yang mencerminkan operasi berkas OS standar, memungkinkan pengujian unit mensimulasikan sistem berkas di memori tanpa menyentuh disk fisik:

```go
import "github.com/krewire/krewire/pkg/fs"

// Di lingkungan produksi: disk OS nyata
filesystem := fs.NewOS()

// Di pengujian unit: mock in-memory yang cepat
filesystem := fs.NewMemory()
```
