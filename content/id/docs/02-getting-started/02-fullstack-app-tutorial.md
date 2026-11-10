---
title: "Membangun Aplikasi Fullstack dalam 15 Menit"
description: "Tutorial dari nol hingga mahir: scaffold aplikasi dengan kiw new --app, simpan data di SQLite melalui pkg/orm beserta migrasi kiw migrate, render halaman .kiw dengan komponen Krewire UI, tangani form tervalidasi yang dilindungi CSRF, dan rilis satu binary."
date: "2026-11-09"
---

# Membangun Aplikasi Fullstack dalam 15 Menit

<nav class="track-crumbs" aria-label="Jalur pemula">
<span class="track-crumbs-label">Jalur pemula</span>
<ol>
<li><a href="/">Pilih jalur</a></li>
<li><a href="/getting-started">Memulai</a></li>
<li><a href="/getting-started/quickstart">1 · Quickstart</a></li>
<li aria-current="step"><span>2 · Aplikasi fullstack</span></li>
<li><a href="/getting-started/dsl">3 · DSL .kiw</a></li>
</ol>
</nav>

Anda akan membangun **todo-app**, sebuah pengelola tugas kecil: halaman yang menampilkan daftar tugas yang tersimpan di SQLite, form untuk menambah tugas dengan validasi di sisi server, tombol untuk menyelesaikan dan menghapus tugas, proteksi CSRF di setiap form, dan satu binary yang dapat Anda salin ke server.

<div class="overview-grid">

  <div class="chapter-card">
    <p class="chapter-card-desc"><b>Waktu:</b> sekitar 15 menit.<br><b>Yang Anda perlukan:</b> Go 1.27.1+, <code>kiw</code> (lihat <a href="/getting-started/quickstart">Quickstart</a>), dan akses jaringan untuk mengunduh modul Go.<br><b>Yang Anda pelajari:</b> <code>kiw new --app</code>, <code>kiw make</code>, <code>pkg/orm</code>, <code>kiw migrate</code>, view <code>.kiw</code>, <code>pkg/validation</code>, <code>web.CSRF</code>, <code>kiw build</code>.</p>
  </div>

</div>

Krewire menyediakan baterai fullstack lengkap: router HTTP, dependency injection, multi-driver database ORM, migrasi skema, validasi, middleware CSRF, dan tampilan komponen `.kiw`.

---

## Langkah 1: Scaffold Aplikasi

```bash
kiw new todo-app --app
cd todo-app
```

```text
INFO scaffolded Krewire project name=todo-app variant=app files=17
created todo-app/go.mod
created todo-app/krewire.yaml
created todo-app/main.go
created todo-app/internal/app/app.go
created todo-app/internal/config/config.go
created todo-app/internal/http/http.go
created todo-app/web/layouts/shell.go
created todo-app/web/pages/pages.go
created todo-app/web/components/components.go
created todo-app/web/theme/theme.go
created todo-app/assets/embed.go
created todo-app/assets/public/app.css
created todo-app/assets/public/app.js
...
```

Bagian-bagian yang akan Anda ubah:

```text
todo-app/
├── main.go                 # loads krewire.yaml, builds the app, runs it — unchanged
├── krewire.yaml            # name, kind: app, dev.port: 8080 — gains a database section
├── internal/config/config.go # mirrors krewire.yaml in Go — gains a Database field
├── internal/app/app.go     # wiring: providers registered in the app container
├── internal/http/http.go   # a sample JSON API (/api/status) — unchanged
├── web/layouts/shell.go    # the HTML shell around every page — unchanged
└── web/pages/pages.go      # the pages and their data
```

Jalankan sekali untuk melihat hasil scaffold:

```bash
kiw dev
```

Buka <http://localhost:8080>, lalu hentikan server dengan `Ctrl+C`.

---

## Langkah 2: Definisikan Model Data dan Migrasi

### 2.1 Arahkan aplikasi ke database

Tambahkan bagian `database:` ke `krewire.yaml`. Aplikasi membacanya untuk membuka database, dan `kiw migrate` membaca key yang sama:

