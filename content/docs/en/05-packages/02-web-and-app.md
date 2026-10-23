---
title: "Web & Application (web, app)"
description: "Core HTTP routing, middleware pipeline, RFC 7807 problem details, dependency injection container, and server lifecycle supervision."
date: "2026-10-08"
---

# Web & Application (`web`, `app`)

Import paths:
- `github.com/krewire/krewire/packages/web`
- `github.com/krewire/krewire/packages/app`

The **`web`** and **`app`** packages provide the high-performance HTTP web engine and dependency injection container powering the `app`, `service`, and `site` workloads.

---

## 1. Web Engine (`packages/web`)

The `web` package wraps Go's `net/http` with expressive routing, parameter extraction, middleware chains, and JSON serialization.

### Features
- **Zero Allocations Routing:** High-performance radix tree router supporting route groups, path parameters (`/users/:id`), and wildcards (`/static/*`).
- **Standard Library Compatibility:** Middleware and handlers use standard `http.Handler` and `http.HandlerFunc`.
- **RFC 7807 Problem Details:** Standardized JSON error responses for APIs.
- **Render Helpers:** First-class helpers for `web.JSON()`, `web.HTML()`, and `web.Redirect()`.

### Example

```go
package main

import (
	"net/http"
	"github.com/krewire/krewire/packages/web"
)

func main() {
	r := web.NewRouter()

	// Global Middleware
	r.Use(web.Recovery())
	r.Use(web.Logger())

	// Route Groups
	api := r.Group("/api/v1")
	{
		api.Get("/items", func(w http.ResponseWriter, r *http.Request) {
			items := []string{"book", "keyboard", "display"}
			web.JSON(w, http.StatusOK, items)
		})

		api.Get("/items/:id", func(w http.ResponseWriter, r *http.Request) {
			id := web.Param(r, "id")
			web.JSON(w, http.StatusOK, map[string]string{"id": id})
		})
	}

	http.ListenAndServe(":8080", r)
}
```

---

## 2. Application Container (`packages/app`)

The `app` package manages application lifecycle, dependency injection, and graceful process supervision.

### Features
- **Container Assembly:** Register singleton and transient services with type safety.
- **Graceful Shutdown:** Intercepts OS signals (`SIGINT`, `SIGTERM`), draining active HTTP connections with a configurable timeout.
- **Health Probes:** Built-in hooks for database pings and readiness verification.

### Example

```go
application := app.New(app.Config{
    Name:    "storefront",
    Version: "0.1.0",
})

// Register shared database connection
application.Provide(func() (*sql.DB, error) {
    return sql.Open("postgres", os.Getenv("DATABASE_URL"))
})

// Run server with signal draining
if err := application.RunServer(":8080", router); err != nil {
    log.Fatal(err)
}
```
