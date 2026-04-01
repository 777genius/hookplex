---
title: "worktree_remove"
description: "Capability reference for worktree_remove"
canonicalId: "capability:worktree_remove"
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

# worktree_remove

Эта capability показывает одно и то же поведение поперёк платформ. Открывайте её, когда важнее понять само действие, а не читать каждый platform tree отдельно.

## Коротко

- Платформ с этой capability: 1
- Связанных events: 1
- Текущий maturity: beta

## С чем читать вместе

- [Модель поведения платформ](/ru/concepts/platform-behavior-model), если нужно решить, подходит ли вам capability-first view.
- [Обещания поддержки по путям](/ru/reference/support-promise-by-path), если важнее promise и operational cost, а не само имя capability.

## Платформы

- [`claude`](/ru/api/platform-events/claude)

## Связанные runtime events

- `claude/WorktreeRemove`

## Таблица покрытия

| Platform | Event | Maturity | Contract |
| --- | --- | --- | --- |
| claude | WorktreeRemove | beta | runtime-supported but not stable |