```yaml
# yaml-language-server: $schema=https://krewire.com/schemas/krewire.json
project:
  name: todo-app
  kind: app
  version: 0.1.0

dev:
  port: 8080

database:
  driver: sqlite
  dsn: todo.db
```

Cerminkan bagian ini di `internal/config/config.go`: tambahkan field `Database` ke `Metadata`, dan tipe baru di bawah `Dev`:

```go
// Metadata mirrors krewire.yaml.
type Metadata struct {
	Project  Project  `yaml:"project"`
	Dev      Dev      `yaml:"dev"`
	Database Database `yaml:"database"`
}

// Database holds the database section of krewire.yaml, the same keys
// `kiw migrate` reads. APP_DATABASE_DSN overrides dsn for both.
type Database struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn" validate:"required"`
}
```

`LoadMetadata` hasil scaffold sudah menimpa nilai dengan environment, sehingga `APP_DATABASE_DSN` menimpa `dsn` baik untuk aplikasi maupun untuk `kiw migrate`. `validate:"required"` membuat aplikasi menolak start tanpa DSN, alih-alih diam-diam memakai database sementara di memori.

Tidak ada yang perlu dipasang untuk SQLite: driver Go murni disertakan di inti sebagai `pkg/orm/sqlite` (tanpa compiler C, ramah `CGO_ENABLED=0`). Cukup import di Langkah 2.4.

### 2.2 Buat model

```bash
kiw make model Task
```

Perintah ini membuat `internal/models/task.go`. Ganti isinya dengan:

```go
package models

import (
	"time"
)

// Task is one row of the tasks table.
type Task struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Done      bool      `json:"done" db:"done"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// TableName maps Task to the "tasks" table (pkg/orm would default to "task").
func (Task) TableName() string { return "tasks" }
```

`pkg/orm` memetakan field yang diekspor ke kolom snake_case (`CreatedAt` → `created_at`), memperlakukan `ID` bertipe integer sebagai primary key auto-increment, dan menggunakan `TableName()` jika model mendefinisikannya.

### 2.3 Tulis migrasi

```bash
kiw make migration create_tasks
```

`kiw` membuat `migrations/<timestamp>_create_tasks.sql` dengan bagian `-- +up` dan `-- +down` yang masih kosong. Isi keduanya:

```sql
-- Migration: create_tasks
-- Created at: 2026-11-09T12:38:54Z

-- +up
CREATE TABLE IF NOT EXISTS tasks (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  title      TEXT     NOT NULL,
  done       INTEGER  NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL
);

-- +down
DROP TABLE IF EXISTS tasks;
```

Sematkan folder migrasi ke dalam binary dengan `migrations/embed.go`:

```go
// Package migrations embeds the SQL files created by `kiw make migration`.
package migrations

import "embed"

//go:embed *.sql
var Files embed.FS
```

Terapkan migrasi dan lihat hasilnya:

```bash
kiw migrate up
kiw migrate status
```

```text
applied 20261109123854_create_tasks (284µs)
STATE    MIGRATION                    APPLIED AT
applied  20261109123854_create_tasks  2026-11-09T15:56:47Z
```

`kiw migrate up` membuat `todo.db` dan mencatat migrasi di tabel `schema_migrations`, beserta checksum berkasnya.

### 2.4 Buka database dan terapkan migrasi

Buat `internal/db/db.go`:

```go
// Package db opens the SQLite database and applies the embedded migrations.
package db

import (
	"context"

	"github.com/krewire/krewire/pkg/orm"
	"github.com/krewire/krewire/pkg/orm/migrate"
	"github.com/krewire/krewire/pkg/orm/sqlite" // bundled pure-Go SQLite driver

	"todo-app/migrations"
)

// Open opens (or creates) the SQLite file at path and applies every pending
// migration, so a fresh deployment starts with the right schema.
func Open(ctx context.Context, path string) (*orm.DB, error) {
	db, err := sqlite.Open(path)
	if err != nil {
		return nil, err
	}
	if err := migrate.Up(ctx, db, migrations.Files); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}
```

