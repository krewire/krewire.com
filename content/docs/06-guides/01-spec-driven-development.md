---
title: "Spec-Driven Development"
description: "The spec-first methodology: how requirements are authored, assigned unique identifiers, cited in Go unit tests, and verified by quality gates."
date: "2026-10-08"
---

# Spec-Driven Development

Krewire mandates a **Spec-First** development workflow. Code and tests never float free of documented intent. Before implementing any non-trivial capability, an engineer or autonomous agent authors a formal specification defining unambiguous, verifiable requirements.

---

## 1. Specification Anatomy & Identifiers

Specifications reside in the owning repository under `docs/specs/` and follow the filename convention:

```text
{ProjectId}-{Scope}-{SpecID}-{slug}.md
```

Examples:
- `KWL-CORE-K1N2Q-core-business-rules.md` (Libraries / Packages)
- `KWN-BUILD-1QGI2-project-building.md` (CLI Tool)
- `KWM-BUILDER-FX9H2-mdbind-site-builder.md` (Mdbind Engine)
- `KRW-SITE-X7K9Q-landing-site.md` (Krewire.com Site)

### SpecID Prefixes by Project

| Prefix | Owning Repository | Scope |
| :--- | :--- | :--- |
| **`KWL-*`** | `krewire/packages` | Modular domain libraries |
| **`KWN-*`** | `krewire/tools/kiw` | Developer CLI |
| **`KWM-*`** | `mdbind` | Standalone book engine |
| **`KWG-*`** | `krewire/templates/boost` | Boost AI agent templates |
| **`KRW-*`** | `krewire.com` & sites | Websites & documentation portals |

---

## 2. Requirement Numbering

Every requirement in a specification receives a deterministic ID formatted as:

```text
REQ-{Category}-{Number}
```

For example:
- `REQ-BUILD-001`: The compiler MUST emit static HTML without external network access.
- `REQ-SEC-012`: Cookies MUST include `SameSite=Strict` and `HttpOnly` attributes.

---

## 3. Test Traceability Citations

When writing Go tests, developers cite the governing SpecID and Requirement ID in comment annotations immediately preceding the test function:

```go
package ssg_test

import (
	"testing"
	"github.com/krewire/krewire/packages/testing/assert"
)

// Tests for KWL-SSG-8F12X
// Spec: KWL-SSG-8F12X REQ-BUILD-001 Scope: BUILD
func TestCompiler_EmitsStaticHTMLWithoutNetwork(t *testing.T) {
	output, err := CompileFixture("testdata/landing")

	assert.NoError(t, err)
	assert.Contains(t, output, "<h1>Welcome</h1>")
}
```

This bidirectional traceability guarantees:
1. Every line of test code verifies an explicit, agreed requirement.
2. Requirements without tests fail CI auditing.
3. Feature regressions and scope drift are detected instantly.
