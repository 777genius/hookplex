package pluginkitai

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/777genius/plugin-kit-ai/sdk/claude"
	"github.com/777genius/plugin-kit-ai/sdk/codex"
	"github.com/777genius/plugin-kit-ai/sdk/gemini"
	"github.com/777genius/plugin-kit-ai/sdk/internal/descriptors/gen"
	"github.com/777genius/plugin-kit-ai/sdk/internal/runtime"
	"github.com/777genius/plugin-kit-ai/sdk/internal/runtime/process"
)

// IO aliases the runtime I/O contract used by the SDK app host.
type IO = runtime.IO

// Env aliases the runtime environment reader used by invocation resolution.
type Env = runtime.Env

// Logger aliases the structured logger interface accepted by the SDK app host.
type Logger = runtime.Logger

// Middleware aliases the SDK middleware function signature.
type Middleware = runtime.Middleware

// Next aliases the middleware continuation function.
type Next = runtime.Next

// Handled aliases the typed handler result container.
type Handled = runtime.Handled

// InvocationContext aliases the metadata that accompanies a decoded invocation.
type InvocationContext = runtime.InvocationContext

// Result aliases the low-level runtime result written back to the host process.
type Result = runtime.Result

// NopLogger aliases the logger implementation that drops all log records.
type NopLogger = runtime.NopLogger

// CapabilityID aliases the normalized cross-platform capability identifier.
type CapabilityID = runtime.CapabilityID

// SupportStatus aliases the support-level enum used by generated support entries.
type SupportStatus = runtime.SupportStatus

// MaturityLevel aliases the API maturity enum exposed by support metadata.
type MaturityLevel = runtime.MaturityLevel

// TransportMode aliases the runtime transport mode enum for supported hooks.
type TransportMode = runtime.TransportMode

// SupportEntry aliases a generated public support-matrix row.
type SupportEntry = runtime.SupportEntry

// Config configures a root SDK app instance before handlers are registered.
type Config struct {
	// Name is the human-readable app label used in diagnostics and examples.
	Name string
	// Args overrides the process argv used to resolve the current invocation.
	Args []string
	// IO overrides the stdin/stdout/stderr implementation used by Run.
	IO IO
	// Env overrides environment lookups used during invocation resolution.
	Env Env
	// Logger overrides structured logging emitted by the runtime engine.
	Logger Logger
}

// App owns middleware, handler registration, and invocation dispatch.
type App struct {
	mu      sync.Mutex
	runDone bool
	name    string
	args    []string
	io      runtime.IO
	env     runtime.Env
	logger  runtime.Logger

	handlers *runtime.HandlerRegistry
	mws      []runtime.Middleware
	custom   *customRegistry
}

// New builds an App with sane defaults for argv, process I/O, env, and logging.
func New(cfg Config) *App {
	args := cfg.Args
	if len(args) == 0 {
		args = os.Args
	}
	io := cfg.IO
	if io == nil {
		io = process.IO{}
	}
	env := cfg.Env
	if env == nil {
		env = process.Env{}
	}
	logger := cfg.Logger
	if logger == nil {
		logger = runtime.NopLogger{}
	}
	return &App{
		name:     cfg.Name,
		args:     append([]string(nil), args...),
		io:       io,
		env:      env,
		logger:   logger,
		handlers: runtime.NewHandlerRegistry(),
		custom:   newCustomRegistry(),
	}
}

// Supported returns a copy of the generated public support matrix entries.
func Supported() []SupportEntry {
	entries := gen.AllSupportEntries()
	out := make([]SupportEntry, len(entries))
	copy(out, entries)
	return out
}

// Use appends middleware that wraps all subsequent handler dispatch.
func (a *App) Use(mw Middleware) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.runDone {
		panic("plugin-kit-ai: Use after Run")
	}
	if mw == nil {
		return
	}
	a.mws = append(a.mws, mw)
}

// Claude returns a registrar for Claude-specific hook handlers.
func (a *App) Claude() *claude.Registrar {
	return claude.NewRegistrar(registrarBackend{app: a})
}

// Codex returns a registrar for Codex-specific event handlers.
func (a *App) Codex() *codex.Registrar {
	return codex.NewRegistrar(registrarBackend{app: a})
}

// Gemini returns a registrar for Gemini-specific hook handlers.
func (a *App) Gemini() *gemini.Registrar {
	return gemini.NewRegistrar(registrarBackend{app: a})
}

// Run dispatches the current process invocation with context.Background().
func (a *App) Run() int {
	return a.RunContext(context.Background())
}

