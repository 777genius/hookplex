---
title: "plugin-kit-ai bundle install"
description: "Устанавливает локальный экспортированный Python/Node bundle в целевой каталог."
canonicalId: "command:plugin-kit-ai:bundle:install"
surface: "cli"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "cli:plugin-kit-ai bundle install"
translationRequired: false
---
<DocMetaCard surface="cli" stability="public-stable" maturity="stable" source-ref="cli:plugin-kit-ai bundle install" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/cli/plugin-kit-ai" />

# plugin-kit-ai bundle install

Сгенерировано из реального Cobra command tree.

Устанавливает локальный экспортированный Python/Node bundle в целевой каталог.

## plugin-kit-ai bundle install

Устанавливает локальный экспортированный Python/Node bundle в целевой каталог.

### Описание

Устанавливает локальный `.tar.gz` bundle, созданный через `plugin-kit-ai export`, в целевой каталог.

Эта стабильная handoff-поверхность поддерживает только локальные экспортированные Python/Node bundle для `codex-runtime` или `claude`.
Команда безопасно распаковывает содержимое bundle, печатает следующие шаги и не расширяет binary-only сценарий установки `plugin-kit-ai install`.

```
plugin-kit-ai bundle install &lt;bundle.tar.gz&gt; [flags]
```

### Опции

```
      --dest string   целевой каталог для распакованного содержимого bundle
  -f, --force         перезаписывает существующий целевой каталог
  -h, --help          справка по install
```

### См. также

* plugin-kit-ai bundle	 - Инструменты bundle-экспорта для переносимых архивов интерпретируемого runtime.
