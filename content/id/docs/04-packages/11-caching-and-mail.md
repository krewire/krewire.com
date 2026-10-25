---
title: "Sistem Cache & Email"
description: "Caching performa tinggi in-memory dan Redis serta pengiriman email transaksional dengan multi-driver."
date: "2026-10-24"
---

# Sistem Cache & Email (`pkg/cache`, `pkg/mail`)

## 1. Multi-Driver Cache (`pkg/cache`, `pkg/redis`)

`pkg/cache` menghadirkan abstraksi caching key-value terpadu dengan dukungan TTL, pembersihan kedaluwarsa berkala, dan invalidasi terdistribusi:

- **`memory`**: Cache in-process thread-safe berkecepatan tinggi menggunakan mutex Go.
- **`redis`**: Engine caching terdistribusi menggunakan protokol asli RESP (`pkg/redis`) dengan connection pooling dan pipelining perintah.

```go
cacheStore, _ := cache.New(cache.Config{
    Driver: "redis",
    Host:   "127.0.0.1:6379",
})

// Menyimpan entri dengan TTL 10 menit
_ = cacheStore.Set(ctx, "session:usr_123", userData, 10*time.Minute)

// Mengambil entri dari cache
val, err := cacheStore.Get(ctx, "session:usr_123")
```

## 2. Mailer Transaksional Multi-Driver (`pkg/mail`)

`pkg/mail` mengirimkan email transaksional ke berbagai penyedia tanpa perlu mengubah kode logika bisnis:

| Driver | Sangat Cocok Untuk | Autentikasi |
|---|---|---|
| `memory` | Pengujian unit dan integrasi | Penangkapan pesan di dalam memori |
| `smtp` | Self-hosted Postfix/Exim atau Google Workspace | Kredensial SMTP TLS |
| `resend` | API email modern untuk pengembang | Kunci API via HTTP |
| `sendgrid` | Pengiriman transaksional volume tinggi | Token Bearer API |
| `postmark` | Keterkiriman ketat & reset kata sandi | Token API Server |

```go
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
```
