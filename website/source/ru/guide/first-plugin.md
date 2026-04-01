---
title: "Соберите первый плагин"
description: "Минимальный пошаговый сценарий от init до строгой проверки готовности."
canonicalId: "page:guide:first-plugin"
section: "guide"
locale: "ru"
generated: false
translationRequired: true
---

# Соберите первый плагин

Этот гайд использует самый сильный путь по умолчанию и специально держит сценарий узким:

- target: `codex-runtime`
- runtime: `go`
- readiness gate: `validate --strict`

## 1. Установите CLI

```bash
brew install 777genius/homebrew-plugin-kit-ai/plugin-kit-ai
plugin-kit-ai version
```

## 2. Создайте проект

```bash
plugin-kit-ai init my-plugin
cd my-plugin
```

Путь `init` по умолчанию уже является рекомендуемой стартовой точкой для продакшена.

## 3. Сгенерируйте target-файлы

```bash
plugin-kit-ai render .
```

Не редактируйте сгенерированные target-файлы вручную как главный источник истины. Держите исходное состояние проекта в package-standard layout.

## 4. Прогоните проверку готовности

```bash
plugin-kit-ai validate . --platform codex-runtime --strict
```

Используйте это как главную проверку готовности для локального проекта.

## 5. Когда менять путь

Переходите на другой путь только когда это действительно нужно:

- выбирайте `claude` для плагинов Claude
- выбирайте `--runtime node --typescript` для основного стабильного пути без Go
- выбирайте `--runtime python`, когда проект остаётся локальным для репозитория, а команда осознанно Python-first
- выбирайте `codex-package`, `gemini`, `opencode` или `cursor`, только если ваша модель поставки действительно требует эти target’ы

## Следующие шаги

- Прочитайте [Выбор runtime](/ru/concepts/choosing-runtime), прежде чем уходить с пути по умолчанию.
- Используйте [Стартовые шаблоны](/ru/guide/starter-templates), когда нужен проверенный пример репозитория.
- Откройте [Справочник CLI](/ru/api/cli/), когда нужно точное поведение команд.
