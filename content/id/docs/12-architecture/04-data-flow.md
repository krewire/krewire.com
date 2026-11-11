---
title: "Alur Data: SSG vs Runtime"
description: "Bagaimana data mencapai view .kiw di Krewire: evaluasi saat build untuk site dan book, rendering saat request untuk app, handler Go yang menyuntikkan struct dan hasil query, serta skrip klien yang berkomunikasi dengan endpoint JSON."
date: "2026-11-09"
---

# Alur Data: SSG vs Runtime

<nav class="track-crumbs" aria-label="Jalur Pendalaman Arsitektur">
<span class="track-crumbs-label">Jalur pendalaman</span>
<ol>
<li><a href="/architecture">Ringkasan</a></li>
<li><a href="/architecture/layered-pyramid">1 · Piramida</a></li>
<li><a href="/architecture/ecosystem-dag">2 · DAG ekosistem</a></li>
<li><a href="/architecture/scope-levels">3 · Level cakupan</a></li>
<li aria-current="step"><span>4 · Alur data</span></li>
</ol>
</nav>

Berkas `.kiw` tidak pernah mengambil data sendiri. Berkas ini adalah template yang menerima sebuah **nilai data** lalu mengubahnya menjadi HTML. Yang membedakan antar-workload adalah *kapan* nilai tersebut dihitung dan *siapa* yang menghitungnya:

- Pada **`site` atau `book`**, data dievaluasi **satu kali, saat build**, oleh `kiw build` (`pkg/web/ssg`, `pkg/book`). Hasilnya berupa berkas biasa; tidak ada server yang menjalankan kode Anda.
- Pada **`app`**, data dihitung **per request** oleh fungsi Go yang Anda daftarkan pada `web.App` (`pkg/web`, dirakit dengan `pkg/app`). Renderer yang sama mengubahnya menjadi HTML pada setiap request.
- Pada keduanya, **skrip klien** dapat berkomunikasi dengan endpoint JSON setelah halaman dimuat — pada `app`, endpoint tersebut adalah handler Go Anda sendiri.

---

## 1. Siklus Hidup Request-Response

```text
BUILD TIME — site, book (runs once, on your machine or in CI)

  krewire.yaml ─┐
  content/*.md ─┤
  lang/*.json  ─┼──► kiw build ──► page data = yaml values + frontmatter
  pages/*.kiw  ─┤    (ssg.LoadFromDir)        + Collections + i18n
  components/  ─┤                                   │
  layouts/     ─┘                                   ▼
                                     ssg render: page .kiw + layout + kui
                                                    │
                                                    ▼
                                  .krewire/build/*.html ──► any static host


REQUEST TIME — app (runs on every request, inside your binary)

  Browser ── GET / ──► web.App router
                          │  middleware: Recover → AccessLog → CSRF → …
                          ▼
                  PageSpec.DataFunc(r) ──► your Go code ──► SQL, services
                          │  returns a struct / map / rows
                          ▼
                  ssg render: Root .kiw view + layout + kui   (same engine)
                          │  EmitProps → data-props JSON on the root element
                          ▼
  Browser ◄── HTML ───────┘
     │
     └── Scripts: fetch /api/… ──► Router().Get / Post ──► web.JSON(…) ──► Browser
```

Renderer di bagian tengah sama untuk kedua kolom. `web.App` membangun sebuah `ssg.Site` dari halaman, komponen, dan layout-nya, sehingga halaman dengan data tetap dirender identik byte demi byte, baik saat disajikan secara langsung maupun saat diekspor.

| | SSG saat build (`site`, `book`) | SSR saat request (`app`) |
| :--- | :--- | :--- |
| Siapa yang menghitung data | `kiw build` | Kode Go Anda pada setiap request (`PageSpec.DataFunc`) atau sekali saat startup (`PageSpec.Data`) |
| Sumber data | `krewire.yaml`, frontmatter halaman, `content/<collection>/*.md`, kamus `lang/` | Apa pun yang dapat dijangkau Go: database (`pkg/orm`), layanan, request (query, cookie, sesi, token CSRF) |
| Output | Berkas `.html` statis di `.krewire/build/` | Respons HTML per request, dari satu binary Go |
| Server yang dibutuhkan | Tidak ada — host statis apa pun | Binary aplikasi |
| Konten per pengguna | Hanya di sisi klien, setelah dimuat | Ya — dihitung per request |
| Kesegaran data | Hingga build berikutnya | Setiap request |
| Interaktivitas | Skrip klien yang memanggil API eksternal atau terpisah | Skrip klien yang memanggil rute JSON di binary yang sama |

