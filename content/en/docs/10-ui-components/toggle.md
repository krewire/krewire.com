---
title: "Toggle"
description: "Binary toggle switch for instant setting changes."
date: "2026-11-09"
---

# Toggle

Category: **Actions & Buttons**  
Component File: `components/Toggle.kiw`  
Source: `github.com/krewire/ui/components/Toggle.kiw`

Binary toggle switch for instant setting changes.

---

## 1. Copy & Own Installation

Copy the `Toggle.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Toggle.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Toggle />` directly within your `.kiw` page or layout templates:

```html
<Toggle class="my-custom-toggle">
  <!-- Slot content or properties -->
</Toggle>
```

---

## 3. Supported Props & Attributes

- `activeIcon`: Configures the `activeIcon` attribute or behavior.
- `disabled`: Configures the `disabled` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `inactiveIcon`: Configures the `inactiveIcon` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `onClick`: Configures the `onClick` attribute or behavior.
- `pressed`: Configures the `pressed` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
