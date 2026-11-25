# `todo-app` (`examples/todo-app`)

## Purpose & Workspace Role

`todo-app` is the companion full-stack Go tutorial application in `krewire.com/examples/todo-app/`. It demonstrates how a production-structured Krewire `app` workload combines `app.Container` dependency injection, pure-Go SQLite persistence via `krewire/orm` and `hub/sqlite`, embedded SQL migrations, server-side `.kiw` rendering, and CSRF-protected HTML form workflows into a single static binary.

---

## Architectural Position

- **Ecosystem Layer**: Layer 5 example consuming Layer 2 libraries (`github.com/krewire/krewire/app`, `config`, `orm`, `validation`, `web`, `web/ssg`, `github.com/krewire/hub/sqlite`, and `github.com/krewire/ui/core`, `kiw`).
- **Pure-Go Invariant**: Uses `github.com/krewire/hub/sqlite` (pure-Go SQLite driver) so the entire application compiles with `CGO_ENABLED=0`.
- **Package Structure**:
  - `main.go` — Thin entrypoint loading `krewire.yaml` via `config.LoadMetadata` and calling `app.New(meta).Run(meta.Addr())`.
  - `internal/config/` — Typed `Metadata` (`Project`, `Dev`) with `rconfig.Load`, `rconfig.Override`, and `validation.Struct`.
  - `internal/db/` — `Open(ctx, path)` opens SQLite (`MaxOpenConns: 1`) and runs `-- +up` sections from embedded `migrations.Files` inside `schema_migrations`.
  - `internal/models/` — `Task` struct (`ID`, `Title`, `Done`, `CreatedAt`) with `TableName() string`.
  - `internal/tasks/` — `Store` (`List`, `Create`, `Toggle`, `Delete`), `BoardData`, and POST handlers (`CreateHandler`, `ToggleHandler`, `DeleteHandler`) using Post/Redirect/Get and `validation.Struct`.
  - `internal/http/` — `Handler` and `Provider` mounting `/api/status` and `/api/contact`.
  - `internal/app/` — `New(meta)` wires `configProvider`, `dbProvider`, `webProvider` (CSRF middleware, static assets, theme, layout, `.kiw` views, routes), and `http.Provider` via `rvapp.NewApp(...).Build()`.
  - `web/` — `theme.New()`, `layouts.Shell()`, `components.Hero()`, `pages.Register`, and `views.Register` (parsing `ui.ComponentsFS` and embedded `*.kiw` templates).
  - `migrations/` & `assets/` — `embed.FS` packages for SQL migrations and `public/` static assets.

---

## Core Concepts & APIs

- **`app.New(meta *config.Metadata) (*rvweb.App, error)`**: Composes the dependency injection container (`rvapp.Container`) and returns the ready `*web.App`.
- **`tasks.Store`**: Wraps `*orm.DB` to query and mutate `models.Task` rows and serve `BoardData` (including `web.CSRFFrom(r.Context())`) to `TaskBoard.kiw`.
- **`views.Register(a *web.App) error`**: Parses Krewire UI's `.kiw` component catalog alongside `TaskBoard.kiw` and registers the combined CSS at `assets/style.css`.

---

## Realistic Usage Examples

### Running and Testing the Todo App

```bash
cd /home/krewire/workspace/krewire/krewire.com/examples/todo-app

# Run tests
CGO_ENABLED=0 go test -v ./...

# Run locally on port 8080 (creates todo.db in working directory)
CGO_ENABLED=0 go run .

# Override database path and listen address
TODO_DB=/tmp/my-todo.db APP_ADDR=:9090 CGO_ENABLED=0 go run .
```

---

## Verification

```bash
cd /home/krewire/workspace/krewire/krewire.com/examples/todo-app
gofmt -l .
go vet ./...
CGO_ENABLED=0 go test -v ./...
CGO_ENABLED=0 go build .
```
