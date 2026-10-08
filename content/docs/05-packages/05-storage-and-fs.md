---
title: "Storage & Filesystem (storage, fs, file)"
description: "Backend-independent object storage, atomic file operations, and virtual testable filesystem boundaries."
date: "2026-10-08"
---

# Storage & Filesystem (`storage`, `fs`, `file`)

Import paths:
- `github.com/krewire/krewire/packages/storage`
- `github.com/krewire/krewire/packages/file`
- `github.com/krewire/krewire/packages/fs`

Filesystem and persistent object access are core requirements for modern web applications. Krewire provides clean, testable boundaries that decouple business logic from local disks or cloud storage providers.

---

## 1. Object Storage Abstraction (`packages/storage`)

`packages/storage` defines a uniform interface for storing, reading, and deleting binary objects, with swappable drivers for:
- **In-Memory Store:** Instant, deterministic test execution without side-effects.
- **Local Disk Store:** Storing uploads and generated artifacts on the server filesystem.
- **Cloud S3 Store:** Storing assets in Amazon S3, MinIO, or Cloudflare R2.

```go
package main

import (
	"context"
	"strings"
	"github.com/krewire/krewire/packages/storage"
)

func main() {
	// Initialize local or cloud storage driver
	store := storage.NewLocalDriver("/var/data/uploads")

	ctx := context.Background()

	// Write object
	err := store.Put(ctx, "avatar_12.png", strings.NewReader("image-bytes"), "image/png")

	// Read object
	reader, err := store.Get(ctx, "avatar_12.png")

	// Delete object
	err = store.Delete(ctx, "avatar_12.png")
}
```

---

## 2. Atomic File Replacement (`packages/file`)

Writing files directly to disk risks leaving corrupted or partial files if a crash occurs mid-write. `packages/file` provides crash-safe atomic write operations:

```go
import "github.com/krewire/krewire/packages/file"

// Writes to a temporary sibling file first, then atomically renames
err := file.WriteAtomic("config.json", data, 0644)
```

---

## 3. Testable Virtual Filesystem (`packages/fs`)

`packages/fs` exposes a clean interface mirroring standard OS file operations, allowing unit tests to simulate filesystems in memory without touching disk:

```go
import "github.com/krewire/krewire/packages/fs"

// Production: real OS disk
filesystem := fs.NewOS()

// Unit Testing: fast in-memory mock
filesystem := fs.NewMemory()
```
