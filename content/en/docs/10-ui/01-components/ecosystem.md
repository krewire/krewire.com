---
title: "Ecosystem"
description: "Interactive grid component displaying Krewire's 10-repository architecture."
date: "2026-11-09"
---

# Ecosystem

Category: **Data Display & Content**  
Component File: `components/Ecosystem.kiw`  
Source: `github.com/krewire/ui/components/Ecosystem.kiw`

Interactive grid component displaying Krewire's 10-repository architecture.

---

## 1. Copy & Own Installation

Copy the `Ecosystem.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Ecosystem.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Ecosystem />` directly within your `.kiw` page or layout templates:

```html
<Ecosystem class="my-custom-ecosystem">
  <!-- Slot content or properties -->
</Ecosystem>
```

---

## 3. Supported Props & Attributes

- `badge`: Configures the `badge` attribute or behavior.
- `header`: Configures the `header` attribute or behavior.
- `lead`: Configures the `lead` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
