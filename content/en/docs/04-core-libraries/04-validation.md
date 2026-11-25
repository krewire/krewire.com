---
title: "Declarative Struct Validation"
description: "Reflection-based struct and scalar validation using validate tags in pure Go in validation."
date: "2026-10-25"
---

# Declarative Struct Validation (`validation`)

Import path: `github.com/krewire/krewire/validation`

The `validation` package validates Go structs and scalar fields using `validate` struct tags. It is framework-agnostic and relies 100% on the Go standard library, allowing HTTP endpoints, CLI tools, and domain entities to share one unified validation engine.

---

## 1. Supported Rules

- **`required`**: Field must not be zero-valued, empty, or nil.
- **`min=N`**: Minimum numeric value or minimum string/slice length.
- **`max=N`**: Maximum numeric value or maximum string/slice length.
- **`len=N`**: Exact string or slice length.
- **`email`**: Valid RFC 5322 email address format.
- **`oneof=a b c`**: Value must match one of the whitespace-delimited options.
- **`pattern=regex`**: Value must match compiled regular expression.

---

## 2. Core API Reference

- **`validation.Struct(v any) error`**: Validates a struct or pointer to a struct. Returns `nil` when all rules pass, or a `*validation.ValidationError` detailing failing fields.
- **`validation.Field(value any, tag string) error`**: Validates a single scalar value against a rule string.

---

## 3. Usage Example

```go
package main

import (
	"errors"
	"fmt"

	"github.com/krewire/krewire/validation"
)

type UserRegistration struct {
	Username string `json:"username" validate:"required,min=3,max=32,pattern=^[a-zA-Z0-9_-]+$"`
	Email    string `json:"email"    validate:"required,email"`
	Role     string `json:"role"     validate:"required,oneof=developer admin observer"`
	Age      int    `json:"age"      validate:"min=13,max=120"`
}

func main() {
	user := UserRegistration{
		Username: "al",
		Email:    "invalid-email",
		Role:     "superuser",
		Age:      10,
	}

	err := validation.Struct(user)
	if err != nil {
		var valErr *validation.ValidationError
		if errors.As(err, &valErr) {
			for _, f := range valErr.Fields {
				fmt.Printf("Field %q failed rule %q (got: %v, param: %q)\n",
					f.Field, f.Rule, f.Value, f.Param)
			}
		}
	}
}
```

---

## 4. Alignment with the Three Pillars (3S)

- **Secure (S1)**: Invariants are validated at the perimeter before untrusted client payloads reach databases, command executors, or file storage.
- **Sustainable (S2)**: Pure Go standard library reflection; zero third-party dependencies, zero CGO.
- **Scalable (S3)**: Single validation schema shared identically across web handlers, CLI arguments, and background jobs.
