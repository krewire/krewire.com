---
title: "Application Container & Assembly"
description: "Typed dependency injection container, provider assembly, lifecycle management, and graceful process supervision in pkg/app."
date: "2026-10-31"
---

# Application Container & Assembly (`pkg/app`)

Import path: `github.com/krewire/krewire/pkg/app`

The `pkg/app` package provides Krewire's typed dependency container, provider assembly system, and graceful process supervision (`KWF-D1CNT`). It serves as the single composition root of an application—every dependency is declared with strict Go type safety, resolved in deterministic order, and observable via structured logging and tracer hooks.

---

## Core Features

- **Typed Dependency Container:** Generic-based `Provide[T]`, `Singleton[T]`, and `Resolve[T]` API eliminating reflection errors at runtime.
- **Provider Assembly:** Two-phase lifecycle (`Register` followed by `Boot`) managed by `NewApp`, locking container modifications before boot.
- **Test-Time Overrides:** `(*App).Test(overrides...)` enables rapid swapping of mock dependencies and test doubles without polluting production setup.
- **Lifecycle Supervision & Graceful Draining:** Built-in signal handling (`SIGINT`, `SIGTERM`) coordinating shutdown across concurrent HTTP listeners and background workers.
- **Observability & Tracing:** Built-in hooks for `slog.Logger`, resolution cycle detection (`CycleError`), and custom performance tracers.

---

## API Overview

### 1. Container Operations

```go
// Create container with options
c := app.New(app.WithLogger(logger), app.WithTrace())

// Register transient factory (fresh instance per resolution)
app.Provide[*UserService](c, func() *UserService {
    return &UserService{}
})

// Register singleton factory (computed once and cached)
app.Singleton[*sql.DB](c, func() (*sql.DB, error) {
    return sql.Open("pgx", os.Getenv("DATABASE_URL"))
})

// Resolve typed dependency
db, err := app.Resolve[*sql.DB](c)
```

### 2. Provider Assembly (`app.Provider`)

Providers structure modular subsystems across domains:

```go
type Provider interface {
    Register(c *app.Container) error
    Boot(c *app.Container) error
}
```

- **`Register` Phase:** Bind factories and singletons into the container.
- **`Boot` Phase:** Triggered once all providers have registered and the container is sealed; resolve cross-service wiring and mount routes.

```go
a := app.NewApp(
    &DatabaseProvider{},
    &UserServiceProvider{},
    &HttpServerProvider{},
)

container, err := a.Build()
if err != nil {
    log.Fatalf("failed to assemble application: %v", err)
}
```

---

## Complete Example

The following example demonstrates building a clean composition root with custom providers, resolving typed services, and executing with graceful lifecycle management:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/krewire/krewire/pkg/app"
)

// Service defines a business logic interface.
type GreeterService struct {
	Greeting string
}

// GreeterProvider binds the service into the app container.
type GreeterProvider struct{}

func (p *GreeterProvider) Register(c *app.Container) error {
	return app.Singleton[*GreeterService](c, func() *GreeterService {
		return &GreeterService{Greeting: "Hello from Krewire"}
	})
}

func (p *GreeterProvider) Boot(c *app.Container) error {
	svc, err := app.Resolve[*GreeterService](c)
	if err != nil {
		return err
	}
	fmt.Printf("[Boot] Greeter initialized: %s\n", svc.Greeting)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Assemble application providers
	application, err := app.Bootstrap(ctx, &GreeterProvider{})
	if err != nil {
		log.Fatalf("bootstrap error: %v", err)
	}
	defer application.Stop(ctx)

	// 2. Resolve typed dependency
	greeter, err := app.Resolve[*GreeterService](application.Container())
	if err != nil {
		log.Fatalf("resolve error: %v", err)
	}

	// 3. Define HTTP handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s at %s\n", greeter.Greeting, time.Now().Format(time.RFC3339))
	})

	// 4. Run server with graceful interrupt handling
	runner := app.HttpRunner(mux, ":8080")
	if err := application.Run(ctx, runner); err != nil {
		log.Fatalf("server terminated: %v", err)
	}
}
```

---

## Design Boundaries

`pkg/app` owns application composition, registration order, and runner supervision. Domain business rules belong in `pkg/model`, HTTP routing belongs in `pkg/web`, and runtime engine mechanics belong in `pkg/lifecycle`. `pkg/app` never imports `tools/kiw`.
