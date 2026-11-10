# AGENTS.md — krewire.com Landing & Docs Portal

Agent guide for AI agents working inside the `krewire.com` repository
(`github.com/krewire/krewire.com`).

`krewire.com` is the official public website and documentation portal for Krewire.
It is 100% dogfooded using the ecosystem itself: `project.kind: site`, `.kiw` DSL,
`pages/`, `components/`, `layouts/`, and built with `kiw build`.

## Language

- **Write all stored artifacts in English** — pages, components, docs content, commit messages.
- If the human writes in Indonesian or another language, reply in the human's language in chat, but keep all repository files in English.

## Specs & Progress Tracking (GitHub Issues SSOT)

To keep progress tracking fully automated, resilient against markdown parsing drift, and deterministic for both AI agents and human developers:

1. **GitHub Issues as Single Source of Truth:**
   - Active specifications and their implementation progress are tracked as GitHub Issues under `krewire/krewire.com` with label `spec`.
   - Title format: `[Spec: KRW-<Code>] <Title>` (e.g. `[Spec: KRW-X7K9Q] Krewire Site — Landing + Documentation`).
   - Every issue body contains the structured requirement checklist with `- [ ]` and `- [x]` checkboxes citing `ReqID` and priority (e.g. `- [ ] KRW-LAND-001 [Must] ...`).

2. **Automated Agent & Developer Progress Workflow:**
   - **Discovering tasks:** Query active specs via CLI/API:
     ```bash
     gh issue list -R krewire/krewire.com --label spec --json number,title,url
     gh issue view <number> -R krewire/krewire.com --json title,body
     ```
   - **Verifying completion:** Build and inspect static output:
     ```bash
     kiw build
     kiw serve
     ```
   - **Updating status:** When a requirement is implemented and verified, check off the requirement box (`- [ ]` → `- [x]`) in the GitHub Issue body. Commits must reference the issue number and requirement ID (`feat(site): implement <ReqID> (ref #<IssueNumber>, req <ReqID>)`). Close the issue once all requirements are completed.

## Quality Gates

Before finishing any task, run:
```bash
kiw build
```
`kiw build` must exit 0 and write the site to `.krewire/build/` (the directory the
Docker image copies). Verify the landing page renders, theme toggling works, and
both locales are emitted under `.krewire/build/docs/` and `.krewire/build/id/docs/`.

Docs content is Markdown under `content/<locale>/docs/` and is rendered by
`pkg/book`; the `site` half owns `pages/index.kiw` and `layouts/`.

## Commit Discipline

- **Conventional Commits**: `feat|fix|docs|style|refactor|chore(site): desc`.
- Atomic commits (≤~300 LOC).
- Run `git` inside `krewire.com/`.

## Deployment Pathways (Production vs Hotfix)

- **Standard Production Release**: Triggered exclusively by pushing stable SemVer Git tags (`vX.Y.Z`) on `main`. Workflow: `.github/workflows/deploy.yaml`.
- **Emergency Hotfix Deployment**: Bypasses the release protocol and Git tags entirely. Push directly to the dedicated `hotfix` branch (`git checkout hotfix && git merge main --ff-only && git push origin hotfix`). Workflow: `.github/workflows/hotfix.yaml`. Never tag hotfixes or deploy them from `main`.

## Release checklist: public status content

Before tagging a site release (`vX.Y.Z`), review the status pages under
`content/<locale>/docs/11-project/` (English and Indonesian together) so that
every public claim stays verifiable:

1. **Maturity & roadmap** (`06-maturity-and-roadmap.md`): re-check tags (`git tag` in each repo), the upstream `krewire/docs/workloads/maturity-matrix.md`, open issues, and live endpoints; move roadmap items between shipped / in progress / proposed / deferred; update the "last verified" date.
2. **Core & platform services** (`03-core-and-platform-services.md`): re-check repository visibility, product availability, and pricing. Never publish pricing, tiers, or availability that is not live and owned.
3. **Openness** (`02-openness-and-self-hosting.md`): re-check licenses and network calls made by the core CLI.
4. **Impact report** (`05-impact-report.md`): add the new period's values with definition, period, and source; keep unavailable data labelled "not yet measured".
5. **Contributing** (`04-contributing.md`): re-check channels, labels, governance, and security-reporting routes.
6. Run `python3 scripts/check-i18n-parity.py`, `python3 scripts/sync-i18n.py --check`, and `kiw build`.
