---
title: "Политика версий и совместимости"
description: "Поймите, как plugin-kit-ai трактует stable и beta поверхности, каналы установки, runtime helpers и публичные release baselines."
canonicalId: "page:reference:version-and-compatibility"
section: "reference"
locale: "ru"
generated: false
translationRequired: true
---

# Политика версий и совместимости

Открывайте эту страницу, когда нужен один компактный публичный ответ на практический вопрос: что команде считать текущим baseline, что действительно несёт compatibility promise, а что нужно читать как guidance, а не как гарантию.

## Выбор за 60 секунд

- Нужен текущий user-facing baseline: начните с самого свежего подходящего [release note](/ru/releases/), сейчас это [v1.0.6](/ru/releases/v1-0-6).
- Нужна самая безопасная compatibility rule: stable paths несут обычное обещание, beta paths поддерживаются, но не заморожены.
- Нужна install rule: Homebrew, npm и PyPI wrappers устанавливают CLI; это не runtime API и не SDK contract.
- Нужна support rule: свяжите эту страницу с [Границей поддержки](/ru/reference/support-boundary) и [Поддержкой target’ов](/ru/reference/target-support).

## Что помогает решить эта страница

- какой version signal должен определять решения команды
- является ли изменение обычным обновлением, миграцией или сменой пути
- читаете ли вы release note как guidance, compatibility promise или и то и другое

## Правило публичного baseline

- Текущий публичный baseline задаётся самым свежим релизом, который относится к вашему user-facing сценарию.
- Сейчас главным baseline в публичных docs является [v1.0.6](/ru/releases/v1-0-6).
- Для старых repo или более узких поверхностей вроде Go SDK сначала читайте релиз, который относится именно к этой поверхности, а не предполагайте, что последний релиз одинаково касается всего.

## Stable, Beta, Experimental

- `public-stable`: обычное production-ожидание
- `public-beta`: поддерживается, но ещё двигается
- `public-experimental`: полезно для ранних пользователей, но это не долгосрочный compatibility promise

Практическое правило простое: не делайте beta или experimental частью долгоживущего team standard, если не готовы принимать churn.

## Каналы установки против runtime-контрактов

- Homebrew, npm и PyPI — это install channels для CLI.
- Они не превращают wrapper packages в runtime API.
- Публичные runtime- и SDK-контракты живут в guides, reference rules и generated [API](/ru/api/) surfaces.

## Совместимость runtime helper'ов

- Shared runtime helpers вроде `plugin-kit-ai-runtime` нужно читать через release notes и delivery guidance, а не как отдельное вечное обещание.
- Если текущая guidance меняется вокруг `--runtime-package`, используйте самый свежий release note вместе с [Выбором модели поставки](/ru/guide/choose-delivery-model).
- Не считайте, что «вышла новая версия пакета» автоматически означает «все repo должны обновиться прямо сейчас».

## Что означают release notes

- Release note говорит, что изменилось для пользователей, что не изменилось и какая рекомендация стала сильнее.
- Release note не является универсальным обещанием parity для всех target, runtime и language surfaces.
- Безопасный способ читать релизы такой:
  1. откройте самый свежий подходящий note
  2. проверьте migration callouts
  3. вернитесь в guide или reference за точной механикой

## Что команде реально стоит pin'ить

- Фиксируйте workflow вокруг `doctor`, `render` и `validate --strict`.
- Фиксируйте выбранный runtime и target path в документации repo и в CI.
- Фиксируйте версии shared runtime helper'ов там, где этого требует выбранная delivery model.
- Держите команду на одном опубликованном baseline, а не на смеси из нескольких наполовину запомненных release states.

## Чего не нужно предполагать

- не нужно считать, что у всех target одинаковый compatibility promise
- не нужно считать wrapper packages SDK'ами
- не нужно считать, что beta convenience flow несёт такой же долгий контракт, как stable path
- не нужно считать, что один release note заменяет проверку support и target boundaries

## С чего лучше начать

- Нужен словарь зрелости: [Модель стабильности](/ru/concepts/stability-model)
- Нужна точная support line: [Граница поддержки](/ru/reference/support-boundary)
- Нужна точная матрица target’ов: [Поддержка target’ов](/ru/reference/target-support)
- Нужен текущий журнал изменений: [Релизы](/ru/releases/)
- Нужен rollout path для живых repo: [Плейбук обновлений и миграции](/ru/guide/upgrade-and-migration-playbook)
