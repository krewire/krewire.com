---
title: "Installation"
description: "Comprehensive installation guide for the krew CLI developer toolchain on Linux, macOS, and Windows."
date: "2026-09-29"
---

# Installation

The **`krew`** CLI is the unified developer multi-tool and driver for the Krewire ecosystem. A single binary powers project scaffolding, code generation, local development servers with live reload, quality gates, migrations, and production compilation across all eight workloads.

---

## 1. System Requirements

Before installing `krew`, verify that your development environment satisfies the baseline requirements:

| Requirement | Supported Specifications | Notes |
| :--- | :--- | :--- |
| **Operating System** | Linux (Ubuntu, Debian, Fedora, Arch, Alpine), macOS (12+), Windows (10/11) | Native static binaries available for all platforms. |
| **Architecture** | `amd64` (x86_64), `arm64` (Apple Silicon, ARM64 servers) | Verified multi-arch static builds with zero native C dependencies (`CGO_ENABLED=0`). |
| **Go Runtime** | **Go 1.27.1 or higher** | Required when building from source or running Go workloads. |
| **Version Control** | `git` (2.30+) | Required for template bootstrapping and module fetching. |

Check your current Go version:

```bash
go version
```

> [!NOTE]
> If you do not have Go installed, download and install the official package from [golang.org/dl](https://golang.org/dl).

---

## 2. Installation Methods

Choose the installation method best suited to your platform:

### Method A: Automated Shell Installer (Linux & macOS)

The recommended installation method for Linux and macOS systems uses the automated installation script:

```bash
curl -fsSL https://krewire.com/install.sh | sh
```

*(Alternatively: `curl -fsSL https://krewire.com/scripts/install.sh | sh`)*

#### How the Installer Works:
1. **Platform Detection:** Inspects your kernel (`Linux`/`Darwin`) and CPU architecture (`x86_64`/`arm64`).
2. **Release Resolution:** Queries GitHub Releases for precompiled binary archives. If pre-built binaries are not available for your platform, it automatically falls back to `go install github.com/krewire/krew/cmd/krew@latest`.
3. **Checksum Verification:** Downloads the official `krew_checksums.txt` manifest and cryptographically validates the SHA256 hash before extraction.
4. **Atomic Placement:** Places `krew` into `/usr/local/bin` using `install -m 755` (or falls back to `~/.local/bin` if running without root permissions). It also automatically creates a `kiw` alias symlink.
5. **Secure Cleanup:** Traps exit signals to ensure temporary directories are immediately erased.

To customize the installation directory:

```bash
curl -fsSL https://krewire.com/install.sh | INSTALL_DIR="$HOME/bin" sh
```

---

### Method B: Automated PowerShell Installer (Windows)

On Windows 10/11, open PowerShell (Administrator or standard user) and run:

```powershell
irm https://krewire.com/install.ps1 | iex
```

*(Alternatively: `irm https://krewire.com/scripts/install.ps1 | iex`)*

#### How the PowerShell Installer Works:
1. Detects `amd64` or `arm64` Windows architecture.
2. Downloads and validates the precompiled binary archive or compiles via `go install`.
3. Installs `krew.exe` and `kiw.exe` into `$env:LOCALAPPDATA\Programs\krew`.
4. Adds the installation directory to your User `PATH` environment variable automatically.

---

### Method C: Install via Go Toolchain (Cross-Platform)

If you already have Go installed and `$GOPATH/bin` in your system `PATH`, install `krew` directly from source:

```bash
go install github.com/krewire/krew/cmd/krew@latest
```

To pin a specific version (e.g. `v0.2.3`):

```bash
go install github.com/krewire/krew/cmd/krew@v0.2.3
```

#### Configuring Your PATH:

Ensure your Go binary directory is included in your shell profile:

**Bash / Zsh (`~/.bashrc` or `~/.zshrc`):**
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

**Fish (`~/.config/fish/config.fish`):**
```fish
set -gx PATH $PATH (go env GOPATH)/bin
```

**Windows PowerShell:**
```powershell
$env:Path += ";$((go env GOPATH))\bin"
```

---

## 3. Verifying the Installation

After installation completes, verify that the binary is accessible and operational:

### 1. Version Inspection

```bash
krew version
```

Output:
```text
krew v0.2.3 (darwin/arm64) commit 7a8f9c1 built 2026-11-20
```

### 2. Interactive Help & Command Map

```bash
krew help
```

---

## 4. Next Steps

With `krew` installed, you can initialize a new project in under 30 seconds:

```bash
# Scaffold a new fullstack project
krew new my-app

# Navigate to the workspace
cd my-app

# Launch the live-reload development server
krew dev

# Compile into a production standalone static bundle or binary
krew build
```
