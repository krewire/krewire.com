---
title: "Card"
description: "Elevated card surface with header, media, body, and footer slots."
date: "2026-11-09"
---

# Card

Category: **Data Display & Content**

The `<Card />` component provides an elevated surface container with signature pop-brutalist borders (`2.5px solid var(--border)`), hard offset box-shadows, and slots for headers, media, body content, and action footers.

---

## Live Preview

<div data-component-preview="Card" data-slug="card"></div>

The frames above load the real `Card.kiw` from Krewire UI, prebuilt by `kiw build` for every variant, in the light and dark themes side by side. Use the toggles to switch variants, then copy the component source or the usage snippet. **Open in playground** shows it next to the other core components.

---

## Examples & Usage

### Basic Card with Titles and Badge

```html
<Card title="Database Cluster" subtitle="PostgreSQL 17.2 · Production" badge="Healthy">
  <p>Cluster is running at optimal load with 6 worker replicas.</p>
  <div slot="footer" class="flex justify-end gap-2">
    <Button variant="ghost" size="sm">View Logs</Button>
    <Button variant="outline" size="sm">Backup</Button>
  </div>
</Card>
```

### Featured Card with Custom Header and Media

```html
<Card variant="featured">
  <div slot="header" class="p-2 border-b">
    <h4 class="font-bold">Enterprise Cloud Gateway</h4>
  </div>
  <div slot="media">
    <img src="/assets/images/architecture.png" alt="Architecture Diagram" />
  </div>
  <p>High-throughput edge API gateway with token verification and rate limiting.</p>
  <div slot="footer">
    <Button variant="primary" size="md">Configure Gateway</Button>
  </div>
</Card>
```

---

## Supported Props

### Props
- `title`: Header title text.
- `subtitle`: Descriptive secondary text below the title.
- `badge`: Micro status badge text displayed in the header.
- `variant`: Card visual variant (e.g. `default`, `featured`).
- `class`: Additional CSS classes appended to card root.

### Slots
- `header` / `head`: Custom markup replacing the default title container.
- `media` / `image`: Visual image or banner media slot.
- `body` / default: Main card content body.
- `footer` / `actions`: Bottom action container.
