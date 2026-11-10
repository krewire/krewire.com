---
title: "Build a Fullstack App in 15 Minutes"
description: "Zero-to-hero tutorial: scaffold an app with kiw new --app, store data in SQLite through pkg/orm with kiw migrate migrations, render .kiw pages with Krewire UI components, handle validated CSRF-protected forms, and ship a single binary."
date: "2026-11-09"
---

# Build a Fullstack App in 15 Minutes

<nav class="track-crumbs" aria-label="Beginner track">
<span class="track-crumbs-label">Beginner track</span>
<ol>
<li><a href="/">Choose a path</a></li>
<li><a href="/getting-started">Getting Started</a></li>
<li><a href="/getting-started/quickstart">1 · Quickstart</a></li>
<li aria-current="step"><span>2 · Fullstack app</span></li>
<li><a href="/getting-started/dsl">3 · The .kiw DSL</a></li>
</ol>
</nav>

You will build **todo-app**, a small task manager: a page that lists tasks stored in SQLite, a form to add a task with server-side validation, buttons to complete and delete tasks, CSRF protection on every form, and a single binary you can copy to a server.

<div class="overview-grid">

  <div class="chapter-card">
    <p class="chapter-card-desc"><b>Time:</b> about 15 minutes.<br><b>You need:</b> Go 1.27.1+, <code>kiw</code> (see the <a href="/getting-started/quickstart">Quickstart</a>), and network access to download Go modules.<br><b>You learn:</b> <code>kiw new --app</code>, <code>kiw make</code>, <code>pkg/orm</code>, <code>kiw migrate</code>, <code>.kiw</code> views, <code>pkg/validation</code>, <code>web.CSRF</code>, <code>kiw build</code>.</p>
  </div>

</div>

Krewire provides the fullstack battery: HTTP router, dependency injection, multi-driver database ORM, schema migrations, validation, CSRF middleware, and `.kiw` component views.

---

## Step 1: Scaffold the App

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

The parts you will touch:

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

Run it once to see the scaffold:

```bash
kiw dev
```

Open <http://localhost:8080>, then stop the server with `Ctrl+C`.

---

## Step 2: Define the Data Model and Migration

### 2.1 Point the app at a database

Add a `database:` section to `krewire.yaml`. The app reads it to open the database, and `kiw migrate` reads the same keys:

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

Mirror the section in `internal/config/config.go`: add the `Database` field to `Metadata`, and the new type below `Dev`:

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

The scaffolded `LoadMetadata` already overlays the environment, so `APP_DATABASE_DSN` overrides `dsn` for the app and for `kiw migrate` alike. `validate:"required"` makes the app refuse to start without a DSN instead of quietly using a throwaway in-memory database.

There is nothing to install for SQLite: the pure-Go driver ships with core as `pkg/orm/sqlite` (no C compiler, `CGO_ENABLED=0` friendly). Importing it in Step 2.4 is enough.

### 2.2 Generate the model

```bash
kiw make model Task
```

This creates `internal/models/task.go`. Replace its contents with:

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

`pkg/orm` maps exported fields to snake_case columns (`CreatedAt` → `created_at`), treats an integer `ID` as an auto-increment primary key, and uses `TableName()` when the model defines it.

### 2.3 Write the migration

```bash
kiw make migration create_tasks
```

`kiw` creates `migrations/<timestamp>_create_tasks.sql` with empty `-- +up` and `-- +down` sections. Fill them in:

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

Embed the migrations folder into the binary with `migrations/embed.go`:

```go
// Package migrations embeds the SQL files created by `kiw make migration`.
package migrations

import "embed"

//go:embed *.sql
var Files embed.FS
```

Apply the migration and look at the result:

```bash
kiw migrate up
kiw migrate status
```

```text
applied 20261109123854_create_tasks (284µs)
STATE    MIGRATION                    APPLIED AT
applied  20261109123854_create_tasks  2026-11-09T15:56:47Z
```

`kiw migrate up` created `todo.db` and recorded the migration in the `schema_migrations` table, together with a checksum of the file.

