---
title: "Cache Multi-Driver"
description: "Caching performa tinggi in-memory dan Redis terdistribusi dengan TTL, pembersihan kedaluwarsa berkala, dan connection pooling."
date: "2026-10-25"
---

# Cache Multi-Driver (`pkg/cache`, `pkg/redis`)

`pkg/cache` menghadirkan abstraksi caching key-value terpadu dengan dukungan TTL, pembersihan kedaluwarsa berkala, dan invalidasi terdistribusi di seluruh ekosistem Krewire.

## Driver Cache yang Didukung

- **`memory`**: Cache in-process thread-safe berkecepatan tinggi menggunakan mutex Go. Sangat cocok untuk aplikasi instans tunggal, CLI tool, dan unit test.
- **`redis`**: Engine caching terdistribusi menggunakan protokol asli RESP (`pkg/redis`) dengan connection pooling, pipelining perintah, dan auto-reconnect.

## Contoh Penggunaan

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/krewire/krewire/pkg/cache"
)

func main() {
    ctx := context.Background()

    // Inisialisasi store cache Redis
    cacheStore, err := cache.New(cache.Config{
        Driver: "redis",
        Host:   "127.0.0.1:6379",
    })
    if err != nil {
        panic(err)
    }

    // Menyimpan entri dengan TTL 10 menit
    _ = cacheStore.Set(ctx, "session:usr_123", "active", 10*time.Minute)

    // Mengambil entri dari cache
    val, err := cacheStore.Get(ctx, "session:usr_123")
    if err == nil {
        fmt.Println("Nilai cache:", val)
    }
}
```
