---
title: "Что можно построить"
description: "Широкий публичный обзор реальных сценариев и форм продукта, которые поддерживает plugin-kit-ai."
canonicalId: "page:guide:what-you-can-build"
section: "guide"
locale: "ru"
generated: false
translationRequired: true
---

# Что можно построить

Эта страница объясняет главное обещание продукта: начинайте с одного репозитория, а затем расширяйте тот же репозиторий на новые поддерживаемые выходы по мере роста продукта.

<MermaidDiagram
  :chart="`
flowchart TD
  Product[One repo] --> Runtime[Runtime plugins]
  Product --> Multi[More supported outputs later]
  Product --> Bundle[Portable bundle handoff]
  Product --> Shared[Shared runtime package]
  Product --> Package[Package and extension targets]
  Product --> Workspace[Workspace config targets]
`"
/>

## 1. Один repo, много поддерживаемых выходов

Это главное обещание продукта.

- Начинайте с одного plugin repo.
- Добавляйте только те outputs, которые реально нужны дальше.
- Держите один процесс через `render`, `validate` и CI.
- Не ожидайте одинаковой глубины runtime-поддержки у всех target’ов.

Безопасная mental model здесь такая:

- один repo
- один процесс
- много поддерживаемых выходов
- разная глубина поддержки по target’ам

## 2. Начните с самого сильного первого репозитория

Большинству команд стоит начинать с Codex runtime на Go.

Такой первый репозиторий даёт:

- самый сильный старт для продакшена
- один процесс работы с репозиторием вместо ручного редактирования target-файлов
- ясный путь через `render` и `validate --strict`

Если стек уже определён заранее, та же модель первого репозитория поддерживает:

- Go для самого сильного стандартного продакшен-контракта
- Node/TypeScript для основного стабильного non-Go пути
- Python для команд, которые осознанно остаются на локальном Python runtime

## 3. Добавляйте Claude, когда hooks действительно нужны

Используйте Claude-путь, когда Claude hooks действительно являются требованием продукта.

Это правильный выбор, если:

- вам нужны hooks именно Claude
- стабильного подмножества Claude достаточно для вашего плагина
- нужен более сильный и предсказуемый процесс авторинга, чем при ручной правке native files

## 4. Расширяйте тот же репозиторий дальше

После первого рабочего репозитория тот же проект можно расширить до:

- выходов для Claude hooks
- выходов для Codex package
- packaging для Gemini
- workspace/config outputs для OpenCode и Cursor
- portable bundle delivery для поддерживаемых Python и Node репозиториев

В этом и состоит реальная cross-target история: один репозиторий, один процесс, больше поддерживаемых выходов со временем.

## 5. Репозитории плагинов, готовые для команды

`plugin-kit-ai` — это не только scaffolding. Это ещё и путь к репозиторию, который другой коллега может понять, проверить и использовать без скрытых договорённостей.

Это означает, что система поддерживает:

- строгие проверки готовности
- понятные сценарии для CI
- явный выбор пути и target’а
- предсказуемый handoff между авторами и downstream-пользователями

## 6. Portable bundle handoff для Python и Node

Для поддерживаемых Python и Node путей можно выйти за пределы локального authoring и собирать portable bundle artifacts для handoff.

Это важно, когда:

- модель поставки требует скачиваемые артефакты вместо live repo
- нужен более чистый сценарий установки для downstream-пользователей Python и Node путей
- вы используете bundle publish/fetch flow как часть release handoff

Подробный public flow описан в [Bundle handoff](/ru/guide/bundle-handoff).

## 7. Shared runtime package

Python и Node helper-логика может жить либо:

- в vendored helper files внутри repo
- в общем `plugin-kit-ai-runtime` package

Это даёт поддерживаемый путь для:

- reusable runtime helpers на несколько repo
- более чистые обновления зависимостей
- стандартизированного helper API без ручного копирования scaffolded files

## 8. Targets для package, extension и workspace-config

Не каждая публичная форма — это локальный runtime-плагин внутри репозитория.

`plugin-kit-ai` также покрывает:

- packaging-oriented lanes
- extension-style targets
- workspace-config integration targets

Эти target’ы важны, когда конечный продукт — это packaging или configuration, а не исполняемый плагин.

Перед выбором этих путей прочитайте [Package и workspace targets](/ru/guide/package-and-workspace-targets).

## 9. Читайте в таком порядке

Если вы ещё решаете, что именно делать:

1. прочитайте эту страницу
2. используйте [Быстрый старт](/ru/guide/quickstart) или [Выбор starter repo](/ru/guide/choose-a-starter)
3. прочитайте [Один проект, несколько target’ов](/ru/guide/one-project-multiple-targets), когда нужна честная картина расширения
4. идите в [Модель target’ов](/ru/concepts/target-model), только когда уже нужен точный technical split

Свяжите эту страницу с [Примерами и рецептами](/ru/guide/examples-and-recipes), [Выбором starter repo](/ru/guide/choose-a-starter), [Выбором delivery model](/ru/guide/choose-delivery-model), [Bundle handoff](/ru/guide/bundle-handoff), [Package и workspace targets](/ru/guide/package-and-workspace-targets) и [API поверхностями](/ru/api/).
