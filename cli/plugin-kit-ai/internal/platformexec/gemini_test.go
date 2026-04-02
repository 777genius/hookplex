package platformexec

import (
	"path/filepath"
	"testing"
)

func TestValidateGeminiHookEntrypoints(t *testing.T) {
	body := []byte(`{
  "hooks": {
    "SessionStart": [{"matcher":"*","hooks":[{"type":"command","command":"./bin/demo GeminiSessionStart"}]}],
    "SessionEnd": [{"matcher":"*","hooks":[{"type":"command","command":"./bin/demo GeminiSessionEnd"}]}],
    "BeforeTool": [{"matcher":"*","hooks":[{"type":"command","command":"./bin/demo GeminiBeforeTool"}]}],
    "AfterTool": [{"matcher":"*","hooks":[{"type":"command","command":"./bin/demo GeminiAfterTool"}]}]
  }
}`)
	mismatches, err := validateGeminiHookEntrypoints(body, "./bin/demo")
	if err != nil {
		t.Fatal(err)
	}
	if len(mismatches) != 0 {
		t.Fatalf("mismatches = %v", mismatches)
	}
}

func TestValidateGeminiHookEntrypointsMismatch(t *testing.T) {
	body := []byte(`{
  "hooks": {
    "SessionStart": [{"matcher":"resume","hooks":[{"type":"command","command":"./bin/other GeminiSessionStart"}]}]
  }
}`)
	mismatches, err := validateGeminiHookEntrypoints(body, "./bin/demo")
	if err != nil {
		t.Fatal(err)
	}
	if len(mismatches) == 0 {
		t.Fatal("expected mismatches")
	}
}

func TestGeminiExtensionDirBase(t *testing.T) {
	t.Parallel()
	cwd := t.TempDir()
	got := geminiExtensionDirBase(filepath.Join(cwd, "."))
	if got != filepath.Base(cwd) {
		t.Fatalf("base = %q, want %q", got, filepath.Base(cwd))
	}
}
