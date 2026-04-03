package pluginkitairepo_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestContractClarity_GeminiRuntimeDocsStayAligned(t *testing.T) {
	root := RepoRoot(t)
	pluginKitAIBin := buildPluginKitAI(t)

	matrixBody, err := os.ReadFile(filepath.Join(root, "docs", "generated", "support_matrix.md"))
	if err != nil {
		t.Fatal(err)
	}
	matrix := string(matrixBody)
	for _, row := range []string{
		"| gemini | SessionStart | runtime_supported | beta | runtime-supported but not stable | false |",
		"| gemini | SessionEnd | runtime_supported | beta | runtime-supported but not stable | false |",
		"| gemini | BeforeModel | runtime_supported | beta | runtime-supported but not stable | false |",
		"| gemini | AfterModel | runtime_supported | beta | runtime-supported but not stable | false |",
		"| gemini | BeforeToolSelection | runtime_supported | beta | runtime-supported but not stable | false |",
		"| gemini | BeforeAgent | runtime_supported | beta | runtime-supported but not stable | false |",
		"| gemini | AfterAgent | runtime_supported | beta | runtime-supported but not stable | false |",
		"| gemini | BeforeTool | runtime_supported | beta | runtime-supported but not stable | false |",
		"| gemini | AfterTool | runtime_supported | beta | runtime-supported but not stable | false |",
	} {
		mustContain(t, matrix, row)
	}

	targetMatrixBody, err := os.ReadFile(filepath.Join(root, "docs", "generated", "target_support_matrix.md"))
	if err != nil {
		t.Fatal(err)
	}
	targetMatrix := string(targetMatrixBody)
	mustContain(t, targetMatrix, "| gemini | extension_package | mcp_extension | optional | extension | copy install | link | restart required | ~/.gemini/extensions/<name> | runtime-supported beta extension target |")

	cmd := exec.Command(pluginKitAIBin, "capabilities", "--mode", "runtime", "--format", "json", "--platform", "gemini")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("capabilities json: %v\n%s", err, out)
	}
	var entries []map[string]any
	if err := json.Unmarshal(out, &entries); err != nil {
		t.Fatalf("parse capabilities json: %v\n%s", err, out)
	}
	byKey := map[string]map[string]any{}
	for _, entry := range entries {
		key := entry["platform"].(string) + "/" + entry["event"].(string)
		byKey[key] = entry
	}
	for _, key := range []string{
		"gemini/SessionStart",
		"gemini/SessionEnd",
		"gemini/BeforeModel",
		"gemini/AfterModel",
		"gemini/BeforeToolSelection",
		"gemini/BeforeAgent",
		"gemini/AfterAgent",
		"gemini/BeforeTool",
		"gemini/AfterTool",
	} {
		assertCapabilityContract(t, byKey, key, "beta", "runtime-supported but not stable")
	}

	productionDoc, err := os.ReadFile(filepath.Join(root, "docs", "PRODUCTION.md"))
	if err != nil {
		t.Fatal(err)
	}
	supportDoc, err := os.ReadFile(filepath.Join(root, "docs", "SUPPORT.md"))
	if err != nil {
		t.Fatal(err)
	}
	sdkReadme, err := os.ReadFile(filepath.Join(root, "sdk", "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	sdkStability, err := os.ReadFile(filepath.Join(root, "sdk", "STABILITY.md"))
	if err != nil {
		t.Fatal(err)
	}
	repoTestsReadme, err := os.ReadFile(filepath.Join(root, "repotests", "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	geminiStarterReadme, err := os.ReadFile(filepath.Join(root, "cli", "plugin-kit-ai", "internal", "scaffold", "templates", "gemini.README.go.md.tmpl"))
	if err != nil {
		t.Fatal(err)
	}

	mustContain(t, string(productionDoc), "`public-beta` Go runtime lane for `SessionStart`, `SessionEnd`, `BeforeModel`, `AfterModel`, `BeforeToolSelection`, `BeforeAgent`, `AfterAgent`, `BeforeTool`, and `AfterTool`")
	mustContain(t, string(productionDoc), "make test-gemini-runtime")
	mustContain(t, string(productionDoc), "make test-gemini-runtime-live")
	mustContain(t, string(productionDoc), "audited 9-hook public-beta Gemini Go runtime lane")

	mustContain(t, string(supportDoc), "`github.com/777genius/plugin-kit-ai/sdk/gemini`")
	mustContain(t, string(supportDoc), "`(*plugin-kit-ai.App).Gemini`")
	mustContain(t, string(supportDoc), "`public-beta` Go runtime lane for `SessionStart`, `SessionEnd`, `BeforeModel`, `AfterModel`, `BeforeToolSelection`, `BeforeAgent`, `AfterAgent`, `BeforeTool`, and `AfterTool`")
	mustContain(t, string(supportDoc), "[GEMINI_RUNTIME_AUDIT.md](./GEMINI_RUNTIME_AUDIT.md)")
	mustContain(t, string(supportDoc), "exported Gemini runtime event/response/helper surfaces for:")
	mustNotContain(t, string(supportDoc), "- Gemini:\n  - `SessionStart`")

	mustContain(t, string(sdkReadme), "`gemini/SessionStart`")
	mustContain(t, string(sdkReadme), "`gemini/BeforeToolSelection`")
	mustContain(t, string(sdkReadme), "`gemini/BeforeTool`")
	mustContain(t, string(sdkReadme), "`gemini/AfterTool`")
	mustContain(t, string(sdkReadme), "`gemini.BeforeToolSelectionForceAny(...)`")
	mustContain(t, string(sdkReadme), "`gemini.AfterToolTailCallValue(...)`")
	mustContain(t, string(sdkReadme), "runtime-supported beta lane")
	mustContain(t, string(sdkReadme), "[../../docs/GEMINI_RUNTIME_AUDIT.md](../../docs/GEMINI_RUNTIME_AUDIT.md)")
	mustNotContain(t, string(sdkReadme), "Gemini now has a promoted stable subset")

	mustContain(t, string(sdkStability), "`(*plugin-kit-ai.App).Gemini`")
	mustContain(t, string(sdkStability), "approved exported Gemini event and response types for:")
	mustContain(t, string(sdkStability), "approved exported Gemini helper constructors for the current 9-hook beta runtime surface")
	mustContain(t, string(sdkStability), "`BeforeToolSelectionForceAny`")
	mustContain(t, string(sdkStability), "`AfterToolTailCallValue`")
	mustNotContain(t, string(sdkStability), "approved exported Gemini helper constructors for the stable subset")

	mustContain(t, string(repoTestsReadme), "`PLUGIN_KIT_AI_RUN_GEMINI_RUNTIME_LIVE=1`")
	mustContain(t, string(repoTestsReadme), "`PLUGIN_KIT_AI_E2E_GEMINI`")
	mustContain(t, string(repoTestsReadme), "make test-gemini-runtime")
	mustContain(t, string(repoTestsReadme), "make test-gemini-runtime-live")
	mustContain(t, string(repoTestsReadme), "current 9-hook public-beta runtime")
	mustContain(t, string(repoTestsReadme), "blocked-tool control semantics")
	mustContain(t, string(repoTestsReadme), "blocked-model control semantics")
	mustContain(t, string(repoTestsReadme), "`mode:\"NONE\"` semantics")
	mustNotContain(t, string(repoTestsReadme), "production-ready runtime")

	mustContain(t, string(geminiStarterReadme), "Runtime claim: `runtime-supported beta extension target`")
	mustContain(t, string(geminiStarterReadme), "audited 9-hook public-beta Go runtime")
	mustContain(t, string(geminiStarterReadme), "make test-gemini-runtime")
	mustContain(t, string(geminiStarterReadme), "make test-gemini-runtime-live")
	mustContain(t, string(geminiStarterReadme), "`plugin-kit-ai capabilities --mode runtime --platform gemini`")
	mustNotContain(t, string(geminiStarterReadme), "production-ready for the supported Gemini Go runtime")
}
