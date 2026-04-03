#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

has_error=0

check_pattern() {
  local label="$1"
  local pattern="$2"
  shift 2

  local args=(
    -n
    --hidden
    --glob=!**/node_modules/**
    --glob=!**/.git/**
    --glob=!scripts/check-legacy-boundary.sh
  )

  for exclude in "$@"; do
    args+=("--glob=!${exclude}")
  done

  local matches
  matches="$(rg "${args[@]}" -- "${pattern}" . || true)"
  if [[ -n "${matches}" ]]; then
    echo "forbidden ${label} references found:" >&2
    echo "${matches}" >&2
    has_error=1
  fi
}

check_pattern "codex-native alias" 'codex-native' 'repotests/plugin_manifest_lifecycle_integration_test.go'
check_pattern "Cursor legacy rules import" '\.cursorrules' 'docs/research/**'
check_pattern "OpenCode env-config compatibility" 'OPENCODE_CONFIG(_DIR)?'
check_pattern "legacy Gemini binary aliases" 'PLUGIN_KIT_AI_GEMINI_BIN|GEMINI_BIN'
check_pattern "Gemini migratedTo field" 'migratedTo|migrated_to' 'docs/research/**'
check_pattern "deleted maintainer docs tree" 'maintainer-docs' 'website/tools/quality/check-output.mjs'
check_pattern "removed migration guide slug" 'migrate-existing-config'

if [[ "${has_error}" -ne 0 ]]; then
  echo "legacy boundary check failed" >&2
  exit 1
fi

echo "legacy boundary intact"
