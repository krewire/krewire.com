---
title: "Contributing & Community"
description: "How to participate in the Krewire ecosystem — code, documentation, translation, examples, testing, triage, and community support — plus workflow, standards, channels, and an honest account of governance today."
date: "2026-11-09"
---

# Contributing & Community

Krewire is meant to be built by its users, not only consumed by them. This page is the entry point for anyone who wants to participate — whether you write Go, write docs, translate, test, or help other people.

---

## Ways to contribute

You do not need to write code to make a meaningful contribution.

| Contribution | What it looks like | Where |
| :--- | :--- | :--- |
| **Code** | Fix a bug, improve a package, or extend `kiw`. | [`krewire/krewire`](https://github.com/krewire/krewire), [`krewire/krew`](https://github.com/krewire/krew), [`krewire/ui`](https://github.com/krewire/ui) |
| **Documentation** | Correct, clarify, or extend a docs page. Docs live in `content/<locale>/docs/` as Markdown. | [`krewire/krewire.com`](https://github.com/krewire/krewire.com) |
| **Translation** | Keep English and Indonesian pages consistent; translate new pages. | [`krewire/krewire.com`](https://github.com/krewire/krewire.com) (`content/id/`, `lang/id/`) |
| **Examples & starters** | Improve project templates or write a worked example. | `templates/` in [`krewire/krewire`](https://github.com/krewire/krewire) |
| **Testing** | Try a workload end to end on your platform and report what broke, with steps to reproduce. | Issues in the relevant repository |
| **Issue triage** | Reproduce reported bugs, ask for missing details, and suggest labels (`bug`, `documentation`, `good first issue`, …). | Issues in any public repository |
| **Community support** | Answer questions and share what you built. | [Discussions on `krewire/krewire.com`](https://github.com/krewire/krewire.com/discussions) |

Contributions are currently focused on the open-source core repositories (`krewire`, `krew`, `ui`, `krewire.com`).

---

## Finding a first contribution

1. Look for issues labelled [`good first issue`](https://github.com/search?q=org%3Akrewire+is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22&type=issues) or [`help wanted`](https://github.com/search?q=org%3Akrewire+is%3Aissue+is%3Aopen+label%3A%22help+wanted%22&type=issues) across the organization. **As of 2026-11-09 no open issue carries either label**, so this list may be empty.
2. If nothing is labelled, pick something you noticed yourself — a confusing docs page, a missing translation, a failing example — and open an issue describing it before starting larger work.
3. Documentation fixes are the fastest way in: every docs page on this site maps to one Markdown file in `content/en/docs/` and its Indonesian twin in `content/id/docs/`.

---

## Workflow and standards

These come from each repository's `CONTRIBUTING.md` and `AGENTS.md`:

1. **Branch** from `main` with a short-lived, descriptive branch (for example `fix/ssg-scoped-css`).
2. **Commit** using [Conventional Commits](https://www.conventionalcommits.org/) (`feat`, `fix`, `docs`, `refactor`, `test`, `chore`), in small, focused increments.
3. **Run the quality gates** before opening a pull request:

   ```bash
   gofmt -l .        # must print nothing
   go vet ./...
   go test ./...
   kiw compat        # after version.go changes
   ```

   For this website: `kiw build` must succeed, `python3 scripts/check-i18n-parity.py` must pass for docs, and `python3 scripts/sync-i18n.py` must be run after changing `lang/*/site.json`.
4. **Open a pull request** against `main`. Changes are reviewed before merge.
5. **Write in English** in repositories (code, commits, docs); Indonesian content belongs in the `id` locale files.

> For local developer setup guidance, refer to the public [Installation](/getting-started/installation) guide.

---

## Community channels

| Channel | Use it for | Status (2026-11-09) |
| :--- | :--- | :--- |
| [GitHub Issues](https://github.com/krewire) (each public repository) | Bugs, feature requests, docs problems. | Active |
| [GitHub Discussions on `krewire/krewire.com`](https://github.com/krewire/krewire.com/discussions) | Questions, ideas, show-and-tell. | Enabled; no discussions yet |
| Email: [hello@krewire.com](mailto:hello@krewire.com) | General contact. | Published contact |
| Instagram: [@krewire](https://instagram.com/krewire) | Announcements. | Social channel, not a support channel |

Security vulnerabilities should **not** be reported in public issues; see [Maturity, Roadmap & Support](/project/maturity-and-roadmap) for the current reporting route.

---

## Governance and maintainership today

To avoid implying structures that do not exist, here is the current state plainly:

- **Decision-making:** each repository is maintained by the Krewire Team; changes are decided through issue discussion and pull-request review by those maintainers.
- **Formal governance:** no governance charter, steering committee, RFC voting process, or contributor ladder is in operation.
- **Code of conduct:** none has been adopted yet. Until one is, contributors are asked to be respectful and constructive.
- **Sponsorship:** no sponsorship program is active; GitHub Sponsors is not enabled for the organization.
- **How responsibilities can grow:** contributors who make sustained, high-quality contributions may be invited by maintainers to help with triage or review. A written maintainership process is **not yet defined**; when it is, it will be published on this page.

Some community pages (for example the community portal source in [`krewire/krewire.github.io`](https://github.com/krewire/krewire.github.io)) describe contributor ladders or sponsorship tiers. As of 2026-11-09 these are not in operation; this page reflects the current state.

*Content owner: Krewire Team. Review this page whenever governance, channels, or contribution guidelines change, and at every tagged site release.*
