package hookplexrepo_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestHookplexInitGeneratesBuildableModule(t *testing.T) {
	for _, platform := range []string{"claude", "codex"} {
		t.Run(platform, func(t *testing.T) {
			root := RepoRoot(t)
			cliDir := filepath.Join(root, "cli", "hookplex")
			sdkDir := filepath.Join(root, "sdk", "hookplex")

			binDir := t.TempDir()
			bin := filepath.Join(binDir, "hookplex")
			build := exec.Command("go", "build", "-o", bin, "./cmd/hookplex")
			build.Dir = cliDir
			if out, err := build.CombinedOutput(); err != nil {
				t.Fatalf("build hookplex: %v\n%s", err, out)
			}

			plugRoot := t.TempDir()
			run := exec.Command(bin, "init", "genplug", "--platform", platform, "-o", plugRoot, "--extras")
			if out, err := run.CombinedOutput(); err != nil {
				t.Fatalf("hookplex init: %v\n%s", err, out)
			}

			replaceArg := "github.com/hookplex/hookplex/sdk=" + sdkDir
			modEdit := exec.Command("go", "mod", "edit", "-replace", replaceArg)
			modEdit.Dir = plugRoot
			if out, err := modEdit.CombinedOutput(); err != nil {
				t.Fatalf("go mod edit: %v\n%s", err, out)
			}

			validate := exec.Command(bin, "validate", plugRoot, "--platform", platform)
			validate.Env = append(os.Environ(), "GOWORK=off")
			if out, err := validate.CombinedOutput(); err != nil {
				t.Fatalf("hookplex validate: %v\n%s", err, out)
			}

			tidy := exec.Command("go", "mod", "tidy")
			tidy.Dir = plugRoot
			tidy.Env = append(os.Environ(), "GOWORK=off")
			if out, err := tidy.CombinedOutput(); err != nil {
				t.Fatalf("go mod tidy in generated module: %v\n%s", err, out)
			}

			test := exec.Command("go", "test", "./...")
			test.Dir = plugRoot
			test.Env = append(os.Environ(), "GOWORK=off")
			if out, err := test.CombinedOutput(); err != nil {
				t.Fatalf("go test in generated module: %v\n%s", err, out)
			}

			vet := exec.Command("go", "vet", "./...")
			vet.Dir = plugRoot
			vet.Env = append(os.Environ(), "GOWORK=off")
			if out, err := vet.CombinedOutput(); err != nil {
				t.Fatalf("go vet in generated module: %v\n%s", err, out)
			}
		})
	}
}
