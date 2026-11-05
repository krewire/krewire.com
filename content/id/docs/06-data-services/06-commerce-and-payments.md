---
title: "Commerce & Gateway Pembayaran Indonesia"
description: "Panduan integrasi untuk Midtrans QRIS, Virtual Account (BCA, Mandiri, BRI, BNI), dan kartu kredit menggunakan pkg/commerce."
date: "2026-10-24"
---

# Commerce & Gateway Pembayaran Indonesia (`pkg/commerce`)

`pkg/commerce` memberdayakan aplikasi Krewire dan Krewire Cloud untuk menerima metode pembayaran lokal Indonesia secara bawaan melalui Midtrans.

## Jalur Pembayaran yang Didukung

- **Dynamic QRIS**: Pembuatan kode QR instan yang kompatibel dengan GoPay, OVO, DANA, ShopeePay, BCA Mobile, serta seluruh aplikasi perbankan nasional.
- **Bank Virtual Account (VA)**: Rekonsiliasi otomatis real-time untuk Bank Central Asia (BCA), Bank Mandiri, Bank Rakyat Indonesia (BRI), Bank Negara Indonesia (BNI), dan PermataBank.
- **Kartu Kredit**: Deteksi fraud berbasis 3D-Secure 2.0 dan penagihan kartu berulang (recurring billing).

## Contoh Checkout

```go
package main

import (
    "context"
    "fmt"
    "os"
    
    "github.com/krewire/krewire/pkg/commerce"
)

func main() {
    client := commerce.NewMidtransClient(commerce.Config{
        ServerKey:   os.Getenv("MIDTRANS_SERVER_KEY"),
        ClientKey:   os.Getenv("MIDTRANS_CLIENT_KEY"),
        Environment: commerce.EnvProduction, // atau EnvSandbox
    })

    // Menagih pelanggan via QRIS
    charge, err := client.ChargeQRIS(context.Background(), commerce.ChargeRequest{
        OrderID:       "ORDER-2026-9021",
        GrossAmount:   150000, // IDR 150.000
        CustomerEmail: "alice@example.id",
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Pindai QRIS untuk membayar: %s\n", charge.QRString)
}
```