`sqlite.Open` membuat berkas jika belum ada dan mengaktifkan WAL, foreign key, serta busy timeout 5 detik. `migrate.Up` menerapkan setiap migrasi yang belum tercatat di `schema_migrations`, masing-masing dalam transaksinya sendiri dan di bawah lock, sehingga `kiw migrate` dan beberapa instance aplikasi bisa start bersamaan dan setiap migrasi tetap berjalan sekali. Binary yang di-deploy pun memigrasikan databasenya sendiri saat startup.

**Alur kerja migrasi mulai sekarang:**

1. `kiw migrate create add_due_date` (sama dengan `kiw make migration add_due_date`) — berkas baru bertimestamp akan terurut setelah berkas-berkas sebelumnya.
2. Tulis perubahannya di `-- +up` (misalnya `ALTER TABLE tasks ADD COLUMN due_date DATETIME;`) dan pembaliknya di `-- +down`.
3. `kiw migrate up` menerapkannya; aplikasi juga akan menerapkannya saat start berikutnya.
4. `kiw migrate status` menampilkan setiap migrasi sebagai `pending`, `applied`, `modified`, atau `missing`. `kiw migrate down` membatalkan migrasi terakhir (`down 2`: dua terakhir) dan `kiw migrate redo` membatalkan lalu menerapkannya lagi — berguna saat Anda mengulang-ulang migrasi baru.
5. Jangan pernah mengedit migrasi yang sudah diterapkan di tempat lain: runner membandingkan checksum dan berhenti dengan error, dan `kiw migrate status` menampilkan berkas itu sebagai `modified`. Tulis perubahannya di migrasi baru.

> **Jalan pintas untuk prototipe:** `db.AutoMigrate(ctx, &models.Task{})` membuat tabel langsung dari struct (`CREATE TABLE IF NOT EXISTS`). Perintah ini tidak pernah mengubah tabel yang sudah ada, jadi beralihlah ke berkas migrasi begitu Anda memiliki data yang perlu dipertahankan.

---

## Langkah 3: Tulis Query CRUD

Buat `internal/tasks/tasks.go` berisi store. Daftar impornya sudah mencakup apa yang ditambahkan Langkah 5 ke berkas yang sama, sehingga proyek dapat dikompilasi kembali di akhir Langkah 5:

```go
// Package tasks holds the task store and the HTML form handlers.
package tasks

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/krewire/krewire/pkg/orm"
	"github.com/krewire/krewire/pkg/validation"
	"github.com/krewire/krewire/pkg/web"

	"todo-app/internal/models"
)

// Store wraps the database with the queries the app needs.
type Store struct{ DB *orm.DB }

// List returns open tasks first, newest first.
func (s *Store) List(ctx context.Context) ([]models.Task, error) {
	var out []models.Task
	err := s.DB.Model(&models.Task{}).WithContext(ctx).
		OrderBy("done", "ASC").OrderBy("id", "DESC").
		FindAll(&out)
	return out, err
}

// Create inserts a task and fills in its ID.
func (s *Store) Create(ctx context.Context, t *models.Task) error {
	_, err := s.DB.Model(t).WithContext(ctx).Insert(t)
	return err
}

// Toggle flips the done flag of one task.
func (s *Store) Toggle(ctx context.Context, id int64) error {
	var t models.Task
	if err := s.DB.Model(&t).WithContext(ctx).Find(&t, id); err != nil {
		return err
	}
	t.Done = !t.Done
	_, err := s.DB.Model(&t).WithContext(ctx).Update(&t)
	return err
}

// Delete removes one task.
func (s *Store) Delete(ctx context.Context, id int64) error {
	_, err := s.DB.Model(&models.Task{}).WithContext(ctx).Delete(&models.Task{}, id)
	return err
}

```

API query-nya terbaca seperti SQL: `Model(&T{})` memilih tabel, `Where`, `OrderBy`, `Limit`, dan `Offset` mempersempitnya, lalu `FindAll`, `Find`, `Insert`, `Update`, `Delete`, atau `Count` menjalankannya. Nilai selalu dikirim sebagai parameter terikat (bound parameter), tidak pernah sebagai SQL yang digabungkan.

---

## Langkah 4: Bangun Halaman dalam `.kiw` dengan Krewire UI

### 4.1 View

