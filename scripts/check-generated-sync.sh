#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

before="$(mktemp)"
after="$(mktemp)"
trap 'rm -f "$before" "$after"' EXIT

generated_files=(
  "cli/hookplex/internal/scaffold/platforms_gen.go"
  "cli/hookplex/internal/validate/rules_gen.go"
  "sdk/hookplex/internal/descriptors/gen/completeness_gen_test.go"
  "sdk/hookplex/internal/descriptors/gen/registry_gen.go"
  "sdk/hookplex/internal/descriptors/gen/resolvers_gen.go"
  "sdk/hookplex/internal/descriptors/gen/support_gen.go"
  "docs/generated/support_matrix.md"
)

for f in "${generated_files[@]}"; do
  shasum "$f"
done >"$before"

GOCACHE="${GOCACHE:-/tmp/hookplex-gocache}" go run ./cmd/hookplex-gen >/tmp/hookplex-gen.out 2>/tmp/hookplex-gen.err

for f in "${generated_files[@]}"; do
  shasum "$f"
done >"$after"

if ! diff -u "$before" "$after"; then
  echo "generated files drifted; rerun generation and review tracked changes" >&2
  exit 1
fi

echo "generated artifacts in sync"
