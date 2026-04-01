---
title: "Зачем plugin-kit-ai"
description: "Какую проблему решает plugin-kit-ai, кому он подходит и когда это плохой выбор."
canonicalId: "page:concepts:why-plugin-kit-ai"
section: "concepts"
locale: "ru"
generated: false
translationRequired: true
---

# Зачем plugin-kit-ai

`plugin-kit-ai` решает довольно конкретную задачу: командам нужны реальные репозитории плагинов с ясной границей поддержки, а не куча вручную правленных target-файлов и одноразовых вспомогательных скриптов.

## Что он даёт

- единая управляемая модель проекта вместо drift в target files
- один source of truth, который может рендерить несколько target’ов без распада на россыпь вручную поддерживаемых repo
- сильный путь по умолчанию на Go и стабильные локальные пути для Python и Node
- предсказуемые `render` и `validate` шаги
- generated API и support metadata, привязанные к реальным исходным данным

## Кому это подходит

- авторам плагинов, которым нужна более сильная структура, чем у ad-hoc local scripts
- командам, которые мигрируют с native target files на управляемую модель проекта
- maintainers, которым важны drift detection, strict validation и явные публичные границы

## Когда это плохой выбор

Скорее всего это не ваш инструмент, если:

- вам нужен только маленький одноразовый local script без намерения поддерживать структуру
- вы хотите универсальное управление зависимостями для всех interpreted runtime ecosystems
- вы хотите, чтобы каждый target и каждая hook family имели одинаковые гарантии стабильности

## Что в этом главное

`plugin-kit-ai` делает ставку на управляемость, а не на ad-hoc гибкость.

Практически это означает:

- один source of truth вместо drift между target files
- один понятный workflow через `render`, `validate` и CI
- один repo, который может расти до нескольких target’ов без потери структуры
- явную границу поддержки, по которой команде проще принимать инженерные решения

Прочитайте [Модель управляемого проекта](/ru/concepts/managed-project-model), если вам нужно самое короткое определение продукта.
Прочитайте [Один проект, несколько target’ов](/ru/guide/one-project-multiple-targets), если вам нужно product-level объяснение этой идеи.
Свяжите эту страницу с [Выбором runtime](/ru/concepts/choosing-runtime) и [Границей поддержки](/ru/reference/support-boundary).
