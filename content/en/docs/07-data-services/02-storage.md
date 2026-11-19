---
title: "Multi-Driver Object Storage"
description: "Uniform object storage abstraction with swappable drivers for memory, local disk, Redis, and Amazon S3 in pkg/storage."
date: "2026-10-25"
---

# Multi-Driver Object Storage (`pkg/storage`)

Import path: `github.com/krewire/krewire/pkg/storage`

The `pkg/storage` package defines a decoupled, high-performance object storage abstraction for binary payloads. By programming to the `storage.Store` interface rather than cloud-specific SDKs, applications remain fully portable and can transition smoothly from a single minimal VPS to distributed cloud environments.

---

## 1. Core Interfaces

```go
package storage

import (
	"io"
	"io/fs"
)

// Store stores opaque binary objects addressed by a unique key.
type Store interface {
	Put(key string, data []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Stat(key string) (ObjectInfo, error)
}

// Provider is an alias for Store per Krewire architecture conventions.
type Provider = Store

// StreamStore provides chunked streaming for large objects.
type StreamStore interface {
	Store
	PutStream(key string, r io.Reader) error
	GetStream(key string) (io.ReadCloser, error)
}

// ObjectInfo represents metadata for a stored binary object.
type ObjectInfo struct {
	Key  string
	Size int64
	Mode fs.FileMode
}
```

---

## 2. Pluggable Storage Drivers

1. **Local Filesystem (`storage.NewLocal`)**:
   Stores files on the host disk with automatic directory creation and path-traversal sanitization (`SafePath`).
   ```go
   store, err := storage.NewLocal("/var/data/uploads", nil)
   ```
2. **In-Memory Store (`storage.NewMemory`)**:
   Thread-safe in-memory map driver for fast, deterministic unit tests with zero disk side-effects.
   ```go
   testStore := storage.NewMemory()
   ```
3. **Cloud S3 / MinIO (`storage.NewS3`)**:
   Native S3-compatible cloud object storage driver (AWS S3, MinIO, Cloudflare R2, Wasabi).
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
   Binary payload store backed by Redis or Valkey clusters for ephemeral caching.

---

## 3. Usage Example: Decoupled Media Service

```go
package main

import (
	"fmt"
	"log"

	"github.com/krewire/krewire/pkg/storage"
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
	// Production on local host VPS
	store, err := storage.NewLocal("/var/data/media", nil)
	if err != nil {
		log.Fatalf("failed to init storage: %v", err)
	}

	service := &MediaService{store: store}
	if err := service.SaveAvatar("usr_101", []byte{0x89, 0x50, 0x4E, 0x47}); err != nil {
		log.Fatal(err)
	}
}
```

---

## 4. Alignment with the Three Pillars (3S)

- **Secure (S1)**: `storage.NewLocal` enforces strict path traversal guards, ensuring keys cannot break out of the designated root directory.
- **Sustainable (S2)**: Pure Go implementation with zero external runtime daemons required for local or in-memory operations.
- **Scalable (S3)**: Swapping storage backends (Local $\to$ S3 $\to$ Redis) requires zero modifications to business domain code.
