# codex-python-starter

Copy-first starter for Python teams that want the stable `codex-runtime` Notify path with the smallest possible repo-local setup.

## Who It Is For

- Teams wiring a local Codex plugin into an existing repo
- Python users who want `requirements.txt` plus a repo-local `.venv`
- Users who want the stable interpreted subset, not the Go-first production lane

## Prerequisites

- `plugin-kit-ai` installed
- Python `3.10+`
- Codex local runtime lane

## Runtime

- Platform: `codex-runtime`
- Runtime: `python`
- Entrypoint: `./bin/codex-python-starter`
- Execution mode: `launcher`
- Status: `public-stable`, repo-local interpreted subset

## First Run

```bash
plugin-kit-ai doctor .
plugin-kit-ai bootstrap .
plugin-kit-ai validate . --platform codex-runtime --strict
```

This starter keeps one canonical Python env story:

- `requirements.txt`
- repo-local `.venv`

`plugin-kit-ai bootstrap .` creates `.venv` when needed and installs `requirements.txt`.
If you prefer `uv`, `poetry`, or `pipenv`, keep using the stable runtime lane, but this starter stays opinionated on `requirements.txt` plus `.venv`.

## Local Smoke

```bash
./bin/codex-python-starter notify '{"client":"codex-tui"}'
```

## Stable Default

- `Notify`

Treat `plugin-kit-ai validate --strict` as the CI-grade readiness gate for this runtime lane.
This starter is for repo-local integration, not the official packaged Codex bundle lane.

## Target Files

- `launcher.yaml`: runtime and entrypoint for local Notify integration
- `targets/codex-runtime/package.yaml`: authored Codex runtime metadata
- `.codex/config.toml`: rendered managed Codex config

## Ship It

This starter already includes `.github/workflows/bundle-release.yml`.

```bash
plugin-kit-ai doctor .
plugin-kit-ai bootstrap .
plugin-kit-ai validate . --platform codex-runtime --strict
plugin-kit-ai bundle publish . --platform codex-runtime --repo owner/repo --tag v1.0.0
plugin-kit-ai bundle fetch owner/repo --tag v1.0.0 --platform codex-runtime --runtime python --dest ./handoff-plugin
```
