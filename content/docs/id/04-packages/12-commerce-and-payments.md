---
title: "Commerce & Indonesian Payment Gateway"
description: "Integration guide for Midtrans QRIS, Virtual Accounts (BCA, Mandiri, BRI, BNI), and credit cards using pkg/commerce."
date: "2026-10-24"
---

# Commerce & Indonesian Payment Gateway (`pkg/commerce`)

`pkg/commerce` empowers Krewire applications and Krewire Cloud to accept Indonesian local payment methods natively through Midtrans.

## Supported Payment Rails

- **Dynamic QRIS**: Instant QR code generation compatible with GoPay, OVO, DANA, ShopeePay, BCA Mobile, and all national banking apps.
- **Bank Virtual Accounts (VA)**: Real-time automated reconciliation for Bank Central Asia (BCA), Bank Mandiri, Bank Rakyat Indonesia (BRI), Bank Negara Indonesia (BNI), and PermataBank.
- **Credit Cards**: 3D-Secure 2.0 fraud detection and recurring card charging.

## Checkout Example

```go
package main

import (
    "context"
    
    "github.com/krewire/krewire/pkg/commerce"
)

func main() {
    client := commerce.NewMidtransClient(commerce.Config{
        ServerKey:   os.Getenv("MIDTRANS_SERVER_KEY"),
        ClientKey:   os.Getenv("MIDTRANS_CLIENT_KEY"),
        Environment: commerce.EnvProduction, // or EnvSandbox
    })

    // Charge customer via QRIS
    charge, err := client.ChargeQRIS(context.Background(), commerce.ChargeRequest{
        OrderID:     "ORDER-2026-9021",
        GrossAmount: 150000, // IDR 150,000
        CustomerEmail: "alice@example.id",
    })

    fmt.Printf("Scan QRIS to pay: %s\n", charge.QRString)
}
```
