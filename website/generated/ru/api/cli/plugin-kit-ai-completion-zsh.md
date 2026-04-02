---
title: "plugin-kit-ai completion zsh"
description: "Generate the autocompletion script for zsh"
canonicalId: "command:plugin-kit-ai:completion:zsh"
surface: "cli"
section: "api"
locale: "ru"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "cli:plugin-kit-ai completion zsh"
translationRequired: false
---
<DocMetaCard surface="cli" stability="public-stable" maturity="stable" source-ref="cli:plugin-kit-ai completion zsh" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/cli/plugin-kit-ai" />

# plugin-kit-ai completion zsh

Сгенерировано из реального Cobra command tree.

Generate the autocompletion script for zsh

## plugin-kit-ai completion zsh

Generate the autocompletion script for zsh

### Описание

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" &gt;&gt; ~/.zshrc

To load completions in your current shell session:

	source &lt;(plugin-kit-ai completion zsh)

To load completions for every new session, execute once:

#### Linux:

	plugin-kit-ai completion zsh &gt; "${fpath[1]}/_plugin-kit-ai"

#### macOS:

	plugin-kit-ai completion zsh &gt; $(brew --prefix)/share/zsh/site-functions/_plugin-kit-ai

You will need to start a new shell for this setup to take effect.


```
plugin-kit-ai completion zsh [flags]
```

### Опции

```
  -h, --help              справка по zsh
      --no-descriptions   disable completion descriptions
```

### См. также

* plugin-kit-ai completion	 - Генерирует скрипт автодополнения для указанной оболочки.
