#!/usr/bin/env bash
# S4: every package must report 100.0% statement coverage from real execution.
set -euo pipefail
cd "$(dirname "$0")/.."
out=$(go test -cover ./... 2>&1)
echo "$out"
bad=$(echo "$out" | grep -v 'coverage: 100.0% of statements' || true)
if [ -n "$bad" ]; then echo "coverage-gate: FAIL (not 100%)"; exit 1; fi
echo "coverage-gate: OK"
