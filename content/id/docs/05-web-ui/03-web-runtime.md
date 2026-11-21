---
title: "Web Runtime & WASM (pkg/runtime)"
description: "Menulis aplikasi web reaktif di sisi klien menggunakan Go yang dikompilasi ke WebAssembly: diffing virtual DOM, hooks, registri komponen, penanda hidrasi SSR, dan hidrasi pulau multi-framework."
date: "2026-10-25"
---

# Web Runtime & WASM (`pkg/runtime`)

Jalur impor: `github.com/krewire/krewire/pkg/runtime`

Pustaka `pkg/runtime` menghadirkan mesin eksekusi WebAssembly (WASM) di sisi klien, algoritma diffing dan rekonsiliasi Virtual DOM (VDOM), penjadwal siklus hidup komponen Fiber, serta kontrak hidrasi pulau (island hydration) progresif lintas framework.

Pustaka ini memungkinkan pengembang menulis widget sisi klien yang interaktif dan reaktif dalam 100% Go murni yang dikompilasi ke WebAssembly (`GOOS=js GOARCH=wasm`), atau menyematkan pulau frontend modular yang ringan (Svelte 5, Vue 3) sembari menjaga pemuatan awal halaman tetap statis dan bebas JavaScript secara bawaan.

---

## 1. Tata Letak Subpackage

```text
pkg/runtime/
├── vdom/         # Virtual DOM immutable: VNode, El, Text, Component, Diff, RenderHTML
├── component/    # Model komponen, penjadwal Fiber, hooks state, registri event
├── mount/        # Kontrak hidrasi progresif (data-kiw-mount, load, idle, visible)
├── boot/         # Memindai penanda data-kiw-mount dan menghidrasi komponen terdaftar
├── schedule/     # Pengelompokan render per-frame (requestAnimationFrame pada js/wasm)
└── build/        # BuildWASM: pembungkus kompilator pure-Go wasm dengan hashing konten
```

---

## 2. Hidrasi Pulau Progresif (`pkg/runtime/mount`)

Krewire menerapkan **Arsitektur Kepulauan (Islands Architecture)**: halaman dirender secara statis di server (SSG/SSR) tanpa beban JavaScript berlebih. Pulau-pulau interaktif dihidrasi secara selektif di peramban klien berdasarkan visibilitas layar atau periode senggang peramban.

### Kontrak Atribut HTML Hidrasi

Titik pemasangan (mount point) memancarkan penanda HTML standar:

```html
<div data-kiw-mount="Counter"
     data-kiw-framework="wasm"
     data-kiw-hydrate="visible"
     data-kiw-props='{"initial":10}'>
  <!-- HTML cadangan hasil render server -->
  <button class="btn">Hitungan: 10</button>
</div>
```

### Strategi Hidrasi (`mount.Hydrate`)

- **`load`** (`mount.Load`): Menghidrasi segera setelah modul klien atau bundle WASM selesai dimuat.
- **`idle`** (`mount.Idle`): Menghidrasi pada periode senggang peramban melalui `window.requestIdleCallback`.
- **`visible`** (`mount.Visible`): Menghidrasi hanya saat pulau masuk ke viewport peramban melalui `IntersectionObserver`.

### Dukungan Pulau Multi-Framework

Melalui `pkg/runtime/mount` dan ekosistem plugin Krewire, titik pemasangan mendukung:
- **`wasm`** (Bawaan): Pulau WebAssembly hasil kompilasi Go murni tanpa dependensi runtime node/npm.
- **`svelte`**: Pulau komponen reaktif Svelte 5 yang dipasang melalui `krew add svelte`.
- **`vue`**: Pulau komponen reaktif Vue 3 yang dipasang melalui `krew add vue`.

---

## 3. Mesin Virtual DOM (`pkg/runtime/vdom`)

```go
import "github.com/krewire/krewire/pkg/runtime/vdom"

tree := vdom.El("section", map[string]string{"class": "hero"},
	vdom.El("h1", nil, vdom.Text("Halo dari Go")),
	vdom.Component("Counter", "counter-1", map[string]string{"initial": "0"}),
)

html := vdom.RenderHTML(tree)        // Server-side rendering (SSR)
patches := vdom.Diff(oldTree, tree) // Rekonsiliasi di sisi klien
```

