---
title: "Code"
description: "Inline monospace code snippet with high contrast styling."
date: "2026-11-09"
---

# Code

Category: **Data Display & Content**  
Component File: `components/Code.kiw`  
Source: `github.com/krewire/ui/components/Code.kiw`

Inline monospace code snippet with high contrast styling.

---

## 1. Copy & Own Installation

Copy the `Code.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Code.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Code />` directly within your `.kiw` page or layout templates:

```html
<Code class="my-custom-code">
  <!-- Slot content or properties -->
</Code>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
