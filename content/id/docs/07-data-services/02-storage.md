---
title: "Penyimpanan Objek Multi-Driver"
description: "Abstraksi penyimpanan objek terpadu dengan driver yang dapat ditukar untuk memori, disk lokal, dan Amazon S3 / MinIO di pkg/storage."
date: "2026-10-25"
---

# Penyimpanan Objek Multi-Driver (`pkg/storage`)

Jalur impor: `github.com/krewire/krewire/pkg/storage`

`pkg/storage` mendefinisikan antarmuka seragam untuk menyimpan, membaca, dan menghapus objek biner dengan driver yang dapat ditukar antara memori, disk, dan cloud storage.

---

## Driver yang Didukung

- **Penyimpanan In-Memory:** Eksekusi pengujian deterministik instan tanpa efek samping pada disk fisik.
- **Penyimpanan Disk Lokal:** Menyimpan berkas unggahan dan artefak build pada sistem berkas server.
- **Penyimpanan Cloud S3:** Menyimpan aset pada Amazon S3, MinIO, atau Cloudflare R2.

---

## Contoh Penggunaan

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

	// Menyimpan objek
	err := store.Put(ctx, "avatar_12.png", strings.NewReader("image-bytes"), "image/png")

	// Membaca objek
	reader, err := store.Get(ctx, "avatar_12.png")

	// Menghapus objek
	err = store.Delete(ctx, "avatar_12.png")
}
```\n