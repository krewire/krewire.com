---
title: "Utilitas Pengujian & Asersi"
description: "Harness asersi pengujian, pembantu HTTP, asersi komponen UI, dan fixture mock di pkg/testing."
date: "2026-10-25"
---

# Utilitas Pengujian & Asersi (`pkg/testing`)

Jalur impor: `github.com/krewire/krewire/pkg/testing`

`pkg/testing` menyediakan utilitas pengujian komprehensif di seluruh paket dan aplikasi, termasuk helper HTTP dan asersi UI.

---

## Contoh Penggunaan Asersi

```go
package payment_test

import (
	"testing"

	"github.com/krewire/krewire/pkg/testing"
)

func TestPayment_ChargesCardWithIdempotency(t *testing.T) {
	charge, err := ProcessCharge("ch_123", 5000)

	test.NoError(t, err)
	test.Equal(t, "succeeded", charge.Status)
}
```

Disiplin pengujian ini memastikan rangkaian pengujian bersih, mudah dibaca, dan andal untuk memverifikasi perilaku perangkat lunak tanpa dependensi eksternal.\n