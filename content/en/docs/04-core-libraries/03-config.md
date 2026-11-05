---
title: "Hierarchical Configuration"
description: "Hierarchical application configuration loader with CLI flags, environment variables, and dotenv support in pkg/config."
date: "2026-10-25"
---

# Hierarchical Configuration (`pkg/config`)

Import path: `github.com/krewire/krewire/pkg/config`

The `config` package loads typed application settings with deterministic priority:
1. Explicit CLI flags
2. Environment variables (`APP_*` prefix)
3. Local environment file (`.env.local`)
4. Base environment file (`.env`)
5. Configuration files (`krewire.yaml` / JSON)
6. Struct field default tags

---

## Usage Example

```go
package main

import (
	"fmt"
	"github.com/krewire/krewire/pkg/config"
)

type AppConfig struct {
	Port     int    `env:"PORT" default:"8080"`
	DBHost   string `env:"DB_HOST" default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" default:"info"`
}

func main() {
	var cfg AppConfig
	// Loads from .env, env vars, and defaults automatically
	if err := config.Load(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("Server listening on port %d with log level %s\n", cfg.Port, cfg.LogLevel)
}
```\n