---
title: "Impact & Sustainability Report"
description: "An evidence-based report on what Krewire contributes to the developer ecosystem and how the project is sustained: impact dimensions, metric definitions, reporting periods, data sources, and explicitly unavailable data."
date: "2026-11-09"
---

# Impact & Sustainability Report

This report explains what Krewire is contributing and how that work is sustained. It is deliberately small: it only publishes numbers that can be checked, and it says plainly where data does not exist.

| Field | Value |
| :--- | :--- |
| **Edition** | Baseline — first report |
| **Reporting period** | Project start (first public commit, 2026-08-24) to **2026-11-09** |
| **Owner** | Krewire Team |
| **Next update** | At the next tagged release of this website |

---

## How to read this report

Krewire separates three levels, and only treats the last two as impact:

| Level | Meaning | Example |
| :--- | :--- | :--- |
| **Outputs** | Things the project produced. Necessary, but not impact by themselves. | Releases, documentation pages, starters. |
| **Outcomes** | Changes for the people who use what was produced. | Successful first builds, projects in production, external contributors. |
| **Long-term impact** | Durable change in how software is built, aligned with the [mission](/project). | Builders who can still run and afford their software years later. |

**Vanity metrics are excluded as evidence of impact.** GitHub stars, forks, raw download counts, and page views are not reported as impact, because they do not show that anyone succeeded with Krewire.

---

## Outputs (measured)

| Metric | Definition | Value | Period | Source |
| :--- | :--- | :--- | :--- | :--- |
| Public repositories | Repositories in the `krewire` GitHub organization visible without authentication. | 5 | As of 2026-11-09 | GitHub organization API (`public_repos`) |
| Tagged versions of the core | Git tags on `krewire/krewire`. | 1 (`v0.2.0`) | As of 2026-11-09 | [`krewire/krewire` tags](https://github.com/krewire/krewire/tags) |
| Published GitHub Releases | Releases with downloadable assets across public repositories. | 0 | As of 2026-11-09 | GitHub Releases API |
| Workload starters | Workload kinds that `kiw new` can scaffold. | 8 | As of v0.2.0 | `kiw new` flags in `tools/kiw/internal/commands/new.go` |
| Documentation pages | Markdown pages per locale under `content/<locale>/docs/` of this website (English and Indonesian kept 1:1). | 135 per locale | As of 2026-11-09 | `scripts/check-i18n-parity.py` in this repository |

---

## Outcomes (mostly not yet measured)

| Metric | Definition | Value | Source / note |
| :--- | :--- | :--- | :--- |
| Active projects using Krewire | Distinct projects outside the Krewire organization built with `kiw` in the period. | **Not yet measured** | `kiw` has no telemetry; a voluntary, privacy-respecting method has not been defined. |
| First-build success | Share of new users whose first `kiw new` + `kiw build` succeeds. | **Not yet measured** | No onboarding study or opt-in reporting exists. |
| Verified production use | Public deployments built with Krewire, verifiable by URL and repository. | 1 — this website | [krewire.com](https://krewire.com) is built with `kiw build` by this repository's deploy workflow. It is the project's own site: evidence of dogfooding, not of external adoption. |
| External contributors | Commit authors outside the maintainer team, per period. | **Not yet measured** | No public maintainer list exists to separate external from internal authors. For context only: the default branches of the public repositories have 3 distinct commit-author email identities in the period. |
| Recurring contributors | External contributors active in two or more periods. | **Not yet measured** | Requires at least two reporting periods. |
| Maintainer retention & contributor diversity | Maintainers active across periods; spread of contributors across organizations and regions. | **Not yet measured** | No maintainer list or contributor survey exists. |
| Educational / community adoption | Courses, workshops, or community groups using Krewire. | **Not yet measured** | None reported to the project. |

---

## Long-term impact (aspirations, not measured)

The long-term aspirations are listed on [About Krewire](/project). No metric for them exists yet, and they will not be reported as achieved without evidence.

---

## Sustainability

Revenue is **not** the primary success measure of the ecosystem (see [About Krewire](/project)). Financial information is published so that users can judge whether the project can be maintained.

| Item | Status (2026-11-09) | Source / note |
| :--- | :--- | :--- |
| Funding sources | **None reported** | No grants, sponsorships, or revenue have been disclosed. |
| Sponsorship | **Not active** | `krewire/krewire` contains a `FUNDING.yaml` referencing GitHub Sponsors, but GitHub Sponsors is not enabled for the organization. |
| Revenue from platform services | **None** | No platform service is publicly available or priced ([Core & Platform Services](/project/core-and-platform-services)). |
| Operating costs | **Not yet published** | Known infrastructure: this website runs as an nginx container deployed to a VPS by the repository's deploy workflow. The cost has not been disclosed. |
| Operating-cost coverage & funding diversity | **Not yet measurable** | Requires published costs and funding sources. |

---

## Update process

To keep reporting within the project's current capacity:

1. **When:** at each tagged release of this website, or at least every six months.
2. **How:** add the new period's values next to the previous ones in the tables above; keep earlier values for comparison. The report is plain Markdown, so updating it needs no website redesign.
3. **Rules:** every metric keeps its definition, period, and source; unavailable data stays labelled "not yet measured"; estimates are never presented as measurements.
4. **Adding metrics:** propose a new metric in an issue on [`krewire/krewire.com`](https://github.com/krewire/krewire.com/issues) with a definition and a reliable data source before publishing it.
