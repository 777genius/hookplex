---
title: "Диагностика проблем"
description: "Самые частые проблемы при установке, render, validate и bootstrap в plugin-kit-ai проектах."
canonicalId: "page:reference:troubleshooting"
section: "reference"
locale: "ru"
generated: false
translationRequired: true
---

# Диагностика проблем

## Что проверить первым

1. Запустите `plugin-kit-ai doctor <path>`.
2. Убедитесь, что выбран правильный target и runtime path.
3. Снова выполните `render`, прежде чем считать generated files корректными.
4. Воспринимайте `validate --strict` как сигнал, а не как шум.

## CLI установился, но не запускается

Проверьте, что binary действительно находится в shell `PATH`. Если вы ставили CLI через npm или PyPI, убедитесь, что пакет реально скачал опубликованный binary, а не воспринимайте сам пакет как runtime.

## Python или Node runtime-проекты падают слишком рано

Сначала проверьте сам runtime:

- Python runtime-проекты требуют Python `3.10+`
- Node runtime-проекты требуют Node.js `20+`

Используйте `plugin-kit-ai doctor <path>`, прежде чем считать, что сломан сам проект.

## Падает `validate --strict`

Воспринимайте это как сигнал, а не как шум. Смысл strict validation именно в том, чтобы ловить drift и readiness problems до того, как вы объявите проект здоровым.

Частые причины:

- generated artifacts устарели, потому что был пропущен `render`
- выбранная platform не соответствует исходному состоянию проекта
- выбранный runtime-путь требует bootstrap или исправления окружения

Что проверить первым:

- выполните `plugin-kit-ai render <path>`
- снова запустите `plugin-kit-ai validate <path> --platform <target> --strict`
- убедитесь, что target действительно совпадает с проектом, который вы собираете

## `render` выдаёт не то, что ожидалось

Обычно это значит, что исходное состояние проекта и ваша ментальная модель уже разошлись. Проверьте package-standard layout, а не редактируйте generated target files вручную в попытке “починить” output.

## Я не понимаю, какой путь выбрать

Начинайте с пути Go по умолчанию, если нужен самый сильный контракт. Переходите на Node/TypeScript или Python только тогда, когда компромисс локального runtime действительно осознан и нужен.

## У меня уже есть native config files

Используйте migration flow:

```bash
plugin-kit-ai import ./native-plugin --from codex-runtime
plugin-kit-ai normalize ./native-plugin
plugin-kit-ai render ./native-plugin
plugin-kit-ai validate ./native-plugin --platform codex-runtime --strict
```

## Когда перестать дебажить и вернуться к guide-разделу

- Если вы всё ещё не понимаете, нужен ли вам runtime, package или workspace-config output, вернитесь к [Выбору target](/ru/guide/choose-a-target).
- Если вы не уверены, нужно ли вам идти в Go, Node или Python, вернитесь к [Выбору runtime](/ru/concepts/choosing-runtime).
- Если вы не понимаете, начинать со starter или с `init`, вернитесь к [Выбору стартового репозитория](/ru/guide/choose-a-starter).
- Если repo всё ещё работает, но объявленный стандарт уже не совпадает с реальностью, идите в [Сигналы drift baseline](/ru/guide/baseline-drift-signals).

См. [Процесс авторинга](/ru/reference/authoring-workflow), [Частые вопросы](/ru/reference/faq) и [Миграцию существующего native config](/ru/guide/migrate-existing-config).
