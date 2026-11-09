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
