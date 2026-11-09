---
title: "ActionMenu"
description: "Compact contextual dropdown action menu for table rows and cards."
date: "2026-11-09"
---

# ActionMenu

Category: **Actions & Buttons**  
Component File: `components/ActionMenu.kiw`  
Source: `github.com/krewire/ui/components/ActionMenu.kiw`

Compact contextual dropdown action menu for table rows and cards.

---

## 1. Copy & Own Installation

Copy the `ActionMenu.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/ActionMenu.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<ActionMenu />` directly within your `.kiw` page or layout templates:

```html
<ActionMenu class="my-custom-action-menu">
  <!-- Slot content or properties -->
</ActionMenu>
```

---

## 3. Supported Props & Attributes

- `align`: Configures the `align` attribute or behavior.
- `item`: Configures the `item` attribute or behavior.
- `trigger`: Configures the `trigger` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
