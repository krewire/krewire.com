---
title: "Tabs"
description: "Accessible tabbed content switcher supporting persistent active states."
date: "2026-11-09"
---

# Tabs

Category: **Navigation & Structure**  
Component File: `components/Tabs.kiw`  
Source: `github.com/krewire/ui/components/Tabs.kiw`

Accessible tabbed content switcher supporting persistent active states.

---

## Live Preview

<div data-component-preview="Tabs" data-slug="tabs"></div>

The frames above load the real `Tabs.kiw` from Krewire UI, prebuilt by `kiw build` for every variant, in the light and dark themes side by side. Use the toggles to switch variants, then copy the component source or the usage snippet. **Open in playground** shows it next to the other core components.

---

## 1. Copy & Own Installation

Copy the `Tabs.kiw` component directly into your project's `./components` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Tabs.kiw ./components/
```

Once copied, your project owns the component file completely. You can customize the HTML markup, logic, and styling without upstream library lock-in.

---

## 2. Basic Usage

Use `<Tabs />` directly within your `.kiw` page or layout templates:

```html
<Tabs class="my-custom-tabs">
  <!-- Slot content or properties -->
</Tabs>
```

---

## 3. Supported Props & Attributes

- `id`: Configures the `id` attribute or behavior.
- `label`: Configures the `label` attribute or behavior.
- `panels`: Configures the `panels` attribute or behavior.
- `tabs`: Configures the `tabs` attribute or behavior.

- `class`: Additional custom CSS classes to append to the component root.

---

## 4. Visual Tokens & Accessibility

- **Design System Stylesheet**: Inherits variables from `krewire-ui.css` (`--border`, `--radius`, `--pop-shadow`).
- **Dark Mode**: Automatically inverts surface and text colors when `data-theme="dark"` or the `.dark` class is applied to `:root`.
- **Zero-Bloat**: Built natively with pure HTML primitives and scoped styles with zero npm dependencies.
