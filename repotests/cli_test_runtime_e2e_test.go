package pluginkitairepo_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestPluginKitAITestClaudeShellStableFlow(t *testing.T) {
	if !shellRuntimeAvailable() {
		t.Skip("bash runtime not available for shell test flow")
	}

	pluginKitAIBin := buildPluginKitAI(t)
	plugRoot := runtimeProjectRoot(t)
	initCmd := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", "claude", "--runtime", "shell", "-o", plugRoot)
	if out, err := initCmd.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai init claude shell: %v\n%s", err, out)
	}

	writeRuntimeFile(t, plugRoot, filepath.Join("fixtures", "claude", "Stop.json"), `{"session_id":"s","cwd":"/tmp","hook_event_name":"Stop"}`)
	writeRuntimeFile(t, plugRoot, filepath.Join("fixtures", "claude", "PreToolUse.json"), `{"session_id":"s","cwd":"/tmp","hook_event_name":"PreToolUse","tool_name":"Bash","tool_input":{"command":"echo hi"}}`)
	writeRuntimeFile(t, plugRoot, filepath.Join("fixtures", "claude", "UserPromptSubmit.json"), `{"session_id":"s","cwd":"/tmp","hook_event_name":"UserPromptSubmit","prompt":"hello"}`)

	update := exec.Command(pluginKitAIBin, "test", plugRoot, "--platform", "claude", "--all", "--update-golden")
	if out, err := update.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai test claude update-golden: %v\n%s", err, out)
	}

	match := exec.Command(pluginKitAIBin, "test", plugRoot, "--platform", "claude", "--all")
	out, err := match.CombinedOutput()
	if err != nil {
		t.Fatalf("plugin-kit-ai test claude match: %v\n%s", err, out)
	}
	text := string(out)
	for _, want := range []string{
		"PASS claude/Stop",
		"PASS claude/PreToolUse",
		"PASS claude/UserPromptSubmit",
		"golden=matched",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("test output missing %q:\n%s", want, text)
		}
	}
	for _, rel := range []string{
		filepath.Join("goldens", "claude", "Stop.stdout"),
		filepath.Join("goldens", "claude", "PreToolUse.stdout"),
		filepath.Join("goldens", "claude", "UserPromptSubmit.stdout"),
	} {
		if _, err := os.Stat(filepath.Join(plugRoot, rel)); err != nil {
			t.Fatalf("expected golden %s: %v", rel, err)
		}
	}
}

func TestPluginKitAITestCodexShellNotifyJSONFlow(t *testing.T) {
	if !shellRuntimeAvailable() {
		t.Skip("bash runtime not available for shell test flow")
	}

	pluginKitAIBin := buildPluginKitAI(t)
	plugRoot := runtimeProjectRoot(t)
	initCmd := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", "codex-runtime", "--runtime", "shell", "-o", plugRoot)
	if out, err := initCmd.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai init codex shell: %v\n%s", err, out)
	}
	writeRuntimeFile(t, plugRoot, filepath.Join("fixtures", "codex-runtime", "Notify.json"), `{"client":"codex-tui"}`)

	update := exec.Command(pluginKitAIBin, "test", plugRoot, "--platform", "codex-runtime", "--event", "Notify", "--update-golden")
	if out, err := update.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai test codex update-golden: %v\n%s", err, out)
	}

	match := exec.Command(pluginKitAIBin, "test", plugRoot, "--platform", "codex-runtime", "--event", "Notify", "--format", "json")
	out, err := match.CombinedOutput()
	if err != nil {
		t.Fatalf("plugin-kit-ai test codex json: %v\n%s", err, out)
	}
	text := string(out)
	for _, want := range []string{
		`"passed": true`,
		`"event": "Notify"`,
		`"golden_status": "matched"`,
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("json output missing %q:\n%s", want, text)
		}
	}
}
