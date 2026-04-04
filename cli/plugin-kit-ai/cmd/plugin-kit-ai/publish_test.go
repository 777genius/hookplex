package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/777genius/plugin-kit-ai/cli/internal/app"
)

type fakePublishRunner struct {
	result app.PluginPublishResult
	err    error
	opts   app.PluginPublishOptions
}

func (f *fakePublishRunner) Publish(opts app.PluginPublishOptions) (app.PluginPublishResult, error) {
	f.opts = opts
	return f.result, f.err
}

func TestPublishDelegatesToRunner(t *testing.T) {
	t.Parallel()
	runner := &fakePublishRunner{
		result: app.PluginPublishResult{Lines: []string{"published"}},
	}
	cmd := newPublishCmd(runner)
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{".", "--channel", "codex-marketplace", "--dest", "/tmp/market", "--dry-run"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	if runner.opts.Channel != "codex-marketplace" || runner.opts.Dest != "/tmp/market" || runner.opts.Root != "." || !runner.opts.DryRun {
		t.Fatalf("opts = %+v", runner.opts)
	}
	if !strings.Contains(buf.String(), "published") {
		t.Fatalf("output = %s", buf.String())
	}
}

func TestPublishHelpMentionsBoundedChannels(t *testing.T) {
	t.Parallel()
	cmd := newPublishCmd(&fakePublishRunner{})
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"--help"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	output := buf.String()
	for _, want := range []string{
		`publish channel ("codex-marketplace" or "claude-marketplace")`,
		"destination marketplace root directory",
		"preview the materialized publish result without writing changes",
		"codex-marketplace",
		"claude-marketplace",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("help output missing %q:\n%s", want, output)
		}
	}
}
