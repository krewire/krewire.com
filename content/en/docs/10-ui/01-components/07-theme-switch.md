---
title: "ThemeSwitch"
description: "Dark and light mode toggle button with localStorage persistence."
date: "2026-11-09"
---

# ThemeSwitch

Category: **Actions & Buttons**  
Component File: `components/ThemeSwitch.kiw`  
Source: `github.com/krewire/ui/components/ThemeSwitch.kiw`

Dark and light mode toggle button with localStorage persistence.

---

## 1. Copy & Own Installation

Copy the `ThemeSwitch.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/ThemeSwitch.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<ThemeSwitch />` directly within your `.kiw` page or layout templates:

```html
<ThemeSwitch class="my-custom-theme-switch">
  <!-- Slot content or properties -->
</ThemeSwitch>
```

---

## 3. Supported Props & Attributes

- None (accepts default slot content or children).

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