`VNode` adalah deskripsi struktur pohon yang bersifat immutable. Fungsi `RenderHTML` dan `Diff` sama-sama memproses properti melalui `NormalizeProps`, memastikan output server dan rekonsiliasi klien memiliki keselarasan mutlak pada nama atribut, kelas CSS, serta penanda event.

---

## 4. Model Komponen & Hooks Reaktif (`pkg/runtime/component`)

Komponen adalah entitas apa pun yang mengimplementasikan antarmuka `Component`:

```go
type Component interface {
	Render() *vdom.VNode
}
```

Logika berstatus (stateful) menggunakan hooks generik yang terikat ke `Fiber` aktif dan menegakkan urutan eksekusi yang stabil:

```go
package main

import (
	"context"
	"fmt"

	"github.com/krewire/krewire/pkg/runtime/component"
	"github.com/krewire/krewire/pkg/runtime/vdom"
)

type Counter struct{}

func (c *Counter) Render() *vdom.VNode {
	count, setCount := component.UseState(0)

	component.UseEvent("click", "inc", func(_ context.Context, _ component.Event) {
		setCount(count + 1)
	})

	return vdom.El("div", map[string]string{"class": "counter"},
		vdom.El("h2", nil, vdom.Text(fmt.Sprintf("Hitungan: %d", count))),
		vdom.El("button", map[string]string{"data-kiw-on": "click:inc"}, vdom.Text("Tambah")),
	)
}

func init() {
	component.MustRegister("Counter", func(component.Props) component.Component {
		return &Counter{}
	})
}
```

Hooks yang tersedia mencakup `UseState`, `UseRef`, `UseMemo`, `UseEffect`, dan `UseEvent`. Memanggil hook di luar siklus render akan memicu panic dengan galat `*component.HookRuleError`, yang secara langsung mendeteksi kesalahan urutan hooks.

---

## 5. Pemindai Boot Klien (`pkg/runtime/boot`)

Runtime klien secara otomatis melakukan booting dan memasang komponen saat DOM selesai dimuat:

```go
package main

import (
	"github.com/krewire/krewire/pkg/runtime/boot"
	"github.com/krewire/krewire/pkg/runtime/component"
)

func main() {
	// 1. Mendaftarkan komponen Go interaktif
	component.Register("Counter", func() component.Component {
		return &Counter{}
	})

	// 2. Memindai DOM untuk penanda [data-kiw-mount] dan menghidrasi pulau
	boot.Run()
}
```

---

## 6. Membangun Modul Biner WASM (`pkg/runtime/build`)

```go
import "github.com/krewire/krewire/pkg/runtime/build"

artifacts, err := build.BuildWASM(build.Config{
	Entry:  "./frontend",
	OutDir: "public/wasm",
	Name:   "counter",
})
if err != nil {
	return err
}

fmt.Println(artifacts.WASM, artifacts.JS, artifacts.Digest)
```

Dari CLI pengembang, pipeline kompilasi yang sama dapat dipanggil langsung:

```bash
krew build --target wasm
```

---

## 7. Keselarasan dengan Tiga Pilar (3S)

- **Aman (S1 - Secure)**: WASM dieksekusi di dalam sandbox hermetik peramban WebAssembly tanpa akses memori langsung yang tidak aman. VDOM meng-escape otomatis semua simpul teks dan nilai atribut untuk mencegah XSS.
- **Berkelanjutan (S2 - Sustainable)**: Nol dependensi runtime JavaScript untuk beban kerja Go WASM murni. Hidrasi pulau progresif menjamin pengguna hanya mengunduh skrip interaktif untuk bagian halaman yang benar-benar membutuhkan dinamisme.
- **Skalabel (S3 - Scalable)**: Menyatukan Go WASM, Svelte 5, dan Vue 3 di bawah satu kontrak hidrasi deklaratif yang seragam (`data-kiw-hydrate`).
