---
title: "Text"
description: "Typography wrapper for body paragraphs, lead text, and muted captions."
date: "2026-11-09"
---

# Text

Category: **Data Display & Content**  
Component File: `components/Text.kiw`  
Source: `github.com/krewire/ui/components/Text.kiw`

Typography wrapper for body paragraphs, lead text, and muted captions.

---

## 1. Copy & Own Installation

Copy the `Text.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Text.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Text />` directly within your `.kiw` page or layout templates:

```html
<Text class="my-custom-text">
  <!-- Slot content or properties -->
</Text>
```

---

## 3. Supported Props & Attributes

- `color`: Configures the `color` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `weight`: Configures the `weight` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
