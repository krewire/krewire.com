---
title: "Audiences & Use Cases"
description: "Who Krewire is for, the concrete problems it addresses, representative use cases mapped to workloads and documentation, and how to adopt it incrementally alongside your existing stack."
date: "2026-11-09"
---

# Audiences & Use Cases

This page helps you decide quickly whether Krewire is relevant to your work and where to begin. It only maps use cases to capabilities that exist in the public core repository today; maturity per workload is listed in [Maturity, Roadmap & Support](/project/maturity-and-roadmap).

---

## Primary audiences

Krewire's initial focus is on **independent developers** and **small engineering teams**.

### Independent developers

**Problem:** one person maintaining a product often has to learn and keep up to date several toolchains — a JavaScript bundler for the front end, a separate backend framework, a docs generator, a job queue, and deployment scripts.

**What Krewire offers:** one language (Go), one CLI (`kiw`), and one configuration file (`krewire.yaml`) across eight workload kinds, with single-binary or static output that runs on modest hardware.

| Use case | Start with | Documentation |
| :--- | :--- | :--- |
| Publish a landing page, portfolio, or blog | `kiw new my-site --site` | [Static Sites (`site`)](/workloads/site) |
| Publish product documentation or a technical book | `kiw new my-docs --book` | [Documentation Books (`book`)](/workloads/book) |
| Ship a small web application or JSON API as one binary | `kiw new my-app --app` | [Fullstack Monoliths (`app`)](/workloads/app) |
| Build a command-line tool | `kiw new my-tool --cli` | [CLI & TUIs (`cli`)](/workloads/cli) |

### Small engineering teams

**Problem:** small teams pay a high coordination cost when every service uses different conventions, configuration formats, and deployment pipelines.

**What Krewire offers:** shared directory conventions, configuration, and lifecycle commands (`kiw dev`, `kiw build`, `kiw test`, `kiw deploy`) across workloads, so moving between services does not mean switching mental models.

| Use case | Start with | Documentation |
| :--- | :--- | :--- |
| Move slow work (emails, reports) out of the request path | `kiw new jobs --worker` | [Background Workers (`worker`)](/workloads/worker) |
| Add a headless service next to an existing app | `kiw new api --service` | [Microservices (`service`)](/workloads/service) |
| Put one entry point in front of several upstream services | `kiw new edge --gateway` | [API Gateways (`gateway`)](/workloads/gateway) |
| Deploy to a VPS without a container platform | `kiw deploy --target systemd` | [Deployment & Operations](/guides/deployment-and-operations) |
| Keep agent-assisted development consistent | `AGENTS.md` + Boost starter | [Agentic Development](/getting-started/agentic-development) |

---

## Secondary audiences

These audiences benefit from the same open foundation, but the website does not yet tailor dedicated material for them.

| Audience | Representative use case | Where to start |
| :--- | :--- | :--- |
| **Learners & educators** | Read and teach a complete, MIT-licensed Go codebase that spans web, CLI, workers, and infrastructure. | [Philosophy](/overview/philosophy), [Architecture Discipline](/architecture/layered-pyramid), source in `krewire/krewire` |
| **Open-source maintainers** | Build a project's documentation site with `site` or `book`, as this website does. | [Static Sites (`site`)](/workloads/site), [Contributing & Community](/project/contributing) |
| **Organizations** | Evaluate self-hostable, single-binary output and the open-core boundary before committing. | [Openness & Self-Hosting](/project/openness-and-self-hosting), [Core & Platform Services](/project/core-and-platform-services) |

---

## Incremental adoption — no full-stack migration required

Krewire is designed to be adopted one workload at a time. You do **not** need to migrate your whole stack.

1. **Start with one isolated workload.** A docs site (`book`) or a landing page (`site`) has no runtime dependency on the rest of your system.
2. **Co-locate, do not replace.** A Krewire workload can live in its own directory (with its own `go.mod` and `krewire.yaml`) inside an existing Node.js, Laravel, or PHP repository. According to the core interoperability guide, `kiw` detects these host projects (`package.json`, `composer.json`) but never modifies host manifests or dependency directories.
3. **Integrate over standard interfaces.** Workloads communicate over HTTP or standard input/output, and output is a plain Go binary or static files that any web server can serve. This website, for example, is built with `kiw build` and served by a stock nginx container (see the repository's `Dockerfile`).
4. **Grow only when needed.** Add `worker`, `service`, or `gateway` workloads when your requirements call for them; unused packages are not compiled into your output.

```text
existing-project/            # Node.js, Laravel, PHP … (untouched)
├── package.json
├── src/
└── docs/                    # new: Krewire book workload
    ├── krewire.yaml
    └── content/
```

### What Krewire does not do

- It does not transpile JavaScript or PHP into Go.
- It does not require Krewire-hosted services, accounts, or license keys (see [Openness & Self-Hosting](/project/openness-and-self-hosting)).
- It does not yet publish pre-built `kiw` binaries; the install script currently falls back to `go install`, which requires Go 1.27.1+ (see [Installation](/getting-started/installation)).

---

## Is Krewire a good fit?

| Krewire is likely a good fit if… | Krewire may not be a good fit yet if… |
| :--- | :--- |
| You are comfortable with Go, or want to standardize on it. | Your team cannot adopt Go 1.27.1+. |
| You value few dependencies and self-hostable output. | You need a long-term API compatibility guarantee — Krewire is pre-1.0 (v0.2.x). |
| You want to start small (one site, one tool) and grow. | You need a managed hosting service today — none is publicly available yet. |

*Content owner: Krewire Team. Update this page when workloads, starters, or maturity labels change.*
