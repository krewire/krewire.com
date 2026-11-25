---
title: "Multi-Driver Cache"
description: "High-performance in-memory, file, Redis, and Memcached caching with TTL, expiration sweeps, and Remember pattern in cache."
date: "2026-10-25"
---

# Multi-Driver Cache (`cache`)

Import path: `github.com/krewire/krewire/cache`

The `cache` library delivers a unified key-value caching abstraction with support for TTL expiration, LRU eviction bounds, background cleanup sweeps, and the atomic `Remember` pattern across the Krewire ecosystem.

---

## 1. Core Interface

```go
package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	GetString(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, val []byte, ttl time.Duration) error
	SetString(ctx context.Context, key string, val string, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Has(ctx context.Context, key string) (bool, error)
	Clear(ctx context.Context) error
	Remember(ctx context.Context, key string, ttl time.Duration, fn func() ([]byte, error)) ([]byte, error)
	Close() error
}

type Driver = Cache
```

---

## 2. Pluggable Cache Drivers

1. **In-Memory Cache (`cache.NewMemoryCache`)**:
   High-speed in-process cache with mutex synchronization, configurable capacity (default 10,000 items), and LRU eviction.
   ```go
   memCache := cache.NewMemoryCache(1 * time.Minute)
   defer memCache.Close()
   ```
2. **Redis / Valkey / Dragonfly (`cache.NewRedisCache`)**:
   Pure-Go RESP protocol client with automatic connection pooling and reconnect logic:
   ```go
   redisCache, err := cache.NewRedisCache(cache.RedisConfig{
       Address:  "127.0.0.1:6379",
       Password: os.Getenv("REDIS_PASSWORD"),
       DB:       0,
   })
   ```
3. **Filesystem Cache (`cache.NewFileCache`)**:
   Persistent disk-backed cache with atomic file writes and periodic cleanup sweeps.
   ```go
   fileCache, err := cache.NewFileCache("/tmp/krewire-cache", 10*time.Minute)
   ```
4. **Memcached (`cache.NewMemcachedCache`)**:
   Pure-Go Memcached driver supporting multi-server clusters.

---

## 3. Usage Example & Remember Pattern

The `Remember` pattern transparently retrieves cached items or executes a fallback resolver on cache miss:

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/krewire/krewire/cache"
)

func GetUserProfile(ctx context.Context, c cache.Cache, userID string) ([]byte, error) {
	key := fmt.Sprintf("users:%s", userID)

	return c.Remember(ctx, key, 30*time.Minute, func() ([]byte, error) {
		// Executed only on cache miss: fetch from primary database
		return FetchUserFromDatabase(userID)
	})
}
```

---

## 4. Alignment with the Three Pillars (3S)

- **Secure (S1)**: `FileCache` enforces restricted file permissions (`0600`) and rejects path traversal attempts.
- **Sustainable (S2)**: Pure Go without CGO or heavy external Redis client dependencies.
- **Scalable (S3)**: Seamlessly transition from `MemoryCache` on a single VPS to `RedisCache` or `Memcached` in distributed clusters without rewriting business logic.
