---
title: "Badge"
description: "Compact pill badge with color variants for tags, counts, and status indicators."
date: "2026-11-09"
---

# Badge

Category: **Feedback & Status**  
Component File: `components/Badge.kiw`  
Source: `github.com/krewire/ui/components/Badge.kiw`

Compact pill badge with color variants for tags, counts, and status indicators.

---

## Live Preview

<div data-component-preview="Badge" data-slug="badge"></div>

The frames above load the real `Badge.kiw` from Krewire UI, prebuilt by `kiw build` for every variant, in the light and dark themes side by side. Use the toggles to switch variants, then copy the component source or the usage snippet. **Open in playground** shows it next to the other core components.

---

## 1. Copy & Own Installation

Copy the `Badge.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Badge.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Badge />` directly within your `.kiw` page or layout templates:

```html
<Badge class="my-custom-badge">
  <!-- Slot content or properties -->
</Badge>
```

---

## 3. Supported Props & Attributes

- `dot`: Configures the `dot` attribute or behavior.
- `size`: Configures the `size` attribute or behavior.
- `text`: Configures the `text` attribute or behavior.
- `variant`: Configures the `variant` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
