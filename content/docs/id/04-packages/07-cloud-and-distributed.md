---
title: "Cloud & Distributed Systems (cloud, worker, service, infra)"
description: "Reference manual for Krewire Cloud libraries: async background worker queues, microservice registry and RPC, and declarative Infrastructure as Code."
date: "2026-10-10"
---

# Cloud & Distributed Systems (`cloud`)

Import path: `github.com/krewire/krewire/pkg/cloud`

The **`cloud`** suite equips Krewire applications with production-grade distributed systems primitives, background task scheduling, microservice routing, and declarative cloud infrastructure as code.

---

## 1. Background Worker & Job Queues (`cloud/worker`)

Import path: `github.com/krewire/krewire/pkg/cloud/worker`

The `worker` package provides asynchronous background job execution, scheduled cron jobs, configurable retry backoffs, and dead-letter queues (DLQ).

### Job Processing Example

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/krewire/krewire/pkg/cloud/worker"
)

type EmailPayload struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
}

func main() {
	// Initialize runner with concurrency pool and default retry policy
	runner := worker.NewRunner(worker.Config{
		Concurrency: 10,
		MaxRetries:  3,
		Backoff:     500 * time.Millisecond,
	})

	// Register job handler
	runner.Handle("send_email", func(ctx context.Context, job worker.Job) error {
		var payload EmailPayload
		if err := job.Decode(&payload); err != nil {
			return err
		}
		fmt.Printf("Sending email to %s: %s\n", payload.Recipient, payload.Subject)
		return nil
	})

	// Dispatch job
	jobID, err := runner.Dispatch(context.Background(), "send_email", EmailPayload{
		Recipient: "user@example.com",
		Subject:   "Welcome to Krewire!",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Enqueued job: %s\n", jobID)
}
```

---

## 2. Microservice Mesh & Gateways (`cloud/service`)

Import path: `github.com/krewire/krewire/pkg/cloud/service`

The `service` package provides cloud-native microservice routing, dynamic service discovery, configuration synchronization, and distributed tracing.

### Subpackages:
- **`service/registry`**: Service registration and health heartbeats with support for Consul, etcd, NATS, and DNS.
- **`service/messaging`**: Event broker abstraction with publisher/subscriber semantics.
- **`service/tracing`**: OpenTelemetry distributed context propagation across HTTP and gRPC boundaries.
- **`service/resilience`**: Concurrency bulkheads, rate limiters, and circuit breakers.

### Service Discovery Example

```go
package main

import (
	"context"

	"github.com/krewire/krewire/pkg/cloud/service/registry"
)

func main() {
	reg := registry.NewMemoryRegistry()

	// Register instance
	instance := registry.Instance{
		ID:      "auth-srv-01",
		Name:    "auth-service",
		Address: "10.0.1.15:8080",
		Tags:    []string{"v1", "auth"},
	}
	_ = reg.Register(context.Background(), instance)

	// Discover service nodes
	nodes, _ := reg.Lookup(context.Background(), "auth-service")
	for _, node := range nodes {
		println("Available node:", node.Address)
	}
}
```

---

## 3. Infrastructure as Code (`cloud/infra`)

Import path: `github.com/krewire/krewire/pkg/cloud/infra`

The `infra` package enables developers to define cloud infrastructure declaratively in Go. It enforces deterministic dependency resolution and generates execution plans with zero side-effects prior to applying changes.

### Declarative Schema & State

```go
package main

import (
	"github.com/krewire/krewire/pkg/cloud/infra/schema"
	"github.com/krewire/krewire/pkg/cloud/infra/state"
)

func main() {
	// Declare cloud resources
	vpc := schema.Resource{
		Type: "aws_vpc",
		Name: "production-vpc",
		Properties: map[string]any{
			"cidr_block": "10.0.0.0/16",
			"tags": map[string]string{
				"Environment": "production",
			},
		},
	}

	// Verify plan before applying
	plan := schema.NewPlan([]schema.Resource{vpc})
	if plan.HasChanges() {
		println("Plan requires infrastructure changes:")
		println(plan.Summary())
	}
}
```
