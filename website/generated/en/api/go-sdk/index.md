---
title: "Go SDK"
description: "Generated Go SDK package reference"
canonicalId: "page:api:go-sdk:index"
surface: "go-sdk"
section: "api"
locale: "en"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "sdk"
translationRequired: false
---
# Go SDK

The Go SDK is the recommended default path when you want the strongest production contract.

- Open this area when you are building a production-oriented Go plugin.
- This is the best starting point when you want the least downstream runtime friction.
- If you are still choosing between Go, Python, and Node, start with `/guide/what-you-can-build` and `/concepts/choosing-runtime`.

## Best First Stops

- Start with `sdk` if you want the root composition and runtime entry surface.
- Open `claude` when you are building Claude-oriented handlers.
- Open `codex` when you are building Codex-oriented handlers and runtime integration.

| Package | Summary |
| --- | --- |
| [`sdk`](/en/api/go-sdk/sdk) | Root composition and runtime entry package. |
| [`claude`](/en/api/go-sdk/claude) | Public Claude-oriented handlers and event wiring. |
| [`codex`](/en/api/go-sdk/codex) | Public Codex-oriented handlers and runtime integration. |
| [`platformmeta`](/en/api/go-sdk/platformmeta) | Platform metadata and support-oriented helpers. |
