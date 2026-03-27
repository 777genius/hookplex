package pluginkitairepo_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestPluginKitAIInitGoRuntimeLauncherFlow(t *testing.T) {
	for _, platform := range []string{"claude", "codex"} {
		t.Run(platform, func(t *testing.T) {
			root := RepoRoot(t)
			sdkDir := filepath.Join(root, "sdk", "plugin-kit-ai")
			pluginKitAIBin := buildPluginKitAI(t)

			plugRoot := runtimeProjectRoot(t)
			run := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", platform, "--runtime", "go", "-o", plugRoot, "--extras")
			if out, err := run.CombinedOutput(); err != nil {
				t.Fatalf("plugin-kit-ai init: %v\n%s", err, out)
			}

			replaceArg := "github.com/plugin-kit-ai/plugin-kit-ai/sdk=" + sdkDir
			modEdit := exec.Command("go", "mod", "edit", "-replace", replaceArg)
			modEdit.Dir = plugRoot
			modEdit.Env = append(os.Environ(), "GOWORK=off")
			if out, err := modEdit.CombinedOutput(); err != nil {
				t.Fatalf("go mod edit: %v\n%s", err, out)
			}

			validate := exec.Command(pluginKitAIBin, "validate", plugRoot, "--platform", platform)
			validate.Env = append(os.Environ(), "GOWORK=off")
			if out, err := validate.CombinedOutput(); err != nil {
				t.Fatalf("plugin-kit-ai validate: %v\n%s", err, out)
			}

			tidy := exec.Command("go", "mod", "tidy")
			tidy.Dir = plugRoot
			tidy.Env = append(os.Environ(), "GOWORK=off")
			if out, err := tidy.CombinedOutput(); err != nil {
				t.Fatalf("go mod tidy: %v\n%s", err, out)
			}

			binName := "genplug"
			if runtime.GOOS == "windows" {
				binName += ".exe"
			}
			build := exec.Command("go", "build", "-o", filepath.Join("bin", binName), "./cmd/genplug")
			build.Dir = plugRoot
			build.Env = append(os.Environ(), "GOWORK=off")
			if out, err := build.CombinedOutput(); err != nil {
				t.Fatalf("go build generated entrypoint: %v\n%s", err, out)
			}

			entry := filepath.Join(plugRoot, "bin", binName)
			switch platform {
			case "codex":
				cmd := exec.Command(entry, "notify", `{"client":"codex-tui"}`)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("run go codex launcher: %v\n%s", err, out)
				}
				if strings.TrimSpace(string(out)) != "" {
					t.Fatalf("stdout = %q, want empty", string(out))
				}
			case "claude":
				cmd := exec.Command(entry, "Stop")
				cmd.Stdin = strings.NewReader(`{"session_id":"s","transcript_path":"t","cwd":".","permission_mode":"default","hook_event_name":"Stop","stop_hook_active":false,"last_assistant_message":"ok"}`)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("run go claude launcher: %v\n%s", err, out)
				}
				if strings.TrimSpace(string(out)) != "{}" {
					t.Fatalf("stdout = %q, want {}", string(out))
				}
			}
		})
	}
}

