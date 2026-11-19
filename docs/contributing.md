# Contributing to Krewire.com

Thank you for your interest in improving the official Krewire website and documentation!

---

## 1. Contribution Workflow

We follow standard trunk-based development:

1. **Check Specifications & Issues:**
   - Active specifications are tracked as GitHub Issues under `krewire/krewire.com` labeled `spec`.
   - Before starting substantial work, verify active issues or discuss requirements in an issue.
2. **Branching:**
   - Create a short-lived branch off `main` (e.g. `docs/improve-cli-reference`, `fix/navbar-theme-toggle`).
   - For emergency hotfixes only, work directly on the dedicated `hotfix` branch.
3. **Edit & Author:**
   - Follow Krewire `.kiw` component standards and Markdown guidelines.
   - All repository files, components, and commit messages must be in **English**.
4. **Local Verification:**
   - Run `kiw build` to confirm clean compilation into `.krewire/build/`.
   - Run `kiw serve` to preview changes at `http://localhost:8080`.
   - If editing documentation, run `python3 scripts/check-i18n-parity.py` to ensure 1:1 bilingual parity.
   - If editing translation keys, run `python3 scripts/sync-i18n.py --check`.
   - Verify both light and dark themes.
5. **Commit:**
   - Use Conventional Commits (`docs(content): ...`, `feat(site): ...`, `fix(ui): ...`).
   - Keep commits atomic (≤~300 LOC per commit).
   - Reference the issue number and requirement ID when applicable (`feat(site): implement KRW-LAND-001 (ref #12, req KRW-LAND-001)`).
6. **Pull Request:**
   - Open a pull request against `main`.

---

## 2. Documentation Guidelines

When authoring or modifying documentation under `content/`:
- **1:1 Bilingual Parity:**
  Whenever you add, move, rename, or update an English document in `content/en/docs/`, apply the identical structural change to the corresponding Indonesian file in `content/id/docs/`.
- **Clarity & Directness:**
  Keep explanations focused, accurate, and concise. Avoid unnecessary filler or speculation.
- **Frontmatter Invariants:**
  Each markdown file requires `title` and `description` in YAML frontmatter.
- **Code Blocks:**
  Always specify the language identifier for fenced code blocks (e.g. ````go`, ````yaml`, ````bash`, ````html`).
- **No Hardcoded Versions:**
  Do not hardcode versions into documentation prose; dynamic pages receive `.Version` from `krewire.yaml`.

---

## 3. Styling & Component Rules

- Scoped styles must remain in the corresponding `.kiw` file `<style>` blocks.
- Global styles live in `public/assets/krewire.css`.
- Use CSS design tokens defined in the theme (such as `--primary`, `--base-1`, `--border`, `--text-muted`).
- Do not introduce external CSS or JavaScript libraries (no npm, no Tailwind, no Bootstrap).
- When modifying core UI components in `components/`, regenerate static preview assets:
  ```bash
  python3 scripts/gen-ui-previews.py
  ```

---

## 4. Pre-Commit Quality Gate Checklist

Before submitting code, ensure all three quality gates pass:

```bash
# 1. 1:1 bilingual documentation parity across all 134 manuscripts
python3 scripts/check-i18n-parity.py

# 2. Translation key parity and sync verification
python3 scripts/sync-i18n.py --check

# 3. Clean static site generation
kiw build
```

---

## 5. Deployment Pathways

- **Standard Production Release**: Triggered exclusively by pushing stable SemVer Git tags (`vX.Y.Z`) on `main`. Workflow: [`.github/workflows/deploy.yaml`](../.github/workflows/deploy.yaml).
- **Emergency Hotfix Deployment**: Bypasses the release protocol and Git tags entirely. Pushed directly to `hotfix` branch (`git checkout hotfix && git merge main --ff-only && git push origin hotfix`). Workflow: [`.github/workflows/hotfix.yaml`](../.github/workflows/hotfix.yaml).
