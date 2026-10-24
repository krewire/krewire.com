---
title: "Philosophy"
description: "The core philosophy, three pillars, and foundational engineering tenets guiding Krewire's design and digital SDLC ecosystem."
date: "2026-10-14"
---

# Philosophy

Krewire is guided by a fundamental belief: **make it nearly free — in financial cost, operational burden, and cognitive effort — for anyone to build, own, and operate digital products that are secure, sustainable, and scalable.**

Software should turn builders into sovereign producers, not perpetual renters of proprietary platforms. Krewire is delivered as **Modular Libraries**: independent, single-purpose Go packages that you compose behind one configuration file (`krewire.yaml`) and one unified CLI (`kiw`).

---

## 1. The Three Pillars

Every architectural decision, package, and capability in the Krewire ecosystem must satisfy three non-negotiable pillars:

| Pillar | Principle | What it Looks Like in Krewire |
| :--- | :--- | :--- |
| **Secure** | Safety is the default, not an enterprise add-on. | Standard library first, strict downward dependency graphs, secure-by-default HTTP primitives, OWASP/CWE alignment, sandboxed runtimes, and secrets referenced via `${env:NAME}` rather than committed. |
| **Sustainable** | Software a single builder can read, run, and afford decades later. | Single static binaries with embedded templates and assets, zero runtime lock-in, opt-in batteries that cost nothing when unused, and a permanent MIT open-core foundation. |
| **Scalable** | Growth is progressive and additive, never a rewrite. | Reversible provider contracts, standard workload kinds, and the progressive pipeline allowing products to scale from a $5 VPS to hyperscale without rewriting business logic. |

---

## 2. The Seven Foundational Tenets

The three pillars are the outcomes; the following seven tenets are the engineering means through which those outcomes are achieved:

### 1. Ownership is Power

Software you rent eventually governs you. When third-party platforms control the build pipeline, the runtime environment, or the deployment targets, they dictate the margins, the pricing, and the rules.

Krewire builds tools whose output builders own outright:
- **Single-Binary Deploys:** Applications compile down to self-contained, statically linked machine code.
- **Embedded Assets:** Templates, migrations, and static assets reside inside the binary via Go's `embed` package.
- **Self-Sovereign Data:** Databases, file storage, and cache providers are decoupled from vendor lock-in.
- **Zero Runtime Dependencies:** No required external language interpreters, container runtimes, or proprietary cloud engines.

### 2. Cost is a Design Decision

Near-zero cost is engineered, never discounted. Every unnecessary dependency, indirection, and piece of cleverness charges interest forever — in hosting bills, maintenance overhead, security audits, and onboarding time.

Krewire rejects stacks that silently exclude builders through prohibitive operational costs:
- **Standard Library First:** Leverage Go's high-performance standard library (`net/http`, `html/template`, `log/slog`, `context`, `flag`) before introducing external third-party dependencies.
- **Modular Packaging:** Projects import only what they use. If an app only needs static site generation, background queue packages are never linked into the build.
- **Efficient Compute:** Lean memory footprints (10–30 MB) and millisecond boot times allow production workloads to run reliably on affordable hardware.
- **Permanent Open Core:** The foundational tier remains MIT-licensed forever. No paid license is ever a prerequisite to build, deploy, or operate a complete system.

### 3. One Path Beats Many Choices

Fragmentation taxes hardest those who can least afford it — early-stage teams, students, and independent developers. Infinite optionality creates decision fatigue and friction.

Committing to one coherent toolchain transforms fragmented exploration into compounding mastery:
- **One Config:** `krewire.yaml` is the single source of truth for all workload definitions, environment bindings, and build targets.
- **One CLI:** `kiw` handles project creation, local development, testing, compilation, and deployment across all workloads.
- **One Language:** Go powers the entire stack — from backend services and worker queues to static sites and client-side WebAssembly runtimes.
- **Universal Mental Model:** Eight workload kinds share identical directory semantics, configuration syntax, and lifecycle phases.

### 4. Simple Enough to Teach, Boring Enough to Trust

