---
title: "Openness & Self-Hosting"
description: "Which Krewire repositories are public and under which license, which capabilities run without Krewire-hosted services, how to self-host your output, and how Krewire coexists with external runtimes."
date: "2026-11-09"
---

# Openness & Self-Hosting

This page lets you check — without trusting marketing copy — what you can **inspect, use, modify, self-host, and operate independently**. Each statement links to a license, repository, or technical reference. Facts were verified against the repositories on **2026-11-09**.

---

## Repository availability & licensing

| Repository | What it contains | Visibility | License |
| :--- | :--- | :--- | :--- |
| [`krewire/krewire`](https://github.com/krewire/krewire) | Core monorepo: `pkg/*` libraries, the `kiw` CLI, templates. | Public | [MIT](https://github.com/krewire/krewire/blob/main/LICENSE) (package-level `LICENSE` files are also MIT) |
| [`krewire/krew`](https://github.com/krewire/krew) | Autonomous AI agent engine (embed, CLI, HTTP/SSE, MCP). | Public | [MIT](https://github.com/krewire/krew/blob/main/LICENSE) |
| [`krewire/ui`](https://github.com/krewire/ui) | `.kiw` design system and UI components. | Public | [MIT](https://github.com/krewire/ui/blob/main/LICENSE) |
| [`krewire/krewire.com`](https://github.com/krewire/krewire.com) | This website and documentation. | Public | [MIT](https://github.com/krewire/krewire.com/blob/main/LICENSE) |
| [`krewire/krewire.github.io`](https://github.com/krewire/krewire.github.io) | Community portal source. | Public | [MIT](https://github.com/krewire/krewire.github.io/blob/main/LICENSE) |
| `krewire/dash` | Krewire Dash — virtual workspace server. | **Private** | Source not publicly available; no public license grant. |
| `krewire/cloud` | Krewire Cloud — hybrid PaaS control plane. | **Private** | Source not publicly available; no public license grant. |
| `krewire/auth` | Krewire Auth — identity, SSO, MFA, RBAC server. | **Private** | Source not publicly available; no public license grant. |
| `krewire/panel` | Krewire Panel — server administration panel. | **Private** | Source not publicly available; no public license grant. |
| `krewire/internal` | Internal documentation hub (ADRs, roadmaps). | **Private** | Not publicly available. |

The open-source core is everything in the **public** rows. Contribution routes for each public repository are described in [Contributing & Community](/project/contributing).

---

## What works without Krewire-hosted services

The following capabilities of the public core run entirely on your machine or your own infrastructure. None of them requires a Krewire account, license key, or Krewire-operated server.

| Capability | Commands / packages | Network access it needs |
| :--- | :--- | :--- |
| Scaffold projects for all eight workloads | `kiw new`, `kiw init` | None (built-in templates); optional `--template <git URL>` clones from a URL you choose. |
| Local development and preview | `kiw dev`, `kiw serve`, `kiw dashboard` | None. |
| Build sites, books, binaries | `kiw build` | None beyond fetching Go modules you depend on. |
| Test and verify | `kiw test`, `kiw vet`, `kiw compat` | None. |
| Deploy to your own targets | `kiw deploy --target` (`binary`, `systemd`, `gh-pages`, `infra`) | Only the target you choose (for example, pushing to your own GitHub Pages branch). |
| Use the libraries in any Go program | `github.com/krewire/krewire/pkg/...` | None — standard Go packages. |
| Run the AI agent | `krew` | An LLM provider of your choice; with Ollama (`OLLAMA_HOST`) inference can run fully locally ([krew README](https://github.com/krewire/krew#readme)). |

**Network calls that do exist in the core CLI** (verified in source):

- `kiw self-update` queries the GitHub Releases API for `krewire/krewire`.
- The [install script](https://krewire.com/scripts/install.sh) downloads from GitHub Releases when a release exists and otherwise falls back to `go install` (Go module proxy). No GitHub Release has been published yet, so today the fallback path is used and Go 1.27.1+ is required.
- The CLI source declares default endpoints under `api.krewire.com` for future platform integrations; as of v0.2.0 no core `kiw` command calls them.

**One caveat in the libraries:** `pkg/auth.NewClient` is a client for a Krewire Auth server and defaults to `https://auth.krewire.com` when no URL is given. The Auth server is private and no public hosted Auth service is available (see [Core & Platform Services](/project/core-and-platform-services)). For fully independent use, the stateless primitives in the same package — `SignJWT`/JWT parsing and `BasicAuth` middleware — need no server. See [Authentication Primitives](/security-platforms/auth).

---

## Self-hosting your output

Krewire does not host your application; you do. The output of `kiw build` / `kiw deploy` is:

- **Static files** for `site` and `book` — servable by any static web server or GitHub Pages. This website is an example: it is built with `kiw build` and served by a stock `nginx:alpine` container (see this repository's [`Dockerfile`](https://github.com/krewire/krewire.com/blob/main/Dockerfile)).
- **A single Go binary** for `app`, `cli`, `worker`, `service`, and `gateway` — runnable on any Linux/macOS host; `kiw deploy --target systemd` generates a systemd unit.

**Operational requirements:**

| Requirement | Detail |
| :--- | :--- |
| Build toolchain | Go 1.27.1 or newer (all public modules declare `go 1.27.1`). |
| Installer platforms | Linux and macOS on amd64/arm64 ([install script](https://krewire.com/scripts/install.sh)). |
| `systemd` target | A Linux host with systemd. |
| `krew` from source | Its `go.mod` currently uses a local `replace` for the core module, so build it inside a multi-repository Go workspace (`go work`), as this site's CI does. |

See [Deployment & Operations](/guides/deployment-and-operations) for step-by-step guidance.

---

## Optional hosted services

Krewire Dash, Cloud, Auth, and Panel are **optional**. No public core command or library depends on them to build, run, or deploy your software. Their visibility, delivery model, and current availability are documented in [Core & Platform Services](/project/core-and-platform-services). In short: all four are developed in private repositories and none is publicly available as a hosted service as of 2026-11-09.

---

## Interoperability and coexistence

- **Co-location with other stacks.** A Krewire workload can live in its own directory inside a Node.js, Laravel, or PHP repository. `kiw info` detects these host projects but does not modify `package.json`, `composer.json`, `node_modules/`, or `vendor/` ([interoperability guide](https://github.com/krewire/krewire/blob/main/docs/guides/interoperability-and-migration.md), covered by tests such as `TestDetectHost_NodeJS` and `TestDetectHost_Laravel`).
- **Standard interfaces.** Workloads talk over HTTP or standard input/output; output is ordinary binaries and static files.
- **Your infrastructure, your choice.** Deploy targets are your own host (binary/systemd), GitHub Pages, or your own infrastructure via `pkg/cloud/infra`.

*Content owner: Krewire Team. Re-verify the tables above at every tagged release; repository visibility and licenses are the source of truth.*