---

## 2. Saat Build: Bagaimana Situs Mendapatkan Datanya

`kiw build` (melalui `ssg.LoadFromDir`) menggabungkan beberapa sumber ke dalam data setiap halaman:

| Sumber | Tersedia di template sebagai |
| :--- | :--- |
| `krewire.yaml` → `project.version` | `.Version` (`{version}`) |
| Frontmatter halaman (`title`, kunci kustom apa pun) | `.Title`, `.Description`, dan setiap kunci **dengan ejaan aslinya**, misalnya `{{.tagline}}` |
| Koleksi `content/<name>/*.md` | `.Collections.<name>` — daftar berisi `{Title, Date, Permalink, Data, Content, Draft, Slug, Description}` |
| `lang/<locale>/*.json` | fungsi template `t` / `translate` |
| `pages/<dir>/[slug].kiw` | satu halaman per entri koleksi, dengan `.Title`, `.Content`, dan `.Page` |

Sebuah halaman daftar ditambah satu halaman per postingan — dengan `content/posts/hello.md` yang memiliki frontmatter `title`, `slug`, dan `date`:

```html
---
title: Posts
layout: Base
tagline: "Notes from the build pipeline"
---
<h1>{{.Title}}</h1>
<p>{{.tagline}} · site version {version}</p>
<ul>
  {{range .Collections.posts}}
  <li><a href="/posts/{{.Slug}}">{{.Title}}</a> — {{.Date.Format "2 Jan 2006"}}</li>
  {{end}}
</ul>
```

```html
---
layout: Base
---
<article>
  <h1>{{.Title}}</h1>
  <div>{{.Content}}</div>
</article>
```

Jika disimpan sebagai `pages/posts.kiw` dan `pages/posts/[slug].kiw`, `kiw build --target site` akan menulis `posts.html`, `posts/hello.html`, dan `posts/second.html`:

```html
<p>Notes from the build pipeline · site version 0.1.0</p>
<ul>
  <li><a href="/posts/second">Second post</a> — 5 Nov 2026</li>
  <li><a href="/posts/hello">Hello, Krewire</a> — 1 Nov 2026</li>
</ul>
```

Tiga detail yang menghemat waktu debugging:

- Singkatan `{name}` mengubah huruf pertama menjadi kapital (`{tagline}` → `.Tagline`). Data yang dimuat dari YAML mempertahankan ejaannya sendiri, jadi gunakan bentuk eksplisit `{{.tagline}}` dan `{{range .Collections.posts}}` untuk kunci frontmatter dan nama koleksi.
- Situs yang memiliki folder `content/` juga akan diproses oleh pipeline book saat Anda menjalankan `kiw build` biasa. Gunakan `kiw build --target site` jika `content/` hanya berisi koleksi.
- Tidak ada apa pun di sisi ini yang berjalan saat request. Untuk menampilkan data yang berubah di antara build, lakukan build ulang (misalnya secara terjadwal di CI) atau ambil datanya dari peramban (bagian 5).

---

## 3. Saat Request: Handler Go Menyuntikkan Data

Pada `app`, setiap halaman adalah sebuah `web.PageSpec`. `Root` halaman menyebutkan view `.kiw` (yang didaftarkan sebagai komponen), dan datanya berasal dari salah satu dari dua field:

| Field | Dijalankan | Diekspor oleh `App.Export` | Digunakan untuk |
| :--- | :--- | :--- | :--- |
| `Data any` | tidak pernah — tetap sejak startup | **ya**, ke HTML statis | halaman pemasaran, dokumentasi, apa pun yang identik bagi setiap pengunjung |
| `DataFunc func(*http.Request) (any, error)` | pada setiap `GET` | tidak (halaman dinamis dilewati) | data per pengguna, query database, apa pun yang bergantung pada request |

