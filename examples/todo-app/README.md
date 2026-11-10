# todo-app — companion code for "Build a Fullstack App in 15 Minutes"

The finished project from the tutorial at
<https://krewire.com/docs/getting-started/fullstack-app-tutorial>: a task
manager built on a `kiw new todo-app --app` scaffold with SQLite (`pkg/orm` +
the bundled pure-Go driver `pkg/orm/sqlite`), migrations applied with
`pkg/orm/migrate` and `kiw migrate`, a `.kiw` view using Krewire UI
components, server-side validation, and CSRF-protected HTML forms.

Verified with Go 1.27.1 and `github.com/krewire/krewire` at commit `d2f4cb9`
(the version pinned in `go.mod`). That commit is on core `main` but is
**not in a tagged release yet**: `pkg/orm/sqlite`, `pkg/orm/migrate`, and
`kiw migrate` ship with the next core release after v0.2.0. Until
then, build `kiw` from the same commit:

	go install github.com/krewire/krewire/tools/kiw/cmd/kiw@d2f4cb9

## Run it

	git clone --depth 1 https://github.com/krewire/krewire.com
	cd krewire.com/examples/todo-app
	go test ./...
	go run .            # or: kiw dev

Open <http://localhost:8080>. Tasks are stored in `todo.db` in the working
directory (`database.dsn` in `krewire.yaml`); set `APP_DATABASE_DSN` to use
another file. The app applies pending migrations when it starts; inspect or
roll them back with `kiw migrate status`, `kiw migrate down`, and
`kiw migrate redo`.

## Build and deploy

	kiw build                                   # → .krewire/build/bin/todo-app
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 kiw build

Copy `.krewire/build/bin/todo-app` **and** `krewire.yaml` to the server (the
scaffolded `main.go` reads `krewire.yaml` from the working directory at start
up) and run `APP_ADDR=:8080 APP_DATABASE_DSN=/var/lib/todo-app/todo.db ./todo-app`.

## Where things are

	migrations/                SQL created by `kiw make migration create_tasks`, embedded
	internal/db/db.go          opens SQLite (pkg/orm/sqlite) and runs migrate.Up
	internal/config/config.go  the krewire.yaml mirror, plus its database section
	internal/models/task.go    Task model (`kiw make model Task`)
	internal/tasks/tasks.go    queries + form handlers (validation, redirects)
	web/views/TaskBoard.kiw    the page, written in .kiw with Krewire UI components
	web/views/views.go         turns the .kiw files into components for web.App
	web/pages/pages.go         mounts TaskBoard at "/" with per-request data
	internal/app/app.go        wiring: database provider, CSRF middleware, routes
	internal/app/app_test.go   end-to-end test of the form flow