Buat `web/views/TaskBoard.kiw`:

```html
<main class="board">
  <header class="board-head">
    <h1>{title}</h1>
    <Badge variant="primary" text={summary} />
  </header>

  {#if error}
  <Alert variant="error" title="Task not saved" message={error} />
  {/if}

  <form method="post" action="/tasks" class="task-form">
    <input type="hidden" name="csrf_token" value="{token}">
    <Input name="title" id="title" label="New task" placeholder="Write the release notes" value={draft} />
    <Button type="submit" variant="primary" text="Add task" />
  </form>

  <ul class="task-list">
    {#each tasks as task}
    <li class="task-row">
      <form method="post" action="/tasks/{task.ID}/toggle">
        <input type="hidden" name="csrf_token" value="{token}">
        <button type="submit" class="task-check" aria-label="Toggle task">{#if task.Done}✓{:else}○{/if}</button>
      </form>
      <span class="task-title">{task.Title}</span>
      {#if task.Done}<Badge variant="success" size="sm" text="done" />{/if}
      <form method="post" action="/tasks/{task.ID}/delete">
        <input type="hidden" name="csrf_token" value="{token}">
        <Button type="submit" variant="ghost" size="sm" text="Delete" />
      </form>
    </li>
    {:else}
    <li class="task-empty">No tasks yet. Add your first one above.</li>
    {/each}
  </ul>

  <Card title="How this page works" variant="default">
    <p>GET / runs a Go function that queries SQLite and passes the rows to this .kiw view.
    Every form posts back to a Go handler that validates the input and redirects.</p>
  </Card>
</main>

<style>
.board { max-width: 40rem; margin: 0 auto; padding: 2rem 1rem; display: grid; gap: 1.25rem; }
.board-head { display: flex; align-items: center; justify-content: space-between; gap: 1rem; }
.board-head h1 { margin: 0; font-size: 1.75rem; }
.task-form { display: grid; gap: 0.75rem; justify-items: start; }
.task-form .field-group { width: 100%; }
.task-list { list-style: none; margin: 0; padding: 0; display: grid; gap: 0.5rem; }
.task-row { display: flex; align-items: center; gap: 0.75rem; padding: 0.6rem 0.8rem; border: 2px solid var(--border, #0B1F3B); border-radius: 10px; }
.task-row form { margin: 0; }
.task-title { flex: 1; }
.task-check { width: 2rem; height: 2rem; border-radius: 50%; border: 2px solid var(--border, #0B1F3B); background: transparent; color: inherit; cursor: pointer; font-weight: 800; }
.task-empty { opacity: 0.7; padding: 0.6rem 0; }
</style>
```

- `{title}`, `{token}`, `{draft}` membaca field dari data halaman (`Title`, `Token`, `Draft`) — huruf pertamanya otomatis dijadikan kapital.
- `{#each tasks as task} … {:else} … {/each}` melakukan perulangan atas baris-baris data dan menampilkan bagian `{:else}` jika tidak ada data; `{task.ID}` dan `{task.Title}` membaca baris yang sedang diproses.
- `<Badge>`, `<Alert>`, `<Input>`, `<Button>`, dan `<Card>` adalah komponen Krewire UI. Jelajahi props-nya di [katalog UI](/ui/components).

> **Patokan untuk data dinamis:** kirimkan data request ke komponen **sebagai props** (`text={summary}`, `value={draft}`, `message={error}`). Konten *di antara* tag sebuah komponen (body-nya, seperti paragraf di dalam `<Card>`) dirender secara terpisah, tanpa akses ke data halaman — biarkan tetap statis. Itulah sebabnya form dan daftar tugas ditulis sebagai HTML biasa di sekitar komponen self-closing. [Panduan alur data](/architecture/data-flow) menjelaskan mekanismenya.

### 4.2 Muat view `.kiw` ke dalam aplikasi

Buat `web/views/views.go`. Berkas ini mendaftarkan seluruh katalog Krewire UI beserta setiap berkas `.kiw` di folder tersebut sebagai komponen `web.App`:

