#!/usr/bin/env python3
"""
Generate the Krewire UI live-preview pages used by the component playground.

For every component in SPEC this script writes `pages/ui/preview/<slug>.kiw`.
Each page renders the *real* kui component once per variant/size combination
at build time (`kiw build` resolves the tags against the kui catalog), next to
the exact usage snippet that produced it. `public/assets/ui-playground.js`
loads these pages in light/dark iframes and switches between the prebuilt
combinations, so the playground needs no network access and no client-side
template engine.

With `--kui <path-to-krewire/ui checkout>` the script also copies the raw
component sources to `public/ui/source/<Name>.kiw.txt`, which back the
"Copy .kiw Component" button. Use the same kui revision that the `kiw`
binary building the site was compiled against (CI builds kiw with ui@main).

Usage:
    python3 scripts/gen-ui-previews.py                 # regenerate pages
    python3 scripts/gen-ui-previews.py --kui ../ui     # pages + sources
    python3 scripts/gen-ui-previews.py --check         # exit 1 if pages are stale
"""

import html
import os
import shutil
import sys

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
PAGES = os.path.join(ROOT, "pages", "ui", "preview")
SOURCES = os.path.join(ROOT, "public", "ui", "source")

# Each entry: slug, component name, axes (ordered), default selection, and a
# function returning the .kiw usage snippet for one combination. Axis names are
# what the playground shows as toggle groups.
SPEC = [
    {
        "slug": "button",
        "name": "Button",
        "axes": [
            ("variant", ["primary", "secondary", "accent", "error", "outline", "ghost"]),
            ("size", ["sm", "md", "lg"]),
        ],
        "default": {"variant": "primary", "size": "md"},
        "usage": lambda c: f'<Button variant="{c["variant"]}" size="{c["size"]}" text="Save changes" />',
    },
    {
        "slug": "badge",
        "name": "Badge",
        "axes": [
            ("variant", ["neutral", "primary", "secondary", "accent", "success", "warning", "error"]),
            ("size", ["sm", "md"]),
        ],
        "default": {"variant": "primary", "size": "md"},
        "usage": lambda c: f'<Badge variant="{c["variant"]}" size="{c["size"]}" text="Stable" />',
    },
    {
        "slug": "input",
        "name": "Input",
        "axes": [("state", ["default", "error", "disabled"])],
        "default": {"state": "default"},
        "usage": lambda c: {
            "default": '<Input name="email" id="email" type="email" label="Email" placeholder="you@example.com" help="We never share your address." />',
            "error": '<Input name="email" id="email" type="email" label="Email" value="not-an-email" error="Enter a valid email address." />',
            "disabled": '<Input name="email" id="email" type="email" label="Email" value="owner@example.com" disabled=true />',
        }[c["state"]],
    },
    {
        "slug": "card",
        "name": "Card",
        "axes": [("variant", ["default", "featured", "primary", "subtle"])],
        "default": {"variant": "default"},
        "usage": lambda c: (
            f'<Card variant="{c["variant"]}" title="Deploy preview" subtitle="Ready in 42 seconds">\n'
            "  <p>The static build was published to every edge node.</p>\n"
            "</Card>"
        ),
    },
    {
        "slug": "alert",
        "name": "Alert",
        "axes": [("variant", ["info", "success", "warning", "error"])],
        "default": {"variant": "info"},
        "usage": lambda c: f'<Alert variant="{c["variant"]}" title="Build finished" message="3 pages were regenerated." />',
    },
    {
        "slug": "modal",
        "name": "Modal",
        "axes": [("close", ["with-close", "without-close"])],
        "default": {"close": "with-close"},
        "usage": lambda c: (
            "<Button variant=\"primary\" text=\"Open modal\" onClick=\"document.getElementById('confirm-delete').showModal()\" />\n"
            + "<Modal id=\"confirm-delete\" title=\"Delete environment?\""
            + (" hideClose=true" if c["close"] == "without-close" else "")
            + ">\n"
            "  <p>This permanently removes the staging environment.</p>\n"
            "  <slot name=\"footer\">\n"
            "    <Button variant=\"outline\" text=\"Cancel\" onClick=\"this.closest('dialog').close()\" />\n"
            "    <Button variant=\"error\" text=\"Delete\" onClick=\"this.closest('dialog').close()\" />\n"
            "  </slot>\n"
            "</Modal>"
        ),
    },
    {
        "slug": "tabs",
        "name": "Tabs",
        "axes": [("items", ["three-tabs", "two-tabs"])],
        "default": {"items": "three-tabs"},
        "usage": lambda c: (
            '<Tabs label="Project sections">\n'
            + '  <button type="button" class="kiw-tab is-active" role="tab" aria-selected="true">Overview</button>\n'
            + '  <button type="button" class="kiw-tab" role="tab" aria-selected="false">Deployments</button>\n'
            + (
                '  <button type="button" class="kiw-tab" role="tab" aria-selected="false">Settings</button>\n'
                if c["items"] == "three-tabs"
                else ""
            )
            + "</Tabs>"
        ),
    },
]


