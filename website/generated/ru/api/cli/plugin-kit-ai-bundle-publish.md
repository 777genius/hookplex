---
title: "plugin-kit-ai bundle publish"
description: "Публикует экспортированный Python/Node bundle в GitHub Releases."
canonicalId: "command:plugin-kit-ai:bundle:publish"
surface: "cli"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "cli:plugin-kit-ai bundle publish"
translationRequired: false
---
<DocMetaCard surface="cli" stability="public-stable" maturity="stable" source-ref="cli:plugin-kit-ai bundle publish" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/cli/plugin-kit-ai" />

# plugin-kit-ai bundle publish

Сгенерировано из реального Cobra command tree.

Публикует экспортированный Python/Node bundle в GitHub Releases.

## plugin-kit-ai bundle publish

Публикует экспортированный Python/Node bundle в GitHub Releases.

### Описание

Публикует экспортированный Python/Node bundle в GitHub Releases.

Эта стабильная producer-side handoff-поверхность экспортирует bundle, по умолчанию создаёт опубликованный release,
использует `--draft`, если релиз нужно оставить черновиком, загружает сам bundle и соседний `.sha256`-asset,
и остаётся отдельной от binary-only сценария `plugin-kit-ai install`.

```
plugin-kit-ai bundle publish [path] [flags]
```

### Опции

```
      --draft                 оставляет целевой release черновиком вместо публикации
  -f, --force                 заменяет существующие bundle-артефакты с тем же именем
      --github-token string   GitHub token (необязательно; по умолчанию берётся из `GITHUB_TOKEN`)
  -h, --help                  справка по publish
      --platform string       целевая платформа для экспорта и публикации (`codex-runtime` или `claude`)
      --repo string           GitHub owner/repo, куда будут загружены bundle-артефакты
      --tag string            GitHub release tag, который нужно переиспользовать или создать
```

### См. также

* plugin-kit-ai bundle	 - Инструменты bundle-экспорта для переносимых архивов интерпретируемого runtime.
