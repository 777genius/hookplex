---
title: "Словарь терминов"
description: "Канонические термины, которые используются по всей публичной документации plugin-kit-ai."
canonicalId: "page:reference:glossary"
section: "reference"
locale: "ru"
generated: false
translationRequired: true
---

# Словарь терминов

## Authored State

Исходное состояние проекта, которым владеет сам репозиторий. Оно живёт в package-standard layout, а `generate` превращает его в output-файлы для нужного target’а.

## Generated Target Files

Сгенерированные файлы для конкретного target’а. Это не предпочтительный долгосрочный источник истины.

## Path

Практический путь со своими правилами работы. Примеры: стандартный путь на Go, локальный Node/TypeScript runtime path и пути для workspace-config.

## Target

Тип результата или интеграции, в который вы целитесь: `codex-runtime`, `claude`, `codex-package`, `gemini`, `opencode` или `cursor`.

## Runtime Path

Путь, в котором проект напрямую владеет исполняемым поведением плагина. Поэтому выбор runtime, поведение обработчиков и strict validation здесь особенно важны.

## Package Or Extension Path

Путь, сфокусированный на корректной сборке package или extension artifacts, а не на локальном исполняемом плагине.

## Workspace-Config Path

Путь, где основным продуктом является конфигурация под управлением репозитория, а не исполняемый runtime-плагин.

## Wrapper Install Channel

Способ установить CLI, например через Homebrew, npm или PyPI. Это не public runtime API.

## Shared Runtime Package

Зависимость `plugin-kit-ai-runtime`, используемая в одобренных Python и Node сценариях вместо копирования helper-файлов в каждый репозиторий.

## Support Boundary

Публичная граница между тем, что проект считает stable, тем, что остаётся beta, и тем, что сознательно не входит в долгосрочное обещание.

## Readiness Gate

Команда или flow, который нужно воспринимать как публичный сигнал, что репозиторий в порядке. Для большинства проектов это `validate --strict`, часто вместе с `doctor` и `generate`.

## Handoff

Момент, когда repo, artifact или package уже можно передать другому человеку, другой машине или пользователю без скрытых договорённостей.

Свяжите этот словарь с [Моделью target’ов](/ru/concepts/target-model), [Границей поддержки](/ru/reference/support-boundary) и [Готовностью к продакшену](/ru/guide/production-readiness).
