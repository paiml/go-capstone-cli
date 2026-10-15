#!/usr/bin/env bash
# S4 falsifier: a test that calls nothing must report 0.0%.
set -euo pipefail
d=$(mktemp -d)
trap 'rm -rf "$d"' EXIT
mkdir "$d/f"
cp "$(dirname "$0")/../config/config.go" "$d/f/config.go"
printf 'module falsifier\n\ngo 1.21\n' > "$d/go.mod"
printf 'package config\n\nimport "testing"\n\nfunc TestNothing(t *testing.T) {}\n' > "$d/f/nothing_test.go"
out=$(cd "$d" && go test -cover ./... 2>&1)
echo "$out"
echo "$out" | grep -q 'coverage: 0.0% of statements' || { echo "falsifier: FAIL (gate cannot see 0%)"; exit 1; }
echo "falsifier: OK (empty test reports 0.0%)"
