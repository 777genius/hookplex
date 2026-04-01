---
title: "notify"
description: "Capability reference for notify"
canonicalId: "capability:notify"
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

# notify

Эта capability показывает одно и то же поведение поперёк платформ. Открывайте её, когда важнее понять само действие, а не читать каждый platform tree отдельно.

## Коротко

- Платформ с этой capability: 2
- Связанных events: 2
- Текущий maturity: beta

## С чем читать вместе

- [Модель поведения платформ](/ru/concepts/platform-behavior-model), если нужно решить, подходит ли вам capability-first view.
- [Обещания поддержки по путям](/ru/reference/support-promise-by-path), если важнее promise и operational cost, а не само имя capability.

## Платформы

- [`claude`](/ru/api/platform-events/claude)
- [`codex`](/ru/api/platform-events/codex)

## Связанные runtime events

- `claude/Notification`
- `codex/Notify`

## Таблица покрытия

| Platform | Event | Maturity | Contract |
| --- | --- | --- | --- |
| claude | Notification | beta | runtime-supported but not stable |
| codex | Notify | stable | production-ready |
