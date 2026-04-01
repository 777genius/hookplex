---
title: "Установка"
description: "Установка plugin-kit-ai через поддерживаемые каналы."
canonicalId: "page:guide:installation"
section: "guide"
locale: "ru"
generated: false
translationRequired: true
---

# Установка

Homebrew — рекомендуемый путь по умолчанию, если он подходит вашей среде.

## Поддерживаемые каналы

- Homebrew для самого чистого CLI пути.
- npm, если у вас среда уже завязана на npm.
- PyPI / pipx, если у вас среда уже завязана на Python.
- Verified install script как запасной путь.

## Рекомендуемые команды

### Homebrew

```bash
brew install 777genius/homebrew-plugin-kit-ai/plugin-kit-ai
plugin-kit-ai version
```

### npm

```bash
npm i -g plugin-kit-ai
plugin-kit-ai version
```

### PyPI / pipx

```bash
pipx install plugin-kit-ai
plugin-kit-ai version
```

### Verified Script

```bash
curl -fsSL https://raw.githubusercontent.com/777genius/plugin-kit-ai/main/scripts/install.sh | sh
plugin-kit-ai version
```

## Что выбирать большинству людей?

- Выбирайте Homebrew, если вы на macOS и хотите самый гладкий путь по умолчанию.
- Выбирайте npm или pipx только тогда, когда это уже соответствует среде вашей команды.
- Используйте verified script как запасной путь вне сценариев, где всё уже крутится вокруг пакетного менеджера.

## Путь для CI

Для CI лучше использовать dedicated setup action, а не учить каждый workflow вручную скачивать CLI.

## Важная граница

npm и PyPI пакеты — это способы установить CLI binary. Они не считаются публичным runtime API и не являются SDK.

См. [Справочник > Каналы установки](/ru/reference/install-channels) для формальной границы контракта.
