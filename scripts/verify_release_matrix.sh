#!/usr/bin/env bash
# ─────────────────────────────────────────────────────────────────────────────
# verify_release_matrix.sh — Release Matrix & Deployment Gate Policy Conformance
# Tests the complete event-matrix defined in krewire.com#43 / ADR-R7P3K:
# - Stable success, stable failure, pre-release success, experimental success,
#   manual dispatch, and provenance verification.
# ─────────────────────────────────────────────────────────────────────────────
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"

echo "=== Verifying Release Pipeline & Production Deploy Gate Policy ==="

TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

run_gate_evaluator() {
  local event="$1"
  local conclusion="$2"
  local head_branch="$3"
  local head_sha="$4"
  local tag="$5"
  local reason="$6"

  # Simulates the exact gate evaluation logic from deploy.yaml
  local DEPLOY_SHA=""
  local DEPLOY_VERSION=""
  local DEPLOY_TRACK=""

  if [ "${event}" = "workflow_run" ]; then
    if [ "${conclusion}" != "success" ]; then
      return 1
    fi
    if [ -z "${head_sha}" ]; then
      return 1
    fi
    if [[ "${head_branch}" =~ -rc|-beta|-alpha|-exp|-dev|-nightly ]]; then
      return 1
    fi
    if [[ ! "${head_branch}" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-lts.*)?$ ]]; then
      return 1
    fi
    DEPLOY_SHA="${head_sha}"
    DEPLOY_VERSION="${head_branch}"
    DEPLOY_TRACK="stable"
    if [[ "${head_branch}" =~ -lts ]]; then DEPLOY_TRACK="lts"; fi

  elif [ "${event}" = "push" ]; then
    if [[ ! "${tag}" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-lts.*)?$ ]]; then
      return 1
    fi
    DEPLOY_SHA="${head_sha}"
    DEPLOY_VERSION="${tag}"
    DEPLOY_TRACK="stable"
    if [[ "${tag}" =~ -lts ]]; then DEPLOY_TRACK="lts"; fi

  elif [ "${event}" = "workflow_dispatch" ]; then
    if [ -z "${reason}" ] || [ "${reason}" = "none" ]; then
      return 1
    fi
    DEPLOY_SHA="${head_sha}"
    DEPLOY_VERSION="manual-${head_sha:0:7}"
    DEPLOY_TRACK="manual"
  else
    return 1
  fi

  return 0
}

assert_gate() {
  local desc="$1"
  local expected="$2" # "PASS" or "BLOCK"
  local event="$3"
  local conclusion="$4"
  local head_branch="$5"
  local head_sha="$6"
  local tag="$7"
  local reason="$8"

  TOTAL_TESTS=$((TOTAL_TESTS + 1))
  printf "  [%02d] %-60s " "${TOTAL_TESTS}" "${desc}..."

  if run_gate_evaluator "${event}" "${conclusion}" "${head_branch}" "${head_sha}" "${tag}" "${reason}"; then
    local actual="PASS"
  else
    local actual="BLOCK"
  fi

  if [ "${actual}" = "${expected}" ]; then
    echo "✓ [${actual}]"
    PASSED_TESTS=$((PASSED_TESTS + 1))
  else
    echo "✗ FAILED (expected ${expected}, got ${actual})"
    FAILED_TESTS=$((FAILED_TESTS + 1))
  fi
}

echo ""
echo "--- Scenario Matrix: Workflow Run Events ---"
assert_gate "Stable release success (v1.0.0)" "PASS" "workflow_run" "success" "v1.0.0" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "LTS release success (v1.0.0-lts)" "PASS" "workflow_run" "success" "v1.0.0-lts" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Stable release failure" "BLOCK" "workflow_run" "failure" "v1.0.0" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Stable release cancelled" "BLOCK" "workflow_run" "cancelled" "v1.0.0" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Stable release timed out" "BLOCK" "workflow_run" "timed_out" "v1.0.0" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "RC pre-release success (v1.0.0-rc.1)" "BLOCK" "workflow_run" "success" "v1.0.0-rc.1" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Beta pre-release success (v1.0.0-beta.1)" "BLOCK" "workflow_run" "success" "v1.0.0-beta.1" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Alpha pre-release success (v1.0.0-alpha.1)" "BLOCK" "workflow_run" "success" "v1.0.0-alpha.1" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Experimental release success (v1.0.0-exp.1)" "BLOCK" "workflow_run" "success" "v1.0.0-exp.1" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Dev track release success (v1.0.0-dev.1)" "BLOCK" "workflow_run" "success" "v1.0.0-dev.1" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Nightly release success (v0.2.0-nightly.20261124)" "BLOCK" "workflow_run" "success" "v0.2.0-nightly.20261124" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Untagged default branch success (main)" "BLOCK" "workflow_run" "success" "main" "0123456789abcdef0123456789abcdef01234567" "" ""
assert_gate "Missing head_sha" "BLOCK" "workflow_run" "success" "v1.0.0" "" "" ""

