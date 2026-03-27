# codex-basic-prod

Reference Codex plugin repo for the current `plugin-kit-ai` production workflow.

Stable runtime promise:

- `Notify`

## Workflow

```bash
plugin-kit-ai normalize .
plugin-kit-ai render .
plugin-kit-ai render --check .
plugin-kit-ai validate . --platform codex --strict
go test ./...
go build -o bin/codex-basic-prod ./cmd/codex-basic-prod
./bin/codex-basic-prod notify '{"client":"codex-tui"}'
```