func TestPluginKitAIInitNodeRuntimeSupportsTypeScriptBuildThroughLauncher(t *testing.T) {
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("node not in PATH")
	}
	if _, err := exec.LookPath("npm"); err != nil {
		t.Skip("npm not in PATH")
	}

	for _, platform := range []string{"claude", "codex"} {
		t.Run(platform, func(t *testing.T) {
			pluginKitAIBin := buildPluginKitAI(t)
			plugRoot := runtimeProjectRoot(t)
			run := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", platform, "--runtime", "node", "-o", plugRoot, "--extras")
			if out, err := run.CombinedOutput(); err != nil {
				t.Fatalf("plugin-kit-ai init: %v\n%s", err, out)
			}

			validate := exec.Command(pluginKitAIBin, "validate", plugRoot, "--platform", platform)
			if out, err := validate.CombinedOutput(); err != nil {
				t.Fatalf("plugin-kit-ai validate before TS conversion: %v\n%s", err, out)
			}

			writeRuntimeFile(t, plugRoot, filepath.Join("src", "main.ts"), tsPluginSource())
			writeRuntimeFile(t, plugRoot, "tsconfig.json", tsConfig())
			writeRuntimeFile(t, plugRoot, "package.json", tsPackageJSON())
			patchNodeLauncherForDist(t, plugRoot)

			npmInstall := exec.Command("npm", "install")
			npmInstall.Dir = plugRoot
			if out, err := npmInstall.CombinedOutput(); err != nil {
				t.Fatalf("npm install: %v\n%s", err, out)
			}

			npmBuild := exec.Command("npm", "run", "build")
			npmBuild.Dir = plugRoot
			if out, err := npmBuild.CombinedOutput(); err != nil {
				t.Fatalf("npm run build: %v\n%s", err, out)
			}

			entry := filepath.Join(plugRoot, "bin", "genplug")
			if runtime.GOOS == "windows" {
				entry += ".cmd"
			}
			switch platform {
			case "codex":
				cmd := exec.Command(entry, "notify", `{"client":"codex-tui"}`)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("run TS-over-node codex launcher: %v\n%s", err, out)
				}
				if strings.TrimSpace(string(out)) != "" {
					t.Fatalf("stdout = %q, want empty", string(out))
				}
			case "claude":
				cmd := exec.Command(entry, "Stop")
				cmd.Stdin = strings.NewReader(`{"hook_event_name":"Stop"}`)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("run TS-over-node claude launcher: %v\n%s", err, out)
				}
				if strings.TrimSpace(string(out)) != "{}" {
					t.Fatalf("stdout = %q, want {}", string(out))
				}
			}
		})
	}
}

func TestPluginKitAIInitPythonRuntimeLauncherFlow(t *testing.T) {
	if !pythonRuntimeAvailable() {
		t.Skip("python runtime not available for launcher flow")
	}

	for _, platform := range []string{"claude", "codex"} {
		t.Run(platform, func(t *testing.T) {
			pluginKitAIBin := buildPluginKitAI(t)
			plugRoot := runtimeProjectRoot(t)
			run := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", platform, "--runtime", "python", "-o", plugRoot, "--extras")
			if out, err := run.CombinedOutput(); err != nil {
				t.Fatalf("plugin-kit-ai init: %v\n%s", err, out)
			}

			validate := exec.Command(pluginKitAIBin, "validate", plugRoot, "--platform", platform)
			if out, err := validate.CombinedOutput(); err != nil {
				t.Fatalf("plugin-kit-ai validate python runtime: %v\n%s", err, out)
			}

			entry := filepath.Join(plugRoot, "bin", "genplug")
			if runtime.GOOS == "windows" {
				entry += ".cmd"
			}
			switch platform {
			case "codex":
				cmd := exec.Command(entry, "notify", `{"client":"codex-tui"}`)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("run python codex launcher: %v\n%s", err, out)
				}
				if strings.TrimSpace(string(out)) != "" {
					t.Fatalf("stdout = %q, want empty", string(out))
				}
			case "claude":
				cmd := exec.Command(entry, "Stop")
				cmd.Stdin = strings.NewReader(`{"hook_event_name":"Stop"}`)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("run python claude launcher: %v\n%s", err, out)
				}
				if strings.TrimSpace(string(out)) != "{}" {
					t.Fatalf("stdout = %q, want {}", string(out))
				}
			}
		})
	}
}