### 2.4 Open the database and apply migrations

Create `internal/db/db.go`:

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

`sqlite.Open` creates the file when it is missing and turns on WAL, foreign keys, and a 5-second busy timeout. `migrate.Up` applies every migration not yet recorded in `schema_migrations`, each in its own transaction and under a lock, so `kiw migrate` and several app instances can start at the same time and each migration still runs once. A deployed binary therefore migrates its own database on start-up.

**The migration workflow from now on:**

1. `kiw migrate create add_due_date` (the same as `kiw make migration add_due_date`) — a new timestamped file sorts after the previous ones.
2. Write the change in `-- +up` (for example `ALTER TABLE tasks ADD COLUMN due_date DATETIME;`) and its reversal in `-- +down`.
3. `kiw migrate up` applies it; the app would also apply it on its next start.
4. `kiw migrate status` lists each migration as `pending`, `applied`, `modified`, or `missing`. `kiw migrate down` rolls back the last migration (`down 2`: the last two) and `kiw migrate redo` rolls it back and applies it again — handy while you iterate on a new migration.
5. Never edit a migration that is already applied elsewhere: the runner compares checksums and stops with an error, and `kiw migrate status` shows the file as `modified`. Put the change in a new migration instead.

> **Prototyping shortcut:** `db.AutoMigrate(ctx, &models.Task{})` creates a table straight from the struct (`CREATE TABLE IF NOT EXISTS`). It never alters existing tables, so switch to migration files once you have data to keep.

---

## Step 3: Write the CRUD Queries

Create `internal/tasks/tasks.go` with the store. The imports already include what Step 5 adds to this same file, so the project compiles again at the end of Step 5:

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

The query API reads like SQL: `Model(&T{})` picks the table, `Where`, `OrderBy`, `Limit`, and `Offset` narrow it, and `FindAll`, `Find`, `Insert`, `Update`, `Delete`, or `Count` run it. Values always travel as bound parameters, never as concatenated SQL.

---

## Step 4: Build the Page in `.kiw` with Krewire UI

### 4.1 The view

Create `web/views/TaskBoard.kiw`:

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

- `{title}`, `{token}`, `{draft}` read fields of the page data (`Title`, `Token`, `Draft`) — the first letter is capitalized for you.
- `{#each tasks as task} … {:else} … {/each}` loops over the rows and shows the `{:else}` part when there are none; `{task.ID}` and `{task.Title}` read the current row.
- `<Badge>`, `<Alert>`, `<Input>`, `<Button>`, and `<Card>` are Krewire UI components. Browse their props in the [UI catalog](/ui/components).

> **Rule of thumb for dynamic data:** pass request data to components **as props** (`text={summary}`, `value={draft}`, `message={error}`). The content *between* a component's tags (its body, like the paragraph inside `<Card>`) is rendered on its own, without access to the page data — keep it static. That is why the form and the list are plain HTML around self-closing components. The [data-flow guide](/architecture/data-flow) explains the mechanics.

### 4.2 Load `.kiw` views into the app

Create `web/views/views.go`. It registers the whole Krewire UI catalog plus every `.kiw` file in the folder as components of the `web.App`:

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

### 4.3 Mount the page

Replace `web/pages/pages.go`:

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

A page with `DataFunc` is **dynamic**: the function runs on every request, and its result is what `TaskBoard.kiw` sees. (A page with fixed `Data` instead is static and can be exported to HTML at build time.)

---

## Step 5: Handle Forms with Validation and CSRF Protection

### 5.1 Page data and form handlers

Append the rest of `internal/tasks/tasks.go`:

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

- **Validation:** `validation.Struct` checks the `validate:"required,max=120"` tags. Trim input first — `required` alone accepts a string of spaces.
- **Post/Redirect/Get:** every `POST` answers `303 See Other`. On a validation error the message and the typed title travel back as query parameters, and `Board` shows them in the `<Alert>` and the `<Input>`.
- **IDs from the URL:** routes use `{id}` placeholders; the value arrives in `p["id"]`.

