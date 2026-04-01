---
title: "codex"
description: "Generated Go SDK package reference for github.com/777genius/plugin-kit-ai/sdk/codex"
canonicalId: "go-package:github.com/777genius/plugin-kit-ai/sdk/codex"
surface: "go-sdk"
section: "api"
locale: "en"
generated: true
editLink: false
stability: "public-stable"
maturity: "stable"
sourceRef: "sdk/codex"
translationRequired: false
---
<DocMetaCard surface="go-sdk" stability="public-stable" maturity="stable" source-ref="sdk/codex" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/sdk/codex" />

# codex

Generated from the public Go package via gomarkdoc.

**Import path:** `github.com/777genius/plugin-kit-ai/sdk/codex`

```go
import "github.com/777genius/plugin-kit-ai/sdk/codex"
```

## Index

- func RegisterCustomJSON\[T any\]\(r \*Registrar, eventName string, fn func\(\*T\) \*Response\) error
- type NotifyEvent
  - func \(e \*NotifyEvent\) RawJSON\(\) json.RawMessage
- type Registrar
  - func NewRegistrar\(backend runtime.RegistrarBackend\) \*Registrar
  - func \(r \*Registrar\) OnNotify\(fn func\(\*NotifyEvent\) \*Response\)
- type Response
  - func Continue\(\) \*Response


## func RegisterCustomJSON

```go
func RegisterCustomJSONT any *Response) error
```

RegisterCustomJSON registers an experimental future Codex hook whose payload is delivered as a JSON argv argument. The handler remains fully typed.

## type NotifyEvent



```go
type NotifyEvent struct {
    Raw    json.RawMessage
    Client string
}
```

### func \(\*NotifyEvent\) RawJSON

```go
func (e *NotifyEvent) RawJSON() json.RawMessage
```



## type Registrar



```go
type Registrar struct {
    // contains filtered or unexported fields
}
```

### func NewRegistrar

```go
func NewRegistrar(backend runtime.RegistrarBackend) *Registrar
```



### func \(\*Registrar\) OnNotify

```go
func (r *Registrar) OnNotify(fn func(*NotifyEvent) *Response)
```



## type Response



```go
type Response struct{}
```

### func Continue

```go
func Continue() *Response
```
