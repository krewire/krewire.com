---
title: "Kbd"
description: "Keyboard key visual indicator for shortcuts and command hotkeys."
date: "2026-11-09"
---

# Kbd

Category: **Data Display & Content**  
Component File: `components/Kbd.kiw`  
Source: `github.com/krewire/ui/components/Kbd.kiw`

Keyboard key visual indicator for shortcuts and command hotkeys.

---

## 1. Copy & Own Installation

Copy the `Kbd.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Kbd.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Kbd />` directly within your `.kiw` page or layout templates:

```html
<Kbd class="my-custom-kbd">
  <!-- Slot content or properties -->
</Kbd>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