Security, sustainability, and reliability are outcomes of disciplined restraint, not features bolted on after the fact. Security comes from a minimal attack surface and sane defaults; sustainability comes from boring, battle-tested components; scalability comes from honest engineering rather than speculative abstraction.

- **Explicit Over Magic:** No hidden reflection magic, unreadable metaprogramming, or implicit global state.
- **Legible Architecture:** Every package has a single responsibility (SRP) and clear cohesion boundaries.
- **Auditability:** Codebases engineered with Krewire can be read, understood, and audited by a junior engineer or a third-party security reviewer without proprietary training.

### 5. Learn by Shipping, Ship by Learning

True engineering competency is developed on production systems, not synthetic exercises. Learners should build software businesses actually run, and businesses should operate software learners can actually inspect and comprehend.

- **Production-Grade Scaffolding:** `kiw new` scaffolds real architectural foundations with structured error handling, health checks, and logging from minute one.
- **Documentation as a First-Class Product:** Comprehensive, version-controlled guides, architecture decision records (ADRs), and executable examples accompany every package.
- **Community-Driven SDLC:** Fostering an inclusive developer ecosystem where open collaboration bridges the gap between education and industry.

### 6. Openness is How We Work, Not Just a License

Trust is earned in the open. Krewire commits to radical legibility across its codebase, roadmap, and development standards.

- **Transparent Evolution:** Decisions, architectural compromises, and design choices are documented publicly in version-controlled records.
- **Honest Open-Core Ladder:** The license ladder (`free` → `pro` → `team` → `enterprise`) is strictly additive. Higher tiers may wrap, harden, or add compliance automation for enterprise teams, but free tier modules are never crippled or feature-gated to drive upgrades.
- **Shared Standards:** Tools, linters, and quality gates run identically on a developer's local laptop and on centralized CI/CD pipelines.

### 7. Modular Libraries over Monolithic Frameworks

Frameworks impose their conventions, lifecycle traps, and rigid class hierarchies onto your application. When requirements change, teams are forced into catastrophic rewrites.

Krewire approaches architecture through composable, modular libraries:
- **Inversion of Control Stays with You:** You compose Krewire packages into your application's `main.go`, rather than inheriting an opinionated runtime container.
- **Progressive Evolution:** Start with a static documentation portal or landing page (`site`), add server-rendered routes and database models (`app`), extract intensive tasks into background workers (`worker`), and decompose into autonomous microservices (`service`) only when organizational scale demands it.
- **Zero Rewrite Tax:** Components, domain business logic, and configuration carry forward intact across each stage of growth.

---

## 3. Mapping Foundations to the Three Pillars

Each of the seven tenets directly supports and reinforces the primary pillars:

| Foundational Tenet | Secure | Sustainable | Scalable |
| :--- | :---: | :---: | :---: |
| **1. Ownership is Power** | ✓ | ✓ | |
| **2. Cost is a Design Decision** | | ✓ | ✓ |
| **3. One Path Beats Many Choices** | ✓ | ✓ | |
| **4. Simple Enough to Teach, Boring Enough to Trust** | ✓ | ✓ | |
| **5. Learn by Shipping, Ship by Learning** | | ✓ | |
| **6. Openness is How We Work** | | ✓ | |
| **7. Modular Libraries over Monolithic Frameworks** | | ✓ | ✓ |

- **Security** is anchored in stack ownership (you control patch cycles), the single path (drastically smaller dependency surface), and boring, battle-tested components.
- **Sustainability** is preserved through near-zero infrastructure costs, open-source longevity, teachable codebases, and perpetual MIT licensing.
- **Scalability** is delivered through composable modules and the progressive pipeline, enabling teams to grow from a hobby server to global distribution without rewriting code.

---

## 4. Next Steps

With Krewire's philosophical foundations established, continue exploring the architecture:

- Proceed to [**1.3 Krewire Workloads →**](/overview/krewire-workloads) to inspect the 8 execution models.
- Review [**1.1 Why Krewire? ←**](/overview/why-krewire) to see how Krewire resolves toolchain fragmentation.
- Or jump directly into [**2. Getting Started →**](/getting-started).