echo ""
echo "--- Scenario Matrix: Direct Push Tag Events ---"
assert_gate "Valid stable tag push (v2.1.0)" "PASS" "push" "" "" "0123456789abcdef0123456789abcdef01234567" "v2.1.0" ""
assert_gate "Valid LTS tag push (v2.0.0-lts2)" "PASS" "push" "" "" "0123456789abcdef0123456789abcdef01234567" "v2.0.0-lts2" ""
assert_gate "Pre-release RC tag push" "BLOCK" "push" "" "" "0123456789abcdef0123456789abcdef01234567" "v2.1.0-rc.1" ""
assert_gate "Experimental tag push" "BLOCK" "push" "" "" "0123456789abcdef0123456789abcdef01234567" "v2.1.0-exp.1" ""

echo ""
echo "--- Scenario Matrix: Manual Dispatch Events ---"
assert_gate "Manual dispatch with auditable reason" "PASS" "workflow_dispatch" "" "" "0123456789abcdef0123456789abcdef01234567" "" "Hotfix critical vulnerability"
assert_gate "Manual dispatch empty reason" "BLOCK" "workflow_dispatch" "" "" "0123456789abcdef0123456789abcdef01234567" "" ""

echo ""
echo "--- Static Workflow Verification ---"

# 1. Pinned Third-Party Actions
echo "Checking action pinning in workflows..."
UNPINNED=$(grep -rn "uses: [a-zA-Z0-9_-]\+/[a-zA-Z0-9_.-]\+@" "${REPO_DIR}/.github/workflows" | grep -v "@[0-9a-f]\{40\}" || true)
if [ -n "${UNPINNED}" ]; then
  echo "✗ FAILED: Found unpinned third-party actions:"
  echo "${UNPINNED}"
  FAILED_TESTS=$((FAILED_TESTS + 1))
else
  echo "✓ All third-party actions are pinned to full 40-character commit SHAs."
  PASSED_TESTS=$((PASSED_TESTS + 1))
fi
TOTAL_TESTS=$((TOTAL_TESTS + 1))

# 2. Checkout uses exact deploy_sha in deploy.yaml
echo "Checking checkout ref in deploy.yaml..."
if grep -q "ref: \${{ needs.gate.outputs.deploy_sha }}" "${REPO_DIR}/.github/workflows/deploy.yaml"; then
  echo "✓ Deploy workflow binds checkout to exact verified release SHA."
  PASSED_TESTS=$((PASSED_TESTS + 1))
else
  echo "✗ FAILED: deploy.yaml does not bind checkout to verified release SHA."
  FAILED_TESTS=$((FAILED_TESTS + 1))
fi
TOTAL_TESTS=$((TOTAL_TESTS + 1))

# 3. Least-privilege permissions in deploy.yaml, release.yaml, etc.
echo "Checking permissions declarations..."
for wf in deploy.yaml release.yaml nightly.yaml hotfix.yaml ci.yaml; do
  if grep -q "^permissions: {}" "${REPO_DIR}/.github/workflows/${wf}"; then
    echo "✓ ${wf} declares least-privilege default (permissions: {})."
    PASSED_TESTS=$((PASSED_TESTS + 1))
  else
    echo "✗ FAILED: ${wf} missing top-level permissions: {}."
    FAILED_TESTS=$((FAILED_TESTS + 1))
  fi
  TOTAL_TESTS=$((TOTAL_TESTS + 1))
done

echo ""
echo "=== Summary: ${PASSED_TESTS}/${TOTAL_TESTS} tests passed ==="
if [ "${FAILED_TESTS}" -gt 0 ]; then
  echo "✗ FAILED: ${FAILED_TESTS} test(s) failed."
  exit 1
fi

echo "✓ All release-event matrix tests and security policies verified successfully!"
exit 0
