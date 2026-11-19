---
title: "Virtual Filesystem & Atomic I/O"
description: "Testable virtual filesystem boundaries, race-free atomic file writes, and symlink protection in pkg/fs and pkg/file."
date: "2026-10-25"
---

# Virtual Filesystem & Atomic I/O (`pkg/fs`, `pkg/file`)

Import paths:
- `github.com/krewire/krewire/pkg/fs`
- `github.com/krewire/krewire/pkg/file`

Krewire provides clean, testable filesystem abstractions and atomic file writing primitives designed to eliminate partial write corruptions, symlink race conditions, and hostile traversal attacks.

---

## 1. Testable Filesystem Interface (`pkg/fs`)

While the Go standard library `io/fs` provides read-only abstractions, `pkg/fs.FileSystem` provides a fully testable read-write abstraction:

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

### Implementations

- **`fs.OS{}`**: Interacts directly with the underlying operating system disk.
- **`fs.NewMemory()`**: High-performance in-memory filesystem for unit testing, preventing filesystem pollution during automated CI runs.

---

## 2. Race-Free Atomic Writes (`pkg/file`)

Writing directly to destination paths risks leaving corrupted or incomplete files if the process crashes mid-write. `pkg/file.WriteAtomic` guarantees POSIX atomicity and enforces strict defensive security:

### Defensive Protocols in `file.WriteAtomic`

- **Symlink Protection (CWE-59 / CWE-61)**: Rejects writing to any target that is an existing symlink (`ErrSymlinkTarget`).
- **Cryptographic Temp Suffix (CWE-377)**: Generates unpredictable temporary filenames using 16 random bytes from `crypto/rand`.
- **Private Permission Defaults**: Creates parent directories with `0700` (`ModeDirPrivate`) and written files with `0600` (`ModeFilePrivate`).
- **POSIX Atomic Rename**: Uses `os.Rename` so that readers always observe either the old file or the complete new file, never a partial state.

### Code Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/krewire/krewire/pkg/file"
)

func SaveConfiguration(configJSON []byte) error {
	target := "/opt/app/config.json"

	// Pass nil for default OS filesystem or inject a mock fs.FileSystem
	err := file.WriteAtomic(nil, target, configJSON, 0o600)
	if err != nil {
		return fmt.Errorf("failed atomic write: %w", err)
	}

	return nil
}
```

---

## 3. Alignment with the Three Pillars (3S)

- **Secure (S1)**: Prevents symlink hijacking (CWE-59) and temp-file collision attacks (CWE-377) by default.
- **Sustainable (S2)**: Pure Go using stdlib `io/fs`, `os`, and `crypto/rand` with zero CGO or external libraries.
- **Scalable (S3)**: Decoupled virtual filesystem boundary allows isolated memory-backed testing that runs in milliseconds.
