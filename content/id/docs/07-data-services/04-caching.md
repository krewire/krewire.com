---
title: "Cache Multi-Driver"
description: "Caching berkinerja tinggi berbasis memori, berkas fisik, Redis, dan Memcached dengan TTL, pembersihan kedaluwarsa, dan pola Remember di pkg/cache."
date: "2026-10-25"
---

# Cache Multi-Driver (`pkg/cache`)

Jalur impor: `github.com/krewire/krewire/pkg/cache`

Pustaka `pkg/cache` menghadirkan abstraksi penyimpanan cache kunci-nilai (key-value) terpadu dengan dukungan kedaluwarsa TTL, batas eviksi LRU, pembersihan latar belakang otomatis, serta pola atomik `Remember` di seluruh ekosistem Krewire.

---

## 1. Antarmuka Utama

```go
package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	GetString(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, val []byte, ttl time.Duration) error
	SetString(ctx context.Context, key string, val string, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Has(ctx context.Context, key string) (bool, error)
	Clear(ctx context.Context) error
	Remember(ctx context.Context, key string, ttl time.Duration, fn func() ([]byte, error)) ([]byte, error)
	Close() error
}

type Driver = Cache
```

---

## 2. Driver Cache yang Tersedia

1. **Cache Dalam Memori (`cache.NewMemoryCache`)**:
   Cache dalam-proses berkecepatan tinggi dengan sinkronisasi mutex, kapasitas terkonfigurasi (bawaan 10.000 item), dan kebijakan eviksi LRU.
   ```go
   memCache := cache.NewMemoryCache(1 * time.Minute)
   defer memCache.Close()
   ```
2. **Redis / Valkey / Dragonfly (`cache.NewRedisCache`)**:
   Klien protokol RESP dalam Go murni dengan pooling koneksi otomatis dan logika rekoneksi mandiri:
   ```go
   redisCache, err := cache.NewRedisCache(cache.RedisConfig{
       Address:  "127.0.0.1:6379",
       Password: os.Getenv("REDIS_PASSWORD"),
       DB:       0,
   })
   ```
3. **Cache Berbasis Sistem Berkas (`cache.NewFileCache`)**:
   Cache persisten berbasis disk dengan penulisan berkas atomik dan pembersihan berkala.
   ```go
   fileCache, err := cache.NewFileCache("/tmp/krewire-cache", 10*time.Minute)
   ```
4. **Memcached (`cache.NewMemcachedCache`)**:
   Driver Memcached Go murni dengan dukungan kluster multi-server.

---

## 3. Contoh Penggunaan & Pola Remember

Pola `Remember` mengambil entri cache secara transparan atau mengeksekusi fungsi fallback jika data belum tersedia (cache miss):

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/krewire/krewire/pkg/cache"
)

func GetUserProfile(ctx context.Context, c cache.Cache, userID string) ([]byte, error) {
	key := fmt.Sprintf("users:%s", userID)

	return c.Remember(ctx, key, 30*time.Minute, func() ([]byte, error) {
		// Dieksekusi hanya saat cache miss: ambil dari basis data utama
		return FetchUserFromDatabase(userID)
	})
}
```

---

## 4. Keselarasan dengan Tiga Pilar (3S)

- **Aman (S1 - Secure)**: `FileCache` menegakkan izin berkas privat (`0600`) dan menolak upaya manipulasi traversal jalur.
- **Berkelanjutan (S2 - Sustainable)**: Murni Go tanpa CGO atau ketergantungan pada library klien Redis eksternal yang berat.
- **Skalabel (S3 - Scalable)**: Transisi mulus dari `MemoryCache` pada VPS tunggal ke `RedisCache` atau `Memcached` pada kluster terdistribusi tanpa mengubah kode logika bisnis.
