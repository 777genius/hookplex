---
title: "Installation"
description: "Install plugin-kit-ai using supported channels."
canonicalId: "page:guide:installation"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# Installation

Homebrew is the recommended default when it fits your environment.

## Supported Channels

- Homebrew for the cleanest default CLI path.
- npm when your environment is already centered around npm.
- PyPI / pipx when your environment is already centered around Python.
- Verified install script as the fallback path.

## Recommended Commands

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

## Which One Should Most People Use?

- Use Homebrew if you are on macOS and want the smoothest default path.
- Use npm or pipx only when that already matches your team environment.
- Use the verified script when you need a fallback outside package-manager-first setups.

## CI Install Path

For CI, prefer the dedicated setup action instead of teaching every workflow how to download the CLI manually.

## Important Boundary

The npm and PyPI packages are ways to install the CLI binary. They are not the public runtime API and they are not SDKs.

See [Reference > Install Channels](/en/reference/install-channels) for the contract boundary.
