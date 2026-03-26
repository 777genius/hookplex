# Generated Support Matrix

This generated table is the canonical per-event support contract for shipped runtime claims.

| Platform | Event | Status | Maturity | V1 Target | Invocation | Carrier | Transport Modes | Scaffold | Validate | Capabilities | Live Test | Summary |
|----------|-------|--------|----------|-----------|------------|---------|-----------------|----------|----------|--------------|-----------|---------|
| claude | Stop | runtime_supported | beta | true | argv_command_casefold | stdin_json | process | true | true | stop_gate | claude_cli | Claude Stop command hook |
| claude | PreToolUse | runtime_supported | beta | true | argv_command_casefold | stdin_json | process | true | true | tool_gate | claude_cli | Claude PreToolUse command hook |
| claude | UserPromptSubmit | runtime_supported | beta | true | argv_command_casefold | stdin_json | process | true | true | prompt_submit_gate | claude_cli | Claude UserPromptSubmit command hook |
| codex | Notify | runtime_supported | beta | true | argv_command | argv_json | process | true | true | notify | codex_notify | Codex notify hook |