Program berikut — contoh lengkap yang sudah dikompilasi — menunjukkan keduanya, ditambah API JSON yang digunakan oleh skrip klien di bagian 5. Dua view-nya berada di `views/`:

```html
<section>
  <h1>About v{version}</h1>
  <ul>
    {#each features as f}<li>{f}</li>{/each}
  </ul>
</section>
```

```html
<section id="notes">
  <h1>Notes <small>({count})</small></h1>
  <ul id="note-list">
    {#each notes as n}<li>{n.Text}</li>{/each}
  </ul>
  <form id="note-form">
    <input name="text" placeholder="New note" required>
    <button type="submit">Add</button>
  </form>
</section>
```

Dan sisi Go-nya:

```go
// Command notes demonstrates the three ways data reaches a page in a
// Krewire app: fixed data (exportable), per-request data, and JSON fetched
// by a client script.
package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/krewire/krewire/pkg/kiw"
	"github.com/krewire/krewire/pkg/validation"
	"github.com/krewire/krewire/pkg/web"
	"github.com/krewire/krewire/pkg/web/ssg"
)

//go:embed views/*.kiw static
var files embed.FS

// Note is the row type shared by the page, the API, and the script.
type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Store stands in for a database (see the fullstack tutorial for pkg/orm).
type Store struct {
	mu    sync.Mutex
	notes []Note
}

func (s *Store) List() []Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Note(nil), s.notes...)
}

func (s *Store) Add(text string) Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := Note{ID: len(s.notes) + 1, Text: text}
	s.notes = append(s.notes, n)
	return n
}

// AboutPage is fixed at startup, so the page can be exported to HTML.
type AboutPage struct {
	Version  string
	Features []string
}

// NotesPage is computed per request.
type NotesPage struct {
	Notes []Note
	Count int
	Token string
}

const shell = `<!doctype html><html><head><title>{{.Title}}</title></head>
<body><main>{{.Content}}</main></body></html>`

// NewApp wires the pages, the API, and the static script.
func NewApp(st *Store) (*web.App, error) {
	a := web.NewApp()
	a.Use(web.CSRF())
	a.Layout(ssg.Layout{Name: "Shell", Body: shell})

	views, _ := fs.Glob(files, "views/*.kiw")
	for _, p := range views {
		src, _ := files.ReadFile(p)
		mod, err := kiw.ParseKiw(string(src))
		if err != nil {
			return nil, err
		}
		name := strings.TrimSuffix(strings.TrimPrefix(p, "views/"), ".kiw")
		a.Component(ssg.Component{Name: name, Body: mod.Body})
	}

	// 1. Fixed data: rendered the same for everyone, exportable.
	a.Page(web.PageSpec{
		Path: "/about", Title: "About", Layout: "Shell", Root: "About",
		Data: AboutPage{Version: "1.0.0", Features: []string{"SSR", "JSON API", "Static export"}},
	})

	// 2. Per-request data: a handler-style function runs on every GET.
	a.Page(web.PageSpec{
		Path: "/", Title: "Notes", Layout: "Shell", Root: "Notes",
		DataFunc: func(r *http.Request) (any, error) {
			list := st.List()
			return NotesPage{Notes: list, Count: len(list), Token: web.CSRFFrom(r.Context())}, nil
		},
		EmitProps: true,                        // data-props='{"Notes":[…],…}' on the root element
		Scripts:   []string{"/static/notes.js"}, // client script that talks to the API
	})

	// 3. JSON endpoints for the client script.
	a.Router().Get("/api/notes", func(w http.ResponseWriter, r *http.Request, _ web.Params) {
		web.JSON(w, http.StatusOK, st.List())
	})
	a.Router().Post("/api/notes", func(w http.ResponseWriter, r *http.Request, _ web.Params) {
		var in struct {
			Text string `json:"text" validate:"required,max=200"`
		}
		if err := web.ReadJSON(r, &in); err != nil {
			web.Error(w, err)
			return
		}
		if err := validation.Struct(&in); err != nil {
			web.Error(w, web.BadRequest(err.Error()))
			return
		}
		web.JSON(w, http.StatusCreated, st.Add(in.Text))
	})

	static, _ := fs.Sub(files, "static")
	a.Static("/static", static)
	return a, nil
}

func main() {
	a, err := NewApp(&Store{})
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 && os.Args[1] == "export" {
		written, err := a.Export("dist")
		if err != nil {
			panic(err)
		}
		println(strings.Join(written, "\n"))
		return
	}
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	if err := a.Run(addr); err != nil {
		panic(err)
	}
}
```

