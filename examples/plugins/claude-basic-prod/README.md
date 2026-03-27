# claude-basic-prod

Reference Claude plugin repo for the current `plugin-kit-ai` production workflow.

Stable runtime promise:

- `Stop`
- `PreToolUse`
- `UserPromptSubmit`

Additional generated Claude hooks in this example are runtime-supported but not stable in the current public contract.

## Workflow

```bash
plugin-kit-ai normalize .
plugin-kit-ai render .
plugin-kit-ai render --check .
plugin-kit-ai validate . --platform claude --strict
go test ./...
go build -o bin/claude-basic-prod ./cmd/claude-basic-prod
printf '%s' '{"session_id":"s","cwd":"/tmp","hook_event_name":"Stop"}' | ./bin/claude-basic-prod Stop
```
