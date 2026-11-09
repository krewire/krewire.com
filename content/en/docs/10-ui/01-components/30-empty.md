---
title: "Empty"
description: "Empty-state placeholder with icon, message, and call-to-action button."
date: "2026-11-09"
---

# Empty

Category: **Feedback & Status**  
Component File: `components/Empty.kiw`  
Source: `github.com/krewire/ui/components/Empty.kiw`

Empty-state placeholder with icon, message, and call-to-action button.

---

## 1. Copy & Own Installation

Copy the `Empty.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Empty.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Empty />` directly within your `.kiw` page or layout templates:

```html
<Empty class="my-custom-empty">
  <!-- Slot content or properties -->
</Empty>
```

---

## 3. Supported Props & Attributes

- `action`: Configures the `action` attribute or behavior.
- `actionText`: Configures the `actionText` attribute or behavior.
- `description`: Configures the `description` attribute or behavior.
- `href`: Configures the `href` attribute or behavior.
- `icon`: Configures the `icon` attribute or behavior.
- `title`: Configures the `title` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
