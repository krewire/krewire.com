---
title: "Pola Resiliensi & Circuit Breaker"
description: "Komunikasi jaringan toleran kesalahan dengan retry backoff berjitter, propagasi batas waktu, dan circuit breaker di pkg/resilience."
date: "2026-10-25"
---

# Pola Resiliensi & Circuit Breaker (`pkg/resilience`)

Jalur impor: `github.com/krewire/krewire/pkg/resilience`

Sistem terdistribusi dan aplikasi web kerap menghadapi kegagalan jaringan. `pkg/resilience` menyediakan primitif keandalan bawaan untuk mengisolasi kegagalan sementara dan melindungi sistem downstream.

---

## Kebijakan Retry dengan Jittered Backoff

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
		// Memanggil API pihak ketiga
		return callPaymentGateway()
	})
}
```

---

## Circuit Breaker

Mencegah kegagalan bertingkat (*cascading failure*) dengan membuka sirkuit saat dependensi upstream gagal berulang kali:

```go
cb := resilience.NewCircuitBreaker(resilience.CircuitBreakerConfig{
    FailureThreshold: 5,
    ResetTimeout:     10 * time.Second,
})

err := cb.Execute(func() error {
    return queryMicroservice()
})
```\n