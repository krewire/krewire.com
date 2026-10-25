---
title: "Virtual Filesystem & Atomic Files"
description: "Testable virtual filesystem boundaries, atomic file writes, and crash safety in pkg/fs and pkg/file."
date: "2026-10-25"
---

# Virtual Filesystem & Atomic Files (`pkg/fs`, `pkg/file`)

Import paths:
- `github.com/krewire/krewire/pkg/file`
- `github.com/krewire/krewire/pkg/fs`

Krewire provides clean, testable boundaries that decouple business logic from local disks and prevent corrupt files during unexpected process termination.

---

## Atomic File Replacement (`pkg/file`)

Writing files directly to disk risks leaving corrupted or partial files if a crash occurs mid-write. `pkg/file` provides crash-safe atomic write operations:

```go
import "github.com/krewire/krewire/pkg/file"

// Writes to a temporary sibling file first, then atomically renames
err := file.WriteAtomic("config.json", data, 0644)
```

---

## Testable Virtual Filesystem (`pkg/fs`)

`pkg/fs` exposes a clean interface mirroring standard OS file operations, allowing unit tests to simulate filesystems in memory without touching disk:

```go
import "github.com/krewire/krewire/pkg/fs"

// Production: real OS disk
filesystem := fs.NewOS()

// Unit Testing: fast in-memory mock
filesystem := fs.NewMemory()
```\n