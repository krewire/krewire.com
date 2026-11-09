---
title: "LangSwitch"
description: "Language selector switch for multi-lingual internationalized applications."
date: "2026-11-09"
---

# LangSwitch

Category: **Actions & Buttons**  
Component File: `components/LangSwitch.kiw`  
Source: `github.com/krewire/ui/components/LangSwitch.kiw`

Language selector switch for multi-lingual internationalized applications.

---

## 1. Copy & Own Installation

Copy the `LangSwitch.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/LangSwitch.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<LangSwitch />` directly within your `.kiw` page or layout templates:

```html
<LangSwitch class="my-custom-lang-switch">
  <!-- Slot content or properties -->
</LangSwitch>
```

---

## 3. Supported Props & Attributes

- `code`: Configures the `code` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
