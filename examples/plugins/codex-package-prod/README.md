# codex-package-prod

Reference Codex package repo for the official `codex-package` lane.

This example covers the official Codex plugin bundle only. It does not include `launcher.yaml` or repo-local notify integration.
It demonstrates first-class package metadata in `targets/codex-package/package.yaml`, structured prompt UX in `targets/codex-package/interface.json`, and an optional app manifest in `targets/codex-package/app.json`.

## Workflow

```bash
plugin-kit-ai normalize .
plugin-kit-ai render .
plugin-kit-ai render --check .
plugin-kit-ai validate . --platform codex-package --strict
```
