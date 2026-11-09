---
title: "Callout"
description: "Bordered callout container for documentation tips, warnings, and notes."
date: "2026-11-09"
---

# Callout

Category: **Feedback & Status**  
Component File: `components/Callout.kiw`  
Source: `github.com/krewire/ui/components/Callout.kiw`

Bordered callout container for documentation tips, warnings, and notes.

---

## 1. Copy & Own Installation

Copy the `Callout.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Callout.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Callout />` directly within your `.kiw` page or layout templates:

```html
<Callout class="my-custom-callout">
  <!-- Slot content or properties -->
</Callout>
```

---

## 3. Supported Props & Attributes

- `title`: Configures the `title` attribute or behavior.
- `type`: Configures the `type` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
