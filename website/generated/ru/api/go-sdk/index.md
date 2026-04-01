---
title: "Go SDK"
description: "Generated Go SDK package reference"
canonicalId: "page:api:go-sdk:index"
surface: "go-sdk"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "sdk"
translationRequired: false
---
# Go SDK

Go SDK — рекомендуемый путь по умолчанию, когда нужен самый сильный и предсказуемый контракт для продакшена.

- Открывайте эту зону, когда строите production-oriented plugin на Go.
- Это лучший старт, если вы хотите минимальную зависимость от внешних runtime на машинах пользователей.
- Если вы ещё выбираете между Go, Python и Node, начните с `/guide/what-you-can-build` и `/concepts/choosing-runtime`.

| Package | Summary |
| --- | --- |
| [`sdk`](/ru/api/go-sdk/sdk) | Корневой composition/runtime entry package. |
| [`claude`](/ru/api/go-sdk/claude) | Публичные Claude-oriented handlers и event wiring. |
| [`codex`](/ru/api/go-sdk/codex) | Публичные Codex-oriented handlers и runtime integration. |
| [`platformmeta`](/ru/api/go-sdk/platformmeta) | Platform metadata и support-oriented helpers. |
