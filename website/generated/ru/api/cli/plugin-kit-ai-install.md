---
title: "plugin-kit-ai install"
description: "Устанавливает бинарник плагина из GitHub Releases с проверкой через `checksums.txt`."
canonicalId: "command:plugin-kit-ai:install"
surface: "cli"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "cli:plugin-kit-ai install"
translationRequired: false
---
<DocMetaCard surface="cli" stability="public-stable" maturity="stable" source-ref="cli:plugin-kit-ai install" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/cli/plugin-kit-ai" />

# plugin-kit-ai install

Сгенерировано из реального Cobra command tree.

Устанавливает бинарник плагина из GitHub Releases с проверкой через `checksums.txt`.

## plugin-kit-ai install

Устанавливает бинарник плагина из GitHub Releases с проверкой через `checksums.txt`.

### Описание

Скачивает `checksums.txt` и release-артефакт для ваших `GOOS/GOARCH`, проверяет `SHA256` и записывает бинарник в `--dir`.
По умолчанию используется каталог `bin`. Выбор артефакта такой: (1) один архив GoReleaser `*_GOOS_GOARCH.tar.gz` с извлечением файла из корня архива; или (2) сырой бинарник с именем вида `*-GOOS-GOARCH` либо `*.exe` на Windows.

Используйте ровно один из флагов `--tag` или `--latest`. Draft-релизы не принимаются; для prerelease нужен `--pre`.
Необязательный `--output-name` задаёт имя устанавливаемого файла.

Эта команда устанавливает сторонние бинарники плагинов, а не сам CLI `plugin-kit-ai`.

```
plugin-kit-ai install [owner/repo] [flags]
```

### Опции

```
      --dir string            каталог для установленного бинарника (создаётся при отсутствии) (по умолчанию `bin`)
  -f, --force                 перезаписывает существующий бинарник
      --github-token string   GitHub token (необязательно; по умолчанию берётся из `GITHUB_TOKEN`)
      --goarch string         переопределяет целевой `GOARCH` (по умолчанию: `GOARCH` хоста)
      --goos string           переопределяет целевой `GOOS` (по умолчанию: `GOOS` хоста)
  -h, --help                  справка по install
      --latest                устанавливает из `GitHub releases/latest` (без prerelease) вместо `--tag`
      --output-name string    записывает бинарник под этим именем в `--dir` (по умолчанию: имя из архива)
      --pre                   разрешает GitHub prerelease-релизы (не stable)
      --tag string            Git release tag (обязателен, если не указан `--latest`), например `v0.1.0`
```

### См. также

* plugin-kit-ai	 - CLI plugin-kit-ai для создания проектов и служебных операций вокруг AI-плагинов.
