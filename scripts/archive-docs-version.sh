#!/usr/bin/env bash
# Automated Documentation Release Archiver
# Snapshots content/<locale>/docs into content/<locale>/<version>/ for multi-version documentation browsing.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

VERSION="${1:-$(grep 'version:' "${ROOT_DIR}/krewire.yaml" | head -n 1 | awk '{print $2}' | tr -d '"' | tr -d "'")}"

if [[ -z "${VERSION}" ]]; then
  echo "Error: Version could not be determined from krewire.yaml or arguments." >&2
  exit 1
fi

echo "Archiving documentation snapshot for version: ${VERSION}"

LOCALES=("en" "id")

for LOC in "${LOCALES[@]}"; do
  SRC_DIR="${ROOT_DIR}/content/${LOC}/docs"
  DEST_DIR="${ROOT_DIR}/content/${LOC}/${VERSION}"

  if [[ -d "${SRC_DIR}" ]]; then
    mkdir -p "${DEST_DIR}"
    cp -R "${SRC_DIR}/"* "${DEST_DIR}/"
    echo "  ✓ Archived ${LOC} docs to: content/${LOC}/${VERSION}"
  else
    echo "  ⚠ Warning: Source directory ${SRC_DIR} not found, skipping."
  fi
done

echo "Documentation archive complete for ${VERSION}."
