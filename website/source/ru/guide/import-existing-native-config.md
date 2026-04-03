---
title: "Импорт существующего native config"
description: "Перевод hand-managed native config в package-standard модель проекта."
canonicalId: "page:guide:import-existing-native-config"
section: "guide"
locale: "ru"
generated: false
translationRequired: true
---

# Импорт существующего native config

Используйте этот путь, когда у вас уже есть native Claude, Codex, Gemini, OpenCode или Cursor config и вы хотите привести его к package-standard модели проекта.

## Форма импорта

```bash
plugin-kit-ai import ./native-plugin --from codex-runtime
plugin-kit-ai normalize ./native-plugin
plugin-kit-ai render ./native-plugin
plugin-kit-ai validate ./native-plugin --platform codex-runtime --strict
```

## Цель импорта

Цель не в том, чтобы держать native files как долгосрочный источник истины. Цель — перейти к модели проекта, которой управляет сам репозиторий, и позволить `render` детерминированно выпускать target artifacts.

## Хорошая дисциплина импорта

- import один раз, чтобы зафиксировать модель проекта
- normalize, когда нужно привести проект к package-standard shape
- render, чтобы заново выпустить target artifacts из исходного состояния проекта
- validate строго, прежде чем доверять импортированному проекту

## Когда это действительно полезно

- в команде уже есть drift в native config
- нужен единый источник истины на уровне репозитория
- хотите перестать вручную править target artifacts так, как будто это и есть модель проекта

См. [Границу поддержки](/ru/reference/support-boundary) и [CLI Reference](/ru/api/cli/) для формальных правил этого flow.
