---
title: "Testing Utilities & Assertions"
description: "Testing assertion harness, HTTP helpers, UI component assertions, and mock fixtures in pkg/testing."
date: "2026-10-25"
---

# Testing Utilities & Assertions (`pkg/testing`)

Import path: `github.com/krewire/krewire/pkg/testing`

`pkg/testing` provides comprehensive testing utilities across all packages and applications, including HTTP helpers and UI assertions.

---

## Assertion Harness Example

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

This discipline ensures clean, readable, and robust test suites that verify behavior across software releases without external dependencies.\n