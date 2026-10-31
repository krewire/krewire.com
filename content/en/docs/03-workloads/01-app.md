---
title: "Fullstack Monolith (app)"
description: "Architecture and implementation guide for the app workload: Go net/http foundation, embedded assets, SSR templating, dependency injection, and production operations."
date: "2026-10-31"
---

# Fullstack Monolith (`app`)

The **`app`** workload is Krewire's foundation for building production web monoliths. It compiles the entire web surface—HTTP routing, server-rendered views (SSR), embedded static assets, typed dependency injection, configuration, and JSON API endpoints—into a **single, self-contained Go binary**.

Built on standard Go primitives (`net/http`, `embed.FS`, `context`, `slog`), the `app` workload requires no external web servers like Nginx or Node.js runtimes for local development or basic deployment.

---

## Architectural Characteristics

1. **Zero Runtime Dependencies:** Everything needed to run the application is compiled into a single static executable.
2. **Standard Library First:** The routing, middleware, and handler interfaces integrate directly with Go's `net/http` ecosystem.
3. **Embedded Static Assets:** CSS, JavaScript, and templates are compiled into the binary using Go's `embed` package, preventing missing-asset errors in production.
4. **Deterministic Dependency Assembly:** Leverages `pkg/app` for typed dependency injection, provider registration (`Register` and `Boot`), and observable lifecycles.
5. **Clean Domain Boundaries:** Built around modular monolith patterns (`KWF-5ZHQV`), separating configuration, HTTP handlers, SSR presentation, and domain logic.

---

## Scaffolding an `app` Project

Use the `kiw new` command (which equips the `app` variant by default or via `--app`), or equip an existing folder with `kiw init --app`:

```bash
kiw new storefront --app
cd storefront
```

The resulting directory tree follows Krewire's canonical modular monolith structure:

```text
storefront/
├── krewire.yaml          # Project configuration (kind: app)
├── go.mod                # Go module definition
├── main.go               # Thin application entry point
├── internal/
│   ├── app/
│   │   └── app.go        # Composition root: wires providers into *web.App
│   ├── config/
│   │   └── config.go     # Typed configuration loaded from krewire.yaml and env
│   └── http/
│       └── http.go       # HTTP JSON API handlers & route mounting
├── web/
│   ├── layouts/
│   │   └── shell.go      # HTML shell layout with theme toggles
│   ├── pages/
│   │   └── pages.go      # SSR page registrations (web.PageSpec)
│   ├── components/
│   │   └── components.go # SSR UI components (e.g. Hero)
│   └── theme/
│       └── theme.go      # Theme palettes (light / dark)
├── assets/
│   ├── embed.go          # Embedded static filesystem (embed.FS)
│   └── public/
│       ├── app.css       # Client styling
│       └── app.js        # Client script bundle
├── README.md             # Project documentation
├── .gitignore            # Git exclusion rules
├── .env.example          # Environment variable template
└── .env                  # Local runtime environment file
```

---

## Application Entry Point (`main.go`)

In Krewire's modular architecture, `main.go` remains an intentionally thin entry point that loads typed configuration, instantiates the application container, and initiates the HTTP runner:

```go
// Thin entry point: load config, build app, run.
package main

import (
	"fmt"
	"os"

	"storefront/internal/app"
	"storefront/internal/config"
)

// Signature is the ecosystem watermark embedded into the binary at compile time via -ldflags.
var Signature = "build with krewire"

func main() {
	meta, err := config.LoadMetadata("krewire.yaml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	a, err := app.New(meta)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := a.Run(meta.Addr()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

---

## Core Components & Architecture

### 1. Composition Root (`internal/app/app.go`)

The composition root wires all subsystem providers (`configProvider`, `webProvider`, `http.Provider`) through `pkg/app`:

```go
package app

import (
	"io/fs"

	rvapp "github.com/krewire/krewire/pkg/app"
	rvweb "github.com/krewire/krewire/pkg/web"
	"storefront/assets"
	"storefront/internal/config"
	"storefront/internal/http"
	"storefront/web/components"
	"storefront/web/layouts"
	"storefront/web/pages"
	"storefront/web/theme"
)

// New assembles the full application and returns the ready *web.App
func New(meta *config.Metadata) (*rvweb.App, error) {
	container, err := rvapp.NewApp(
		&configProvider{meta: meta},
		&webProvider{},
		&http.Provider{},
	).Build()
	if err != nil {
		return nil, err
	}
	return rvapp.Resolve[*rvweb.App](container)
}
```

### 2. HTTP Routing & API Handlers (`internal/http/http.go`)

HTTP API handlers bind input payloads, perform validation with `pkg/validation`, and mount onto the application router:

```go
package http

import (
	"net/http"

	rvapp "github.com/krewire/krewire/pkg/app"
	rvweb "github.com/krewire/krewire/pkg/web"
)

type Handler struct{}

func (h *Handler) Status(w http.ResponseWriter, _ *http.Request, _ rvweb.Params) {
	rvweb.JSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"version": "0.1.0",
	})
}

type Provider struct{}

func (p *Provider) Register(c *rvapp.Container) error {
	return rvapp.Singleton[*Handler](c, func() *Handler { return &Handler{} })
}

func (p *Provider) Boot(c *rvapp.Container) error {
	a, err := rvapp.Resolve[*rvweb.App](c)
	if err != nil {
		return err
	}
	h, err := rvapp.Resolve[*Handler](c)
	if err != nil {
		return err
	}
	a.Router().Get("/api/status", h.Status)
	return nil
}
```

### 3. Server-Side Rendering (`web/pages/pages.go`)

Pages are registered with declarative `web.PageSpec` configurations supporting static data or dynamic per-request resolvers:

```go
package pages

import (
	"net/http"
	"time"

	"github.com/krewire/krewire/pkg/web"
)

func Register(a *web.App) {
	a.Page(web.PageSpec{
		Path:   "/",
		Title:  "Storefront",
		Layout: "Shell",
		Root:   "Hero",
		Data: map[string]any{
			"badge":    "v0.1.0",
			"title":    "Welcome to Storefront",
			"subtitle": "SSR, JSON API, and embedded assets from one Go binary.",
		},
		EmitProps: true,
		Scripts:   []string{"/static/app.js"},
	})

	a.Page(web.PageSpec{
		Path:   "/greet",
		Title:  "Greet",
		Layout: "Shell",
		Root:   "Hero",
		DataFunc: func(r *http.Request) (any, error) {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "world"
			}
			return map[string]any{
				"badge":    time.Now().Format("15:04:05"),
				"title":    "Hello, " + name,
				"subtitle": "This page is rendered per request.",
			}, nil
		},
	})
}
```

---

## Configuration (`krewire.yaml`)

Configure server ports, project metadata, and optional client WebAssembly targets in `krewire.yaml`:

```yaml
project:
  name: "storefront"
  kind: app
  version: "0.1.0"

dev:
  port: 8080

build:
  output: ".krewire/build"
```

The runtime address can also be customized via the `APP_ADDR` environment variable or `--addr` CLI flag (e.g. `kiw run --addr :3000`).

---

## Development & Operations

### Local Development (`kiw dev` or `kiw run`)

Start the application with live-recompilation on file changes:

```bash
kiw dev
```

Or run the application directly:

```bash
kiw run
```

### Running Tests (`kiw test`)

Execute package and integration tests:

```bash
kiw test ./...
```

### Building for Production (`kiw build`)

Compile the production binary:

```bash
kiw build
```

The compiled binary will be placed at `.krewire/build/bin/storefront` (or cross-compiled for any platform using standard Go env variables like `GOOS=linux GOARCH=amd64`).
