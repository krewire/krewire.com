---
title: "Resilience Patterns & Circuit Breakers"
description: "Fault-tolerant network communication with jittered retry policies, timeout propagation, and circuit breakers in pkg/resilience."
date: "2026-10-25"
---

# Resilience Patterns & Circuit Breakers (`pkg/resilience`)

Import path: `github.com/krewire/krewire/pkg/resilience`

Distributed systems and web applications fail over network boundaries. `pkg/resilience` provides built-in resilience primitives to isolate transient failures and protect downstream systems.

---

## Retry Policies with Jittered Backoff

```go
package main

import (
	"context"
	"time"
	"github.com/krewire/krewire/pkg/resilience"
)

func main() {
	policy := resilience.RetryPolicy{
		MaxAttempts:    3,
		InitialBackoff: 200 * time.Millisecond,
		MaxBackoff:     2 * time.Second,
		Jitter:         true,
	}

	err := policy.Do(context.Background(), func(ctx context.Context) error {
		// Call third-party network API
		return callPaymentGateway()
	})
}
```

---

## Circuit Breakers

Prevent cascading service outages by tripping open when an upstream dependency fails repeatedly:

```go
cb := resilience.NewCircuitBreaker(resilience.CircuitBreakerConfig{
    FailureThreshold: 5,
    ResetTimeout:     10 * time.Second,
})

err := cb.Execute(func() error {
    return queryMicroservice()
})
```\n