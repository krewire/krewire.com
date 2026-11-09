---
title: "Popover"
description: "Floating contextual popover container anchored to target elements."
date: "2026-11-09"
---

# Popover

Category: **Overlays & Dialogs**  
Component File: `components/Popover.kiw`  
Source: `github.com/krewire/ui/components/Popover.kiw`

Floating contextual popover container anchored to target elements.

---

## 1. Copy & Own Installation

Copy the `Popover.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Popover.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Popover />` directly within your `.kiw` page or layout templates:

```html
<Popover class="my-custom-popover">
  <!-- Slot content or properties -->
</Popover>
```

---

## 3. Supported Props & Attributes

- `content`: Configures the `content` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.
- `trigger`: Configures the `trigger` attribute or behavior.
- `triggerText`: Configures the `triggerText` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
