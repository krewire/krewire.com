---
title: "Caching & Mail Systems"
description: "High-performance in-memory and Redis caching alongside multi-driver transactional email dispatch."
date: "2026-10-24"
---

# Caching & Mail Systems (`pkg/cache`, `pkg/mail`)

## 1. Multi-Driver Cache (`pkg/cache`, `pkg/redis`)

`pkg/cache` delivers a unified key-value caching abstraction with support for TTL, expiration sweeps, and distributed invalidation:

- **`memory`**: High-speed, in-process thread-safe cache using Go mutexes.
- **`redis`**: Distributed caching engine using the native RESP protocol (`pkg/redis`) with connection pooling and pipeline commands.

```go
cacheStore, _ := cache.New(cache.Config{
    Driver: "redis",
    Host:   "127.0.0.1:6379",
})

// Set with 10-minute TTL
_ = cacheStore.Set(ctx, "session:usr_123", userData, 10*time.Minute)

// Get cached entry
val, err := cacheStore.Get(ctx, "session:usr_123")
```

## 2. Multi-Driver Transactional Mailer (`pkg/mail`)

`pkg/mail` dispatches transactional emails across different providers without rewriting business code:

| Driver | Best For | Authentication |
|---|---|---|
| `memory` | Unit and integration testing | In-memory message capture |
| `smtp` | Self-hosted Postfix/Exim or Google Workspace | SMTP TLS credentials |
| `resend` | Modern developer email API | API Key via HTTP |
| `sendgrid` | High-volume transactional delivery | Bearer API token |
| `postmark` | Strict deliverability & password resets | Server API Token |

```go
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
```
