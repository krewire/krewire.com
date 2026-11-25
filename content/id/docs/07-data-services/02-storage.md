---
title: "Penyimpanan Objek Multi-Driver"
description: "Abstraksi penyimpanan objek terpadu dengan driver yang dapat ditukar untuk memori, disk lokal, Redis, dan Amazon S3 di storage."
date: "2026-10-25"
---

# Penyimpanan Objek Multi-Driver (`storage`)

Jalur impor: `github.com/krewire/krewire/storage`

Paket `storage` mendefinisikan abstraksi penyimpanan objek biner berkinerja tinggi yang ter-decoupling. Melalui antarmuka `storage.Store`, aplikasi tidak terikat secara kaku pada SDK vendor cloud tertentu, sehingga tetap portabel dan dapat bertransisi dengan mulus dari satu VPS hemat hingga lingkungan cloud terdistribusi.

---

## 1. Antarmuka Utama

```go
package storage

import (
	"io"
	"io/fs"
)

// Store menyimpan objek biner opaq yang dialamatkan menggunakan kunci unik.
type Store interface {
	Put(key string, data []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Stat(key string) (ObjectInfo, error)
}

// Provider adalah alias untuk Store sesuai konvensi arsitektur Krewire.
type Provider = Store

// StreamStore menyediakan metode streaming terpotong untuk objek berukuran besar.
type StreamStore interface {
	Store
	PutStream(key string, r io.Reader) error
	GetStream(key string) (io.ReadCloser, error)
}

// ObjectInfo mewakili metadata untuk objek biner yang tersimpan.
type ObjectInfo struct {
	Key  string
	Size int64
	Mode fs.FileMode
}
```

---

## 2. Driver Penyimpanan yang Tersedia

1. **Sistem Berkas Lokal (`storage.NewLocal`)**:
   Menyimpan berkas pada disk lokal server dengan pembuatan direktori otomatis dan sanitasi pencegahan path-traversal (`SafePath`).
   ```go
   store, err := storage.NewLocal("/var/data/uploads", nil)
   ```
2. **Penyimpanan Dalam Memori (`storage.NewMemory`)**:
   Driver map in-memory yang thread-safe untuk pengujian unit yang cepat dan deterministik tanpa efek samping ke disk.
   ```go
   testStore := storage.NewMemory()
   ```
3. **Cloud S3 / MinIO (`storage.NewS3`)**:
   Driver penyimpanan cloud kompatibel S3 (AWS S3, MinIO, Cloudflare R2, Wasabi).
   ```go
   store, err := storage.NewS3(storage.S3Config{
       Endpoint:  "s3.us-east-1.amazonaws.com",
       Bucket:    "user-assets",
       AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
       SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
       UseSSL:    true,
   })
   ```
4. **Redis / Valkey (`storage.NewRedis`)**:
   Penyimpanan payload biner yang didukung oleh kluster Redis atau Valkey untuk kebutuhan caching cepat.

---

## 3. Contoh Penggunaan: Layanan Media Terpisah

```go
package main

import (
	"fmt"
	"log"

	"github.com/krewire/krewire/storage"
)

type MediaService struct {
	store storage.Store
}

func (s *MediaService) SaveAvatar(userID string, payload []byte) error {
	key := fmt.Sprintf("avatars/%s.png", userID)
	return s.store.Put(key, payload)
}

func (s *MediaService) ReadAvatar(userID string) ([]byte, error) {
	key := fmt.Sprintf("avatars/%s.png", userID)
	return s.store.Get(key)
}

func main() {
	// Lingkungan produksi di VPS lokal
	store, err := storage.NewLocal("/var/data/media", nil)
	if err != nil {
		log.Fatalf("gagal inisialisasi penyimpanan: %v", err)
	}

	service := &MediaService{store: store}
	if err := service.SaveAvatar("usr_101", []byte{0x89, 0x50, 0x4E, 0x47}); err != nil {
		log.Fatal(err)
	}
}
```

---

## 4. Keselarasan dengan Tiga Pilar (3S)

- **Aman (S1 - Secure)**: `storage.NewLocal` menegakkan guard traversal path yang ketat, memastikan kunci tidak dapat keluar dari direktori root yang ditentukan.
- **Berkelanjutan (S2 - Sustainable)**: Implementasi murni Go tanpa dependensi daemon runtime eksternal untuk pengoperasian lokal maupun in-memory.
- **Skalabel (S3 - Scalable)**: Pertukaran backend penyimpanan (Lokal $\to$ S3 $\to$ Redis) tidak memerlukan perubahan sama sekali pada kode domain bisnis.
