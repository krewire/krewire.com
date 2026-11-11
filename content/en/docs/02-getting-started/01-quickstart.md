---
title: "Quickstart"
description: "Five-minute beginner quickstart: install kiw, scaffold a site, edit a page, add a component, run the dev server, and build for production."
date: "2026-11-09"
---

# Quickstart

<nav class="track-crumbs" aria-label="Beginner track">
<span class="track-crumbs-label">Beginner track</span>
<ol>
<li><a href="/">Choose a path</a></li>
<li><a href="/getting-started">Getting Started</a></li>
<li aria-current="step"><span>1 · Quickstart</span></li>
<li><a href="/getting-started/fullstack-app-tutorial">2 · Fullstack app</a></li>
<li><a href="/getting-started/dsl">3 · The .kiw DSL</a></li>
</ol>
</nav>

In five minutes you will install the `kiw` CLI, create a static site, change a page, write your first component, and produce a production build. You only need Go 1.27.1+ and a terminal — no Node.js, no npm, and no knowledge of Krewire's internals.

---

## Step 1: Install `kiw`

On Linux or macOS:

```bash
curl -fsSL https://krewire.com/scripts/install.sh | sh
```

On any platform with Go 1.27.1+ (including Windows):

```bash
go install github.com/krewire/krewire/tools/kiw/cmd/kiw@latest
```

Check that the shell can find it:

```bash
kiw version
```

The first line reads `kiw Krewire Devtool`, followed by the ecosystem version and your Go version. If the command is not found, see [Installation → Configuring your PATH](/getting-started/installation).

---

## Step 2: Create a Site

```bash
kiw new my-site --site --title "My Site"
cd my-site
```

`kiw` prints the files it created:

```text
created my-site/krewire.yaml
created my-site/pages/index.kiw
created my-site/layouts/Base.kiw
created my-site/components/Hero.kiw
created my-site/public/favicon.svg
created my-site/README.md
```

What each one is for:

```text
my-site/
├── krewire.yaml         # project name, kind (site), build output, dev port
├── go.mod               # Go module (created by kiw new)
├── pages/index.kiw      # the home page → /
├── layouts/Base.kiw     # the HTML shell every page renders inside
├── components/Hero.kiw  # a reusable component, used as <Hero /> in index.kiw
└── public/favicon.svg   # copied as-is into the build
```

A file in `pages/` becomes a URL, a file in `components/` becomes a tag, and a file in `layouts/` wraps pages. That is the whole mental model for a site.

---

## Step 3: Start the Dev Server

```bash
kiw dev
```

```text
INFO msg="dev server running" url=http://localhost:8080 ...
INFO msg="watching for changes" root=... interval=500ms
```

Open <http://localhost:8080>. Leave the server running for the next steps.

---

## Step 4: Edit the Page

Replace the contents of `pages/index.kiw` with:

```html
---
title: My Site
layout: Base
---

<div class="hero-container">
  <h1>Hello from {{.Title}}</h1>
  <p>I edited this page and the browser reloaded by itself.</p>
  <Badge variant="primary" text="Built with Krewire" />
</div>

<style>
  .hero-container {
    max-width: 680px;
    margin: 4rem auto;
    text-align: center;
  }
  .hero-container h1 {
    font-size: 2.75rem;
    color: #39D353;
  }
</style>
```

Save. The terminal logs `site rebuilt successfully` and the browser shows the new heading.

- The block between `---` lines is **frontmatter**: `title` is available as `{{.Title}}`, `layout: Base` picks `layouts/Base.kiw`.
- `<Badge … />` is a component from **Krewire UI** — the 66 built-in components are available in every site without installing anything. Browse them in the [UI catalog](/ui/components).
- The `<style>` block is **scoped**: its rules only apply to this page's markup.

---

## Step 5: Add Your Own Component

Create `components/Feature.kiw`:

```html
<div class="feature">
  <h3>{title}</h3>
  <p>{body}</p>
</div>

<style>
.feature {
  padding: 1rem 1.25rem;
  border: 1px solid #30363d;
  border-radius: 8px;
  text-align: left;
}
.feature h3 {
  margin: 0 0 0.5rem;
  color: #39D353;
}
</style>
```

`{title}` reads the `title` attribute; `{body}` is whatever you put between the opening and closing tags. Use it in `pages/index.kiw`, below the badge:

```html
  <div class="features">
    <Feature title="Fast">Pages are plain HTML generated at build time.</Feature>
    <Feature title="Simple">One CLI, one config file, no npm.</Feature>
  </div>
```

and add `.features { display: grid; gap: 1rem; margin-top: 2rem; }` to the page's `<style>` block. Save, and the two cards appear.

> **Tip:** `kiw make component Feature` also creates `components/Feature.kiw` for you, as a starting point to edit.

---

## Step 6: Build for Production

Stop the dev server (`Ctrl+C`) and run:

```bash
kiw build
```

```text
INFO building SSG site from file layout root=... output=.../my-site/.krewire/build
created index.html
created favicon.svg
created assets/krewire.css
```

`.krewire/build/` is the complete website: plain HTML and CSS that any static host (Nginx, Caddy, GitHub Pages, an object bucket) can serve. Preview it exactly as it will be served:

```bash
kiw serve --addr :3000
```

and open <http://localhost:3000>.

---

## What You Learned

| Concept | Where it lives |
| :--- | :--- |
| Pages are files | `pages/*.kiw` → one URL each |
| Layouts wrap pages | `layout: Base` → `layouts/Base.kiw` |
| Components are tags | `components/Feature.kiw` → `<Feature>` |
| Built-in UI | `<Badge>`, `<Button>`, `<Card>`, … from Krewire UI |
| One command per step | `kiw new` → `kiw dev` → `kiw build` → `kiw serve` |

## Next Steps

- Build something dynamic: [**2.2 Build a Fullstack App in 15 Minutes →**](/getting-started/fullstack-app-tutorial)
- Learn the full component syntax: [**2.7 The `.kiw` DSL →**](/getting-started/dsl)
- Writing documentation instead? Read the [Book workload](/workloads/book).