Yang terjadi pada `GET /`:

1. Router menjalankan rantai middleware. `web.CSRF()` menerbitkan cookie `XSRF-TOKEN` dan menaruh token di context request.
2. `DataFunc` membangun struct `NotesPage` dari store (di aplikasi nyata, berupa query `pkg/orm` — lihat [tutorial fullstack](/getting-started/fullstack-app-tutorial)) dan membaca token dengan `web.CSRFFrom(r.Context())`.
3. Renderer mengeksekusi `Notes.kiw` dengan struct tersebut sebagai `.`: `{count}` adalah `.Count`, `{#each notes as n}` melakukan perulangan atas `.Notes`.
4. Karena `EmitProps` diaktifkan, struct tersebut juga diserialisasi ke atribut `data-props` pada elemen root view, dan `/static/notes.js` disisipkan sebelum `</body>`.

`go run . export` memanggil `App.Export("dist")`: perintah ini menulis `dist/about.html` (data tetap) dan melewati `/` (karena memiliki `DataFunc`).

### Struct, map, dan hasil query

Nilai Go apa pun dapat digunakan sebagai data halaman:

- **Struct** (disarankan): field dibaca berdasarkan nama — `{count}` membaca `.Count`, `{n.Text}` membaca field `Text` dari variabel perulangan.
- **Map**: `map[string]any{"Count": 3}` berperilaku seperti struct dengan kunci-kunci tersebut. Gunakan huruf kapital pada kunci agar singkatan `{count}` dapat menemukannya.
- **Hasil query**: kirimkan slice yang Anda scan, misalnya `[]models.Task` dari `db.Model(&models.Task{}).FindAll(&out)`, lalu lakukan perulangan dengan `{#each}`.

### Satu aturan untuk komponen: data masuk melalui props

Komponen Krewire UI (atau komponen buatan Anda sendiri) menerima data **sebagai atribut**: `<Badge text={summary} />`, `<Input value={draft} />`, `<Alert message={error} />`. Ekspresi atribut dievaluasi dalam konteks data halaman, sehingga dapat melihat `.Summary`, `.Draft`, `.Error`.

**Body** sebuah komponen — konten di antara `<Card>` dan `</Card>` — dirender secara terpisah, sebelum diserahkan ke komponen, dan tanpa data halaman. Field halaman yang digunakan di dalam body akan dirender kosong. Biarkan body tetap statis, atau render bagian dinamisnya di luar komponen dan kirimkan nilainya sebagai props. [`TaskBoard.kiw` pada tutorial](/getting-started/fullstack-app-tutorial) mengikuti pola ini: form dan daftar tugas ditulis sebagai HTML biasa di sekitar komponen self-closing.

### Dua hal yang (saat ini) masih dihubungkan sendiri oleh aplikasi live

- **View `.kiw`:** scaffold `--app` mendaftarkan komponen yang ditulis sebagai string Go. Untuk menulis view sebagai berkas `.kiw`, urai dengan `kiw.ParseKiw` dan daftarkan `ssg.Component{Name, Body, Style}` — seperti yang dilakukan `main.go` di atas dan `views.go` pada tutorial. Daftarkan katalog Krewire UI dari `ui.ComponentsFS` dengan cara yang sama.
- **CSS komponen:** halaman menautkan `/assets/style.css`, tetapi `web.App` yang berjalan belum menyajikan kumpulan blok `<style>` dari komponen `.kiw`. Gabungkan blok-blok tersebut dan daftarkan dengan `a.Asset("assets/style.css", css)` (lihat tutorial).

---

## 4. Memilih di Antara Keduanya

