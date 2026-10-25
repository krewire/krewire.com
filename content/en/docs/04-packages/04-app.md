---
title: "Application Container & Monolith"
description: "Dependency injection container, application lifecycle management, and graceful process supervision in pkg/app."
date: "2026-10-25"
---

# Application Container & Monolith (`pkg/app`)

Import path: `github.com/krewire/krewire/pkg/app`

The `app` package manages application lifecycle, dependency injection, and graceful process supervision for web monoliths.

---

## Features

- **Container Assembly:** Register singleton and transient services with strict type safety.
- **Graceful Shutdown:** Intercepts OS signals (`SIGINT`, `SIGTERM`), draining active HTTP connections within a configurable timeout.
- **Health Probes:** Built-in hooks for database pings, readiness checks, and liveliness probes.

---

## Usage Example

```go
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/krewire/krewire/pkg/app"
)

func main() {
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
}
```\n