package hookplexrepo_test

import (
	"bufio"
	"encoding/json"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

const rootGoModModuleLine = "module github.com/hookplex/hookplex"

// RepoRoot returns the hookplex monorepo root (directory containing the anchor go.mod).
// Walks up from the caller's file until it finds go.mod with module github.com/hookplex/hookplex.
// Override with HOOKPLEX_REPO_ROOT for debugging.
func RepoRoot(tb testing.TB) string {
	tb.Helper()
	if v := strings.TrimSpace(os.Getenv("HOOKPLEX_REPO_ROOT")); v != "" {
		return v
	}
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		tb.Fatal("runtime.Caller")
	}
	dir := filepath.Dir(file)
	for {
		modPath := filepath.Join(dir, "go.mod")
		if fileExists(modPath) && isAnchorGoMod(modPath) {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			tb.Fatalf("hookplex repo root not found from %s (expected %s in a parent go.mod)", file, rootGoModModuleLine)
		}
		dir = parent
	}
}

func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

func isAnchorGoMod(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	if !s.Scan() {
		return false
	}
	return strings.TrimSpace(s.Text()) == rootGoModModuleLine
}

func buildHookplex(t *testing.T) string {
	t.Helper()
	root := RepoRoot(t)
	cliDir := filepath.Join(root, "cli", "hookplex")
	binDir := t.TempDir()
	name := "hookplex"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	hookplexBin := filepath.Join(binDir, name)
	build := exec.Command("go", "build", "-o", hookplexBin, "./cmd/hookplex")
	build.Dir = cliDir
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build hookplex: %v\n%s", err, out)
	}
	return hookplexBin
}

func runInstall(t *testing.T, hookplexBin, workDir, apiBase string, extraArgs ...string) (exitCode int, output []byte) {
	t.Helper()
	args := append([]string{"install", "o/r", "--github-api-base", apiBase}, extraArgs...)
	cmd := exec.Command(hookplexBin, args...)
	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return ee.ExitCode(), out
		}
		t.Fatalf("install: %v\n%s", err, out)
	}
	return 0, out
}

func runHookplexInstall(t *testing.T, hookplexBin, workDir, ownerRepo string, extraArgs ...string) (exitCode int, output []byte) {
	t.Helper()
	args := append([]string{"install", ownerRepo}, extraArgs...)
	cmd := exec.Command(hookplexBin, args...)
	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return ee.ExitCode(), out
		}
		t.Fatalf("hookplex install: %v\n%s", err, out)
	}
	return 0, out
}

// buildHookplexE2E builds sdk/hookplex/cmd/hookplex-e2e into a temp dir and returns the binary path.
func buildHookplexE2E(t *testing.T) string {
	t.Helper()
	root := RepoRoot(t)
	sdkDir := filepath.Join(root, "sdk", "hookplex")
	binDir := t.TempDir()
	name := "hookplex-e2e"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	out := filepath.Join(binDir, name)
	cmd := exec.Command("go", "build", "-o", out, "./cmd/hookplex-e2e")
	cmd.Dir = sdkDir
	if b, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build hookplex-e2e: %v\n%s", err, b)
	}
	return out
}

func requireBindTests(t *testing.T) {
	t.Helper()
	if os.Getenv("HOOKPLEX_BIND_TESTS") == "1" {
		return
	}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Skipf("requires loopback bind support or HOOKPLEX_BIND_TESTS=1: %v", err)
	}
	_ = ln.Close()
}

type traceRec struct {
	Hook    string `json:"hook"`
	Outcome string `json:"outcome"`
	Client  string `json:"client,omitempty"`
	RawJSON string `json:"raw_json,omitempty"`
}

func readTraceLines(t *testing.T, path string) []string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		t.Fatal(err)
	}
	var lines []string
	s := bufio.NewScanner(strings.NewReader(string(b)))
	for s.Scan() {
		if strings.TrimSpace(s.Text()) != "" {
			lines = append(lines, s.Text())
		}
	}
	return lines
}

func waitForTraceLines(t *testing.T, path string, timeout time.Duration) []string {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for {
		lines := readTraceLines(t, path)
		if len(lines) > 0 || time.Now().After(deadline) {
			return lines
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func traceHas(t *testing.T, lines []string, wantHook, wantOutcome string) bool {
	t.Helper()
	for _, line := range lines {
		var rec traceRec
		if json.Unmarshal([]byte(line), &rec) != nil {
			continue
		}
		if rec.Hook == wantHook && rec.Outcome == wantOutcome {
			return true
		}
	}
	return false
}

func traceFind(t *testing.T, lines []string, wantHook string) (traceRec, bool) {
	t.Helper()
	for _, line := range lines {
		var rec traceRec
		if json.Unmarshal([]byte(line), &rec) != nil {
			continue
		}
		if rec.Hook == wantHook {
			return rec, true
		}
	}
	return traceRec{}, false
}
