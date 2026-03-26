package hookplexrepo_test

import (
	"os/exec"
	"path/filepath"
	"testing"
)

// TestSDKModule runs the full SDK test suite (module sdk/hookplex).
// Repository root is a workspace-only module so `go test ./...` from the hookplex
// repo root satisfies the iter1 DoD while keeping sdk/hookplex/go.mod as in the plan.
func TestSDKModule(t *testing.T) {
	root := RepoRoot(t)
	sdkDir := filepath.Join(root, "sdk", "hookplex")
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = sdkDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go test in sdk/hookplex: %v\n%s", err, out)
	}
}
