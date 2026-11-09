---
title: "Link"
description: "Accessible hyperlink with underline transitions and target security."
date: "2026-11-09"
---

# Link

Category: **Actions & Buttons**  
Component File: `components/Link.kiw`  
Source: `github.com/krewire/ui/components/Link.kiw`

Accessible hyperlink with underline transitions and target security.

---

## 1. Copy & Own Installation

Copy the `Link.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Link.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Link />` directly within your `.kiw` page or layout templates:

```html
<Link class="my-custom-link">
  <!-- Slot content or properties -->
</Link>
```

---

## 3. Supported Props & Attributes

- `external`: Configures the `external` attribute or behavior.
- `href`: Configures the `href` attribute or behavior.
- `id`: Configures the `id` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
