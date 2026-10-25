---
title: "Web Runtime & WASM (pkg/runtime)"
description: "Menulis aplikasi web reaktif di sisi klien menggunakan Go yang dikompilasi ke WebAssembly: diffing virtual DOM, hooks, registri komponen, penanda hidrasi SSR, dan pipeline build WASM yang dapat direproduksi."
date: "2026-10-25"
---

# Web Runtime & WASM (`pkg/runtime`)

Keluarga package `pkg/runtime` mengompilasi kode Go ke WebAssembly (`GOOS=js GOARCH=wasm`) agar komponen pulau interaktif (interactive islands) dapat berbagi model komponen yang sama dengan halaman `.kiw` yang dirender di sisi server (SSR). Ini adalah sebuah **pustaka (library)**, bukan jenis beban kerja (workload kind): antarmuka frontend WASM dikirimkan di dalam proyek `app` atau `site` dan dibangun menggunakan perintah `kiw build --target wasm`.

---

## Tata Letak Subpackage

```text
pkg/runtime/
├── vdom/         # Virtual DOM immutable: VNode, El, Text, Component, Diff, RenderHTML
├── component/    # Model komponen, hooks, event handler, registri komponen
├── mount/        # Kontrak hidrasi SSR yang digunakan bersama oleh pembuat (emitter) dan boot klien
├── boot/         # Memindai penanda data-kiw-mount dan menghidrasi komponen terdaftar
├── schedule/     # Pengelompokan render per-frame (requestAnimationFrame pada js/wasm)
└── build/        # BuildWASM: pembungkus go build dengan hashing dan diagnostik
```

---

## Virtual DOM (`pkg/runtime/vdom`)

```go
import "github.com/krewire/krewire/pkg/runtime/vdom"

tree := vdom.El("section", map[string]string{"class": "hero"},
	vdom.El("h1", nil, vdom.Text("Halo dari Go")),
	vdom.Component("Counter", "counter-1", map[string]string{"initial": "0"}),
)

html := vdom.RenderHTML(tree)        // server-side rendering (SSR)
patches := vdom.Diff(oldTree, tree) // rekonsiliasi sisi klien
```

`VNode` adalah deskripsi struktur pohon yang bersifat immutable. Fungsi `RenderHTML` dan `Diff` sama-sama menyalurkan properti melalui `NormalizeProps`, sehingga output server dan rekonsiliasi klien memiliki kesesuaian identik pada nama atribut, kelas CSS, serta penanda event.

---

## Model Komponen (`pkg/runtime/component`)

Komponen adalah entitas apa pun yang merender pohon virtual:

```go
type Component interface {
	Render() *vdom.VNode
}
```

Logika berstatus (stateful) menggunakan hooks generik yang terikat ke `Fiber` yang sedang dirender dan menegakkan urutan panggilan yang stabil:

```go
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

Hooks yang tersedia: `UseState`, `UseRef`, `UseMemo`, `UseEffect`, `UseEvent`. Partisipasi siklus hidup bersifat pilihan melalui antarmuka kapabilitas `Mounter`, `Unmounter`, dan `ShouldUpdater`. Memanggil hook di luar fungsi render akan memicu panic dengan galat `*component.HookRuleError`, yang secara eksplisit memunculkan kesalahan urutan panggilan alih-alih merusak state secara diam-diam.

---

## Hidrasi SSR (`pkg/runtime/mount` dan `pkg/runtime/boot`)

Server mengeluarkan fragmen HTML yang dibungkus dalam kontainer dengan tiga atribut khusus:

```html
<div data-kiw-mount="Counter" data-kiw-hydrate="load" data-kiw-props='{"initial":0}'>
  <!-- server-rendered HTML -->
</div>
```

`boot` memindai penanda-penanda tersebut pada sisi klien, membuat instansiasi komponen yang terdaftar, dan mengisikan fiber awal dengan DOM hasil render server sehingga proses render pertama melakukan diffing alih-alih rendering ulang penuh. Pembaruan berikutnya digabungkan melalui `schedule`, yang mengelompokkan notifikasi ke dalam satu proses flush per frame animasi (`requestAnimationFrame`).

---

## Membangun Modul WASM (`pkg/runtime/build`)

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

`BuildWASM` membungkus `go build` dengan flag lingkungan `GOOS=js GOARCH=wasm`, menyalin skrip perekat `wasm_exec.js`, menghasilkan hash konten modul, serta menyediakan diagnostik komprehensif saat terjadi ketidakcocokan toolchain maupun versi.

Dari CLI `kiw`, pipeline yang sama dapat dipanggil langsung sebagai target build:

```bash
kiw build --target wasm
```

Modul biner hasil kompilasi dan skrip perekat JS akan ditempatkan di direktori output Anda, siap dipasang ke halaman HTML mana pun yang dilayani oleh workload `app` atau `site`.
