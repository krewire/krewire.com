# todo-app — Companion Code for "Build a Fullstack App in 15 Minutes"

The finished tutorial project from [Build a Fullstack App in 15 Minutes](https://krewire.com/docs/getting-started/fullstack-app-tutorial): a task manager built on a `kiw new todo-app --app` scaffold featuring:
- Embedded SQLite using `orm` with the pure-Go `modernc.org/sqlite` driver
- Single-File `.kiw` component views using Krewire UI components
- Server-side validation via `validation`
- CSRF-protected HTML forms with token verification
- Zero CGO requirement (`CGO_ENABLED=0`)

Verified with Go 1.27.1 and `github.com/krewire/krewire` pinned in `go.mod`.

---

## Running Locally

```bash
# Clone the repository
git clone --depth 1 https://github.com/krewire/krewire.com
cd krewire.com/examples/todo-app

# Run end-to-end tests
go test ./...

# Start the application server
go run .            # or: kiw dev
```

Open `http://localhost:8080` in your browser. Tasks are stored in `todo.db` in the working directory; set `TODO_DB` environment variable to use an alternative SQLite path.

---

## Build and Deploy

```bash
# Standard local binary build → .krewire/build/bin/todo-app
kiw build

# Pure Go cross-compilation for Linux servers
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 kiw build
```

Copy `.krewire/build/bin/todo-app` **and** `krewire.yaml` to the server (the scaffolded `main.go` reads `krewire.yaml` from the current working directory at startup) and execute:

```bash
APP_ADDR=:8080 ./todo-app
```

---

## Project Structure

```text
examples/todo-app/
├── migrations/                # SQL created by `kiw make migration create_tasks`, embedded
├── internal/
│   ├── db/db.go               # SQLite connection and automatic schema migrations
│   ├── models/task.go         # Task model struct and ORM metadata
│   ├── tasks/tasks.go         # Database queries, form handlers, validation, redirects
│   └── app/
│       ├── app.go             # Application wiring: database provider, CSRF middleware, routes
│       └── app_test.go        # End-to-end test of the complete task workflow
├── web/
│   ├── views/
│   │   ├── TaskBoard.kiw      # Single-file component view using Krewire UI
│   │   └── views.go           # Compiles .kiw files into web.App components
│   └── pages/pages.go         # Mounts TaskBoard at "/" with per-request data
├── krewire.yaml               # App configuration (kind: app, port: 8080, db)
├── main.go                    # Application entrypoint
└── go.mod                     # Go dependencies
```