```go
// Package views turns the .kiw files in this directory into components the
// web app can render, the same way `kiw build` loads components/ for a site.
package views

import (
	"embed"
	"io/fs"
	"strings"

	"github.com/krewire/krewire/pkg/kiw"
	"github.com/krewire/krewire/pkg/ui"
	"github.com/krewire/krewire/pkg/web"
	"github.com/krewire/krewire/pkg/web/ssg"
)

//go:embed *.kiw
var files embed.FS

// Register parses the Krewire UI catalog (Card, Input, Button, Badge, Alert,
// ...) and this package's own .kiw views, and registers them as components
// on the app. A view with the same name as a UI component wins.
func Register(a *web.App) error {
	var css strings.Builder
	if err := register(a, ui.ComponentsFS, &css); err != nil {
		return err
	}
	if err := register(a, files, &css); err != nil {
		return err
	}
	// Pages link /assets/style.css automatically, but a live web.App does not
	// serve the component <style> blocks yet, so serve them ourselves.
	a.Asset("assets/style.css", css.String())
	return nil
}

func register(a *web.App, fsys fs.FS, css *strings.Builder) error {
	names, err := fs.Glob(fsys, "*.kiw")
	if err != nil {
		return err
	}
	for _, name := range names {
		src, err := fs.ReadFile(fsys, name)
		if err != nil {
			return err
		}
		mod, err := kiw.ParseKiw(string(src))
		if err != nil {
			return err
		}
		style := strings.Join(mod.Styles, "\n")
		css.WriteString(style + "\n")
		a.Component(ssg.Component{
			Name:  strings.TrimSuffix(name, ".kiw"),
			Body:  mod.Body,
			Style: style,
		})
	}
	return nil
}
```

### 4.3 Pasang halaman

Ganti `web/pages/pages.go`:

```go
// Package pages defines the application's SSR page definitions
// (FRK-STR-006): data + mounts registered at assembly time.
package pages

import (
	"net/http"

	"github.com/krewire/krewire/pkg/web"
)

// Register wires all application pages into the web app. board computes the
// TaskBoard data per request, so "/" is a dynamic (server-rendered) page.
func Register(a *web.App, board func(*http.Request) (any, error)) {
	a.Page(web.PageSpec{
		Path:     "/",
		Title:    "Tasks",
		Layout:   "Shell",
		Root:     "TaskBoard",
		DataFunc: board,
	})
}
```

Halaman dengan `DataFunc` bersifat **dinamis**: fungsinya dijalankan pada setiap request, dan hasilnya itulah yang dilihat oleh `TaskBoard.kiw`. (Sebaliknya, halaman dengan `Data` tetap bersifat statis dan dapat diekspor ke HTML saat build.)

---

## Langkah 5: Tangani Form dengan Validasi dan Proteksi CSRF

### 5.1 Data halaman dan handler form

Tambahkan sisa isi `internal/tasks/tasks.go`:

