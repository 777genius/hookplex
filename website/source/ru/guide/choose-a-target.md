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

## Короткое правило

- выбирайте `codex-runtime`, когда нужен исполняемый плагин с самым сильным путём по умолчанию
- выбирайте `claude`, когда реальным требованием продукта являются Claude hooks
- выбирайте `codex-package`, когда нужен package output для Codex, а не репозиторий с исполняемым плагином
- выбирайте `gemini`, когда нужен пакет расширения для Gemini
- выбирайте `opencode` или `cursor`, когда репозиторий должен владеть workspace configuration, а не исполняемым плагином

## Краткий справочник по target’ам

| Target | Когда выбирать | Чем он не является |
| --- | --- | --- |
| `codex-runtime` | Нужен основной путь для исполняемого плагина | Это не packaging-only target |
| `claude` | Нужны именно Claude hooks | Это не основной путь для Codex |
| `codex-package` | Нужен package output для Codex | Это не runtime-плагин |
| `gemini` | Вы выпускаете пакет расширения для Gemini | Это не основной runtime-путь |
| `opencode` | Нужна repo-owned OpenCode workspace config | Это не runtime-плагин |
| `cursor` | Нужна repo-owned Cursor workspace config | Это не runtime-плагин |

## С чего начинать по цели

- Нужен самый сильный путь по умолчанию для реального plugin repo: начинайте с `codex-runtime`
- Нужны Claude hooks: начинайте с `claude`
- Нужны package или extension artifacts: начинайте с `codex-package` или `gemini`
- Нужна workspace config под управлением repo: начинайте с `opencode` или `cursor`

## Безопасный выбор по умолчанию

Если вы не уверены, начинайте с `codex-runtime` и стандартного Go path.

Это даёт самую чистую стартовую точку для продакшена, прежде чем вы пойдёте в более узкий или специализированный target.

## Что читать дальше

- Прочитайте [Выбор runtime](/ru/concepts/choosing-runtime), если вы выбрали runtime target и ещё решаете между Go, Python, Node и shell.
- Прочитайте [Package и workspace targets](/ru/guide/package-and-workspace-targets), если выбираете между packaging и workspace-config targets.
- Прочитайте [Примеры и рецепты](/ru/guide/examples-and-recipes), если хотите увидеть реальные repos для разных форм продукта.
- Прочитайте [Поддержку target’ов](/ru/reference/target-support), если вам нужна компактная support matrix.
