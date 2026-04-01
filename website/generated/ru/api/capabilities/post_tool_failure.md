---
title: "post_tool_failure"
description: "Capability reference for post_tool_failure"
canonicalId: "capability:post_tool_failure"
surface: "capabilities"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-beta"
maturity: "beta"
sourceRef: "docs/generated/support_matrix.md"
translationRequired: false
---
<DocMetaCard surface="capabilities" stability="public-beta" maturity="beta" source-ref="docs/generated/support_matrix.md" source-href="https://github.com/777genius/plugin-kit-ai/blob/main/docs/generated/support_matrix.md" />

# post_tool_failure

Эта capability показывает одно и то же поведение поперёк платформ. Открывайте её, когда важнее понять само действие, а не читать каждый platform tree отдельно.

## Коротко

- Платформ с этой capability: 1
- Связанных events: 1
- Текущий maturity: beta

## Платформы

- [`claude`](/ru/api/platform-events/claude)

## Связанные runtime events

- `claude/PostToolUseFailure`

## Таблица покрытия

| Platform | Event | Maturity | Contract |
| --- | --- | --- | --- |
| claude | PostToolUseFailure | beta | runtime-supported but not stable |