| Pertanyaan | Jika ya |
| :--- | :--- |
| Apakah kontennya sama untuk setiap pengunjung dan hanya berubah saat Anda menerbitkannya? | `site` / `book`, atau halaman `app` dengan `Data` yang Anda `Export`. |
| Apakah bergantung pada pengguna, sesi, atau baris database yang sering berubah? | Halaman `app` dengan `DataFunc`. |
| Apakah berupa bagian interaktif kecil dari halaman yang selebihnya statis? | Halaman statis + skrip klien yang memanggil endpoint JSON. |
| Apakah harus tetap berfungsi saat JavaScript dinonaktifkan? | Render di server (`DataFunc`) dan gunakan form HTML (post/redirect/get). |

[Pipeline progresif](/overview/why-krewire) dirancang agar Anda dapat memindahkan halaman dari kolom kiri ke kolom kanan tanpa menulis ulang view `.kiw`-nya: yang berubah hanya asal datanya.

---

## 5. Interaksi Sisi Klien: Skrip, `fetch`, dan Event

Skrip klien menyempurnakan halaman yang dirender di server; skrip ini tidak menggantikannya. Pada contoh di atas, `PageSpec.Scripts` menyisipkan `static/notes.js`:

```text
// Progressive enhancement: the list is already server-rendered; this script
// adds notes through the JSON API without a full page reload.
(function () {
  var root = document.querySelector('[data-props]');
  var props = root ? JSON.parse(root.getAttribute('data-props')) : {};
  var list = document.getElementById('note-list');
  var form = document.getElementById('note-form');

  function render(notes) {
    list.textContent = '';
    notes.forEach(function (n) {
      var li = document.createElement('li');
      li.textContent = n.text; // textContent, never innerHTML, for user data
      list.appendChild(li);
    });
  }

  form.addEventListener('submit', function (e) {
    e.preventDefault();
    fetch('/api/notes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'X-CSRF-Token': props.Token },
      body: JSON.stringify({ text: form.elements.text.value })
    })
      .then(function (r) { if (!r.ok) throw new Error(r.status); return fetch('/api/notes'); })
      .then(function (r) { return r.json(); })
      .then(function (notes) { render(notes); form.reset(); })
      .catch(function (err) { console.error('add note failed', err); });
  });
})();
```

Bagaimana bagian-bagiannya saling terhubung:

| Aspek | Mekanisme |
| :--- | :--- |
| State awal | `EmitProps: true` → `data-props='{"Notes":…,"Count":…,"Token":"…"}'` pada elemen root; skrip membacanya dengan `JSON.parse` alih-alih melakukan request pertama. |
| Membaca data | `fetch('/api/notes')` → handler `Router().Get` yang mengembalikan `web.JSON(w, 200, value)`. |
| Menulis data | `fetch` dengan `method: 'POST'` dan body JSON → handler `Router().Post` yang menggunakan `web.ReadJSON` dan `validation.Struct`. |
| CSRF pada `fetch` | Kirim token di header `X-CSRF-Token` (dari `data-props`, atau dari cookie `XSRF-TOKEN`, yang secara default dapat dibaca oleh skrip). Tanpa token, middleware menjawab `403 {"code":"forbidden","message":"csrf token mismatch"}`. |
| Error | `web.Error` menulis `{"code":…,"message":…}` dengan status yang tepat (`400` untuk validasi, `403` untuk CSRF), sehingga skrip dapat bercabang berdasarkan `r.ok`. |
| Render ulang | Perbarui DOM dengan `textContent` (jangan pernah `innerHTML` untuk data pengguna), atau muat ulang halaman untuk mendapatkan render server yang baru. |
| Event | Event DOM biasa (`submit`, `click`) dalam berkas skrip biasa; pada `app`, cantumkan URL-nya di `PageSpec.Scripts` dan sajikan dengan `App.Static`. |

Alur ini telah dijalankan di headless browser terhadap contoh tersebut: mengirim form menghasilkan `POST /api/notes` (`201`), diikuti `GET /api/notes` (`200`), dan daftar diperbarui tanpa memuat ulang halaman.

Untuk `site`, pola skrip yang sama berlaku, dengan JSON yang berasal dari API di tempat lain (`app` lain, `service`, atau endpoint pihak ketiga) — hasil build statis itu sendiri tidak memiliki rute server.

---

**Sebelumnya:** [12.3 Level Cakupan](/architecture/scope-levels) · **Kembali ke:** [Pendalaman Arsitektur](/architecture)
