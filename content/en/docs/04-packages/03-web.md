---
title: "Web Engine & HTTP Routing"
description: "High-performance radix tree HTTP routing, middleware pipeline, RFC 7807 problem details, parameter extraction, and render helpers in pkg/web."
date: "2026-10-25"
---

# Web Engine & HTTP Routing (`pkg/web`)

Import path: `github.com/krewire/krewire/pkg/web`

The `web` package wraps Go's `net/http` with expressive routing, parameter extraction, middleware chains, and JSON serialization without external runtime dependencies.

---

## Features

- **Zero Allocations Routing:** High-performance radix tree router supporting route groups, path parameters (`/users/:id`), and wildcards (`/static/*`).
- **Standard Library Compatibility:** Middleware and handlers use standard Go `http.Handler` and `http.HandlerFunc`.
- **RFC 7807 Problem Details:** Standardized JSON error responses for APIs.
- **Render Helpers:** First-class helpers for `web.JSON()`, `web.HTML()`, and `web.Redirect()`.

---

## Usage Example

```go
package main

import (
	"net/http"
	"github.com/krewire/krewire/pkg/web"
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
```\n