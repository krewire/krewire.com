---
title: "Sistem Berkas Virtual & Penulisan Atomik"
description: "Batas sistem berkas virtual yang dapat diuji, penulisan berkas atomik tahan kerusakan di pkg/fs dan pkg/file."
date: "2026-10-25"
---

# Sistem Berkas Virtual & Penulisan Atomik (`pkg/fs`, `pkg/file`)

Jalur impor:
- `github.com/krewire/krewire/pkg/file`
- `github.com/krewire/krewire/pkg/fs`

Krewire menyediakan batasan sistem berkas virtual yang bersih untuk memisahkan logika bisnis dari disk lokal serta mencegah kerusakan data jika terjadi penghentian proses tak terduga.

---

## Penggantian Berkas Atomik (`pkg/file`)

Menulis berkas langsung ke disk berisiko menghasilkan berkas korup jika terjadi crash di tengah proses penulisan. `pkg/file` menyediakan penulisan atomik yang aman dari kegagalan proses:

```go
import "github.com/krewire/krewire/pkg/file"

// Menulis ke berkas sementara terlebih dahulu, lalu mengganti nama secara atomik
err := file.WriteAtomic("config.json", data, 0644)
```

---

## Sistem Berkas Virtual yang Dapat Diuji (`pkg/fs`)

`pkg/fs` menyediakan antarmuka yang mencerminkan operasi berkas OS standar, memungkinkan unit test menyimulasikan sistem berkas di dalam memori tanpa menyentuh disk fisik:

```go
import "github.com/krewire/krewire/pkg/fs"

// Lingkungan Produksi: disk OS nyata
filesystem := fs.NewOS()

// Pengujian Unit: mock in-memory cepat
filesystem := fs.NewMemory()
```\n