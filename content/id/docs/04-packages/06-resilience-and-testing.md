---
title: "Resiliensi & Jaminan Mutu (resilience, testing)"
description: "Komunikasi jaringan toleran-kegagalan dengan kebijakan retry dan circuit breaker, dipadukan dengan harness assertion pengujian."
date: "2026-10-08"
---

# Resiliensi & Jaminan Mutu (`resilience`, `testing`)

Jalur impor (*import paths*):
- `github.com/krewire/krewire/pkg/resilience`
- `github.com/krewire/krewire/pkg/testing`

Sistem terdistribusi dan aplikasi web pasti menghadapi kegagalan di batas jaringan. Krewire menyediakan primitif resiliensi bawaan untuk mengisolasi kegagalan serta pustaka pengujian lengkap dengan utilitas assertion HTTP dan UI.

---

## 1. Pola Resiliensi (`pkg/resilience`)

### Kebijakan Retry dengan Jittered Backoff

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
		// Memanggil API pihak ketiga yang mungkin tidak stabil
		return callPaymentGateway()
	})
}
```

### Circuit Breakers

Mencegah pemadaman sistem beruntun (*cascading failure*) dengan memutus aliran permintaan saat layanan upstream gagal berulang kali:

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

## 2. Utilitas Pengujian (`pkg/testing`)

Krewire menyediakan utilitas pengujian di seluruh paket dan aplikasi, termasuk helper HTTP dan assertion:

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

Disiplin ini menjamin test suite yang bersih, mudah dibaca, dan tangguh dalam memverifikasi perilaku perangkat lunak di setiap rilis.
