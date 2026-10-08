---
title: "Resilience & Testing (resilience, testing)"
description: "Fault-tolerant network communication with retry policies and circuit breakers, paired with the spec-driven testing harness."
date: "2026-10-08"
---

# Resilience & Testing (`resilience`, `testing`)

Import paths:
- `github.com/krewire/krewire/packages/resilience`
- `github.com/krewire/krewire/packages/testing`

Distributed systems and web applications fail over network boundaries. Krewire provides built-in resilience primitives to isolate failures and a testing library that links test suites directly to formal specifications.

---

## 1. Resilience Patterns (`packages/resilience`)

### Retry Policies with Jittered Backoff

```go
package main

import (
	"context"
	"time"
	"github.com/krewire/krewire/packages/resilience"
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

## 2. Spec-Driven Testing (`packages/testing`)

Krewire mandates **Spec-Driven Testing** across all repositories. Every non-trivial capability is governed by a specification with numbered requirement IDs, cited directly in test cases:

```go
package payment_test

import (
	"testing"
	"github.com/krewire/krewire/packages/testing/assert"
)

// Tests for KWL-PAY-X92KA
// Spec: KWL-PAY-X92KA REQ-001 Scope: PAYMENT
func TestPayment_ChargesCardWithIdempotency(t *testing.T) {
	charge, err := ProcessCharge("ch_123", 5000)

	assert.NoError(t, err)
	assert.Equal(t, "succeeded", charge.Status)
}
```

This discipline ensures that test coverage maps 1:1 to written requirements and prevents specification drift across software releases.
