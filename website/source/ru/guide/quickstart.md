---
title: "Быстрый старт"
description: "Самый быстрый поддерживаемый путь к рабочему проекту на plugin-kit-ai."
canonicalId: "page:guide:quickstart"
section: "guide"
locale: "ru"
generated: false
translationRequired: true
---

# Быстрый старт

Это самый короткий поддерживаемый путь, если вам нужен один plugin repo, который потом можно расширять на другие поддерживаемые target’ы.

Сначала выберите один сильный путь. Расширение на Claude, Codex package, Gemini и другие типы выходов — это уже следующий шаг.

## Если читать только одно

Начинайте с Go по умолчанию, если вы уже заранее не знаете, что вам нужны Claude hooks, Node/TypeScript или Python.

Но не путайте стартовый путь с окончательной границей продукта: первый выбор не запрещает дальнейшее расширение репозитория.

## Рекомендуемый старт по умолчанию

Если у вас нет сильной причины выбрать другой путь, начинайте так:

```bash
brew install 777genius/homebrew-plugin-kit-ai/plugin-kit-ai
plugin-kit-ai version
plugin-kit-ai init my-plugin
cd my-plugin
plugin-kit-ai generate .
plugin-kit-ai validate . --platform codex-runtime --strict
```

Это даёт самый сильный путь по умолчанию: Go-репозиторий для Codex runtime, который проще всего проверять, передавать другим и потом расширять.

## Почему это путь по умолчанию

- один репозиторий с первого дня
- самая чистая история для продакшена
- самая простая база для расширения на другие поддерживаемые выходы

## Что расширяется потом

- Вы всё равно держите один репозиторий и один процесс проверки.
- Из того же repo можно рендерить поддерживаемые выходы для Claude, Codex, Gemini и других target’ов.
- Глубина поддержки зависит от target’а.
- Runtime plugins, package outputs и workspace-config targets не ведут себя одинаково.

## Как выбрать правильный путь

| Цель | Лучший стартовый путь |
| --- | --- |
| Самый сильный путь для продакшена | `codex-runtime` с `--runtime go` |
| Локальный Python plugin | `codex-runtime --runtime python` |
| Локальный TypeScript plugin | `codex-runtime --runtime node --typescript` |

`claude` выбирайте первым только тогда, когда hooks Claude уже являются реальным требованием продукта.

Package, extension и workspace-config targets безопаснее выбирать уже как следующий слой расширения после первого рабочего пути.

## Типовые первые команды

```bash
plugin-kit-ai init my-plugin --platform codex-runtime --runtime node --typescript
plugin-kit-ai doctor ./my-plugin
plugin-kit-ai bootstrap ./my-plugin
plugin-kit-ai generate ./my-plugin
plugin-kit-ai validate ./my-plugin --platform codex-runtime --strict
```

## Что важно знать перед выбором Python или Node

- Python и Node поддерживаются как полноценный путь для стабильного локального сценария.
- Но на машине, которая запускает плагин, всё равно должен быть установлен Python `3.10+` или Node.js `20+`.
- Go остаётся рекомендуемым путём по умолчанию, когда нужен самый чистый production и distribution story.

## Что читать дальше

- Переходите к [Первому плагину](/ru/guide/first-plugin), если хотите самый узкий рекомендуемый tutorial.
- Переходите к [Python runtime](/ru/guide/python-runtime), если команда Python-first и плагин остаётся локальным для репозитория.
- Переходите к [Что можно построить](/ru/guide/what-you-can-build), если хотите увидеть, как тот же репозиторий позже покрывает больше выходов.
- Переходите к [Выбору starter repo](/ru/guide/choose-a-starter), если хотите стартовать не с пустого repo, а с шаблона.
- Переходите к [Один проект, несколько target’ов](/ru/guide/one-project-multiple-targets), когда готовы расширяться дальше первого пути.
- Переходите к [Выбору target](/ru/guide/choose-a-target) только после того, как уже поняли базовую форму продукта.

См. [Выбор runtime](/ru/concepts/choosing-runtime) для модели выбора и [Установку](/ru/guide/installation) для каналов установки CLI.
