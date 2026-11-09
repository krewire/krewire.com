#!/usr/bin/env python3
"""
Sync the landing-page translation dictionaries into the client bundle.

`lang/<locale>/site.json` is the source of truth for landing-page strings.
`public/assets/krewire.js` embeds a copy of those dictionaries
(`var translations = {...};`) so the language toggle works without an extra
network request. This script:

1. verifies that `en` and `id` declare exactly the same keys, and
2. rewrites the embedded `translations` object from the JSON files.

Usage:
    python3 scripts/sync-i18n.py           # verify + rewrite
    python3 scripts/sync-i18n.py --check   # verify only; exit 1 if out of sync
"""

import json
import os
import sys

LOCALES = ("en", "id")
MARKER = "var translations = "


def flatten(obj, prefix=""):
    keys = set()
    for k, v in obj.items():
        path = f"{prefix}.{k}" if prefix else k
        if isinstance(v, dict):
            keys |= flatten(v, path)
        else:
            keys.add(path)
    return keys


def main():
    check_only = "--check" in sys.argv[1:]
    root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

    dicts = {}
    for loc in LOCALES:
        with open(os.path.join(root, "lang", loc, "site.json"), encoding="utf-8") as f:
            dicts[loc] = json.load(f)

    en_keys, id_keys = flatten(dicts["en"]), flatten(dicts["id"])
    if en_keys != id_keys:
        for k in sorted(en_keys - id_keys):
            print(f"ERROR: key missing in id: {k}", file=sys.stderr)
        for k in sorted(id_keys - en_keys):
            print(f"ERROR: key missing in en: {k}", file=sys.stderr)
        sys.exit(1)

    js_path = os.path.join(root, "public", "assets", "krewire.js")
    with open(js_path, encoding="utf-8") as f:
        src = f.read()

    start = src.find(MARKER)
    if start < 0:
        print(f"ERROR: '{MARKER}' not found in {js_path}", file=sys.stderr)
        sys.exit(1)
    start += len(MARKER)
    _, length = json.JSONDecoder().raw_decode(src[start:])

    rendered = json.dumps({loc: dicts[loc] for loc in LOCALES}, indent=2, ensure_ascii=False)
    in_sync = src[start:start + length] == rendered

    if check_only:
        if not in_sync:
            print("ERROR: public/assets/krewire.js translations are out of sync; run scripts/sync-i18n.py", file=sys.stderr)
            sys.exit(1)
        print(f"SUCCESS: {len(en_keys)} landing keys in parity and embedded bundle is in sync.")
        return

    if not in_sync:
        with open(js_path, "w", encoding="utf-8") as f:
            f.write(src[:start] + rendered + src[start + length:])
    print(f"SUCCESS: {len(en_keys)} landing keys in parity; krewire.js {'updated' if not in_sync else 'already in sync'}.")


if __name__ == "__main__":
    main()
