---
title: "About Krewire"
description: "Krewire's mission, intended beneficiaries, impact-first philosophy, the role of financial sustainability, and long-term vision."
date: "2026-11-09"
---

# About Krewire

This chapter explains **why Krewire exists**, who it is for, what is open, how optional platform services relate to the open-source core, how to participate, and how the project reports its maturity and impact. Every statement here describes either a **current commitment** that can be checked against the public repositories, or a **long-term aspiration** that is labelled as such.

---

## Mission

> **Krewire's mission is to make it nearly free — in financial cost, operational burden, and cognitive effort — for anyone to build, own, and operate digital products that are secure, sustainable, and scalable.**

The same sentence anchors the [Philosophy](/overview/philosophy) chapter and the landing page. It is the ecosystem's reason for existing; individual tools and services are means to it, not ends in themselves.

### Who the mission is for

| Beneficiary | Why Krewire is relevant | Priority |
| :--- | :--- | :--- |
| **Independent developers** | One language (Go), one CLI (`kiw`), and one config file (`krewire.yaml`) reduce the number of tools a single person must learn and maintain. | Primary |
| **Small engineering teams** | Shared conventions across eight workload kinds reduce context switching and onboarding cost. | Primary |
| **Learners and educators** | The core source is public and MIT-licensed, so it can be read, taught, and modified. | Secondary |
| **Open-source maintainers** | Public repositories, Conventional Commits, and a documented release contract (`kiw release`, `kiw compat`). | Secondary |
| **Organizations** | Single static binaries and self-hostable output reduce dependence on proprietary runtimes. | Secondary |

See [Audiences & Use Cases](/project/audiences-and-use-cases) for concrete use cases per audience.

---

## Impact first, sustainability in service of the mission

Krewire's primary objective is **long-term impact on how software is built**: making secure defaults, maintainable architecture, and low-cost operation the easy path rather than the expensive one.

- **Impact is the measure of success.** Revenue is not the primary success metric of the ecosystem. How the project intends to measure impact — and which measurements do not exist yet — is published in the [Impact & Sustainability Report](/project/impact-report).
- **Open access is the default.** The core repositories (`krewire`, `krew`, `ui`, `krewire.com`, `krewire.github.io`) are public and MIT-licensed. Building, running, and deploying with the core tooling does not require a Krewire account, license key, or Krewire-hosted service.
- **Financial sustainability exists to fund the mission.** Maintaining an ecosystem costs time and infrastructure. Optional platform services (Dash, Cloud, Auth, Forge) are developed in private repositories and may be offered commercially in the future to help fund maintenance and community growth. This does **not** mean every service must be free, and it does **not** make the open-source core a paid product.

> **Current status (2026-11-09):** no paid plans, prices, or commercial licenses are published, and no funding sources are reported. When that changes, it will be documented in [Core & Platform Services](/project/core-and-platform-services) and the [Impact & Sustainability Report](/project/impact-report).

### Ecosystem mission vs. product value propositions

| | Ecosystem mission | Individual product value propositions |
| :--- | :--- | :--- |
| **Question it answers** | Why does Krewire exist? | Why would someone use this specific tool or service? |
| **Scope** | All repositories, documentation, and community work. | One product (for example, the `kiw` CLI or Krewire Cloud). |
| **Success measure** | Long-term ecosystem impact (see the impact report). | Product-specific usefulness; for optional services, possibly also revenue that funds the ecosystem. |
| **Commercial?** | No. | Open-source tooling: no. Optional platform services: may be, in the future. |

---

## Alignment with the three pillars

The mission is the *why*; the three engineering pillars are the *how*. Every capability is expected to pass all three before it ships (see [Philosophy](/overview/philosophy)):

| Pillar | How it serves the mission |
| :--- | :--- |
| **Secure** | Safety is the default, not a paid add-on, so builders without a security team still start from sane defaults. |
| **Sustainable** | Software a single builder can still read, run, and afford over many years — single binaries, opt-in batteries, MIT-licensed core. |
| **Scalable** | Growth is additive, not a rewrite, so early choices do not force an expensive migration later. |

---

## Long-term vision (aspiration)

The following are **long-term aspirations**, not current claims or measured outcomes:

- A single builder can still read, run, and afford what they built with Krewire a decade from now.
- Learning to build with Krewire builds transferable engineering skill (Go, the standard library, explicit architecture), not dependence on Krewire itself.
- The ecosystem is sustained by a community of contributors and maintainers beyond its original authors.
- Optional platform services fund the maintenance of the open-source core without ever becoming a prerequisite for it.

Progress toward these aspirations is not yet measured. The [Impact & Sustainability Report](/project/impact-report) lists what is measured today and what is not.

---

## What this chapter does not claim

- It does not report adoption numbers, user counts, or social outcomes — none have been measured.
- It does not describe a legal entity, nonprofit structure, or formal governance body — none is defined publicly.
- It does not publish pricing — none is published.

---

## In this chapter

1. [**Audiences & Use Cases**](/project/audiences-and-use-cases) — Who Krewire is for and how to adopt it incrementally.
2. [**Openness & Self-Hosting**](/project/openness-and-self-hosting) — Licenses, source repositories, and what runs without Krewire-hosted services.
3. [**Core & Platform Services**](/project/core-and-platform-services) — The boundary between the open-source core and optional services.
4. [**Contributing & Community**](/project/contributing) — How to participate, beyond code.
5. [**Impact & Sustainability Report**](/project/impact-report) — What is measured, what is not, and how the project is funded.
6. [**Maturity, Roadmap & Support**](/project/maturity-and-roadmap) — Status labels, roadmap, maintenance expectations, and reporting routes.
