---
title: "platformmeta"
description: "Generated Go SDK package reference for github.com/777genius/plugin-kit-ai/sdk/platformmeta"
canonicalId: "go-package:github.com/777genius/plugin-kit-ai/sdk/platformmeta"
surface: "go-sdk"
section: "api"
locale: "en"
generated: true
editLink: false
stability: "public-beta"
maturity: "beta"
sourceRef: "sdk/platformmeta"
translationRequired: false
---
<DocMetaCard surface="go-sdk" stability="public-beta" maturity="beta" source-ref="sdk/platformmeta" source-href="https://github.com/777genius/plugin-kit-ai/tree/main/sdk/platformmeta" />

# platformmeta

Generated from the public Go package via gomarkdoc.

**Import path:** `github.com/777genius/plugin-kit-ai/sdk/platformmeta`

```go
import "github.com/777genius/plugin-kit-ai/sdk/platformmeta"
```

## Index

- func IDs\(\) \[\]string
- type ContextStrategy
- type LauncherMeta
- type LauncherRequirement
- type ManagedArtifactKind
- type ManagedArtifactSpec
- type NativeDocFormat
- type NativeDocRole
- type NativeDocSpec
- type PlatformFamily
- type PlatformProfile
  - func All\(\) \[\]PlatformProfile
  - func Lookup\(name string\) \(PlatformProfile, bool\)
- type SDKMeta
- type ScaffoldMeta
- type SupportStatus
- type SurfaceSupport
- type SurfaceTier
- type TargetContractMeta
- type TemplateFile
- type TransportMode
- type ValidateMeta


## func IDs

```go
func IDs() []string
```



## type ContextStrategy



```go
type ContextStrategy string
```

```go
const (
    ContextStrategyNone              ContextStrategy = ""
    ContextStrategyGeminiPrimaryRoot ContextStrategy = "gemini_primary_root"
)
```

## type LauncherMeta



```go
type LauncherMeta struct {
    Requirement LauncherRequirement
}
```

## type LauncherRequirement



```go
type LauncherRequirement string
```

```go
const (
    LauncherRequired LauncherRequirement = "required"
    LauncherOptional LauncherRequirement = "optional"
    LauncherIgnored  LauncherRequirement = "ignored"
)
```

## type ManagedArtifactKind



```go
type ManagedArtifactKind string
```

```go
const (
    ManagedArtifactStatic          ManagedArtifactKind = "static"
    ManagedArtifactMirror          ManagedArtifactKind = "mirror"
    ManagedArtifactPortableMCP     ManagedArtifactKind = "portable_mcp"
    ManagedArtifactPortableSkills  ManagedArtifactKind = "portable_skills"
    ManagedArtifactSelectedContext ManagedArtifactKind = "selected_context"
)
```

## type ManagedArtifactSpec



```go
type ManagedArtifactSpec struct {
    Kind          ManagedArtifactKind
    Path          string
    ComponentKind string
    SourceRoot    string
    OutputRoot    string
    ContextMode   ContextStrategy
}
```

## type NativeDocFormat



```go
type NativeDocFormat string
```

```go
const (
    NativeDocYAML     NativeDocFormat = "yaml"
    NativeDocJSON     NativeDocFormat = "json"
    NativeDocTOML     NativeDocFormat = "toml"
    NativeDocMarkdown NativeDocFormat = "md"
)
```

## type NativeDocRole



```go
type NativeDocRole string
```

```go
const (
    NativeDocRoleStructured NativeDocRole = "structured"
    NativeDocRoleExtra      NativeDocRole = "extra"
)
```

## type NativeDocSpec



```go
type NativeDocSpec struct {
    Kind        string
    Path        string
    Format      NativeDocFormat
    Role        NativeDocRole
    ManagedKeys []string
}
```

## type PlatformFamily



```go
type PlatformFamily string
```

```go
const (
    FamilyPackagedRuntime  PlatformFamily = "packaged_runtime"
    FamilyExtensionPackage PlatformFamily = "extension_package"
    FamilyCodePlugin       PlatformFamily = "code_plugin"
    FamilyIDEPlugin        PlatformFamily = "ide_plugin"
)
```

## type PlatformProfile



```go
type PlatformProfile struct {
    ID               string
    Contract         TargetContractMeta
    SDK              SDKMeta
    Launcher         LauncherMeta
    NativeDocs       []NativeDocSpec
    SurfaceTiers     []SurfaceSupport
    ManagedArtifacts []ManagedArtifactSpec
    Scaffold         ScaffoldMeta
    Validate         ValidateMeta
}
```

### func All

```go
func All() []PlatformProfile
```



### func Lookup

```go
func Lookup(name string) (PlatformProfile, bool)
```



## type SDKMeta



```go
type SDKMeta struct {
    PublicPackage   string
    InternalPackage string
    InternalImport  string
    Status          SupportStatus
    TransportModes  []TransportMode
    LiveTestProfile string
}
```

## type ScaffoldMeta



```go
type ScaffoldMeta struct {
    RequiredFiles  []string
    OptionalFiles  []string
    ForbiddenFiles []string
    TemplateFiles  []TemplateFile
}
```

## type SupportStatus



```go
type SupportStatus string
```

```go
const (
    StatusRuntimeSupported SupportStatus = "runtime_supported"
    StatusScaffoldOnly     SupportStatus = "scaffold_only"
    StatusDeferred         SupportStatus = "deferred"
)
```

## type SurfaceSupport



```go
type SurfaceSupport struct {
    Kind string
    Tier SurfaceTier
}
```

## type SurfaceTier



```go
type SurfaceTier string
```

```go
const (
    SurfaceTierStable          SurfaceTier = "stable"
    SurfaceTierBeta            SurfaceTier = "beta"
    SurfaceTierPreview         SurfaceTier = "preview"
    SurfaceTierPassthroughOnly SurfaceTier = "passthrough_only"
    SurfaceTierUnsupported     SurfaceTier = "unsupported"
)
```

## type TargetContractMeta



```go
type TargetContractMeta struct {
    PlatformFamily         PlatformFamily
    TargetClass            string
    TargetNoun             string
    ProductionClass        string
    RuntimeContract        string
    InstallModel           string
    DevModel               string
    ActivationModel        string
    NativeRoot             string
    ImportSupport          bool
    RenderSupport          bool
    ValidateSupport        bool
    PortableComponentKinds []string
    TargetComponentKinds   []string
    Summary                string
}
```

## type TemplateFile



```go
type TemplateFile struct {
    Path     string
    Template string
    Extra    bool
}
```

## type TransportMode



```go
type TransportMode string
```

```go
const (
    TransportProcess TransportMode = "process"
    TransportHybrid  TransportMode = "hybrid"
    TransportDaemon  TransportMode = "daemon"
)
```

## type ValidateMeta



```go
type ValidateMeta struct {
    RequiredFiles  []string
    ForbiddenFiles []string
    BuildTargets   []string
}
```
