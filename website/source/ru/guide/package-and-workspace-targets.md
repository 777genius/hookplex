---
title: "Package и workspace targets"
description: "Как использовать Codex package, Gemini, OpenCode и Cursor, не путая их с путями для исполняемых плагинов."
canonicalId: "page:guide:package-and-workspace-targets"
section: "guide"
locale: "ru"
generated: false
translationRequired: true
---

# Package и workspace targets

Не каждый target в `plugin-kit-ai` является путём для исполняемого плагина.

Читайте эту страницу перед выбором `codex-package`, `gemini`, `opencode` или `cursor`, потому что эти targets решают другие задачи, чем `codex-runtime` и `claude`.

## Выбор за 60 секунд

- выбирайте `codex-runtime` или `claude`, когда продуктом является исполняемый плагин
- выбирайте `codex-package` или `gemini`, когда продуктом являются package или extension artifacts
- выбирайте `opencode` или `cursor`, когда продуктом является конфигурация workspace внутри репозитория

## Короткое правило

- runtime target = исполняемое поведение плагина
- package или extension target = артефакт для публикации или установки
- workspace-config target = repo-owned integration files и конфигурация

## Лучшие варианты по умолчанию

- Лучший package-style default: `codex-package`
- Лучший extension-style default: `gemini`
- Лучшие workspace-config defaults: `cursor` или `opencode`

## Лучшие первые примеры

- `codex-package`: [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod)
- `gemini`: [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package)
- `cursor`: [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic)
- `opencode`: [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic)

## Быстрое дерево решений

1. Нужен артефакт для публикации или установки? Выбирайте `codex-package` или `gemini`.
2. Нужно, чтобы repo владел editor или tool integration files? Выбирайте `cursor` или `opencode`.
3. Нужно, чтобы сам repo исполнял plugin logic? Остановитесь и вернитесь к `codex-runtime` или `claude`.

## Codex Package

Используйте `codex-package`, когда конечным результатом должен быть package для Codex, а не репозиторий с исполняемым плагином.

Это полезно, когда:

- packaging и есть реальный контракт поставки
- вам нужно, чтобы исходное состояние проекта оставалось управляемым в одном месте
- не нужно притворяться, что у этого target тот же runtime contract, что и у `codex-runtime`

Лучший первый пример: [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod)

## Gemini

Используйте `gemini`, когда цель — пакет расширения для Gemini CLI.

Этот target специально ориентирован на packaging.

Его правильно воспринимать так:

- это полноценный extension-packaging path через `render`, `import` и `validate`
- это не основной runtime-путь
- его выбирают, когда Gemini extension artifacts и есть конечный продукт

Лучший первый пример: [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package)

## OpenCode

Используйте `opencode`, когда репозиторий должен владеть конфигурацией OpenCode workspace и связанными project assets.

Этот target важен, когда:

- проекту нужен управляемый `opencode.json`
- репозиторий должен владеть workspace-level MCP и config shape
- нужен документированный путь авторинга конфигурации вместо ручной правки файлов

Но не путайте это с самым сильным runtime contract.

Лучший первый пример: [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic)

## Cursor

Используйте `cursor`, когда репозиторий должен управлять конфигурацией Cursor workspace.

Документированный subset включает:

- `.cursor/mcp.json`
- `.cursor/rules/**` в корне проекта
- optional shared root `AGENTS.md`

Это target для workspace-config, а не основной runtime-путь.

Лучший первый пример: [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic)

## Практическое правило выбора

Выбирайте эти targets, когда результатом проекта являются:

- package artifacts
- extension packaging
- workspace config

Не выбирайте их только потому, что название похоже на runtime-путь.

Если вам на самом деле нужно исполняемое поведение плагина, вернитесь к [Выбору runtime](/ru/concepts/choosing-runtime) и начинайте оттуда.

## Чего не нужно ожидать

- Не ожидайте, что `codex-package` или `gemini` будут вести себя как исполняемые runtime plugins.
- Не ожидайте, что `cursor` или `opencode` заменят основной runtime path, если вам на самом деле нужна plugin logic.
- Не выбирайте эти targets только потому, что имя экосистемы знакомо. Выбирайте их только тогда, когда именно output shape является реальным требованием продукта.

## Правило готовности

Для этих targets правило здорового репозитория остаётся тем же:

- исходное состояние проекта живёт в package-standard layout
- rendered files являются outputs
- `render --check` и `validate --strict` остаются главными проверками

## Что читать дальше

- Прочитайте [Выбор target](/ru/guide/choose-a-target), если сначала нужна полная карта target’ов.
- Прочитайте [Примеры и рецепты](/ru/guide/examples-and-recipes), если хотите сравнить runtime, package и workspace examples бок о бок.
- Прочитайте [Поддержку target’ов](/ru/reference/target-support), если нужна компактная support matrix.

## Что читать вместе с этим

Читайте эту страницу вместе с [Моделью target’ов](/ru/concepts/target-model), [Поддержкой target’ов](/ru/reference/target-support) и [Границей поддержки](/ru/reference/support-boundary).
