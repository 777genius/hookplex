---
title: "v1.0.4 Go SDK"
description: "Patch release notes для исправления Go SDK module path."
canonicalId: "page:releases:v1-0-4-go-sdk"
section: "releases"
locale: "ru"
generated: false
translationRequired: true
---

# v1.0.4 Go SDK

Дата релиза: `2026-03-29`

## Кому особенно важно прочитать этот релиз

- тем, кто использует Go SDK как обычный публичный Go module
- тем, кто обновляет старую Go-интеграцию и хочет увидеть первый корректный module-path fix
- тем, кому нужна одна короткая ссылка, чтобы отказаться от `replace`-workaround’ов

## Почему этот patch важен

Этот patch сделал public Go SDK module path корректным для нормального Go consumption.

## Что изменилось

- корень Go SDK module был перенесён с `sdk/plugin-kit-ai/` в `sdk/`
- public module path `github.com/777genius/plugin-kit-ai/sdk` теперь совпадает с реальным layout репозитория
- starter repos, examples и templates перестали учить новичков `replace`-workaround’ам

## Практический вывод

- используйте `github.com/777genius/plugin-kit-ai/sdk@v1.0.4` или новее для нормального Go module consumption
- считайте `v1.0.3` known-bad релизом для Go SDK module path

## Как лучше читать этот релиз

- SDK consumer: как patch, после которого нормальный module story наконец совпадает с публичной рекомендацией.
- Владельцу repo: как релиз, который помогает убрать старые workaround assumptions из starter’ов, examples и внутренней документации.
- Новому Go-пользователю: как cleanup patch, который стоит читать вместе с [v1.0.0](/ru/releases/v1-0-0) для более широкого stable baseline.

## Почему это важно пользователям

Этот patch убрал лишнее трение для обычных Go consumers и сделал рекомендуемый SDK path похожим на нормальный public module, а не на special-case workaround.

## Что читать дальше

- Читайте [v1.0.0](/ru/releases/v1-0-0), если нужен первый стабильный публичный baseline.
- Читайте [Go SDK](/ru/api/go-sdk/), если нужен live generated SDK surface.
- Читайте [Политику версий и совместимости](/ru/reference/version-and-compatibility), если нужна компактная rule про релизы, baselines и совместимость.
