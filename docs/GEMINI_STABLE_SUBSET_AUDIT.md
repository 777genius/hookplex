# Gemini Stable Subset Audit

This note records the current `public-stable` promotion boundary for the Gemini Go runtime lane.

## Promoted Stable Subset

The Gemini Go runtime is considered production-ready only for:

- `SessionStart`
- `SessionEnd`
- `BeforeModel`
- `AfterModel`
- `BeforeToolSelection`
- `BeforeAgent`
- `AfterAgent`
- `BeforeTool`
- `AfterTool`

## Beta Remainder

The following Gemini hooks remain `public-beta`:

- `Notification`
- `PreCompress`

Why they remain beta:

- `Notification` depends on Gemini system-alert behavior that is documented separately from the core runtime hook path and is still less suitable as a stable production gate.
- `PreCompress` is advisory-only and asynchronous, so its production evidence bar is intentionally higher than the current deterministic contract smoke.

## Promotion Evidence

The promoted Gemini stable subset is backed by:

- typed Go SDK surface in `sdk/gemini`
- descriptor-backed runtime metadata and generated support tables
- scaffolded Gemini Go runtime repos via `plugin-kit-ai init --platform gemini --runtime go`
- strict render/validate contract checks
- deterministic repo-local smoke via `make test-gemini-runtime-smoke`
- dedicated opt-in real CLI runtime smoke via `make test-gemini-runtime-live`

The deterministic Gemini smoke now covers:

- lifecycle and advisory input decoding
- runtime control semantics such as `deny`, `continue:false`, `systemMessage`, `clearContext`, `suppressOutput`
- runtime transform semantics such as request/response rewrite, tool selection config, turn-local context injection, tool-input rewrite, tool-result context, and tail tool calls
- tool payload observability including `tool_input`, `tool_response`, `mcp_context`, and `original_request_name`

## Promotion Rule

Future Gemini hooks are not stable by default.

Any new Gemini runtime surface stays `public-beta` until it has:

- descriptor-backed metadata
- scaffold and validate alignment
- deterministic smoke coverage
- production docs alignment
- sufficient live evidence for the intended stable promise
