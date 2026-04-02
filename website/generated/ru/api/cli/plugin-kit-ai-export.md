---
title: "plugin-kit-ai export"
description: "Создаёт переносимый bundle интерпретируемого runtime без смены install-семантики."
canonicalId: "command:plugin-kit-ai:export"
surface: "cli"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "cli:plugin-kit-ai export"
translationRequired: false
---
<DocMetaCard surface="cli" stability="public-stable" maturity="stable" source-ref="cli:plugin-kit-ai export" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/cli/plugin-kit-ai" />

# plugin-kit-ai export

Сгенерировано из реального Cobra command tree.

Создаёт переносимый bundle интерпретируемого runtime без смены install-семантики.

## plugin-kit-ai export

Создаёт переносимый bundle интерпретируемого runtime без смены install-семантики.

### Описание

Создаёт детерминированный переносимый `.tar.gz` bundle для launcher-based проектов с интерпретируемым runtime.

Эта beta-поверхность покрывает ограниченный handoff/export сценарий для runtime-репозиториев на `python`, `node` и `shell`.
Она не расширяет сценарий `plugin-kit-ai install` и не подразумевает packaging для marketplace или поставку с уже предустановленными зависимостями.

```
plugin-kit-ai export [path] [flags]
```

### Опции

```
  -h, --help              справка по export
      --output string     записывает bundle в путь `.tar.gz` (по умолчанию: `&lt;root&gt;/&lt;name&gt;_&lt;platform&gt;_&lt;runtime&gt;_bundle.tar.gz`)
      --platform string   переопределяет целевую платформу (`codex-runtime` или `claude`)
```

### См. также

* plugin-kit-ai	 - CLI plugin-kit-ai для создания проектов и служебных операций вокруг AI-плагинов.
