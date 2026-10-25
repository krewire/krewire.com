---
title: "Mailer Transaksional"
description: "Pengiriman email transaksional terpadu ke berbagai penyedia: SMTP, Resend, SendGrid, Postmark, dan driver pengujian in-memory."
date: "2026-10-25"
---

# Mailer Transaksional (`pkg/mail`)

`pkg/mail` mengirimkan email transaksional ke berbagai penyedia tanpa perlu mengubah kode logika bisnis. Paket ini memisahkan penyusunan pesan dari mekanisme pengiriman teknis.

## Driver Email yang Didukung

| Driver | Sangat Cocok Untuk | Autentikasi |
|---|---|---|
| `memory` | Pengujian unit dan integrasi | Penangkapan pesan di dalam memori |
| `smtp` | Self-hosted Postfix/Exim atau Google Workspace | Kredensial SMTP TLS |
| `resend` | API email modern untuk pengembang | Kunci API via HTTP |
| `sendgrid` | Pengiriman transaksional volume tinggi | Token Bearer API |
| `postmark` | Keterkiriman ketat & reset kata sandi | Token API Server |

## Contoh Penggunaan

```go
package main

import (
    "context"
    "os"

    "github.com/krewire/krewire/pkg/mail"
)

func main() {
    ctx := context.Background()

    mailer := mail.New(mail.Config{
        Driver: "resend",
        APIKey: os.Getenv("RESEND_API_KEY"),
        From:   "notifications@krewire.com",
    })

    err := mailer.Send(ctx, mail.Message{
        To:      []string{"developer@example.com"},
        Subject: "Selamat Datang di Krewire",
        HTML:    "<h1>Selamat Datang!</h1><p>Deployment Anda telah siap digunakan.</p>",
    })
    if err != nil {
        panic(err)
    }
}
```
