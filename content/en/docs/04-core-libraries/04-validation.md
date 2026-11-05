---
title: "Composable Input Validation"
description: "Composable data validation without reflection overhead, providing structured errors for APIs and forms in pkg/validation."
date: "2026-10-25"
---

# Composable Input Validation (`pkg/validation`)

Import path: `github.com/krewire/krewire/pkg/validation`

The `validation` package validates data structures without reflection overhead, providing clear error structures for APIs and web forms.

---

## Common Rules

- **`rules.Required`**: Value must not be empty or zero.
- **`rules.Email`**: RFC 5322 compliant email format.
- **`rules.Min` / `rules.Max`**: Numeric bounds or minimum/maximum string length.
- **`rules.OneOf`**: Value must match one of the allowed options.
- **`rules.Pattern`**: Regular expression validation.

---

## Usage Example

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
```\n