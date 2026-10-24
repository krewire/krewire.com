---
title: "Resilience & Testing (resilience, testing)"
description: "Fault-tolerant network communication with retry policies and circuit breakers, paired with the testing assertion harness."
date: "2026-10-08"
---

# Resilience & Testing (`resilience`, `testing`)

Import paths:
- `github.com/krewire/krewire/pkg/resilience`
- `github.com/krewire/krewire/pkg/testing`

Distributed systems and web applications fail over network boundaries. Krewire provides built-in resilience primitives to isolate failures and a testing library that provides HTTP and UI assertion utilities.

---

## 1. Resilience Patterns (`pkg/resilience`)

### Retry Policies with Jittered Backoff

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
		// Call unreliable third-party payment API
		return callPaymentGateway()
	})
}
```

### Circuit Breakers

Prevent cascading service outages by tripping open when an upstream dependency fails repeatedly:

```go
cb := resilience.NewCircuitBreaker(resilience.CircuitBreakerConfig{
    FailureThreshold: 5,
    ResetTimeout:     10 * time.Second,
})

err := cb.Execute(func() error {
    return queryMicroservice()
})
```

---

## 2. Testing Utilities (`pkg/testing`)

Krewire provides testing utilities across all packages and applications, including HTTP helpers and UI assertions:

```go
package payment_test

import (
	"testing"
	"github.com/krewire/krewire/pkg/testing/assert"
)

func TestPayment_ChargesCardWithIdempotency(t *testing.T) {
	charge, err := ProcessCharge("ch_123", 5000)

	assert.NoError(t, err)
	assert.Equal(t, "succeeded", charge.Status)
}
```

This discipline ensures clean, readable, and robust test suites that verify behavior across software releases.
