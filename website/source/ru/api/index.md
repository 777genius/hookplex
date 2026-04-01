---
title: "API"
description: "Сгенерированный API-справочник для plugin-kit-ai."
canonicalId: "page:api:index"
section: "api"
locale: "ru"
generated: false
translationRequired: true
aside: false
outline: false
---

<div class="docs-hero docs-hero--compact">
  <p class="docs-kicker">СГЕНЕРИРОВАННЫЙ СПРАВОЧНИК</p>
  <h1>API поверхности</h1>
  <p class="docs-lead">
    Этот справочник генерируется из реального CLI, пакетов и структурированных metadata. Он разделён по публичным разделам, чтобы по мере роста проекта API оставался понятным и предсказуемым.
  </p>
</div>

<div class="docs-grid">
  <a class="docs-card" href="./cli/">
    <h2>CLI</h2>
    <p>Команды, экспортированные из живого дерева Cobra.</p>
  </a>
  <a class="docs-card" href="./go-sdk/">
    <h2>Go SDK</h2>
    <p>Публичные Go-пакеты для стабильных путей интеграции.</p>
  </a>
  <a class="docs-card" href="./runtime-node/">
    <h2>Node Runtime</h2>
    <p>Типизированные runtime-helpers для JS и TS.</p>
  </a>
  <a class="docs-card" href="./runtime-python/">
    <h2>Python Runtime</h2>
    <p>Только публичные Python runtime-helpers, без install-wrapper пакетов.</p>
  </a>
  <a class="docs-card" href="./platform-events/">
    <h2>События платформ</h2>
    <p>События и точки входа, сгруппированные по целевым платформам.</p>
  </a>
  <a class="docs-card" href="./capabilities/">
    <h2>Capabilities</h2>
    <p>Взгляд на систему через capabilities, а не только через дерево пакетов.</p>
  </a>
</div>

## Выбор за 60 секунд

- Открывайте `CLI`, когда занимаетесь authoring, validate, bundle или inspect для plugin repo.
- Открывайте `Go SDK`, когда строите самый сильный production-oriented runtime path.
- Открывайте `Node Runtime` или `Python Runtime`, когда уже выбрали поддерживаемый repo-local interpreted runtime path и теперь нужны helper APIs.
- Открывайте `Platform Events`, когда уже знаете target platform и нужен event-level contract.
- Открывайте `Capabilities`, когда хотите сравнивать похожее поведение поперёк платформ, а не читать одну platform tree за раз.

## С чего лучше начать

- Нужна главная пользовательская поверхность: начинайте с [CLI](./cli/).
- Нужен самый сильный production default: начинайте с [Go SDK](./go-sdk/).
- Нужны interpreted runtime helpers: начинайте с [Node Runtime](./runtime-node/) или [Python Runtime](./runtime-python/).
- Нужен event-level platform detail: начинайте с [Platform Events](./platform-events/).
- Нужна cross-platform карта поведения: начинайте с [Capabilities](./capabilities/).

## Как выбрать нужную поверхность

- Открывайте `CLI`, когда нужны команды, флаги и сам рабочий процесс автора плагина.
- Открывайте `Go SDK`, когда строите самый сильный путь для продакшен-runtime-плагина.
- Открывайте `Node Runtime` или `Python Runtime`, когда нужны helper APIs для поддерживаемых локальных Python или Node проектов.
- Открывайте `Platform Events`, когда выбираете события конкретной платформы.
- Открывайте `Capabilities`, когда нужен взгляд поперёк платформ на то, на что plugin может реагировать или что может контролировать.

## Что покрывает эта API-зона

- живое дерево команд Cobra
- публичные Go packages
- shared runtime helper APIs для Node и Python
- события конкретных платформ
- metadata по capabilities поперёк платформ

## Чем эта API-зона не является

- Это не лучший первый вход, если вы ещё выбираете target, runtime или starter.
- Она не заменяет guide-страницы для first-time setup, delivery и team handoff.
- Это generated reference, привязанный к реальным исходным данным, поэтому он лучше всего работает после того, как вы уже понимаете, какая поверхность вам нужна.
