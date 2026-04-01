---
title: "Примеры и рецепты"
description: "Путеводитель по публичным example repos, starter repos, локальным runtime references и skill examples в plugin-kit-ai."
canonicalId: "page:guide:examples-and-recipes"
section: "guide"
locale: "ru"
generated: false
translationRequired: true
---

# Примеры и рецепты

Используйте эту страницу, когда хотите увидеть, как `plugin-kit-ai` выглядит в реальных репозиториях, а не только в абстрактных объяснениях.

## Как правильно читать эту страницу

Примеры сами по себе не являются продуктом. Они показывают, как один managed plugin project может превращаться в разные output shapes.

- Начинайте с production example, если хотите увидеть законченный публичный репозиторий.
- Начинайте со starter repo, если нужен самый быстрый copy-first вход.
- Начинайте с local runtime reference, если хотите понять стабильный Node или Python local-runtime path.
- Открывайте skill examples только тогда, когда расширяете окружающий authoring workflow.

Если вам всё ещё кажется, что `plugin-kit-ai` это просто набор starter repos для Claude или Codex, сначала прочитайте [One Project, Multiple Targets](/ru/guide/one-project-multiple-targets), а потом вернитесь сюда.

## С чего лучше начать

- Нужен самый ясный finished runtime repo: откройте [`codex-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-basic-prod) или [`claude-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/claude-basic-prod).
- Хотите увидеть packaging и workspace-config outputs: откройте [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod), [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package), [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic) или [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic).
- Нужна самая чистая стартовая точка: откройте [каталог canonical starters](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters) и свяжите его со страницей [Выбор стартового репозитория](/ru/guide/choose-a-starter).
- Нужен стабильный interpreted runtime path: откройте [каталог repo-local runtimes](https://github.com/777genius/plugin-kit-ai/tree/main/examples/local).
- Нужны маленькие supporting patterns вокруг authoring workflow: откройте [каталог skill examples](https://github.com/777genius/plugin-kit-ai/tree/main/examples/skills).

## 1. Production plugin examples

Это самые наглядные примеры законченных публичных форм:

- [`codex-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-basic-prod): production repo для Codex runtime
- [`claude-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/claude-basic-prod): production repo для Claude
- [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod): target для Codex package
- [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package): packaging target для Gemini extension
- [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic): workspace-config target для Cursor
- [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic): workspace-config target для OpenCode

Читайте их, когда нужен:

- конкретный layout репозитория
- реальный пример rendered outputs
- честный публичный пример того, как выглядит “здоровый” проект

Эти примеры быстрее всего показывают, что `plugin-kit-ai` шире, чем один runtime. Одна и та же managed model может закончиться runtime plugin, package-style output или workspace-config integration.

## 2. Starter repos

Используйте starter repos, когда хотите начинать не с пустой директории, а с known-good baseline.

Они лучше всего подходят для:

- первого старта
- онбординга команды
- выбора между Go, Python, Node, Claude и Codex

Если вы ещё выбираете стартовую точку, свяжите это с [Выбором стартового репозитория](/ru/guide/choose-a-starter).

С чего лучше начать:

- [README каталога starters](https://github.com/777genius/plugin-kit-ai/blob/main/examples/starters/README.md): самый быстрый обзор всех canonical starters
- [`codex-go-starter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters/codex-go-starter): самый self-contained default для команд Codex
- [`claude-go-starter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters/claude-go-starter): самый self-contained default для команд Claude
- [`codex-python-starter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters/codex-python-starter): стабильный Python copy-first путь
- [`claude-node-typescript-starter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters/claude-node-typescript-starter): стабильный Node/TypeScript copy-first путь

Starter repos это входные точки, а не долгосрочная граница продукта.

## 3. Local runtime references

Каталог `examples/local` показывает локальные Python и Node runtime references.

Он полезен, когда:

- нужно глубже понять interpreted runtime story
- вы хотите сравнить JavaScript, TypeScript и Python local-runtime setups
- нужен конкретный reference сверх starter repos

С чего лучше начать:

- [README каталога local runtimes](https://github.com/777genius/plugin-kit-ai/blob/main/examples/local/README.md)
- [`codex-python-local`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/local/codex-python-local)
- [`codex-node-local`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/local/codex-node-local)
- [`codex-node-typescript-local`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/local/codex-node-typescript-local)

## 4. Skill examples

Каталог `examples/skills` показывает примеры skills и вспомогательных интеграций.

Это не главный entrypoint для большинства авторов плагинов, но он полезен, когда:

- вы хотите встроить docs, review или formatting helpers в более широкий workflow
- нужно понять, как соседние skills могут жить рядом с plugin repos

С чего лучше начать:

- [README skill examples](https://github.com/777genius/plugin-kit-ai/blob/main/examples/skills/README.md)
- [`go-command-lint`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/skills/go-command-lint)
- [`cli-wrapper-formatter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/skills/cli-wrapper-formatter)
- [`docs-only-review`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/skills/docs-only-review)

## Что читать по цели

- Нужен самый сильный runtime example: начните с production example для Codex или Claude, потом прочитайте [Плагин для команды](/ru/guide/team-ready-plugin).
- Нужны packaging или workspace-config examples: начните с примеров для Codex package, Gemini, Cursor или OpenCode, потом прочитайте [Package и workspace targets](/ru/guide/package-and-workspace-targets).
- Нужна чистая стартовая точка, а не finished example: идите в [Стартовые шаблоны](/ru/guide/starter-templates).
- Сначала нужно выбрать сам target: прочитайте [Выбор target](/ru/guide/choose-a-target).
- Нужно понять центральную модель, которая стоит за всеми этими примерами: прочитайте [Managed Project Model](/ru/concepts/managed-project-model).
- Сначала нужен общий обзор продукта: прочитайте [Что можно построить](/ru/guide/what-you-can-build).

## Главное правило

Examples должны прояснять публичный контракт, а не заменять его.

Используйте example repos, чтобы увидеть форму, layout и healthy outputs. Используйте остальную документацию, чтобы понять, что стабильно, что опционально и что проект действительно обещает.
