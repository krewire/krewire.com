---
title: "Testing Utilities & Assertions"
description: "Testing assertion harness, fluent HTTP chain, golden snapshots, UI component assertions, and browser automation in pkg/testing."
date: "2026-10-25"
---

# Testing Utilities & Assertions (`pkg/testing`)

Import path: `github.com/krewire/krewire/pkg/testing`

`pkg/testing` delivers an idiomatic, pure-Go testing toolkit for the Krewire ecosystem. It includes fluent HTTP request builders, chainable response assertion matchers, temporary fixture management, golden file diffing, UI component inspection, and headless browser automation.

---

## 1. Fluent HTTP Testing Chain

The fluent HTTP client simplifies testing web handlers and REST APIs without boilerplate error handling:

```go
package web_test

import (
	"net/http"
	"testing"

	ftest "github.com/krewire/krewire/pkg/testing"
)

func TestProfileEndpoint(t *testing.T) {
	handler := SetupServer()

	// 1. Unauthenticated request rejected
	req := ftest.GET(t, "/api/profile").Request()
	ftest.Do(t, handler, req).
		Status(http.StatusUnauthorized).
		Header("Content-Type", "application/json").
		Contains("missing authentication token")

	// 2. Authenticated JSON request succeeds
	payload := map[string]string{"name": "Alice"}
	req = ftest.JSONRequest(t, "POST", "/api/profile", payload).
		Header("Authorization", "Bearer valid-token").
		Request()

	var res struct {
		Status string `json:"status"`
	}

	ftest.Do(t, handler, req).
		Status(http.StatusCreated).
		Header("Content-Type", "application/json").
		JSON(&res)

	ftest.Equal(t, res.Status, "ok")
}
```

---

## 2. Assertion Primitives

- **`test.Equal(t, got, want)`**: Deep equality check.
- **`test.NoError(t, err)`**: Asserts `err == nil`.
- **`test.HasError(t, err)`**: Asserts `err != nil`.
- **`test.Contains(t, s, substr)`**: String substring check.
- **`test.True(t, cond, msg)`**: Boolean condition assertion.

---

## 3. Golden Snapshot Testing

Golden files ensure rendered HTML and JSON outputs do not suffer unintended regressions:

```go
package compiler_test

import (
	"testing"

	ftest "github.com/krewire/krewire/pkg/testing"
)

func TestTemplateRendering(t *testing.T) {
	output := RenderTemplate("card", data)

	// Compares output with testdata/card.golden
	ftest.AssertGolden(t, "testdata/card.golden", []byte(output))
}
```

---

## 4. Headless Browser Automation (`pkg/testing/browser`)

The `browser` package provides an automated end-to-end testing harness:

```go
package e2e_test

import (
	"net/http/httptest"
	"testing"

	"github.com/krewire/krewire/pkg/testing/browser"
)

func TestWebUI_E2E(t *testing.T) {
	srv := httptest.NewServer(SetupRouter())
	defer srv.Close()

	b := browser.New(t, srv.URL)
	b.Navigate("/")
	b.AssertStatus(200)
	b.AssertTextVisible("Welcome to Krewire")
}
```

---

## 5. Alignment with the Three Pillars (3S)

- **Secure (S1)**: Hermetic in-process execution prevents opening live network ports during automated testing runs.
- **Sustainable (S2)**: Pure Go without requiring Node.js, Jest, Cypress, or complex testing daemons.
- **Scalable (S3)**: Unified testing contracts shared across HTTP APIs, WebAssembly islands, and CLI tools.
