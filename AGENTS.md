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
   - `docs/specs/index.md` links directly to these GitHub Issues.

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
Verify that the output in `site/` builds without errors, theme toggling works, and HTML renders cleanly.

## Commit Discipline

- **Conventional Commits**: `feat|fix|docs|style|refactor|chore(site): desc`.
- Atomic commits (≤~300 LOC).
- Run `git` inside `krewire.com/`.