func TestPluginKitAIInitShellRuntimeLauncherFlow(t *testing.T) {
	if !shellRuntimeAvailable() {
		t.Skip("bash runtime not available for shell launcher flow")
	}

	for _, platform := range []string{"claude", "codex"} {
		t.Run(platform, func(t *testing.T) {
			pluginKitAIBin := buildPluginKitAI(t)
			plugRoot := runtimeProjectRoot(t)
			run := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", platform, "--runtime", "shell", "-o", plugRoot, "--extras")
			if out, err := run.CombinedOutput(); err != nil {
				t.Fatalf("plugin-kit-ai init: %v\n%s", err, out)
			}

			validate := exec.Command(pluginKitAIBin, "validate", plugRoot, "--platform", platform)
			if out, err := validate.CombinedOutput(); err != nil {
				t.Fatalf("plugin-kit-ai validate shell runtime: %v\n%s", err, out)
			}

			entry := filepath.Join(plugRoot, "bin", "genplug")
			if runtime.GOOS == "windows" {
				entry += ".cmd"
			}
			switch platform {
			case "codex":
				cmd := exec.Command(entry, "notify", `{"client":"codex-tui"}`)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("run shell codex launcher: %v\n%s", err, out)
				}
				if strings.TrimSpace(string(out)) != "" {
					t.Fatalf("stdout = %q, want empty", string(out))
				}
			case "claude":
				cmd := exec.Command(entry, "Stop")
				cmd.Stdin = strings.NewReader(`{"hook_event_name":"Stop"}`)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("run shell claude launcher: %v\n%s", err, out)
				}
				if strings.TrimSpace(string(out)) != "{}" {
					t.Fatalf("stdout = %q, want {}", string(out))
				}
			}
		})
	}
}

func writeRuntimeFile(t *testing.T, root, rel, body string) {
	t.Helper()
	full := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	mode := os.FileMode(0o644)
	if strings.HasSuffix(rel, ".sh") {
		mode = 0o755
	}
	if err := os.WriteFile(full, []byte(body), mode); err != nil {
		t.Fatal(err)
	}
}

func pythonRuntimeAvailable() bool {
	if runtime.GOOS == "windows" {
		if _, err := exec.LookPath("python"); err == nil {
			return true
		}
		if _, err := exec.LookPath("python3"); err == nil {
			return true
		}
		return false
	}
	_, err := exec.LookPath("python3")
	return err == nil
}

func patchNodeLauncherForDist(t *testing.T, root string) {
	t.Helper()
	if runtime.GOOS == "windows" {
		writeRuntimeFile(t, root, filepath.Join("bin", "genplug.cmd"), "@echo off\r\nsetlocal\r\nset \"ROOT=%~dp0..\"\r\nnode \"%ROOT%\\dist\\main.js\" %*\r\n")
		return
	}
	writeRuntimeFile(t, root, filepath.Join("bin", "genplug"), "#!/usr/bin/env bash\nset -euo pipefail\nROOT=\"$(CDPATH= cd -- \"$(dirname -- \"$0\")/..\" && pwd)\"\nif ! command -v node >/dev/null 2>&1; then\n  echo \"plugin-kit-ai launcher: node not found\" >&2\n  exit 1\nfi\nNODE=\"$(command -v node)\"\nexec \"$NODE\" \"$ROOT/dist/main.js\" \"$@\"\n")
}

func tsPluginSource() string {
	return `import fs from "node:fs";

function readStdin(): string {
  return fs.readFileSync(0, "utf8");
}

function handleClaude(): number {
  const event = JSON.parse(readStdin());
  void event;
  process.stdout.write("{}");
  return 0;
}

function handleCodex(): number {
  const payload = process.argv[3];
  if (!payload) {
    process.stderr.write("missing notify payload\n");
    return 1;
  }
  const event = JSON.parse(payload);
  void event;
  return 0;
}

function main(): number {
  const hookName = process.argv[2];
  if (!hookName) {
    process.stderr.write("usage: main.ts <hook-name>\n");
    return 1;
  }
  if (hookName === "notify") {
    return handleCodex();
  }
  return handleClaude();
}

process.exit(main());
`
}

func tsConfig() string {
	return `{
  "compilerOptions": {
    "target": "ES2022",
    "module": "NodeNext",
    "moduleResolution": "NodeNext",
    "types": ["node"],
    "outDir": "dist",
    "rootDir": "src",
    "strict": true,
    "skipLibCheck": true
  },
  "include": ["src/**/*.ts"]
}
`
}

func tsPackageJSON() string {
	return `{
  "name": "genplug",
  "version": "0.1.0",
  "private": true,
  "type": "module",
  "scripts": {
    "build": "tsc -p tsconfig.json"
  },
  "devDependencies": {
    "@types/node": "^24.0.0",
    "typescript": "^5.9.0"
  }
}
`
}
