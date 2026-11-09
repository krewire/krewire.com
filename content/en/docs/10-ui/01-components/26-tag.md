---
title: "Tag"
description: "Removable or clickable keyword tag with color variants."
date: "2026-11-09"
---

# Tag

Category: **Feedback & Status**  
Component File: `components/Tag.kiw`  
Source: `github.com/krewire/ui/components/Tag.kiw`

Removable or clickable keyword tag with color variants.

---

## 1. Copy & Own Installation

Copy the `Tag.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Tag.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Tag />` directly within your `.kiw` page or layout templates:

```html
<Tag class="my-custom-tag">
  <!-- Slot content or properties -->
</Tag>
```

---

## 3. Supported Props & Attributes

- `icon`: Configures the `icon` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `removable`: Configures the `removable` attribute or behavior.
- `selected`: Configures the `selected` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
