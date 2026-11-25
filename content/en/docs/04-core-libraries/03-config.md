---
title: "Declarative Configuration"
description: "Type-safe YAML configuration loading with environment variable overlay and struct field binding in config."
date: "2026-10-25"
---

# Declarative Configuration (`config`)

Import path: `github.com/krewire/krewire/config`

The `config` package provides typed configuration loading from YAML files with environment variable overlays for the Krewire ecosystem.

Precedence is strictly:
$$\text{Built-in Struct Defaults} < \text{YAML Configuration File} < \text{Environment Variables}$$

---

## 1. Core API Reference

- **`config.Load(path string, dst any) error`**: Unmarshals the YAML file at `path` into `dst` (must be a non-nil pointer to a struct). If the file does not exist, leaves `dst` untouched and returns `nil`.
- **`config.LoadOrDefault(path string, dst any) error`**: Loads YAML if present; tolerates an absent configuration file without error.
- **`config.Override(dst any, lookup func(string) (string, bool), opts ...Option) error`**: Applies environment variables on top of an already-loaded struct.
- **`config.WithPrefix(prefix string) Option`**: Customizes the prefix for implicit environment keys (default: `APP`).

---

## 2. Struct Tag & Environment Mapping

- Fields tagged `yaml:"key"` bind from YAML configuration files.
- Fields tagged `env:"EXACT_KEY"` bind from the exact environment variable name.
- Untagged fields derive an implicit environment key from their name or YAML tag uppercased with prefix (e.g. `Server.Port` becomes `APP_SERVER_PORT`).

---

## 3. Usage Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/krewire/krewire/config"
	"github.com/krewire/krewire/validation"
)

type ServerConfig struct {
	Host string `yaml:"host" env:"HOST"`
	Port int    `yaml:"port" env:"PORT" validate:"required,min=1,max=65535"`
}

type AppConfig struct {
	AppName string       `yaml:"app_name"`
	Server  ServerConfig `yaml:"server"`
	DBURL   string       `yaml:"db_url" env:"DATABASE_URL"`
}

func main() {
	// 1. Sensible defaults
	cfg := AppConfig{
		AppName: "Krewire Service",
		Server: ServerConfig{
			Host: "127.0.0.1",
			Port: 8080,
		},
	}

	// 2. Load YAML file
	if err := config.Load("krewire.yaml", &cfg); err != nil {
		panic(err)
	}

	// 3. Overlay environment variables
	if err := config.Override(&cfg, os.LookupEnv, config.WithPrefix("APP")); err != nil {
		panic(err)
	}

	// 4. Validate invariants
	if err := validation.Struct(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("%s listening on %s:%d\n", cfg.AppName, cfg.Server.Host, cfg.Server.Port)
}
```

---

## 4. Alignment with the Three Pillars (3S)

- **Secure (S1)**: Sensitive database URLs, tokens, and certificates are injected dynamically via environment variables (`${env:DATABASE_URL}`), preventing hardcoded secrets in repositories.
- **Sustainable (S2)**: Pure Go standard library reflection and lightweight YAML parsing; no heavy external configuration frameworks.
- **Scalable (S3)**: Single unified configuration model shared across HTTP servers, CLI tools, workers, and background daemons.
