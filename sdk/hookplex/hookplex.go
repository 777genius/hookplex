package hookplex

import (
	"context"
	"os"
	"sync"

	"github.com/hookplex/hookplex/sdk/claude"
	"github.com/hookplex/hookplex/sdk/codex"
	"github.com/hookplex/hookplex/sdk/internal/descriptors/gen"
	"github.com/hookplex/hookplex/sdk/internal/runtime"
	"github.com/hookplex/hookplex/sdk/internal/runtime/process"
)

type IO = runtime.IO
type Env = runtime.Env
type Logger = runtime.Logger
type Middleware = runtime.Middleware
type Next = runtime.Next
type Handled = runtime.Handled
type InvocationContext = runtime.InvocationContext
type Result = runtime.Result
type NopLogger = runtime.NopLogger
type CapabilityID = runtime.CapabilityID
type SupportStatus = runtime.SupportStatus
type MaturityLevel = runtime.MaturityLevel
type TransportMode = runtime.TransportMode
type SupportEntry = runtime.SupportEntry

type Config struct {
	Name   string
	Args   []string
	IO     IO
	Env    Env
	Logger Logger
}

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
}

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
	}
}

func Supported() []SupportEntry {
	entries := gen.AllSupportEntries()
	out := make([]SupportEntry, len(entries))
	copy(out, entries)
	return out
}

func (a *App) Use(mw Middleware) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.runDone {
		panic("hookplex: Use after Run")
	}
	if mw == nil {
		return
	}
	a.mws = append(a.mws, mw)
}

func (a *App) Claude() *claude.Registrar {
	return claude.NewRegistrar(registrarBackend{app: a})
}

func (a *App) Codex() *codex.Registrar {
	return codex.NewRegistrar(registrarBackend{app: a})
}

func (a *App) Run() int {
	return a.RunContext(context.Background())
}

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
		Resolver:      gen.ResolveInvocation,
		Lookup:        gen.Lookup,
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
		panic("hookplex: register after Run")
	}
	b.app.handlers.Register(platform, event, handler)
}
