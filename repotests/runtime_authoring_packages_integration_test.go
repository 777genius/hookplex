package pluginkitairepo_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestPythonRuntimePackageContractFiles(t *testing.T) {
	t.Parallel()
	root := RepoRoot(t)

	for _, trackedPath := range []string{
		"python/plugin-kit-ai-runtime/pyproject.toml",
		"python/plugin-kit-ai-runtime/README.md",
		"python/plugin-kit-ai-runtime/src/plugin_kit_ai_runtime/__init__.py",
	} {
		tracked := exec.Command("git", "ls-files", "--error-unmatch", trackedPath)
		tracked.Dir = root
		if out, err := tracked.CombinedOutput(); err != nil {
			t.Fatalf("python runtime package file must be tracked in git (%s): %v\n%s", trackedPath, err, out)
		}
	}

	pyproject := readRepoFile(t, root, "python", "plugin-kit-ai-runtime", "pyproject.toml")
	for _, want := range []string{
		`name = "plugin-kit-ai-runtime"`,
		`requires-python = ">=3.10"`,
		`dynamic = ["version"]`,
		`version = { attr = "plugin_kit_ai_runtime.__version__" }`,
	} {
		mustContain(t, pyproject, want)
	}

	initPy := readRepoFile(t, root, "python", "plugin-kit-ai-runtime", "src", "plugin_kit_ai_runtime", "__init__.py")
	for _, want := range []string{
		`__version__ = "0.0.0-development"`,
		`CLAUDE_STABLE_HOOKS = (`,
		`CLAUDE_EXTENDED_HOOKS = (`,
		"class ClaudeApp:",
		"class CodexApp:",
	} {
		mustContain(t, initPy, want)
	}

	readme := readRepoFile(t, root, "python", "plugin-kit-ai-runtime", "README.md")
	for _, want := range []string{
		"pip install plugin-kit-ai-runtime",
		"shared dependency instead of copying a local helper file",
		"Go is still the recommended path",
		"stable supported lane",
	} {
		mustContain(t, readme, want)
	}

	workflow := readRepoFile(t, root, ".github", "workflows", "pypi-runtime-publish.yml")
	for _, want := range []string{
		"name: PyPI Runtime Publish",
		`workflows: ["Release Assets"]`,
		"plugin-kit-ai-runtime",
		"id-token: write",
		"pypa/gh-action-pypi-publish@release/v1",
	} {
		mustContain(t, workflow, want)
	}
}

func TestPythonRuntimePackageClaudeAndCodexSmoke(t *testing.T) {
	t.Parallel()
	requirePythonRuntime(t)

	root := RepoRoot(t)
	appDir := t.TempDir()
	pkgSrc := filepath.Join(root, "python", "plugin-kit-ai-runtime", "src")

	codexScript := filepath.Join(appDir, "codex_app.py")
	if err := os.WriteFile(codexScript, []byte(`from plugin_kit_ai_runtime import CodexApp, continue_

app = CodexApp()


@app.on_notify
def on_notify(event):
    if event.get("client") != "codex-tui":
        raise RuntimeError(f"unexpected client: {event}")
    return continue_()


raise SystemExit(app.run())
`), 0o644); err != nil {
		t.Fatal(err)
	}

	codex := exec.Command("python3", codexScript, "notify", `{"client":"codex-tui"}`)
	codex.Env = append(os.Environ(), "PYTHONPATH="+pkgSrc)
	var codexStdout bytes.Buffer
	var codexStderr bytes.Buffer
	codex.Stdout = &codexStdout
	codex.Stderr = &codexStderr
	if err := codex.Run(); err != nil {
		t.Fatalf("python runtime package codex smoke: %v\nstderr=%s", err, codexStderr.String())
	}
	if strings.TrimSpace(codexStdout.String()) != "" {
		t.Fatalf("codex stdout = %q, want empty", codexStdout.String())
	}

	claudeScript := filepath.Join(appDir, "claude_app.py")
	if err := os.WriteFile(claudeScript, []byte(`from plugin_kit_ai_runtime import CLAUDE_STABLE_HOOKS, ClaudeApp, allow

app = ClaudeApp(allowed_hooks=CLAUDE_STABLE_HOOKS, usage="claude_app.py <hook-name>")


@app.on_stop
def on_stop(event):
    if event.get("hook_event_name") != "Stop":
        raise RuntimeError(f"unexpected hook payload: {event}")
    return allow()


raise SystemExit(app.run())
`), 0o644); err != nil {
		t.Fatal(err)
	}

	claude := exec.Command("python3", claudeScript, "Stop")
	claude.Env = append(os.Environ(), "PYTHONPATH="+pkgSrc)
	claude.Stdin = strings.NewReader(`{"hook_event_name":"Stop"}`)
	var claudeStdout bytes.Buffer
	var claudeStderr bytes.Buffer
	claude.Stdout = &claudeStdout
	claude.Stderr = &claudeStderr
	if err := claude.Run(); err != nil {
		t.Fatalf("python runtime package claude smoke: %v\nstderr=%s", err, claudeStderr.String())
	}
	if strings.TrimSpace(claudeStdout.String()) != "{}" {
		t.Fatalf("claude stdout = %q, want {}", claudeStdout.String())
	}
}

