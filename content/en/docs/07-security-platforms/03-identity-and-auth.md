---
title: "Enterprise Identity & Auth Platform"
description: "Architecture, PBKDF2 hashing, RFC 6238 TOTP 2FA, enterprise SSO (OIDC/SAML2), RBAC, session stores, CSRF protection, and integration in Krewire Auth."
date: "2026-10-31"
---

# Enterprise Identity & Auth Platform (Krewire Auth)

Krewire Auth is the enterprise identity platform providing full-spectrum identity, Single Sign-On (SSO), multi-factor authentication (MFA), role-based access governance, and agent-level security boundaries across the Krewire ecosystem.

---

## Subsystems & Architecture

### 1. `pkg/identity`
- **Password Hashing:** OWASP-recommended PBKDF2-HMAC-SHA256 (600,000 iterations).
- **Multi-Factor Authentication:** Time-based One-Time Passwords (TOTP RFC 6238) with QR code and `otpauth://` provisioning URIs.
- **Recovery Codes:** 8-count emergency backup recovery codes validated with SHA-256 hashes.
- **Persistent Storage:** Persistent SQL user and session stores supporting PostgreSQL and SQLite.
- **Redis Session Store:** High-throughput `RedisSessionStore` communicating over native RESP protocol with TTL expiration sweeps.

### 2. `pkg/sso`
- **OIDC / OAuth2:** Generic OIDC/OAuth2 with PKCE (RFC 7636) and state verification.
- **Provider Presets:** Pre-configured identity provider drivers for Google Workspace, GitHub, and Microsoft Azure AD / Entra ID.
- **SAML 2.0:** SP metadata generation, clock-skew tolerant cryptographic assertion validation, and signature verification.
- **OIDC Discovery:** Public discovery metadata endpoint (`/.well-known/openid-configuration`) and JWKS key management.

### 3. `pkg/rbac`
- **User Roles:** Hierarchical personas (`superadmin` > `org_admin` > `developer` > `viewer`).
- **Granular Permissions:** Resource-action rules with wildcard matching (`workloads:*`, `projects:deploy`, `secrets:read`).
- **Agent & Tool RBAC:** Dedicated permissions for autonomous AI agents (`AgentPermission`, `ToolPermission`) with policy evaluators preventing unprivileged tool calls.
- **Security Audit Trail:** Structured `AuditLogger` capturing all authentication, authorization, and administrative actions with IP and user context.

### 4. `pkg/middleware`
- **Double-Submit CSRF Cookie:** Cryptographic cookie-to-header token verification adhering to defense-in-depth security standards.
- **Session Context:** Safe extraction helpers (`WithSession`, `SessionFromContext`) integrating seamlessly with HTTP request lifecycles.

### 5. `pkg/webauth`
Ready-to-mount HTTP handlers with clean, responsive dark/light UI:
- `GET /login`, `POST /login`
- `GET /register`, `POST /register`
- `GET /forgot-password`, `POST /forgot-password`
- `GET /reset-password`, `POST /reset-password`
- `GET /2fa/challenge`, `POST /2fa/challenge`
- `GET /2fa/setup`, `POST /2fa/setup`
- `GET /logout`, `POST /logout`

---

## Platform Access & Integration

> **Availability (2026-11-09):** Krewire Auth is developed in a private repository and is **not publicly available** — there is no public sign-up, download, or hosted endpoint, and `auth.krewire.com` currently serves the main krewire.com website. No pricing, plans, or commercial licenses are published. See [Core & Platform Services](/project/core-and-platform-services) for the delivery model, maturity, and update process.

### Public OIDC & OAuth2 Integration

Krewire Auth is designed to integrate through standard protocols. The server answers OIDC discovery at `/.well-known/openid-configuration`. The hosted addresses below are **planned** and are not live:
- **Discovery Endpoint:** `https://auth.krewire.com/.well-known/openid-configuration`
- **JWKS Key Set:** `https://auth.krewire.com/.well-known/jwks.json`
- **Authorization Endpoint:** `https://auth.krewire.com/oauth/authorize`
- **Token Endpoint:** `https://auth.krewire.com/oauth/token`

### Securing Open-Source Workloads Locally

For open-source projects operating independently of the cloud service, Krewire provides native, zero-dependency authentication and security libraries directly in the core monorepo:
- `pkg/sec`: Constant-time comparisons, OWASP security headers, CSRF token validation, and SSRF guards.
- `pkg/auth`: Stateless JWT tokens, PBKDF2 password hashing, and HTTP Basic Auth middleware.

```go
import (
    "github.com/krewire/krewire/pkg/auth"
    "github.com/krewire/krewire/pkg/sec"
)
```

### Standalone Server

The standalone Krewire Auth server (`krewire-auth-server`) is not publicly distributed, and no enterprise licensing terms or pricing are published. Any future availability will be announced on [Core & Platform Services](/project/core-and-platform-services).