### 5.2 Wire everything together

Update `internal/app/app.go`. Compared with the scaffold, it adds a `dbProvider` that opens `meta.Database.DSN`, registers the views, installs `rvweb.CSRF()`, and mounts the three `POST` routes:

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

How CSRF protection works here: on a `GET`, `web.CSRF()` sets an `XSRF-TOKEN` cookie and exposes the same token through `web.CSRFFrom(r.Context())`, which `Board` puts in every form's hidden `csrf_token` field. On a `POST`, the middleware compares the field (or an `X-CSRF-Token` header) with the cookie and answers `403` when they differ — before your handler runs.

---

## Step 6: Run and Test It

```bash
go mod tidy
kiw dev
```

Open <http://localhost:8080>: add a task, submit an empty title to see the validation error, toggle a task, delete one. The data lands in `todo.db` (set `APP_DATABASE_DSN` to use another file).

Check the CSRF guard from a second terminal — a `POST` without the token is rejected:

```bash
curl -i -d title=Hack http://localhost:8080/tasks
```

```text
HTTP/1.1 403 Forbidden
...
{"code":"forbidden","message":"csrf token mismatch"}
```

Lock the behaviour in with a test, `internal/app/app_test.go`:

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

## Step 7: Build a Single Binary and Deploy

```bash
kiw build
```

```text
INFO building binary kind=app entry=. output=.../todo-app/.krewire/build/bin/todo-app
created .../todo-app/.krewire/build/bin/todo-app
```

The binary embeds the views, the Krewire UI components, the migrations, and the static assets. The SQLite driver is pure Go, so you can cross-compile for a Linux server from any machine:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 kiw build
file .krewire/build/bin/todo-app
```

```text
.krewire/build/bin/todo-app: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, ...
```

Deploy by copying **two files** — the scaffolded `main.go` reads `krewire.yaml` from the working directory at start-up:

```bash
scp .krewire/build/bin/todo-app krewire.yaml user@server:/opt/todo-app/
ssh user@server 'cd /opt/todo-app && APP_ADDR=:8080 APP_DATABASE_DSN=/opt/todo-app/data.db ./todo-app'
```

On first start the binary creates the database file and applies the migrations. For a systemd unit, a reverse proxy with TLS, or a container image, continue with [Deployment & Operations](/guides/deployment-and-operations).

---

## The Companion Project

The complete, tested code of this tutorial lives in the krewire.com repository:

```bash
git clone --depth 1 https://github.com/krewire/krewire.com
cd krewire.com/examples/todo-app
go test ./...
go run .
```

Its `go.mod` pins the exact core version this tutorial was verified with.

## Recap

| You used | For |
| :--- | :--- |
| `kiw new --app`, `kiw make model`, `kiw make migration` | Scaffolding the app, the model, and the SQL migration |
| `pkg/orm` + `pkg/orm/sqlite` | SQLite access with a fluent, parameterized query API and the bundled pure-Go driver |
| `kiw migrate` + `pkg/orm/migrate` | Applying, inspecting, and rolling back migrations — from the terminal and on start-up |
| `.kiw` + Krewire UI | The page, with `Badge`, `Alert`, `Input`, `Button`, and `Card` |
| `web.PageSpec{DataFunc}` | Per-request data for the page |
| `pkg/validation` + post/redirect/get | Server-side validation with friendly errors |
| `web.CSRF()` | Token-based CSRF protection on every form |
| `kiw dev`, `kiw test`, `kiw build` | Develop, verify, and ship one binary |

## Next Steps

- Understand when data is resolved at build time and when per request: [Data Flow: SSG vs Runtime](/architecture/data-flow).
- Go deeper into the pieces: [App workload](/workloads/app), [`pkg/web`](/web-ui/web), [Database & ORM](/data-services/database-and-orm), [Security controls](/security-platforms/security).
- Learn every `.kiw` feature: [**2.7 The `.kiw` DSL →**](/getting-started/dsl)
