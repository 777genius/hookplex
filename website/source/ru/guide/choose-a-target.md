---
title: "Выбор target"
description: "Практический публичный гид по выбору между Codex runtime, Claude, Codex package, Gemini, OpenCode и Cursor."
canonicalId: "page:guide:choose-a-target"
section: "guide"
locale: "ru"
generated: false
translationRequired: true
---

# Выбор target

Используйте эту страницу, когда вы уже понимаете, что хотите работать с `plugin-kit-ai`, но ещё выбираете target под свой продукт.

## Выбор за 60 секунд

- выбирайте `codex-runtime`, когда нужен исполняемый плагин с самым сильным путём по умолчанию
- выбирайте `claude`, когда реальным требованием продукта являются Claude hooks
- выбирайте `codex-package`, когда нужен package output для Codex, а не репозиторий с исполняемым плагином
- выбирайте `gemini`, когда нужен пакет расширения для Gemini
- выбирайте `opencode` или `cursor`, когда репозиторий должен владеть workspace configuration, а не исполняемым плагином

## Главное различие

- Выбирайте **runtime target**, когда репозиторий должен напрямую исполнять plugin logic.
- Выбирайте **package или extension target**, когда итогом должен быть артефакт для установки или публикации.
- Выбирайте **workspace-config target**, когда репозиторий должен владеть integration files и конфигурацией, а не исполняемым поведением плагина.

## Лучшие варианты по умолчанию

- Лучший default для реального исполняемого плагина: `codex-runtime`
- Лучший default, когда первым реальным требованием являются Claude hooks: `claude`
- Лучший package-style target: `codex-package`
- Лучший extension-style target: `gemini`
- Лучшие workspace-config targets: `cursor` или `opencode`

## Краткий справочник по target’ам

| Target | Когда выбирать | Чем он не является |
| --- | --- | --- |
| `codex-runtime` | Нужен основной путь для исполняемого плагина | Это не packaging-only target |
| `claude` | Нужны именно Claude hooks | Это не основной путь для Codex |
| `codex-package` | Нужен package output для Codex | Это не runtime-плагин |
| `gemini` | Вы выпускаете пакет расширения для Gemini | Это не основной runtime-путь |
| `opencode` | Нужна repo-owned OpenCode workspace config | Это не runtime-плагин |
| `cursor` | Нужна repo-owned Cursor workspace config | Это не runtime-плагин |

## Быстрое дерево решений

1. Нужно, чтобы репозиторий исполнял plugin logic? Выбирайте `codex-runtime` или `claude`.
2. Нужен package или extension artifact вместо исполняемого плагина? Выбирайте `codex-package` или `gemini`.
3. Нужны repo-owned editor или tool integration files? Выбирайте `cursor` или `opencode`.
4. Не уверены? Начинайте с `codex-runtime`.

## С чего начинать по цели

- Нужен самый сильный путь по умолчанию для реального plugin repo: начинайте с `codex-runtime`, потом прочитайте [Выбор стартового репозитория](/ru/guide/choose-a-starter).
- Нужны Claude hooks: начинайте с `claude`, потом прочитайте [Стартовые шаблоны](/ru/guide/starter-templates).
- Нужны package или extension artifacts: начинайте с `codex-package` или `gemini`, потом прочитайте [Package и workspace targets](/ru/guide/package-and-workspace-targets).
- Нужна workspace config под управлением repo: начинайте с `opencode` или `cursor`, потом прочитайте [Примеры и рецепты](/ru/guide/examples-and-recipes).

## Лучший первый пример для каждого target

- `codex-runtime`: [`codex-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-basic-prod)
- `claude`: [`claude-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/claude-basic-prod)
- `codex-package`: [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod)
- `gemini`: [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package)
- `cursor`: [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic)
- `opencode`: [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic)

## Безопасный выбор по умолчанию

Если вы не уверены, начинайте с `codex-runtime` и стандартного Go path.

Это даёт самую чистую стартовую точку для продакшена, прежде чем вы пойдёте в более узкий или специализированный target.

## Что читать дальше

- Прочитайте [Managed Project Model](/ru/concepts/managed-project-model), если хотите сначала зафиксировать центральную модель продукта.
- Прочитайте [Выбор runtime](/ru/concepts/choosing-runtime), если вы выбрали runtime target и ещё решаете между Go, Python, Node и shell.
- Прочитайте [Package и workspace targets](/ru/guide/package-and-workspace-targets), если выбираете между packaging и workspace-config targets.
- Прочитайте [Примеры и рецепты](/ru/guide/examples-and-recipes), если хотите увидеть реальные repos для разных форм продукта.
- Прочитайте [Поддержку target’ов](/ru/reference/target-support), если вам нужна компактная support matrix.
