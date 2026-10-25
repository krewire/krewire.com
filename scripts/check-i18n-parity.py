#!/usr/bin/env python3
"""
Verify 1:1 documentation file parity between English (`en`) and Indonesian (`id`)
in content/<locale>/docs.
"""

import os
import sys

def main():
    repo_root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    en_dir = os.path.join(repo_root, "content", "en", "docs")
    id_dir = os.path.join(repo_root, "content", "id", "docs")

    if not os.path.isdir(en_dir) or not os.path.isdir(id_dir):
        print(f"ERROR: directories not found: {en_dir} or {id_dir}", file=sys.stderr)
        sys.exit(1)

    en_files = set()
    for root, _, files in os.walk(en_dir):
        for f in files:
            if f.endswith(".md"):
                en_files.add(os.path.relpath(os.path.join(root, f), en_dir))

    id_files = set()
    for root, _, files in os.walk(id_dir):
        for f in files:
            if f.endswith(".md"):
                id_files.add(os.path.relpath(os.path.join(root, f), id_dir))

    missing_in_id = en_files - id_files
    missing_in_en = id_files - en_files

    if missing_in_id:
        print(f"ERROR: Missing in Indonesian (id): {sorted(missing_in_id)}", file=sys.stderr)
    if missing_in_en:
        print(f"ERROR: Missing in English (en): {sorted(missing_in_en)}", file=sys.stderr)

    if missing_in_id or missing_in_en:
        sys.exit(1)

    # Check for identical untranslated files
    untranslated = []
    for rel in sorted(en_files):
        en_path = os.path.join(en_dir, rel)
        id_path = os.path.join(id_dir, rel)
        with open(en_path, "r", encoding="utf-8") as fe, open(id_path, "r", encoding="utf-8") as fi:
            if fe.read().strip() == fi.read().strip():
                untranslated.append(rel)

    if untranslated:
        print(f"WARNING: The following {len(untranslated)} files are identical to English (untranslated):", file=sys.stderr)
        for u in untranslated:
            print(f" - {u}", file=sys.stderr)
        sys.exit(1)

    print(f"SUCCESS: 1:1 bilingual documentation parity verified across all {len(en_files)} documents!")

if __name__ == "__main__":
    main()
