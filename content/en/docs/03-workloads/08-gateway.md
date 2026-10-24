---
title: "API Gateway (gateway)"
description: "Front one or more upstream services with a declarative route table, per-route middleware, rate limiting, timeouts, and service-discovery load balancing."
date: "2026-10-25"
---

# API Gateway (`gateway`)

The **`gateway`** workload is a standalone reverse proxy that fronts one or more
upstream services. It owns a declarative route table, applies middleware (CORS,
tracing, timeouts, rate limiting, circuit breaking), and resolves upstream
addresses through the service registry.

Powered by `pkg/web/gateway`, the same route table runs from the project binary
or straight from the CLI with `kiw gateway`.

---

## Architectural Characteristics

1. **Declarative Route Table:** exact and prefix routes, per-route method
   matching, and either a fixed `target_url` or a registry-resolved `service`.
2. **Per-Route Middleware:** CORS, tracing, timeout, rate limiting, and circuit
   breaking are registered by name and composed per route instead of globally.
3. **Service Discovery:** round-robin load balancing over registered endpoints;
   an unavailable upstream returns `502` with a structured Problem JSON body.
4. **One Route Table, Two Runners:** `kiw dev` / `kiw run` execute the project
   binary; `kiw gateway` runs the identical `gateway:` block without a rebuild.

---

## Scaffolding a `gateway` Project

```bash
kiw new my-gateway --gateway
cd my-gateway
```

Directory layout:

```text
my-gateway/
├── krewire.yaml          # Project configuration and route table
├── go.mod
├── main.go               # Loads krewire.yaml and serves the gateway
├── internal/
│   └── config/           # Typed loader for the gateway: block
└── README.md
```

---

## Route Table Implementation (`main.go`)

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krewire/krewire/pkg/cloud/service/registry"
	"github.com/krewire/krewire/pkg/web/gateway"
	"my-gateway/internal/config"
)

func main() {
	cfg, err := config.Load("krewire.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "gateway: load config: %v\n", err)
		os.Exit(1)
	}

	routes, err := cfg.Routes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gateway: %v\n", err)
		os.Exit(1)
	}

	gw := gateway.New(
		gateway.WithRegistry(registry.NewMemoryRegistry()),
		gateway.WithMiddleware("cors", gateway.CORSMiddleware("*")),
		gateway.WithMiddleware("tracing", gateway.TracingMiddleware()),
	)
	for _, route := range routes {
		gw.AddRoute(route)
	}

	server := &http.Server{Addr: cfg.Addr(), Handler: gw}
	slog.Info("gateway listening", "addr", server.Addr, "routes", len(routes))

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintf(os.Stderr, "gateway error: %v\n", err)
			os.Exit(1)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	slog.Info("shutting down gateway...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
}
```

---

## Configuration (`krewire.yaml`)

```yaml
project:
  name: "my-gateway"
  kind: gateway
  version: "0.1.0"

gateway:
  host: ""
  port: 8080
  routes:
    - path: /api
      prefix: true
      strip_prefix: true
      target_url: http://localhost:8081
      timeout: 10s
    - path: /users
      service: user-service
      middleware: [cors, tracing]
      rate_limit: 100
```

Route fields map to `gateway.Route`: `path`, `method`, `service`, `target_url`,
`prefix`, `strip_prefix`, `middleware`, `rate_limit`, and `timeout`.

`rate_limit` (requests per second) and `timeout` are enforced by the gateway
itself — a route is rate limited and bounded even with no middleware attached.
A request that exceeds its budget gets `429` with `Retry-After`; a slow upstream
gets `504 Gateway Timeout`. `GATEWAY_ADDR` overrides `gateway.host` +
`gateway.port`.

---

## Middleware

Middleware is registered on the gateway under a name and selected per route
through the `middleware:` list. Names are yours to choose; the package provides
the constructors:

- `gateway.CORSMiddleware(allowedOrigins ...string)`
- `gateway.TracingMiddleware()` — propagates `X-Request-ID`
- `gateway.TimeoutMiddleware(d)` — sets a request deadline
- `gateway.RateLimitMiddleware(rps)` — spaces requests 1/rps apart
- `gateway.CircuitBreakerMiddleware(breaker)` — fail fast on a degraded upstream

A route naming a middleware the gateway has not registered skips it, so the
wiring stays explicit in `main.go`.

---

## Running the Gateway

```bash
# Development: rebuild and restart on change
kiw dev

# Production binary in .krewire/build/bin/<name>
kiw build

# Run the shared route table without building the project binary
kiw gateway --port 8080 --routes "/api=user-service,/pay=payment-service"
```