```go
// BoardData is the data the TaskBoard.kiw view renders.
type BoardData struct {
	Title   string
	Tasks   []models.Task
	Summary string // e.g. "2 open"
	Token   string // CSRF token for the hidden form fields
	Error   string
	Draft   string
}

// Board loads the data for GET /. Validation errors come back from the POST
// handler as query parameters (post/redirect/get).
func (s *Store) Board(r *http.Request) (any, error) {
	list, err := s.List(r.Context())
	if err != nil {
		return nil, err
	}
	remaining := 0
	for _, t := range list {
		if !t.Done {
			remaining++
		}
	}
	return BoardData{
		Title:   "Tasks",
		Tasks:   list,
		Summary: strconv.Itoa(remaining) + " open",
		Token:   web.CSRFFrom(r.Context()),
		Error:   r.URL.Query().Get("error"),
		Draft:   r.URL.Query().Get("title"),
	}, nil
}

// newTask is the form input, validated with pkg/validation struct tags.
type newTask struct {
	Title string `validate:"required,max=120"`
}

// CreateHandler handles POST /tasks.
func (s *Store) CreateHandler(w http.ResponseWriter, r *http.Request, _ web.Params) {
	in := newTask{Title: strings.TrimSpace(r.PostFormValue("title"))}
	if err := validation.Struct(&in); err != nil {
		msg := "Title is required."
		if in.Title != "" {
			msg = "Title must be 120 characters or fewer."
		}
		q := url.Values{"error": {msg}, "title": {in.Title}}
		web.Redirect(w, r, "/?"+q.Encode(), http.StatusSeeOther)
		return
	}
	t := &models.Task{Title: in.Title, CreatedAt: time.Now().UTC()}
	if err := s.Create(r.Context(), t); err != nil {
		web.Error(w, err)
		return
	}
	web.Redirect(w, r, "/", http.StatusSeeOther)
}

// ToggleHandler handles POST /tasks/{id}/toggle.
func (s *Store) ToggleHandler(w http.ResponseWriter, r *http.Request, p web.Params) {
	id, err := strconv.ParseInt(p["id"], 10, 64)
	if err != nil {
		web.Error(w, web.BadRequest("invalid task id"))
		return
	}
	if err := s.Toggle(r.Context(), id); err != nil {
		web.Error(w, err)
		return
	}
	web.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeleteHandler handles POST /tasks/{id}/delete.
func (s *Store) DeleteHandler(w http.ResponseWriter, r *http.Request, p web.Params) {
	id, err := strconv.ParseInt(p["id"], 10, 64)
	if err != nil {
		web.Error(w, web.BadRequest("invalid task id"))
		return
	}
	if err := s.Delete(r.Context(), id); err != nil {
		web.Error(w, err)
		return
	}
	web.Redirect(w, r, "/", http.StatusSeeOther)
}
```

- **Validasi:** `validation.Struct` memeriksa tag `validate:"required,max=120"`. Pangkas (trim) input terlebih dahulu — `required` saja masih menerima string yang hanya berisi spasi.
- **Post/Redirect/Get:** setiap `POST` dijawab dengan `303 See Other`. Saat terjadi error validasi, pesan dan judul yang diketik dikirim balik sebagai parameter query, lalu `Board` menampilkannya di `<Alert>` dan `<Input>`.
- **ID dari URL:** rute menggunakan placeholder `{id}`; nilainya tersedia di `p["id"]`.

### 5.2 Hubungkan semuanya

Perbarui `internal/app/app.go`. Dibandingkan dengan hasil scaffold, berkas ini menambahkan `dbProvider` yang membuka `meta.Database.DSN`, mendaftarkan view, memasang `rvweb.CSRF()`, dan memasang tiga rute `POST`:

```go
// Package app composes the application: it wires every
// dependency through the app container and returns the ready
// *web.App. cmd only calls this and runs.
package app

import (
	"context"
	"io/fs"

	rvapp "github.com/krewire/krewire/pkg/app"
	rvweb "github.com/krewire/krewire/pkg/web"
	"todo-app/assets"
	"todo-app/internal/config"
	"todo-app/internal/db"
	"todo-app/internal/http"
	"todo-app/internal/tasks"
	"todo-app/web/components"
	"todo-app/web/layouts"
	"todo-app/web/pages"
	"todo-app/web/theme"
	"todo-app/web/views"
)

type Config struct {
	Meta *config.Metadata
}

// webProvider assembles the SSR side of the app: theme, layout, pages, and
// embedded static assets.
type webProvider struct{}

// Register binds the web app singleton.
func (p *webProvider) Register(c *rvapp.Container) error {
	return rvapp.Singleton[*rvweb.App](c, func() *rvweb.App { return rvweb.NewApp() })
}

// Boot configures the app once the container can resolve it.
func (p *webProvider) Boot(c *rvapp.Container) error {
	a, err := rvapp.Resolve[*rvweb.App](c)
	if err != nil {
		return err
	}
	store, err := rvapp.Resolve[*tasks.Store](c)
	if err != nil {
		return err
	}
	a.Theme(theme.New())
	a.Layout(layouts.Shell())
	a.Component(components.Hero())
	if err := views.Register(a); err != nil {
		return err
	}

	// CSRF: GET requests receive an XSRF-TOKEN cookie, every POST must echo
	// it back in the csrf_token form field (or X-CSRF-Token header).
	a.Use(rvweb.CSRF())

	sub, err := fs.Sub(assets.Files, "public")
	if err != nil {
		return err
	}
	a.Static("/static", sub)
	pages.Register(a, store.Board)
	a.Router().Post("/tasks", store.CreateHandler)
	a.Router().Post("/tasks/{id}/toggle", store.ToggleHandler)
	a.Router().Post("/tasks/{id}/delete", store.DeleteHandler)
	return nil
}

// dbProvider opens the SQLite database, applies migrations, and binds the
// task store into the container.
type dbProvider struct{ path string }

// Register opens the database and binds the store.
func (p *dbProvider) Register(c *rvapp.Container) error {
	conn, err := db.Open(context.Background(), p.path)
	if err != nil {
		return err
	}
	return rvapp.Singleton[*tasks.Store](c, func() *tasks.Store {
		return &tasks.Store{DB: conn}
	})
}

// Boot does nothing for the database.
func (p *dbProvider) Boot(c *rvapp.Container) error { return nil }

// New assembles the full application and returns the ready *web.App
func New(meta *config.Metadata) (*rvweb.App, error) {
	container, err := rvapp.NewApp(
		&configProvider{meta: meta},
		&dbProvider{path: meta.Database.DSN},
		&webProvider{},
		&http.Provider{},
	).Build()
	if err != nil {
		return nil, err
	}
	return rvapp.Resolve[*rvweb.App](container)
}

type configProvider struct {
	meta *config.Metadata
}

// Register binds the loaded config into the container.
func (p *configProvider) Register(c *rvapp.Container) error {
	return rvapp.Singleton[*Config](c, func() *Config {
		return &Config{Meta: p.meta}
	})
}

// Boot does nothing for config.
func (p *configProvider) Boot(c *rvapp.Container) error { return nil }
```

