# todo-app — companion code for "Build a Fullstack App in 15 Minutes"

The finished project from the tutorial at
<https://krewire.com/docs/getting-started/fullstack-app-tutorial>: a task
manager built on a `kiw new todo-app --app` scaffold with SQLite (`pkg/orm` +
the pure-Go `modernc.org/sqlite` driver), a `.kiw` view using Krewire UI
components, server-side validation, and CSRF-protected HTML forms.

Verified with Go 1.27.1 and `github.com/krewire/krewire` at commit `33fe553`
(the version pinned in `go.mod`).

## Run it

	git clone --depth 1 https://github.com/krewire/krewire.com
	cd krewire.com/examples/todo-app
	go test ./...
	go run .            # or: kiw dev

Open <http://localhost:8080>. Tasks are stored in `todo.db` in the working
directory; set `TODO_DB` to use another file.

## Build and deploy

	kiw build                                   # → .krewire/build/bin/todo-app
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 kiw build

Copy `.krewire/build/bin/todo-app` **and** `krewire.yaml` to the server (the
scaffolded `main.go` reads `krewire.yaml` from the working directory at start
up) and run `APP_ADDR=:8080 ./todo-app`.

## Where things are

	migrations/                SQL created by `kiw make migration create_tasks`, embedded
	internal/db/db.go          opens SQLite and applies the "-- +up" sections once
	internal/models/task.go    Task model (`kiw make model Task`)
	internal/tasks/tasks.go    queries + form handlers (validation, redirects)
	web/views/TaskBoard.kiw    the page, written in .kiw with Krewire UI components
	web/views/views.go         turns the .kiw files into components for web.App
	web/pages/pages.go         mounts TaskBoard at "/" with per-request data
	internal/app/app.go        wiring: database provider, CSRF middleware, routes
	internal/app/app_test.go   end-to-end test of the form flow
