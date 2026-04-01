---
title: "sdk"
description: "Generated Go SDK package reference for github.com/777genius/plugin-kit-ai/sdk"
canonicalId: "go-package:github.com/777genius/plugin-kit-ai/sdk"
surface: "go-sdk"
section: "api"
locale: "en"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "sdk"
translationRequired: false
---
<DocMetaCard surface="go-sdk" stability="public-stable" maturity="stable" source-ref="sdk" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/sdk" />

# sdk

Generated from the public Go package via gomarkdoc.

**Import path:** `github.com/777genius/plugin-kit-ai/sdk`

```go
import "github.com/777genius/plugin-kit-ai/sdk"
```

## Index

- type App
  - func New\(cfg Config\) \*App
  - func \(a \*App\) Claude\(\) \*claude.Registrar
  - func \(a \*App\) Codex\(\) \*codex.Registrar
  - func \(a \*App\) Run\(\) int
  - func \(a \*App\) RunContext\(ctx context.Context\) int
  - func \(a \*App\) Use\(mw Middleware\)
- type CapabilityID
- type Config
- type Env
- type Handled
- type IO
- type InvocationContext
- type Logger
- type MaturityLevel
- type Middleware
- type Next
- type NopLogger
- type Result
- type SupportEntry
  - func Supported\(\) \[\]SupportEntry
- type SupportStatus
- type TransportMode


## type App



```go
type App struct {
    // contains filtered or unexported fields
}
```

### func New

```go
func New(cfg Config) *App
```



### func \(\*App\) Claude

```go
func (a *App) Claude() *claude.Registrar
```




**Example**

```go
package main

import (
	pluginkitai "github.com/777genius/plugin-kit-ai/sdk"
	"github.com/777genius/plugin-kit-ai/sdk/claude"
)

func main() {
	app := pluginkitai.New(pluginkitai.Config{Name: "demo"})
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response {
		return claude.Allow()
	})
	_ = app
}
```


### func \(\*App\) Codex

```go
func (a *App) Codex() *codex.Registrar
```




**Example**

```go
package main

import (
	pluginkitai "github.com/777genius/plugin-kit-ai/sdk"
	"github.com/777genius/plugin-kit-ai/sdk/codex"
)

func main() {
	app := pluginkitai.New(pluginkitai.Config{Name: "demo"})
	app.Codex().OnNotify(func(*codex.NotifyEvent) *codex.Response {
		return codex.Continue()
	})
	_ = app
}
```


### func \(\*App\) Run

```go
func (a *App) Run() int
```



### func \(\*App\) RunContext

```go
func (a *App) RunContext(ctx context.Context) int
```



### func \(\*App\) Use

```go
func (a *App) Use(mw Middleware)
```



## type CapabilityID



```go
type CapabilityID = runtime.CapabilityID
```

## type Config



```go
type Config struct {
    Name   string
    Args   []string
    IO     IO
    Env    Env
    Logger Logger
}
```

## type Env



```go
type Env = runtime.Env
```

## type Handled



```go
type Handled = runtime.Handled
```

## type IO



```go
type IO = runtime.IO
```

## type InvocationContext



```go
type InvocationContext = runtime.InvocationContext
```

## type Logger



```go
type Logger = runtime.Logger
```

## type MaturityLevel



```go
type MaturityLevel = runtime.MaturityLevel
```

## type Middleware



```go
type Middleware = runtime.Middleware
```

## type Next



```go
type Next = runtime.Next
```

## type NopLogger



```go
type NopLogger = runtime.NopLogger
```

## type Result



```go
type Result = runtime.Result
```

## type SupportEntry



```go
type SupportEntry = runtime.SupportEntry
```

### func Supported

```go
func Supported() []SupportEntry
```



## type SupportStatus



```go
type SupportStatus = runtime.SupportStatus
```

## type TransportMode



```go
type TransportMode = runtime.TransportMode
```