Cara kerja proteksi CSRF di sini: pada `GET`, `web.CSRF()` menetapkan cookie `XSRF-TOKEN` dan mengekspos token yang sama melalui `web.CSRFFrom(r.Context())`, yang dimasukkan `Board` ke field tersembunyi `csrf_token` di setiap form. Pada `POST`, middleware membandingkan field tersebut (atau header `X-CSRF-Token`) dengan cookie dan menjawab `403` jika keduanya berbeda — sebelum handler Anda dijalankan.

---

## Langkah 6: Jalankan dan Uji

```bash
go mod tidy
kiw dev
```

Buka <http://localhost:8080>: tambahkan tugas, kirim judul kosong untuk melihat error validasi, ubah status tugas, dan hapus salah satunya. Data tersimpan di `todo.db` (atur `APP_DATABASE_DSN` untuk menggunakan berkas lain).

Periksa pelindung CSRF dari terminal kedua — `POST` tanpa token akan ditolak:

```bash
curl -i -d title=Hack http://localhost:8080/tasks
```

```text
HTTP/1.1 403 Forbidden
...
{"code":"forbidden","message":"csrf token mismatch"}
```

Kunci perilaku ini dengan tes, `internal/app/app_test.go`:

```go
package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"todo-app/internal/config"
)

// TestTaskFlow drives the real handler: GET issues a CSRF token, a POST
// without it is rejected, a valid POST stores a task that GET then lists.
func TestTaskFlow(t *testing.T) {
	a, err := New(&config.Metadata{
		Project:  config.Project{Name: "todo-app", Kind: "app"},
		Database: config.Database{Driver: "sqlite", DSN: filepath.Join(t.TempDir(), "test.db")},
	})
	if err != nil {
		t.Fatal(err)
	}
	h := a.Handler()

	get := httptest.NewRecorder()
	h.ServeHTTP(get, httptest.NewRequest(http.MethodGet, "/", nil))
	if get.Code != http.StatusOK {
		t.Fatalf("GET / = %d", get.Code)
	}
	cookie := get.Result().Cookies()[0]
	token := regexp.MustCompile(`name="csrf_token" value="([^"]+)"`).FindStringSubmatch(get.Body.String())[1]

	post := func(form url.Values) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.AddCookie(cookie)
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		return rec
	}

	if rec := post(url.Values{"title": {"no token"}}); rec.Code != http.StatusForbidden {
		t.Fatalf("POST without CSRF token = %d, want 403", rec.Code)
	}
	if rec := post(url.Values{"csrf_token": {token}, "title": {"  "}}); !strings.Contains(rec.Header().Get("Location"), "error=") {
		t.Fatalf("blank title should redirect with an error, got %q", rec.Header().Get("Location"))
	}
	if rec := post(url.Values{"csrf_token": {token}, "title": {"Write the tutorial"}}); rec.Code != http.StatusSeeOther {
		t.Fatalf("valid POST = %d, want 303", rec.Code)
	}

	list := httptest.NewRecorder()
	h.ServeHTTP(list, httptest.NewRequest(http.MethodGet, "/", nil))
	if !strings.Contains(list.Body.String(), "Write the tutorial") {
		t.Fatal("new task not listed on GET /")
	}
}
```

```bash
kiw test
```

```text
INFO running tests module=todo-app
INFO go test args="test ./..."
ok  	todo-app/internal/app	0.552s
```

---

## Langkah 7: Build Satu Binary dan Deploy

```bash
kiw build
```

```text
INFO building binary kind=app entry=. output=.../todo-app/.krewire/build/bin/todo-app
created .../todo-app/.krewire/build/bin/todo-app
```

Binary ini menyematkan view, komponen Krewire UI, migrasi, dan aset statis. Driver SQLite-nya Go murni, sehingga Anda dapat melakukan cross-compile untuk server Linux dari mesin mana pun:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 kiw build
file .krewire/build/bin/todo-app
```