func TestNPMRuntimePackageContractFiles(t *testing.T) {
	t.Parallel()
	root := RepoRoot(t)

	for _, trackedPath := range []string{
		"npm/plugin-kit-ai-runtime/package.json",
		"npm/plugin-kit-ai-runtime/README.md",
		"npm/plugin-kit-ai-runtime/index.js",
		"npm/plugin-kit-ai-runtime/index.d.ts",
	} {
		tracked := exec.Command("git", "ls-files", "--error-unmatch", trackedPath)
		tracked.Dir = root
		if out, err := tracked.CombinedOutput(); err != nil {
			t.Fatalf("npm runtime package file must be tracked in git (%s): %v\n%s", trackedPath, err, out)
		}
	}

	packageJSON := readRepoFile(t, root, "npm", "plugin-kit-ai-runtime", "package.json")
	for _, want := range []string{
		`"name": "plugin-kit-ai-runtime"`,
		`"version": "0.0.0-development"`,
		`"node": ">=20"`,
		`"types": "./index.d.ts"`,
		`"default": "./index.js"`,
	} {
		mustContain(t, packageJSON, want)
	}

	indexJS := readRepoFile(t, root, "npm", "plugin-kit-ai-runtime", "index.js")
	for _, want := range []string{
		"export const CLAUDE_STABLE_HOOKS",
		"export const CLAUDE_EXTENDED_HOOKS",
		"export class ClaudeApp",
		"export class CodexApp",
	} {
		mustContain(t, indexJS, want)
	}

	indexDTS := readRepoFile(t, root, "npm", "plugin-kit-ai-runtime", "index.d.ts")
	for _, want := range []string{
		"export declare const CLAUDE_STABLE_HOOKS",
		"export declare class ClaudeApp",
		"export declare class CodexApp",
	} {
		mustContain(t, indexDTS, want)
	}

	readme := readRepoFile(t, root, "npm", "plugin-kit-ai-runtime", "README.md")
	for _, want := range []string{
		"npm i plugin-kit-ai-runtime",
		"shared dependency instead of copying a local helper file",
		"Go is still the recommended path",
		"stable supported lane",
	} {
		mustContain(t, readme, want)
	}

	workflow := readRepoFile(t, root, ".github", "workflows", "npm-runtime-publish.yml")
	for _, want := range []string{
		"name: NPM Runtime Publish",
		`workflows: ["Release Assets"]`,
		"plugin-kit-ai-runtime",
		"NPM_TOKEN",
		"npm publish --access public",
	} {
		mustContain(t, workflow, want)
	}
}

func TestNPMRuntimePackageClaudeAndCodexSmoke(t *testing.T) {
	t.Parallel()
	requireNodeRuntime(t)

	root := RepoRoot(t)
	pkgRoot := filepath.Join(root, "npm", "plugin-kit-ai-runtime")
	appDir := t.TempDir()
	nodeModules := filepath.Join(appDir, "node_modules", "plugin-kit-ai-runtime")
	copyTree(t, pkgRoot, nodeModules)

	if err := os.WriteFile(filepath.Join(appDir, "package.json"), []byte("{\"type\":\"module\"}\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	codexScript := filepath.Join(appDir, "codex_app.mjs")
	if err := os.WriteFile(codexScript, []byte(`import { CodexApp, continue_ } from "plugin-kit-ai-runtime";

const app = new CodexApp().onNotify((event) => {
  if (event.client !== "codex-tui") {
    throw new Error(`+"`unexpected client: ${JSON.stringify(event)}`"+`);
  }
  return continue_();
});

process.exit(app.run());
`), 0o644); err != nil {
		t.Fatal(err)
	}

	codex := exec.Command("node", codexScript, "notify", `{"client":"codex-tui"}`)
	codex.Dir = appDir
	var codexStdout bytes.Buffer
	var codexStderr bytes.Buffer
	codex.Stdout = &codexStdout
	codex.Stderr = &codexStderr
	if err := codex.Run(); err != nil {
		t.Fatalf("npm runtime package codex smoke: %v\nstderr=%s", err, codexStderr.String())
	}
	if strings.TrimSpace(codexStdout.String()) != "" {
		t.Fatalf("codex stdout = %q, want empty", codexStdout.String())
	}

	claudeScript := filepath.Join(appDir, "claude_app.mjs")
	if err := os.WriteFile(claudeScript, []byte(`import { CLAUDE_STABLE_HOOKS, ClaudeApp, allow } from "plugin-kit-ai-runtime";

const app = new ClaudeApp({
  allowedHooks: [...CLAUDE_STABLE_HOOKS],
  usage: "claude_app.mjs <hook-name>"
}).onStop((event) => {
  if (event.hook_event_name !== "Stop") {
    throw new Error(`+"`unexpected hook payload: ${JSON.stringify(event)}`"+`);
  }
  return allow();
});

process.exit(app.run());
`), 0o644); err != nil {
		t.Fatal(err)
	}

	claude := exec.Command("node", claudeScript, "Stop")
	claude.Dir = appDir
	claude.Stdin = strings.NewReader(`{"hook_event_name":"Stop"}`)
	var claudeStdout bytes.Buffer
	var claudeStderr bytes.Buffer
	claude.Stdout = &claudeStdout
	claude.Stderr = &claudeStderr
	if err := claude.Run(); err != nil {
		t.Fatalf("npm runtime package claude smoke: %v\nstderr=%s", err, claudeStderr.String())
	}
	if strings.TrimSpace(claudeStdout.String()) != "{}" {
		t.Fatalf("claude stdout = %q, want {}", claudeStdout.String())
	}
}
