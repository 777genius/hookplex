---
title: "Как получить помощь и внести вклад"
description: "Поймите, где просить помощь, где сообщать о проблемах и как вносить улучшения в docs, examples и код plugin-kit-ai."
canonicalId: "page:reference:get-help-and-contribute"
section: "reference"
locale: "ru"
generated: false
translationRequired: true
---

# Как получить помощь и внести вклад

Открывайте эту страницу, когда нужен публичный путь для помощи, bug report, docs feedback или аккуратного вклада в проект.

## Выбор за 60 секунд

- Нужно сообщить о баге или пробеле в документации: откройте [Issues](https://github.com/777genius/plugin-kit-ai/issues) или [New Issue](https://github.com/777genius/plugin-kit-ai/issues/new/choose).
- Хотите сразу предложить понятный и локальный fix: откройте [Pull Requests](https://github.com/777genius/plugin-kit-ai/pulls).
- Нужно сообщить о security-проблеме: следуйте [SECURITY.md](https://github.com/777genius/plugin-kit-ai/blob/main/SECURITY.md) и не создавайте публичный issue.
- Нужно сначала понять публичный контракт: прочитайте [Границу поддержки](/ru/reference/support-boundary) и [Поддержку target’ов](/ru/reference/target-support).

## Для чего особенно полезна эта страница

- чтобы быстро выбрать правильный канал для bug report, docs fix или contribution
- чтобы привязывать community feedback к публичному контракту проекта
- чтобы не отправлять правильный вопрос не в тот канал

## Когда использовать Issues

- для публичных багов в CLI, SDK, runtime helpers или generated docs
- для неясных guide-страниц, сломанных examples и недостающей migration guidance
- для feature ideas, которые сначала требуют обсуждения
- для расхождений между docs и реальным support boundary

## Когда использовать Pull Requests

- для локальных и понятных fixes с ясной user-facing целью
- для docs-уточнений, которым не нужно длинное предварительное обсуждение
- для улучшений examples или reference-страниц, которые явно совпадают с публичной моделью продукта
- для усиления тестов или validation вокруг уже принятого пути

## Хорошие первые области для вклада

- улучшение onboarding и choice pages
- улучшение examples, guidance по starter’ам и migration guidance
- уточнение generated API overview copy, когда меняется source surface
- правка wording в docs и reference, когда пользовательский путь читается двусмысленно

## Как держать contribution здоровым

- держите один смысловой вопрос в одном pull request
- обновляйте docs, когда меняется публичное поведение
- синхронизируйте generated docs, если pipeline показывает drift
- не превращайте user-facing fix в широкий несвязанный refactor

## С чего лучше начать перед contribution

- Нужна модель продукта: [Managed Project Model](/ru/concepts/managed-project-model)
- Нужен официальный контракт репозитория: [Стандарт репозитория](/ru/reference/repository-standard)
- Нужен канонический authoring flow: [Процесс авторинга](/ru/reference/authoring-workflow)
- Нужен путь внедрения в команду: [Внедрение в команду](/ru/guide/team-adoption)

## Если не уверены, с чего начать

1. Сначала воспроизведите проблему или пробел по публичной документации.
2. Проверьте, это действительно docs problem, misunderstanding support boundary или кодовый баг.
3. Открывайте самый маленький полезный issue или pull request, после которого публичный контракт становится яснее.

## Security

Для нераскрытых уязвимостей не открывайте публичный issue. Используйте приватный путь из [SECURITY.md](https://github.com/777genius/plugin-kit-ai/blob/main/SECURITY.md).

## Репозиторный guide по contribution

Общие правила вклада в сам репозиторий описаны в [CONTRIBUTING.md](https://github.com/777genius/plugin-kit-ai/blob/main/CONTRIBUTING.md).
