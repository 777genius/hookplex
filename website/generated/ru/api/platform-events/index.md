---
title: "Platform Events"
description: "Generated platform event reference"
canonicalId: "page:api:platform-events:index"
surface: "platform-events"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "docs/generated/support_matrix.md"
translationRequired: false
---
# События платформ

Эта зона показывает event surfaces по платформам и помогает не смешивать stable lane с более широкой beta runtime coverage.

- Открывайте её, когда уже знаете target и хотите увидеть event-level contract.
- Используйте `Capabilities`, когда нужен cross-platform взгляд вместо platform-first view.

## С чего лучше начать

- Нужна платформа как основная ось выбора: начинайте отсюда.
- Нужно сравнить поведение между платформами: переходите в `Capabilities`.
- Если target ещё не выбран, сначала вернитесь в guides.

## Когда не нужно начинать отсюда

- Если вы ещё не выбрали target, сначала прочитайте `/guide/choose-a-target` и `/concepts/target-model`.

## С чем читать вместе

- [Модель поведения платформ](/ru/concepts/platform-behavior-model), если нужно понять, когда читать platform-first, а когда capability-first.
- [Обещания поддержки по путям](/ru/reference/support-promise-by-path), если сначала нужен decision sheet по силе promise.

## Карта платформ

| Платформа | Events | Stable | Beta | Capabilities |
| --- | --- | --- | --- | --- |
| claude | 18 | 3 | 15 | 18 |
| codex | 1 | 1 | 0 | 1 |

## Переходите по платформам

- [`claude`](/ru/api/platform-events/claude)
- [`codex`](/ru/api/platform-events/codex)