// RunContext dispatches the current process invocation using the supplied context.
func (a *App) RunContext(ctx context.Context) int {
	a.mu.Lock()
	args := append([]string(nil), a.args...)
	io := a.io
	env := a.env
	logger := a.logger
	mws := append([]runtime.Middleware(nil), a.mws...)
	handlers := a.handlers
	a.runDone = true
	a.mu.Unlock()

	engine := runtime.Engine{
		Args:          args,
		IO:            io,
		Env:           env,
		Logger:        logger,
		Resolver:      a.resolveInvocation,
		Lookup:        a.lookupDescriptor,
		BuildEnvelope: process.BuildEnvelope,
		Handlers:      handlers,
		Middleware:    append([]runtime.Middleware{runtime.RecoveryMiddleware(logger)}, mws...),
	}
	res := engine.Dispatch(ctx)
	if res.Stderr != "" {
		_ = io.WriteStderr(res.Stderr)
	}
	if res.ExitCode == 0 && len(res.Stdout) > 0 {
		if err := io.WriteStdout(res.Stdout); err != nil {
			_ = io.WriteStderr(err.Error() + "\n")
			return 1
		}
	}
	return res.ExitCode
}

type registrarBackend struct {
	app *App
}

func (b registrarBackend) Register(platform runtime.PlatformID, event runtime.EventID, handler runtime.TypedHandler) {
	b.app.mu.Lock()
	defer b.app.mu.Unlock()
	if b.app.runDone {
		panic("plugin-kit-ai: register after Run")
	}
	b.app.handlers.Register(platform, event, handler)
}

func (b registrarBackend) RegisterCustom(rawName string, desc runtime.Descriptor, handler runtime.TypedHandler) error {
	b.app.mu.Lock()
	defer b.app.mu.Unlock()
	if b.app.runDone {
		panic("plugin-kit-ai: register after Run")
	}
	if err := b.app.custom.Register(rawName, desc); err != nil {
		return err
	}
	b.app.handlers.Register(desc.Platform, desc.Event, handler)
	return nil
}

type customRegistry struct {
	byRaw  map[string]runtime.Invocation
	byDesc map[customKey]runtime.Descriptor
}

type customKey struct {
	platform runtime.PlatformID
	event    runtime.EventID
}

func newCustomRegistry() *customRegistry {
	return &customRegistry{
		byRaw:  make(map[string]runtime.Invocation),
		byDesc: make(map[customKey]runtime.Descriptor),
	}
}

func (r *customRegistry) Register(rawName string, desc runtime.Descriptor) error {
	name := strings.TrimSpace(rawName)
	if name == "" {
		return fmt.Errorf("custom hook name required")
	}
	if desc.Platform == "" || desc.Event == "" {
		return fmt.Errorf("custom hook descriptor requires platform and event")
	}
	if desc.Decode == nil || desc.Encode == nil {
		return fmt.Errorf("custom hook descriptor requires decode and encode")
	}
	if _, ok := gen.Lookup(desc.Platform, desc.Event); ok {
		return fmt.Errorf("custom hook %s/%s conflicts with built-in descriptor", desc.Platform, desc.Event)
	}
	if _, err := gen.ResolveInvocation([]string{"plugin-kit-ai", name}, nil); err == nil {
		return fmt.Errorf("custom hook name %q conflicts with built-in invocation", name)
	}
	rawKey := strings.ToLower(name)
	if inv, ok := r.byRaw[rawKey]; ok {
		return fmt.Errorf("custom hook name %q already registered for %s/%s", name, inv.Platform, inv.Event)
	}
	key := customKey{platform: desc.Platform, event: desc.Event}
	if _, ok := r.byDesc[key]; ok {
		return fmt.Errorf("custom hook descriptor already registered for %s/%s", desc.Platform, desc.Event)
	}
	r.byRaw[rawKey] = runtime.Invocation{Platform: desc.Platform, Event: desc.Event, RawName: name}
	r.byDesc[key] = desc
	return nil
}

func (a *App) resolveInvocation(args []string, env runtime.Env) (runtime.Invocation, error) {
	if inv, err := gen.ResolveInvocation(args, env); err == nil {
		return inv, nil
	}
	if len(args) < 2 {
		return runtime.Invocation{}, fmt.Errorf("usage: <binary> <hookName>")
	}
	raw := args[1]
	a.mu.Lock()
	defer a.mu.Unlock()
	if inv, ok := a.custom.byRaw[strings.ToLower(raw)]; ok {
		inv.RawName = raw
		return inv, nil
	}
	return runtime.Invocation{}, fmt.Errorf("unknown invocation %q", raw)
}

func (a *App) lookupDescriptor(platform runtime.PlatformID, event runtime.EventID) (runtime.Descriptor, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if desc, ok := a.custom.byDesc[customKey{platform: platform, event: event}]; ok {
		return desc, true
	}
	return gen.Lookup(platform, event)
}
