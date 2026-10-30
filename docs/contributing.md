# Contributing to Krewire.com

Thank you for your interest in improving the official Krewire website and documentation!

---

## Contribution Workflow

We follow standard trunk-based development:

1. **Check Issues**: Find an open task or spec in `krewire/krewire.com` issues before starting work.
2. **Branch**: Create a short-lived branch off `main` (e.g., `docs/add-kiw-guide`, `fix/mobile-nav`).
3. **Edit**: Modify templates, content, or assets.
4. **Verify Locally**:
   - Run `kiw build` and ensure clean completion without warnings or errors.
   - Run `kiw serve` and inspect changes at `http://localhost:8080`.
   - Test both English and Indonesian versions if modifying documentation.
   - Test both light and dark themes.
5. **Commit**: Use Conventional Commits (`docs(content): ...`, `feat(site): ...`, `fix(ui): ...`).
6. **Pull Request**: Open a pull request against `main`.

---

## Content Guidelines

When writing or editing documentation under `content/`:
- **Clarity & Brevity**: Keep explanations focused, accurate, and direct. Avoid unnecessary complexity.
- **Bilingual Consistency**: Whenever updating English documentation in `content/en/docs/`, update or flag the corresponding Indonesian doc in `content/id/docs/`.
- **Code Blocks**: Always specify the language identifier for fenced code blocks (e.g. ````go`, ````yaml`, ````bash`).
- **No Hardcoded Links/Versions**: Do not hardcode version numbers in prose; refer to ecosystem conventions or use template variables where appropriate.

---

## Styling & Component Rules

- Scoped styles must remain in the corresponding `.kiw` file `<style>` blocks.
- Global styles live in `public/assets/krewire.css`.
- Use CSS design tokens defined in the theme (such as `--primary`, `--base-1`, `--text-muted`).
- Do not introduce heavy third-party CSS or JS libraries.