def combos(axes):
    """Yield every combination of axis values as an ordered dict."""
    out = [{}]
    for name, values in axes:
        out = [dict(c, **{name: v}) for c in out for v in values]
    return out


def unique_ids(snippet, suffix):
    """Make element ids unique per rendered block (one page holds every combo)."""
    for ident in ("email", "confirm-delete"):
        snippet = snippet.replace(f'id="{ident}"', f'id="{ident}-{suffix}"')
        snippet = snippet.replace(f"getElementById('{ident}')", f"getElementById('{ident}-{suffix}')")
    return snippet


def render_page(spec):
    axes = spec["axes"]
    lines = [
        "---",
        f'title: "{spec["name"]} preview"',
        "layout: Preview",
        "---",
        f"<!-- Generated by scripts/gen-ui-previews.py — do not edit by hand. -->",
        f'<div class="pv-grid" data-pv-component="{spec["name"]}" data-pv-axes="{",".join(a for a, _ in axes)}">',
    ]
    for combo in combos(axes):
        suffix = "-".join(combo[a] for a, _ in axes)
        usage = spec["usage"](combo)
        attrs = " ".join(f'data-{a}="{combo[a]}"' for a, _ in axes)
        default = combo == spec["default"]
        lines.append(f'<div class="pv" data-pv {attrs}{" data-default" if default else ""}>')
        lines.append('  <div class="pv-stage">')
        for row in unique_ids(usage, suffix).splitlines():
            lines.append("    " + row)
        lines.append("  </div>")
        # The snippet the visitor copies keeps the plain ids; HTML-escaping
        # stops the .kiw desugarer from treating it as a component call.
        lines.append(f'  <pre class="pv-usage" hidden>{html.escape(usage, quote=False)}</pre>')
        lines.append("</div>")
    lines.append("</div>")
    return "\n".join(lines) + "\n"


def main():
    args = sys.argv[1:]
    check = "--check" in args
    kui = None
    if "--kui" in args:
        kui = args[args.index("--kui") + 1]

    stale = []
    os.makedirs(PAGES, exist_ok=True)
    for spec in SPEC:
        path = os.path.join(PAGES, spec["slug"] + ".kiw")
        body = render_page(spec)
        current = open(path, encoding="utf-8").read() if os.path.exists(path) else None
        if current != body:
            stale.append(os.path.relpath(path, ROOT))
            if not check:
                with open(path, "w", encoding="utf-8") as f:
                    f.write(body)

    if kui:
        os.makedirs(SOURCES, exist_ok=True)
        for spec in SPEC:
            src = os.path.join(kui, "components", spec["name"] + ".kiw")
            shutil.copyfile(src, os.path.join(SOURCES, spec["name"] + ".kiw.txt"))

    if check and stale:
        for p in stale:
            print(f"ERROR: {p} is out of date; run scripts/gen-ui-previews.py", file=sys.stderr)
        sys.exit(1)
    for p in stale:
        print(f"wrote {p}")


if __name__ == "__main__":
    main()