```text
.krewire/build/bin/todo-app: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, ...
```

Deploy dengan menyalin **dua berkas** — `main.go` hasil scaffold membaca `krewire.yaml` dari direktori kerja saat startup:

```bash
scp .krewire/build/bin/todo-app krewire.yaml user@server:/opt/todo-app/
ssh user@server 'cd /opt/todo-app && APP_ADDR=:8080 APP_DATABASE_DSN=/opt/todo-app/data.db ./todo-app'
```

Saat pertama kali dijalankan, binary akan membuat berkas database dan menerapkan migrasi. Untuk unit systemd, reverse proxy dengan TLS, atau image kontainer, lanjutkan ke [Deployment & Operasi](/guides/deployment-and-operations).

---

## Proyek Pendamping

Kode lengkap dan teruji dari tutorial ini ada di repositori krewire.com:

```bash
git clone --depth 1 https://github.com/krewire/krewire.com
cd krewire.com/examples/todo-app
go test ./...
go run .
```

`go.mod`-nya mengunci versi inti persis yang digunakan untuk memverifikasi tutorial ini.

## Rangkuman

| Yang Anda gunakan | Untuk |
| :--- | :--- |
| `kiw new --app`, `kiw make model`, `kiw make migration` | Scaffold aplikasi, model, dan migrasi SQL |
| `pkg/orm` + `pkg/orm/sqlite` | Akses SQLite dengan API query fluent yang terparameterisasi dan driver Go murni bawaan |
| `kiw migrate` + `pkg/orm/migrate` | Menerapkan, memeriksa, dan membatalkan migrasi — dari terminal maupun saat startup |
| `.kiw` + Krewire UI | Halaman, dengan `Badge`, `Alert`, `Input`, `Button`, dan `Card` |
| `web.PageSpec{DataFunc}` | Data per request untuk halaman |
| `pkg/validation` + post/redirect/get | Validasi sisi server dengan pesan error yang ramah |
| `web.CSRF()` | Proteksi CSRF berbasis token di setiap form |
| `kiw dev`, `kiw test`, `kiw build` | Mengembangkan, memverifikasi, dan merilis satu binary |

## Langkah Selanjutnya

- Pahami kapan data di-resolve saat build dan kapan per request: [Alur Data: SSG vs Runtime](/architecture/data-flow).
- Pelajari bagian-bagiannya lebih dalam: [Workload App](/workloads/app), [`pkg/web`](/web-ui/web), [Database & ORM](/data-services/database-and-orm), [Kontrol keamanan](/security-platforms/security).
- Pelajari setiap fitur `.kiw`: [**2.7 DSL `.kiw` →**](/getting-started/dsl)
