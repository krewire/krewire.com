---
title: "Model Workload & Primitif Domain"
description: "Model domain fondasi Krewire: delapan jenis beban kerja, metadata proyek, hierarki scope, aturan opt-in, taksonomi error, resolusi lingkungan, dan kontrak versi."
date: "2026-10-25"
---

# Model Workload & Primitif Domain (`model`, `errs`)

Jalur impor:

```text
github.com/krewire/krewire/model
github.com/krewire/krewire/errs
github.com/krewire/krewire/env
github.com/krewire/krewire/version
```

Sebagai paket fondasi inti, `model` **hanya boleh bergantung pada pustaka standar Go** dan saudara internalnya `errs` — tidak boleh bergantung pada paket tingkat lebih tinggi dan tidak boleh mengimpor `krew`.

---

## Model Deklaratif & Batasan (`model`)

Paket `model` mendefinisikan model fundamental, cakupan hierarki arsitektur, tangga lisensi, dan aturan validasi invarian untuk ekosistem Krewire:

```go
package model

type Project struct {
	Name       string `json:"name"`
	ModulePath string `json:"modulePath"`
	ConfigPath string `json:"configPath"`
}

// Validate memvalidasi invarian proyek (nama kebab-case, path modul yang aman).
func (p Project) Validate() error
```

`model` juga mendeklarasikan hierarki `Scope` (`Workspace` → `Module` → `Domain` → `Service` → `Unit`), tangga lisensi open-core `Tier` (`free`, `pro`, `team`, `enterprise`), primitif `DomainEvent` yang imutabel, serta kontrol keamanan dasar.

---

## Penanganan Error & Kode Keluar (`errs`)

Krewire menstandardisasi galat di seluruh paket dengan taksonomi terstruktur dan kode keluar POSIX:

```go
import "github.com/krewire/krewire/errs"

// Kode keluar
errs.ExitCodeSuccess // 0
errs.ExitCodeFailure // 1
errs.ExitCodeUsage   // 2

// Pembuatan error
err := errs.UsageError("krew build: target tidak dikenal")
err = errs.FailureError("kompilasi gagal")
err = errs.NewError("custom", errs.ExitCode(3))

// Konteks yang tetap bertahan dalam rantai error
err = errs.WithHint(err, "jalankan 'krew init' terlebih dahulu")
err = errs.WithAttrs(err, errs.Attr{Key: "path", Value: "krewire.yaml"})
err = errs.WithStack(err)

fmt.Println(errs.FormatTree(err))
```

---

## Resolusi Lingkungan (`env`)

`env` memiliki kosakata lingkungan; resolusinya diperoleh dari konfigurasi (`env:` pada `krewire.yaml`, `KIW_ENV`, atau flag `--env`).

```go
import "github.com/krewire/krewire/env"

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

## Kontrak Versi (`version`)

Versi semantik diurai, dibandingkan, dan divalidasi:

```go
import "github.com/krewire/krewire/version"

v := version.MustParseVersion("0.1.0")
fmt.Println(v.String()) // 0.1.0

if !v.IsCompatible(version.MustParseVersion("0.1.0")) {
	return fmt.Errorf("versi tidak kompatibel %s", v)
}

// Gate kompatibilitas seluruh ekosistem yang digunakan oleh `krew compat`.
err := version.CheckEcosystemCompatibility(required, actual)

// Watermark build yang dicetak oleh `krew version`.
sig := version.BuildSignature()
```\n