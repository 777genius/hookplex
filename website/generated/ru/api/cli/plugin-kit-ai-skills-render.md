---
title: "plugin-kit-ai skills render"
description: "Render Claude/Codex artifacts from canonical SKILL.md files"
canonicalId: "command:plugin-kit-ai:skills:render"
surface: "cli"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "cli:plugin-kit-ai skills render"
translationRequired: false
---
<DocMetaCard surface="cli" stability="public-stable" maturity="stable" source-ref="cli:plugin-kit-ai skills render" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/cli/plugin-kit-ai" />

# plugin-kit-ai skills render

Сгенерировано из реального Cobra command tree.

Render Claude/Codex artifacts from canonical SKILL.md files

## plugin-kit-ai skills render

Render Claude/Codex artifacts from canonical SKILL.md files

```
plugin-kit-ai skills render [path] [flags]
```

### Examples

```
  plugin-kit-ai skills render . --target all
  plugin-kit-ai skills render ./examples/skills/cli-wrapper-formatter --target codex
```

### Опции

```
  -h, --help            справка по render
      --target string   render target ("all", "claude", "codex") (default "all")
```

### См. также

* plugin-kit-ai skills	 - Экспериментальные инструменты для авторинга skills.
