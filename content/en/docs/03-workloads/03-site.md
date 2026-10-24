---
title: "Static Site Generator (site)"
description: "High-performance static site compilation with file-based routing, single-file .kiw components, compile-time scoped CSS, and zero client-side JavaScript."
date: "2026-10-08"
---

# Static Site Generator (`site`)

The **`site`** workload provides an ultra-fast static site generator designed for marketing websites, landing pages, blogs, and public documentation portals.

Powered by `packages/web/ssg` and the `.kiw` DSL, the `site` compiler transforms Single-File Components (SFC) into optimized, semantic HTML and CSS with **zero runtime JavaScript overhead by default**.

---

## Architectural Characteristics

1. **File-Based Routing:** Pages map directly from the filesystem hierarchy:
   - `pages/index.kiw` → `/index.html` (served at `/`)
   - `pages/about.kiw` → `/about.html` (served at `/about`)
   - `pages/blog/[slug].kiw` → Dynamic slug generation from Markdown frontmatter.
2. **Compile-Time Scoped CSS:** Styles declared in `<style>` blocks are automatically hashed and scoped with `data-kiw-component` attributes. No CSS collisions or complex naming conventions needed.
3. **No Build Step Bloat:** Compiles dozens of pages in under 15ms using pure Go. No Node.js or `node_modules` required.
4. **Tailwind CSS Plugin:** Native integration with Tailwind CSS: when `tailwind.config.js` exists, `kiw` automatically triggers minified utility generation.

---

## Scaffolding a `site` Project

```bash
kiw new my-site --site
cd my-site
```

Standard directory layout:

```text
my-site/
├── krewire.yaml          # Project configuration
├── layouts/
│   └── Base.kiw          # HTML document layout
├── pages/
│   ├── index.kiw         # Home page
│   └── about.kiw         # About page
├── components/
│   ├── Navbar.kiw        # Reusable navigation
│   └── Footer.kiw        # Reusable footer
├── public/               # Raw static files copied verbatim
│   ├── favicon.svg
│   └── robots.txt
└── assets/
    └── tailwind.css      # Optional Tailwind stylesheet
```

---

## Writing a Component (`pages/index.kiw`)

```html
---
title: "Modern Web Engineering in Go"
layout: Base
---

<section class="hero">
  <div class="badge">Ecosystem Release v0.1.0</div>
  <h1>Modular Go Libraries. Zero Bloat.</h1>
  <p>Compile production web surfaces directly with native Go tooling.</p>
  <a href="/about" class="btn">Learn More →</a>
</section>

<style>
.hero {
  padding: 4rem 1.5rem;
  text-align: center;
  background: var(--bg);
}
.badge {
  display: inline-block;
  padding: 0.25rem 0.75rem;
  background: var(--primary);
  color: var(--primary-content);
  font-weight: 700;
  border-radius: 999px;
  font-size: 0.85rem;
  margin-bottom: 1.5rem;
}
h1 {
  font-size: 2.75rem;
  font-weight: 900;
  letter-spacing: -0.03em;
  color: var(--fg);
}
p {
  color: var(--muted);
  font-size: 1.15rem;
  max-width: 600px;
  margin: 1rem auto 2rem;
}
.btn {
  display: inline-block;
  background: var(--primary);
  color: var(--primary-content);
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-weight: 700;
  text-decoration: none;
}
</style>
```

---

## Layout Document (`layouts/Base.kiw`)

The layout wraps pages with HTML metadata, global styles, and slots:

```html
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <link rel="stylesheet" href="/assets/style.css">
</head>
<body>
  <Navbar />
  <main>{{.Content}}</main>
  <Footer />
</body>
</html>
```

---

## Development & Deployment

### Start Dev Server with Live Reload

```bash
kiw dev
```

Visit `http://localhost:8080`. Edits to `.kiw` files, layouts, and components trigger hot reloading in `< 10ms`.

### Build Static Site

```bash
kiw build
```

The output in `.krewire/build` is completely static and can be deployed directly to GitHub Pages, Cloudflare Pages, AWS S3, or Nginx.
