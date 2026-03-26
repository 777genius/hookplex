package hookplex

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/hookplex/hookplex/sdk/claude"
	"github.com/hookplex/hookplex/sdk/codex"
)

type testIO struct {
	in  []byte
	out bytes.Buffer
	err bytes.Buffer
}

func (t *testIO) ReadStdin(ctx context.Context) ([]byte, error) {
	return t.in, ctx.Err()
}

func (t *testIO) WriteStdout(b []byte) error {
	_, err := t.out.Write(b)
	return err
}

func (t *testIO) WriteStderr(s string) error {
	_, err := t.err.WriteString(s)
	return err
}

type testEnv map[string]string

func (e testEnv) LookupEnv(k string) (string, bool) {
	v, ok := e[k]
	return v, ok
}

func TestApp_ClaudeStop(t *testing.T) {
	iox := &testIO{in: []byte(`{"session_id":"s","cwd":"/","hook_event_name":"Stop"}`)}
	app := New(Config{
		Name: "t",
		Args: []string{"hookplex", "Stop"},
		IO:   iox,
		Env:  testEnv{},
	})
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response {
		return claude.Allow()
	})
	if c := app.Run(); c != 0 {
		t.Fatalf("exit %d stderr=%q", c, iox.err.String())
	}
	if got := iox.out.String(); got != "{}" {
		t.Fatalf("stdout = %q", got)
	}
}

func TestApp_CodexNotify(t *testing.T) {
	iox := &testIO{}
	app := New(Config{
		Name: "t",
		Args: []string{"hookplex", "notify", `{"client":"codex-tui","ignored":true}`},
		IO:   iox,
		Env:  testEnv{},
	})
	app.Codex().OnNotify(func(e *codex.NotifyEvent) *codex.Response {
		if e.Client != "codex-tui" {
			t.Fatalf("client = %q", e.Client)
		}
		if string(e.RawJSON()) == "" {
			t.Fatal("raw json missing")
		}
		return codex.Continue()
	})
	if c := app.Run(); c != 0 {
		t.Fatalf("exit %d stderr=%q", c, iox.err.String())
	}
	if iox.out.Len() != 0 {
		t.Fatalf("stdout should be empty, got %q", iox.out.String())
	}
}

func TestApp_UnknownInvocation(t *testing.T) {
	iox := &testIO{}
	app := New(Config{
		Name: "t",
		Args: []string{"hookplex", "bogus"},
		IO:   iox,
		Env:  testEnv{},
	})
	if c := app.Run(); c != 1 {
		t.Fatalf("exit %d stderr=%q", c, iox.err.String())
	}
	if got := iox.err.String(); got != "unknown invocation \"bogus\"\n" {
		t.Fatalf("stderr = %q", got)
	}
}

func TestApp_CodexNotifyMissingPayload(t *testing.T) {
	iox := &testIO{}
	app := New(Config{
		Name: "t",
		Args: []string{"hookplex", "notify"},
		IO:   iox,
		Env:  testEnv{},
	})
	if c := app.Run(); c != 1 {
		t.Fatalf("exit %d stderr=%q", c, iox.err.String())
	}
	if got := iox.err.String(); got != "decode codex notify input: missing JSON payload argument\n" {
		t.Fatalf("stderr = %q", got)
	}
}

func TestApp_CodexNotifyMalformedPayload(t *testing.T) {
	iox := &testIO{}
	app := New(Config{
		Name: "t",
		Args: []string{"hookplex", "notify", "{"},
		IO:   iox,
		Env:  testEnv{},
	})
	if c := app.Run(); c != 1 {
		t.Fatalf("exit %d stderr=%q", c, iox.err.String())
	}
	if got := iox.err.String(); !strings.Contains(got, "decode codex notify input:") {
		t.Fatalf("stderr = %q", got)
	}
}

func TestApp_RegisterAfterRunPanics(t *testing.T) {
	iox := &testIO{in: []byte(`{"session_id":"s","cwd":"/","hook_event_name":"Stop"}`)}
	app := New(Config{
		Name: "t",
		Args: []string{"hookplex", "Stop"},
		IO:   iox,
		Env:  testEnv{},
	})
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response { return claude.Allow() })
	_ = app.Run()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
	}()
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response { return claude.Allow() })
}

func TestApp_LastRegistrationWins(t *testing.T) {
	iox := &testIO{in: []byte(`{"session_id":"s","cwd":"/","hook_event_name":"Stop"}`)}
	app := New(Config{
		Name: "t",
		Args: []string{"hookplex", "Stop"},
		IO:   iox,
		Env:  testEnv{},
	})
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response { return claude.Block("first") })
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response { return claude.Allow() })
	if c := app.Run(); c != 0 {
		t.Fatalf("exit %d stderr=%q", c, iox.err.String())
	}
	if got := iox.out.String(); got != "{}" {
		t.Fatalf("stdout = %q", got)
	}
}

func TestApp_MiddlewareRuns(t *testing.T) {
	iox := &testIO{in: []byte(`{"session_id":"s","cwd":"/","hook_event_name":"Stop"}`)}
	app := New(Config{
		Name: "t",
		Args: []string{"hookplex", "Stop"},
		IO:   iox,
		Env:  testEnv{},
	})
	var ran bool
	app.Use(func(next Next) Next {
		return func(ctx InvocationContext) Handled {
			ran = true
			return next(ctx)
		}
	})
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response { return claude.Allow() })
	if c := app.Run(); c != 0 {
		t.Fatalf("exit %d", c)
	}
	if !ran {
		t.Fatal("middleware did not run")
	}
}
