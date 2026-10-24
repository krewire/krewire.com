---
title: "Config & Validation (config, validation)"
description: "Reference guide for hierarchical configuration loading, environment overrides, dotenv support, and composable validation rules."
date: "2026-10-10"
---

# Config & Validation (`config`, `validation`)

Import paths:
- `github.com/krewire/krewire/pkg/config`
- `github.com/krewire/krewire/pkg/validation`

Robust configuration loading and strict input validation prevent runtime errors and security vulnerabilities before requests reach core domain logic.

---

## 1. Hierarchical Configuration (`packages/config`)

The `config` package loads typed application settings with deterministic priority:
1. Explicit CLI flags
2. Environment variables (`APP_*` prefix)
3. Local environment file (`.env.local`)
4. Base environment file (`.env`)
5. Configuration files (`krewire.yaml` / JSON)
6. Struct field default tags

### Example

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
```

---

## 2. Composable Input Validation (`packages/validation`)

The `validation` package validates data structures without reflection overhead, providing clear error structures for APIs and web forms.

### Common Rules:
- **`rules.Required`**: Value must not be empty or zero.
- **`rules.Email`**: RFC 5322 compliant email format.
- **`rules.Min` / `rules.Max`**: Numeric bounds or minimum/maximum string length.
- **`rules.OneOf`**: Value must match one of the allowed options.
- **`rules.Pattern`**: Regular expression validation.

### Example

```go
package main

import (
	"fmt"
	"github.com/krewire/krewire/pkg/validation"
	"github.com/krewire/krewire/pkg/validation/rules"
)

type RegisterRequest struct {
	Username string
	Email    string
	Age      int
}

func (r RegisterRequest) Validate() validation.Errors {
	v := validation.New()

	v.Check("username", r.Username, rules.Required(), rules.Min(3), rules.Max(32))
	v.Check("email", r.Email, rules.Required(), rules.Email())
	v.Check("age", r.Age, rules.Min(18))

	return v.Errors()
}

func main() {
	req := RegisterRequest{Username: "al", Email: "invalid-email", Age: 16}
	errs := req.Validate()

	if errs.HasErrors() {
		for field, msg := range errs.All() {
			fmt.Printf("- %s: %s\n", field, msg)
		}
	}
}
```
