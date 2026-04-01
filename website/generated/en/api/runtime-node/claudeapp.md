---
title: "Claude App"
description: "Generated Node runtime reference for ClaudeApp"
canonicalId: "node-runtime:ClaudeApp"
surface: "runtime-node"
section: "api"
locale: "en"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "npm/plugin-kit-ai-runtime"
translationRequired: false
---
<DocMetaCard surface="runtime-node" stability="public-stable" maturity="stable" source-ref="npm/plugin-kit-ai-runtime" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/npm/plugin-kit-ai-runtime" />

# Claude App

Generated via TypeDoc and typedoc-plugin-markdown.

Defined in: index.d.ts:11

## Constructors

### Constructor

&gt; **new ClaudeApp**(`options`): `ClaudeApp`

Defined in: index.d.ts:12

#### Parameters

##### options

###### allowedHooks

`string`[] \| readonly `string`[]

###### usage

`string`

#### Returns

`ClaudeApp`

## Methods

### on()

&gt; **on**(`hookName`, `handler`): `this`

Defined in: index.d.ts:13

#### Parameters

##### hookName

`string`

##### handler

`ClaudeHandler`

#### Returns

`this`

***

### onPreToolUse()

&gt; **onPreToolUse**(`handler`): `this`

Defined in: index.d.ts:15

#### Parameters

##### handler

`ClaudeHandler`

#### Returns

`this`

***

### onStop()

&gt; **onStop**(`handler`): `this`

Defined in: index.d.ts:14

#### Parameters

##### handler

`ClaudeHandler`

#### Returns

`this`

***

### onUserPromptSubmit()

&gt; **onUserPromptSubmit**(`handler`): `this`

Defined in: index.d.ts:16

#### Parameters

##### handler

`ClaudeHandler`

#### Returns

`this`

***

### run()

&gt; **run**(): `number`

Defined in: index.d.ts:17

#### Returns

`number`
