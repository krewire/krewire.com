---
title: "Button"
description: "Core interactive button component supporting primary, secondary, accent, outline, and ghost variants."
date: "2026-11-09"
---

# Button (`Button.kiw`)

Category: **Actions & Buttons**  
Component File: `components/Button.kiw`  
Source: `github.com/krewire/ui/components/Button.kiw`

The `<Button />` component is the fundamental interactive element of the Krewire UI design system. It renders either an HTML `<button>` or an accessible `<a>` link styled with pop-brutalist solid borders, hard offset shadows, and active click translations.

---

## 1. Copy & Own Installation

Copy `Button.kiw` directly into your project's `components/` directory:

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Button.kiw ./components/
```

---

## 2. Examples & Usage

### 2.1 Color Variants

```html
<Button variant="primary">Primary Action</Button>
<Button variant="secondary">Secondary Action</Button>
<Button variant="accent">Destructive Accent</Button>
<Button variant="outline">Outline Border</Button>
<Button variant="ghost">Ghost Button</Button>
```

### 2.2 Sizing Scale

```html
<Button variant="primary" size="sm">Small (sm)</Button>
<Button variant="primary" size="md">Medium (md - default)</Button>
<Button variant="primary" size="lg">Large (lg)</Button>
```

### 2.3 With Icons and Links

```html
<!-- Button with prepended icon -->
<Button variant="primary" icon="🚀">Launch Cluster</Button>

<!-- Button rendered as link -->
<Button href="/docs" variant="outline" iconRight="→">
  Read Documentation
</Button>
```

### 2.4 Disabled and Click Handlers

```html
<Button variant="primary" disabled=true>Disabled State</Button>
<Button variant="accent" onClick="handleDeleteEnvironment()">Delete</Button>
```

---

## 3. Supported Props

- `variant`: Button style variant (`primary`, `secondary`, `accent`, `outline`, `ghost`, default: `primary`).
- `size`: Sizing token (`sm`, `md`, `lg`, default: `md`).
- `href`: When provided, renders an accessible `<a>` tag with button styling.
- `target`: Target frame when `href` is present (e.g. `_blank`).
- `icon`: Prepend icon or emoji string.
- `iconRight`: Append icon or emoji string.
- `text` / `label`: Button text (alternative to slot body).
- `type`: Button HTML type attribute (`button`, `submit`, `reset`, default: `button`).
- `disabled`: Disables interaction and reduces opacity.
- `onClick`: JavaScript click event handler string.
- `id`: Optional element ID.
- `ariaLabel`: Accessible screen-reader label.
- `class`: Additional CSS classes appended to root.

---

## 4. Pop-Brutalist Visual Tokens

- **Border**: `var(--pop-border-thin, 2px solid var(--border))`
- **Drop Shadow**: `var(--pop-shadow-sm, 2px 2px 0px var(--border))`
- **Hover Transition**: `translate(-1.5px, -1.5px)` with `var(--pop-shadow, 4px 4px 0px var(--border))`
- **Active State**: `translate(0px, 0px)` with `1px 1px 0px var(--border)`
