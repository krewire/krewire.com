---
title: "Transactional Mailer"
description: "Unified transactional email dispatch across multiple providers: SMTP, Resend, SendGrid, Postmark, and in-memory test driver."
date: "2026-10-25"
---

# Transactional Mailer (`pkg/mail`)

`pkg/mail` dispatches transactional emails across different providers without rewriting business code. It decouples message composition from delivery mechanics.

## Supported Mail Drivers

| Driver | Best For | Authentication |
|---|---|---|
| `memory` | Unit and integration testing | In-memory message capture |
| `smtp` | Self-hosted Postfix/Exim or Google Workspace | SMTP TLS credentials |
| `resend` | Modern developer email API | API Key via HTTP |
| `sendgrid` | High-volume transactional delivery | Bearer API token |
| `postmark` | Strict deliverability & password resets | Server API Token |

## Usage Example

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
        Subject: "Welcome to Krewire",
        HTML:    "<h1>Welcome!</h1><p>Your deployment is ready.</p>",
    })
    if err != nil {
        panic(err)
    }
}
```
