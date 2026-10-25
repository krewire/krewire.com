---
title: "Supervisi Siklus Hidup & Kernel"
description: "Bidang kontrol imperatif Krewire: supervisor kernel, registri modul, pengurutan dependensi teratur, dan graceful shutdown."
date: "2026-10-25"
---

# Supervisi Siklus Hidup & Kernel (`pkg/lifecycle`)

Jalur impor: `github.com/krewire/krewire/pkg/lifecycle`

`lifecycle.Kernel` adalah bidang kontrol imperatif ekosistem Krewire: mengikat `model.Project` yang telah tervalidasi ke modul-modul terdaftar dan mengeksekusi beban kerja yang dideklarasikan.

---

## Contoh Eksekusi Kernel

```go
package main

import (
	"context"
	"os"

	"github.com/krewire/krewire/pkg/lifecycle"
	"github.com/krewire/krewire/pkg/model"
)

func main() {
	ctx := context.Background()

	k, err := lifecycle.New(project)
	if err != nil {
		panic(err)
	}

	k.Use(bootModule, webModule)          // modul diinisialisasi sesuai urutan dependensi
	k.RegisterHandler(model.KindApp, run) // satu handler per jenis workload

	if err := k.Boot(ctx); err != nil {
		panic(err)
	}
	defer k.Shutdown(ctx)

	os.Exit(k.Execute(ctx, workload).Int())
}
```

---

## Tipe Pendukung

- **`Registry`** — `NewRegistry`, `Register`, `Ordered`, `Resolve`; urutan eksekusi diselesaikan dari method `DependsOn() []string` masing-masing modul, bukan dari iterasi map acak.
- **`Supervisor`** — `NewSupervisor`, `Add`, `Start`, `Stop`, `Health`, `Reload`, `WithShutdownTimeout`; menyupervisi proses jangka panjang dan background service.
- **Antarmuka Kapabilitas** — Antarmuka `Module`, `Startable`, dan `Depender` yang memungkinkan subsistem modular terhubung ke fase inisialisasi dan terminasi.\n