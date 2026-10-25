---
title: "Model, Siklus Hidup & Error (pkg/model, pkg/lifecycle, pkg/errs)"
description: "Control plane fondasi Krewire: primitif model workload, supervisi siklus hidup, taksonomi error, deteksi environment, dan kontrak versi."
date: "2026-10-25"
---

# Model, Siklus Hidup & Error

Jalur impor (*import paths*):

```text
github.com/krewire/krewire/pkg/model
github.com/krewire/krewire/pkg/lifecycle
github.com/krewire/krewire/pkg/errs
github.com/krewire/krewire/pkg/env
github.com/krewire/krewire/pkg/version
```

Kelima paket ini membentuk fondasi arsitektur ekosistem. Mereka mendefinisikan primitif domain, jenis-jenis workload, supervisi siklus hidup (*lifecycle supervision*), taksonomi error, resolusi environment, dan kontrak versi semantik.

Sesuai ADR `ADR-B7D3A`, `pkg/model` **hanya boleh bergantung pada pustaka standar Go (stdlib)** dan paket saudara internalnya `errs` — tidak boleh bergantung pada paket tingkat yang lebih tinggi dan tidak boleh mengimpor `kiw`.

---

## Tata Letak Paket

```text
pkg/
├── model/        # Jenis workload, scope proyek, tier lisensi, aturan opt-in
├── lifecycle/    # Kernel supervisor, registri modul, graceful shutdown
├── errs/         # Klasifikasi error dan kode keluar POSIX
├── env/          # Target environment (local, production, testing)
└── version/      # Parsing versi semantik dan validasi kompatibilitas
```

---

## Model Workload (`pkg/model`)

Delapan jenis workload resmi dideklarasikan sebagai konstanta bertipe (*typed constants*):

```go
package model

type Kind string

const (
	KindApp     Kind = "app"
	KindCLI     Kind = "cli"
	KindSite    Kind = "site"
	KindBook    Kind = "book"
	KindWorker  Kind = "worker"
	KindService Kind = "service"
	KindGateway Kind = "gateway"
	KindInfra   Kind = "infra"
)

// AllKinds mendaftar setiap Kind yang valid dalam urutan kanonikal.
var AllKinds = []Kind{KindApp, KindCLI, KindSite, KindBook, KindWorker, KindService, KindGateway, KindInfra}

// IsValid melaporkan apakah k merupakan salah satu dari 8 jenis yang dikenali.
func (k Kind) IsValid() bool

// ParseKind mem-parsing string s menjadi Kind, menghasilkan usage error pada input yang tidak dikenal.
func ParseKind(s string) (Kind, error)
```

`AllKinds` adalah satu-satunya pemilik daftar jenis workload: seluruh paket lainnya (termasuk shape detector milik `kiw`) menurunkan nilainya dari sini tanpa mendefinisikan ulang kedelapan nama tersebut.

`pkg/model` juga mendeklarasikan `Project` (nama, module path, kind, path konfigurasi) dengan method `Validate`, tangga cakupan `Scope` (`workspace` → `module` → `domain` → `service` → `unit`), tangga lisensi `Tier`, serta aturan kendala `OptInRule` / `ViolatesOptIn` yang mencegah penautan paket tambahan pada workload yang tidak memilihnya.

---

## Supervisi Siklus Hidup (`pkg/lifecycle`)

`lifecycle.Kernel` adalah pengendali imperatif: ia mengikat `model.Project` yang telah divalidasi ke modul-modul terdaftar dan mengeksekusi workload yang dideklarasikan.

```go
k, err := lifecycle.New(project)
if err != nil {
	return err
}

k.Use(bootModule, webModule)          // modul diinisialisasi sesuai urutan dependensi
k.RegisterHandler(model.KindApp, run) // satu handler per jenis workload

if err := k.Boot(ctx); err != nil {
	return err
}
defer k.Shutdown(ctx)

os.Exit(k.Execute(ctx, workload).Int())
```

Tipe pendukung:

- **`Registry`** — `NewRegistry`, `Register`, `Ordered`, `Resolve`; urutan eksekusi diselesaikan dari `DependsOn() []string` masing-masing modul, bukan dari iterasi map acak.
- **`Supervisor`** — `NewSupervisor`, `Add`, `Start`, `Stop`, `Health`, `Reload`, `WithShutdownTimeout`; mengawasi workload yang berjalan lama (*long-running*).
- **`Module` / `Startable` / `Depender`** — antarmuka kemampuan opsional (*opt-in capability interfaces*).

---

## Penanganan Error & Kode Keluar (`pkg/errs`)

Krewire menstandarisasi penanganan error di seluruh paket dengan taksonomi terstruktur dan kode keluar POSIX:

```go
import "github.com/krewire/krewire/pkg/errs"

// Kode keluar standar
errs.ExitCodeSuccess // 0
errs.ExitCodeFailure // 1
errs.ExitCodeUsage   // 2

// Konstruksi error
err := errs.UsageError("kiw build: unknown target")
err = errs.FailureError("build failed")
err = errs.NewError("custom", errs.ExitCode(3))

// Konteks yang tetap bertahan dalam rantai error
err = errs.WithHint(err, "jalankan 'kiw init' terlebih dahulu")
err = errs.WithAttrs(err, errs.Attr{Key: "path", Value: "krewire.yaml"})
err = errs.WithStack(err)

fmt.Println(errs.FormatTree(err))
```

Fungsi `errs.HintOf`, `errs.AttrsOf`, dan `errs.StackOf` membaca kembali konteks tersebut, sehingga satu handler di level teratas mampu merender pohon diagnostik lengkap yang mudah dibaca.

---

## Resolusi Lingkungan (`pkg/env`)

`pkg/env` memiliki kosakata environment; resolusi nilai environment itu sendiri berasal dari konfigurasi (`env:` di `krewire.yaml`, variabel `KIW_ENV`, atau argumen `--env`).

```go
import "github.com/krewire/krewire/pkg/env"

current, err := env.ParseEnv(cfg.Env) // input kosong menghasilkan env.DefaultEnv
if err != nil {
	return err
}

switch current {
case env.EnvLocal:
	// diagnostik debug dan live-reload
case env.EnvProduction:
	// TLS ketat dan cookie aman
case env.EnvTesting:
	// fixture deterministik
}
```

`env.Envs()` mengekspos urutan kanonikal (`local`, `production`, `testing`).

---

## Kontrak Versi (`pkg/version`)

Versi semantik di-parsing, dibandingkan, dan divalidasi:

```go
import "github.com/krewire/krewire/pkg/version"

v := version.MustParseVersion("0.1.0")
fmt.Println(v.String()) // 0.1.0

if !v.IsCompatible(version.MustParseVersion("0.1.0")) {
	return fmt.Errorf("versi tidak kompatibel %s", v)
}

// Gerbang kompatibilitas seluruh ekosistem yang digunakan oleh `kiw compat`.
err := version.CheckEcosystemCompatibility(required, actual)

// Watermark build yang dicetak oleh perintah `kiw version`.
sig := version.BuildSignature()
```
