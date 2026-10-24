---
title: "Microservice (service)"
description: "Cloud-native microservices with service discovery, API gateway routing, rate limiting, circuit breaking, and OpenTelemetry distributed tracing."
date: "2026-10-08"
---

# Microservice (`service`)

The **`service`** workload equips developers to build high-throughput, cloud-native microservices and distributed APIs.

Powered by `pkg/cloud/service`, `pkg/sec`, and `pkg/resilience`, the `service` workload includes production-grade discovery, distributed tracing, gateway routing, and telemetry out of the box.

---

## Architectural Characteristics

1. **Service Discovery:** Native drivers for Consul, etcd, Kubernetes DNS, and AWS Cloud Map.
2. **Telemetry & Observability:** Built-in OpenTelemetry (OTel) instrumentation for distributed tracing, metrics, and structured logging with W3C `traceparent` propagation.
3. **Traffic Resilience:** Token-bucket rate limiting, circuit breakers, timeout propagation, and concurrency bulkheads.
4. **Health & Liveness Check Probes:** Automatic `/healthz` and `/readyz` endpoints adhering to cloud container conventions.

---

## Scaffolding a `service` Project

```bash
kiw new auth-service --service
cd auth-service
```

Directory layout:

```text
auth-service/
├── krewire.yaml          # Project configuration
├── go.mod
├── cmd/
│   └── service/
│       └── main.go       # Service entry point
└── internal/
    ├── api/              # HTTP / gRPC endpoint handlers
    ├── discovery/        # Service registration logic
    └── domain/           # Core domain business logic
```

---

## Service Implementation

```go
package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/krewire/krewire/pkg/cloud/service"
	"github.com/krewire/krewire/pkg/web"
)

func main() {
	svc := service.New(service.Config{
		Name:    "auth-service",
		Version: "0.1.0",
		Port:    ":8080",
	})

	router := web.NewRouter()

	// Health and readiness endpoints
	router.Get("/healthz", service.HealthHandler())
	router.Get("/readyz", service.ReadinessHandler())

	// Protected business route
	router.Post("/v1/tokens/verify", func(w http.ResponseWriter, r *http.Request) {
		web.JSON(w, http.StatusOK, map[string]any{
			"valid":     true,
			"verified": time.Now().UTC(),
		})
	})

	// Start service with automated discovery registration
	if err := svc.Start(context.Background(), router); err != nil {
		os.Exit(1)
	}
}
```

---

## Configuration (`krewire.yaml`)

```yaml
project:
  name: "auth-service"
  kind: service
  version: "0.1.0"

service:
  name: "auth-service"
  port: 8080
```

`name` and `port` are the only `service:` keys. Service discovery
(`pkg/cloud/service/registry`), rate limiting and circuit breaking
(`pkg/cloud/service/resilience`), and tracing (`pkg/cloud/service/tracing`) are
composed in Go code, where the registry and transport are wired.

---

## Running and Verifying

Run the service locally:

```bash
kiw run
```

Inspect registered endpoints and service health status:

```bash
curl http://localhost:8080/healthz
```
