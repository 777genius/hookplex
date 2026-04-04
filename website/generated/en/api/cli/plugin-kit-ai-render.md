---
title: "plugin-kit-ai render"
description: "Compile native target artifacts from the package graph"
canonicalId: "command:plugin-kit-ai:render"
surface: "cli"
section: "api"
locale: "en"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "cli:plugin-kit-ai render"
translationRequired: false
---
<DocMetaCard surface="cli" stability="public-stable" maturity="stable" source-ref="cli:plugin-kit-ai render" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/cli/plugin-kit-ai" />

# plugin-kit-ai render

Generated from the live Cobra command tree.

Compile native target artifacts from the package graph

## plugin-kit-ai render

Compile native target artifacts from the package graph

### Synopsis

Compile native target artifacts from the package graph discovered via plugin.yaml and standard directories.

Claude and Codex runtime/package lanes render their managed native artifacts from the package graph.
Gemini rendering always produces the native extension package artifacts and may also carry the optional Go runtime lane when the authored project includes it; that lane now exposes a production-ready 9-hook runtime surface, but it still does not imply blanket runtime parity for future hooks beyond the promoted contract.
OpenCode rendering is workspace-config-only: it produces opencode.json plus mirrored skills, commands, agents, themes, local plugin code, and plugin-local package metadata without introducing a launcher/runtime contract.
Cursor rendering is workspace-config-only: it produces .cursor/mcp.json, mirrored .cursor/rules/**, and optional root AGENTS.md without introducing a launcher/runtime contract.

```
plugin-kit-ai render [path] [flags]
```

### Options

```
      --check           fail if generated artifacts are out of date
  -h, --help            help for render
      --target string   render target ("all", "claude", "codex-package", "codex-runtime", "gemini", "opencode", or "cursor") (default "all")
```

### SEE ALSO

* plugin-kit-ai	 - plugin-kit-ai CLI - scaffold and tooling for AI plugins
