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

1. Look for issues labelled `good first issue` or `help wanted` across the organization.
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

- **GitHub Issues:** Bug reports and feature proposals across the public repositories.
- **GitHub Discussions:** Questions, ideas, and architecture sharing on `krewire/krewire.com`.
- **Email:** General inquiries at [hello@krewire.com](mailto:hello@krewire.com).

Security vulnerabilities should **not** be reported in public issues; see [Maturity, Roadmap & Support](/project/maturity-and-roadmap) for the security reporting route.

---

## Governance and maintainership

- **Decision-making:** Repositories are maintained by the Krewire Team; changes are reviewed and merged via pull requests.
- **Standards:** All code must pass automated quality gates (`gofmt`, `go vet`, `go test ./...`) with CGO disabled (`CGO_ENABLED=0`).
- **Growth:** Sustained contributors may be invited to maintain and review specific domains and packages.
