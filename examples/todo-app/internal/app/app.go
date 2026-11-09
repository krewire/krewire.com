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
		&dbProvider{path: config.Env("TODO_DB", "todo.db")},
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
