---
title: "Multi-Driver Object Storage"
description: "Uniform object storage abstraction with swappable drivers for memory, local disk, and Amazon S3 / MinIO in pkg/storage."
date: "2026-10-25"
---

# Multi-Driver Object Storage (`pkg/storage`)

Import path: `github.com/krewire/krewire/pkg/storage`

`pkg/storage` defines a uniform interface for storing, reading, and deleting binary objects, with swappable drivers across memory, disk, and cloud object stores.

---

## Supported Drivers

- **In-Memory Store:** Instant, deterministic test execution without disk side-effects.
- **Local Disk Store:** Storing uploads and generated artifacts on the local server filesystem.
- **Cloud S3 Store:** Storing assets in Amazon S3, MinIO, or Cloudflare R2.

---

## Usage Example

```go
package main

import (
	"context"
	"strings"
	"github.com/krewire/krewire/pkg/storage"
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
```\n