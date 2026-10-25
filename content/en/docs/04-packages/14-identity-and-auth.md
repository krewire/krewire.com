---
title: "Enterprise Identity & Auth Platform"
description: "Architecture, PBKDF2 hashing, RFC 6238 TOTP 2FA, enterprise SSO (OIDC/SAML2), RBAC, and web UI handlers in krewire/auth."
date: "2026-10-24"
---

# Enterprise Identity & Auth Platform (`github.com/krewire/auth`)

`krewire/auth` is the dedicated enterprise monorepo providing full-spectrum identity, Single Sign-On (SSO), multi-factor authentication, and access governance for the Krewire ecosystem.

## Subsystems

1. **`pkg/identity`**:
   - OWASP-recommended PBKDF2-HMAC-SHA256 password hashing (600,000 iterations).
   - Time-based One-Time Passwords (TOTP RFC 6238) with `otpauth://` provisioning URIs.
   - 8-count emergency backup recovery codes with SHA-256 hash validation.
   - Time-limited, single-use password reset token store.
2. **`pkg/sso`**:
   - Generic OIDC/OAuth2 with PKCE (RFC 7636).
   - Pre-configured identity provider presets: Google Workspace, GitHub, Microsoft Azure AD / Entra ID.
   - SAML 2.0 SP metadata generation and cryptographic assertion validation.
3. **`pkg/rbac`**:
   - Hierarchical roles (`superadmin` > `org_admin` > `developer` > `viewer`).
   - Granular `resource:action` permissions with wildcard support (`workloads:*`, `projects:deploy`).
   - Multi-tenant boundary isolation check (`evaluator.Can(ctx, user, action, tenant)`).
4. **`pkg/webauth`**:
   - Ready-to-mount HTTP handlers with clean, responsive dark/light UI:
     - `GET /login`, `POST /login`
     - `GET /register`, `POST /register`
     - `GET /forgot-password`, `POST /forgot-password`
     - `GET /reset-password`, `POST /reset-password`
     - `GET /2fa/challenge`, `POST /2fa/challenge`
     - `GET /2fa/setup`, `POST /2fa/setup`
     - `GET /logout`, `POST /logout`
