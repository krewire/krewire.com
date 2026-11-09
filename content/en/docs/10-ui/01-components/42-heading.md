---
title: "Heading"
description: "Semantic heading typography wrapper supporting h1 through h6."
date: "2026-11-09"
---

# Heading

Category: **Data Display & Content**  
Component File: `components/Heading.kiw`  
Source: `github.com/krewire/ui/components/Heading.kiw`

Semantic heading typography wrapper supporting h1 through h6.

---

## 1. Copy & Own Installation

Copy the `Heading.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Heading.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Heading />` directly within your `.kiw` page or layout templates:

```html
<Heading class="my-custom-heading">
  <!-- Slot content or properties -->
</Heading>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `level`: Configures the `level` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
