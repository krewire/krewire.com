---
title: "Stack"
description: "Directional layout container with configurable gap spacing."
date: "2026-11-09"
---

# Stack

Category: **Layout Primitives**  
Component File: `components/Stack.kiw`  
Source: `github.com/krewire/ui/components/Stack.kiw`

Directional layout container with configurable gap spacing.

---

## 1. Copy & Own Installation

Copy the `Stack.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Stack.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Stack />` directly within your `.kiw` page or layout templates:

```html
<Stack class="my-custom-stack">
  <!-- Slot content or properties -->
</Stack>
```

---

## 3. Supported Props & Attributes

- `align`: Configures the `align` attribute or behavior.
- `direction`: Configures the `direction` attribute or behavior.
- `gap`: Configures the `gap` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `justify`: Configures the `justify` attribute or behavior.
- `wrap`: Configures the `wrap` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
