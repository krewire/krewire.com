---
title: "Model Workload & Primitif Domain"
description: "Model domain fondasi Krewire: delapan jenis beban kerja, metadata proyek, hierarki scope, aturan opt-in, taksonomi error, resolusi lingkungan, dan kontrak versi."
date: "2026-10-25"
---

# Model Workload & Primitif Domain (`pkg/model`, `pkg/errs`)

Jalur impor:

```text
github.com/krewire/krewire/pkg/model
github.com/krewire/krewire/pkg/errs
github.com/krewire/krewire/pkg/env
github.com/krewire/krewire/pkg/version
```

Sesuai ADR `ADR-B7D3A`, `pkg/model` **hanya boleh bergantung pada pustaka standar Go** dan saudara internalnya `errs` — tidak boleh bergantung pada paket tingkat lebih tinggi dan tidak boleh mengimpor `kiw`.

---

## Model Workload (`pkg/model`)

Delapan jenis workload resmi dideklarasikan sebagai konstanta bertipe:

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

// AllKinds mencantumkan setiap Kind valid dalam urutan kanonikal.
var AllKinds = []Kind{KindApp, KindCLI, KindSite, KindBook, KindWorker, KindService, KindGateway, KindInfra}

// IsValid melaporkan apakah k merupakan salah satu dari delapan kind yang dikenal.
func (k Kind) IsValid() bool

// ParseKind mengurai s sebagai Kind, mengembalikan usage error jika input tidak dikenal.
func ParseKind(s string) (Kind, error)
```

`AllKinds` adalah pemilik tunggal (SSOT) dari daftar kind: paket lain (termasuk detektor bentuk pada `kiw`) menurunkan nilainya dari sini alih-alih mendeklarasikan ulang kedelapan nama tersebut.

`pkg/model` juga mendeklarasikan `Project` (nama, path modul, jenis beban kerja, path konfigurasi) dengan method `Validate`, tangga `Scope` (`workspace` → `module` → `domain` → `service` → `unit`), tangga open-core `Tier`, serta batasan generik `OptInRule` / `ViolatesOptIn`.

---

## Penanganan Error & Kode Keluar (`pkg/errs`)

Krewire menstandardisasi galat di seluruh paket dengan taksonomi terstruktur dan kode keluar POSIX:

```go
import "github.com/krewire/krewire/pkg/errs"

// Kode keluar
errs.ExitCodeSuccess // 0
errs.ExitCodeFailure // 1
errs.ExitCodeUsage   // 2

// Pembuatan error
err := errs.UsageError("kiw build: target tidak dikenal")
err = errs.FailureError("kompilasi gagal")
err = errs.NewError("custom", errs.ExitCode(3))

// Konteks yang tetap bertahan dalam rantai error
err = errs.WithHint(err, "jalankan 'kiw init' terlebih dahulu")
err = errs.WithAttrs(err, errs.Attr{Key: "path", Value: "krewire.yaml"})
err = errs.WithStack(err)

fmt.Println(errs.FormatTree(err))
```

---

## Resolusi Lingkungan (`pkg/env`)

`pkg/env` memiliki kosakata lingkungan; resolusinya diperoleh dari konfigurasi (`env:` pada `krewire.yaml`, `KIW_ENV`, atau flag `--env`).

```go
import "github.com/krewire/krewire/pkg/env"

current, err := env.ParseEnv(cfg.Env) // input kosong menghasilkan env.DefaultEnv
if err != nil {
	return err
}

switch current {
case env.EnvLocal:
	// diagnostik debug dan hot reload
case env.EnvProduction:
	// TLS ketat dan cookie aman
case env.EnvTesting:
	// fixture deterministik
}
```

---

## Kontrak Versi (`pkg/version`)

Versi semantik diurai, dibandingkan, dan divalidasi:

```go
import "github.com/krewire/krewire/pkg/version"

v := version.MustParseVersion("0.1.0")
fmt.Println(v.String()) // 0.1.0

if !v.IsCompatible(version.MustParseVersion("0.1.0")) {
	return fmt.Errorf("versi tidak kompatibel %s", v)
}

// Gate kompatibilitas seluruh ekosistem yang digunakan oleh `kiw compat`.
err := version.CheckEcosystemCompatibility(required, actual)

// Watermark build yang dicetak oleh `kiw version`.
sig := version.BuildSignature()
```\n