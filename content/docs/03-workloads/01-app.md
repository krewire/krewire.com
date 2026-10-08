---
title: "Fullstack Monolith (app)"
description: "Architecture and implementation guide for the app workload: Go net/http foundation, embedded assets, SSR templating, dependency injection, and production operations."
date: "2026-10-08"
---

# Fullstack Monolith (`app`)

The **`app`** workload is Krewire's foundation for building production web monoliths. It compiles the entire web surface—HTTP routing, server-rendered views, static assets, database access, background tasks, and API endpoints—into a **single, self-contained Go binary**.

Built on standard Go primitives (`net/http`, `embed.FS`, `context`, `slog`), the `app` workload requires no external web servers like Nginx or Node.js runtimes for local development or basic deployment.

---

## Architectural Characteristics

1. **Zero Runtime Dependencies:** Everything needed to run the application is compiled into a single static executable.
2. **Standard Library First:** The routing and middleware system implements Go's standard `http.Handler` and `http.HandlerFunc` interfaces.
3. **Embedded Static Assets:** CSS, JavaScript, icons, and templates are compiled into the binary using Go's `embed` package, preventing missing-asset errors in production.
4. **Clean Domain Boundaries:** Built around modular monolith patterns (`KWF-5ZHQV`), allowing future extraction into microservices or background workers without code rewrites.

---

## Scaffolding an `app` Project

Use the `kiw new` command with the `--app` flag:

```bash
kiw new storefront --app
cd storefront
```

The resulting directory tree follows the canonical Go layout:

```text
storefront/
├── krewire.yaml          # Project configuration
├── go.mod                # Go module definition
├── cmd/
│   └── server/
│       └── main.go       # Application entry point
├── internal/
│   ├── handlers/         # HTTP request handlers
│   ├── models/           # Domain models & entities
│   ├── services/         # Business logic & operations
│   └── storage/          # Database migrations & repositories
├── views/                # Server-rendered HTML templates
│   ├── layouts/
│   └── pages/
└── public/               # Static assets (CSS, JS, images)
```

---

## Application Entry Point (`main.go`)

A minimal, production-grade `app` workload uses `packages/app` for lifecycle supervision and `packages/web` for HTTP routing:

```go
package main

import (
	"context"
	"embed"
	"net/http"
	"os"

	"github.com/krewire/krewire/packages/app"
	"github.com/krewire/krewire/packages/web"
	"github.com/krewire/krewire/packages/sec"
)

//go:embed public/* views/*
var embeddedAssets embed.FS

func main() {
	// Initialize the Krewire application container
	application := app.New(app.Config{
		Name:    "storefront",
		Version: "0.1.0",
	})

	// Create the web engine
	router := web.NewRouter()

	// Apply default security headers and middleware
	router.Use(sec.SecurityHeaders())
	router.Use(web.Recovery())
	router.Use(web.Logger())

	// Register routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Welcome to Storefront</h1>"))
	})

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		web.JSON(w, http.StatusOK, map[string]string{
			"status": "healthy",
			"version": "0.1.0",
		})
	})

	// Mount embedded static assets
	router.ServeFiles("/static/*", http.FS(embeddedAssets))

	// Run with graceful shutdown
	if err := application.RunServer(":8080", router); err != nil {
		os.Exit(1)
	}
}
```

---

## Configuration (`krewire.yaml`)

Configure server port bindings, build outputs, and runtime flags in `krewire.yaml`:

```yaml
project:
  name: "storefront"
  kind: app
  version: "0.1.0"

dev:
  port: "8080"
  watch:
    - "cmd"
    - "internal"
    - "views"
    - "public"

build:
  output: "bin/server"
  tags: ["netgo"]
```

---

## Development & Operations

### Local Development (`kiw dev`)

Start the application in development mode with automatic file watching and recompilation:

```bash
kiw dev
```

Whenever `.go`, `.html`, or asset files change, `kiw` automatically recompiles and restarts the server process within milliseconds.

### Building for Production (`kiw build`)

Compile a production binary:

```bash
kiw build
```

To cross-compile for Linux servers:

```bash
GOOS=linux GOARCH=amd64 kiw build
```

The resulting binary in `bin/server` is completely standalone and ready for deployment on any Linux VPS or container environment.
