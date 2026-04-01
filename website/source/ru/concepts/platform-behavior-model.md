---
title: "Модель поведения платформ"
description: "Как мыслить platform-first и capability-first views, не путая их с выбором target family и силой обещаний поддержки."
canonicalId: "page:concepts:platform-behavior-model"
section: "concepts"
locale: "ru"
generated: false
translationRequired: true
---

# Модель поведения платформ

Открывайте эту страницу, когда общая модель target’ов уже понятна, а теперь нужна чистая mental model для чтения `Platform Events` и `Capabilities`.

## Короткое правило

- используйте **platform-first view**, когда платформа уже выбрана и нужны её точные events
- используйте **capability-first view**, когда надо сравнить похожее поведение между платформами
- не используйте ни один из этих views для выбора target family; target надо выбрать раньше

## Два способа читать один и тот же слой

`plugin-kit-ai` показывает platform behavior через два дополняющих друг друга взгляда:

- `Platform Events`: сначала по платформам, потом по events
- `Capabilities`: сначала по поведению, потом по покрытию платформ

Ни один из этих взглядов не “правильнее”. Они отвечают на разные вопросы.

## Platform-first view

Начинайте с платформы, когда ваш вопрос звучит так:

- «Какие события вообще есть у Claude?»
- «Что умеет Codex на уровне events?»
- «Какие события на этой платформе stable, а какие beta?»

Это правильный взгляд, когда целевая платформа уже известна и вам нужен точный event surface.

## Capability-first view

Начинайте с capabilities, когда ваш вопрос звучит так:

- «На каких платформах вообще есть notify-подобное поведение?»
- «Это cross-platform behavior или только platform-specific вещь?»
- «Сколько платформ вообще покрывают эту capability?»

Это правильный взгляд, когда важнее сравнить само поведение, а не читать одно platform tree.

## Чем этот слой не является

- Это не то же самое, что выбор target family.
- Он не заменяет promise-by-path framing.
- Он не означает, что Claude и Codex обещают одинаковую глубину и одинаковую стабильность похожего поведения.

Platform layer описывает behavior surfaces, а не равные support guarantees.

## Безопасный порядок чтения

1. Сначала выберите target family.
2. Потом поймите promise по пути.
3. И только затем открывайте platform behavior layer через:
   - `Platform Events` для platform-first чтения
   - `Capabilities` для behavior-first чтения

## Частая ошибка

Люди часто воспринимают одно общее capability name как доказательство одинаковой поддержки между платформами. Это слишком сильный вывод.

Capability name говорит лишь о том, что существует связанное поведение. Но это **не** означает автоматически:

- одинаковую maturity
- одинаковую глубину event surface
- одинаковую операционную цену и силу обещаний

## С чем читать вместе

- [Выбор target](/ru/guide/choose-a-target), если target family ещё не выбран
- [Обещания поддержки по путям](/ru/reference/support-promise-by-path), если нужен framing по силе обещаний и цене
- [Platform Events](/ru/api/platform-events/), если платформа уже известна
- [Capabilities](/ru/api/capabilities/), если главная задача — сравнить поведение

## Финальное правило

Используйте `Platform Events`, чтобы ответить на вопрос «что есть на этой платформе?», и `Capabilities`, чтобы ответить на вопрос «где это поведение вообще существует между платформами?».
