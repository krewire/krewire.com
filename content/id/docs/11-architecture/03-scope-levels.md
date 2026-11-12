---
title: "Level Cakupan"
description: "Lima level cakupan arsitektur Krewire — Workspace, Module, Domain, Service, Unit — sebagaimana dikodifikasi oleh model.Scope, serta bagaimana spesifikasi, tes, dan keputusan ekstraksi menggunakannya."
date: "2026-11-09"
---

# Level Cakupan

<nav class="track-crumbs" aria-label="Jalur Pendalaman Arsitektur">
<span class="track-crumbs-label">Jalur pendalaman</span>
<ol>
<li><a href="/architecture">Ringkasan</a></li>
<li><a href="/architecture/layered-pyramid">1 · Piramida</a></li>
<li><a href="/architecture/ecosystem-dag">2 · DAG ekosistem</a></li>
<li aria-current="step"><span>3 · Level cakupan</span></li>
<li><a href="/architecture/data-flow">4 · Alur data</a></li>
</ol>
</nav>

Lapisan menjawab pertanyaan "siapa boleh mengimpor siapa". **Level cakupan** menjawab pertanyaan lain: "seberapa besar hal yang sedang kita bicarakan?" Setiap spesifikasi, tes, dan dokumen desain di repositori Krewire menyebutkan cakupan yang ditargetkannya, sehingga reviewer langsung tahu apakah sebuah perubahan menyentuh satu fungsi atau seluruh workspace.

## 1. Lima Level

```text
Workspace  ───►  Module  ───►  Domain  ───►  Service  ───►  Unit
 (level 0)      (level 1)     (level 2)     (level 3)     (level 4)
```

| Level | Cakupan | Apa itu | Contoh |
| :--- | :--- | :--- | :--- |
| 0 | **Workspace** | Workspace Go multi-repositori yang berakar pada berkas `go.work`, beserta binary `kiw` dan repositori-repositori sejajarnya. | `krewire/` + `ui/` + `krew/` di-checkout berdampingan |
| 1 | **Module** | Modul Go — unit yang Anda beri versi, rilis, dan impor. | `github.com/krewire/krewire`, `github.com/krewire/ui` |
| 2 | **Domain** | Bounded context di dalam modul: sekumpulan layanan yang kohesif, biasanya `internal/<domain>/`. Sebelum diekstraksi, domain adalah sekelompok paket; setelah diekstraksi, domain dapat menjadi modulnya sendiri. | `internal/catalog`, `internal/user` |
| 3 | **Service** | Unit runtime yang dapat di-deploy, diimplementasikan sebagai paket Go `main`. | `cmd/api/`, `service/billing/` |
| 4 | **Unit** | Potongan kode terkecil: paket, berkas, tipe, atau fungsi. | `pkg/ui`, `ui.Button` |

Nama-nama lama tidak lagi diterima: gunakan **Module** alih-alih "Project", dan **Unit** alih-alih "Package" atau "Func".

## 2. Cakupan sebagai Kode (`pkg/model`)

Hierarki ini bukan sekadar dokumentasi — ia adalah sebuah tipe di paket inti hermetis `pkg/model`:

```go
package main

import (
	"fmt"

	"github.com/krewire/krewire/pkg/model"
)

func main() {
	for _, s := range model.AllScopes {
		fmt.Println(s, s.Level()) // Workspace 0 … Unit 4
	}

	s, err := model.ParseScope(" domain ") // case-insensitive, trims spaces
	fmt.Println(s, err)                    // Domain <nil>

	fmt.Println(model.ScopeService.Less(model.ScopeUnit)) // true: Service is the wider scope

	_, err = model.ParseScope("package")
	fmt.Println(err) // unknown scope "package": want one of Workspace, Module, Domain, Service, Unit
}
```

| API | Perilaku |
| :--- | :--- |
| `model.AllScopes` | Kelima cakupan dalam urutan kanonikal. |
| `Scope.IsValid()` | Melaporkan apakah nilai tersebut merupakan salah satu dari lima cakupan. |
| `Scope.Level()` | `0` (Workspace) hingga `4` (Unit); `-1` untuk cakupan yang tidak valid. |
| `Scope.Less(other)` | Apakah cakupan tersebut berada di urutan sebelum (lebih luas dari) `other`. |
| `model.ParseScope(s)` | Mengurai nama tanpa membedakan huruf besar-kecil; nama yang tidak dikenal mengembalikan error penggunaan. |

## 3. Menerapkan Level Cakupan dalam Praktik

- **Spesifikasi dan issue** menyatakan cakupannya, sehingga "tambahkan varian `Button`" (Unit) tidak pernah di-review seperti "pisahkan domain catalog menjadi modul tersendiri" (Domain → Module).
- **Tes** diorganisasi berdasarkan cakupan: unit test berada di samping kode, tes service dijalankan terhadap paket `main` yang sedang berjalan, dan pemeriksaan tingkat workspace seperti `layering_test.go` mencakup lintas modul.
- **Keputusan ekstraksi** memindahkan kode *naik* satu level setiap kali: sebuah Unit berkembang menjadi Domain di dalam monolit; Domain menjadi Service (dengan `main` sendiri) ketika perlu diskalakan atau di-deploy secara independen; Service menjadi Module tersendiri ketika tim lain harus memberinya versi secara terpisah. Setiap langkah mempertahankan kodenya; hanya pengemasannya yang berubah — inilah [pipeline progresif](/overview/why-krewire) dari satu `app` menuju workload `service` dan `worker`.

---

**Sebelumnya:** [11.2 DAG Ekosistem](/architecture/ecosystem-dag) · **Berikutnya:** [**11.4 Alur Data: SSG vs Runtime →**](/architecture/data-flow)
