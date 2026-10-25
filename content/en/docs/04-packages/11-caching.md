---
title: "Multi-Driver Cache"
description: "High-performance in-memory and Redis distributed caching with TTL, expiration sweeps, and connection pooling."
date: "2026-10-25"
---

# Multi-Driver Cache (`pkg/cache`, `pkg/redis`)

`pkg/cache` delivers a unified key-value caching abstraction with support for TTL, expiration sweeps, and distributed invalidation across the Krewire ecosystem.

## Supported Cache Drivers

- **`memory`**: High-speed, in-process thread-safe cache using Go mutexes. Ideal for single-instance applications, CLI tools, and unit testing.
- **`redis`**: Distributed caching engine using the native RESP protocol (`pkg/redis`) with connection pooling, pipelining, and automatic reconnection.

## Usage Example

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

    // Initialize Redis cache store
    cacheStore, err := cache.New(cache.Config{
        Driver: "redis",
        Host:   "127.0.0.1:6379",
    })
    if err != nil {
        panic(err)
    }

    // Set with 10-minute TTL
    _ = cacheStore.Set(ctx, "session:usr_123", "active", 10*time.Minute)

    // Get cached entry
    val, err := cacheStore.Get(ctx, "session:usr_123")
    if err == nil {
        fmt.Println("Cached value:", val)
    }
}
```
