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
	mustContain(t, matrix, "| gemini | SessionStart | runtime_supported | beta | runtime-supported but not stable | false |")
	mustContain(t, matrix, "| gemini | Notification | runtime_supported | beta | runtime-supported but not stable | false |")
	mustContain(t, matrix, "| gemini | BeforeToolSelection | runtime_supported | beta | runtime-supported but not stable | false |")
	mustContain(t, matrix, "| gemini | BeforeTool | runtime_supported | beta | runtime-supported but not stable | false |")
	mustContain(t, matrix, "| gemini | AfterTool | runtime_supported | beta | runtime-supported but not stable | false |")

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
	assertCapabilityContract(t, byKey, "gemini/SessionStart", "beta", "runtime-supported but not stable")
	assertCapabilityContract(t, byKey, "gemini/Notification", "beta", "runtime-supported but not stable")
	assertCapabilityContract(t, byKey, "gemini/BeforeToolSelection", "beta", "runtime-supported but not stable")
	assertCapabilityContract(t, byKey, "gemini/BeforeTool", "beta", "runtime-supported but not stable")
	assertCapabilityContract(t, byKey, "gemini/AfterTool", "beta", "runtime-supported but not stable")

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

	mustContain(t, string(productionDoc), "public-beta` Go runtime lane for `SessionStart`, `SessionEnd`, `Notification`, `PreCompress`, `BeforeModel`, `AfterModel`, `BeforeToolSelection`, `BeforeAgent`, `AfterAgent`, `BeforeTool`, and `AfterTool`")
	mustContain(t, string(productionDoc), "make test-gemini-runtime-smoke")
	mustContain(t, string(productionDoc), "make test-gemini-runtime-live")
	mustContain(t, string(productionDoc), "still not production-ready")

	mustContain(t, string(supportDoc), "`github.com/777genius/plugin-kit-ai/sdk/gemini`")
	mustContain(t, string(supportDoc), "`(*plugin-kit-ai.App).Gemini`")
	mustContain(t, string(supportDoc), "public-beta` Go runtime lane for `SessionStart`, `SessionEnd`, `Notification`, `PreCompress`, `BeforeModel`, `AfterModel`, `BeforeToolSelection`, `BeforeAgent`, `AfterAgent`, `BeforeTool`, and `AfterTool`")

	mustContain(t, string(sdkReadme), "`gemini/SessionStart` (`public-beta`)")
	mustContain(t, string(sdkReadme), "`gemini/Notification` (`public-beta`)")
	mustContain(t, string(sdkReadme), "`gemini/BeforeToolSelection` (`public-beta`)")
	mustContain(t, string(sdkReadme), "`gemini/BeforeTool` (`public-beta`)")
	mustContain(t, string(sdkReadme), "`gemini/AfterTool` (`public-beta`)")
	mustContain(t, string(sdkReadme), "`gemini.BeforeToolSelectionForceAny(...)`")
	mustContain(t, string(sdkReadme), "`gemini.AfterToolTailCallValue(...)`")

	mustContain(t, string(sdkStability), "`(*plugin-kit-ai.App).Gemini`")
	mustContain(t, string(sdkStability), "`NotificationMessage`")
	mustContain(t, string(sdkStability), "`BeforeToolSelectionForceAny`")
	mustContain(t, string(sdkStability), "`AfterToolTailCallValue`")

	mustContain(t, string(repoTestsReadme), "`PLUGIN_KIT_AI_RUN_GEMINI_RUNTIME_LIVE=1`")
	mustContain(t, string(repoTestsReadme), "`PLUGIN_KIT_AI_E2E_GEMINI`")
	mustContain(t, string(repoTestsReadme), "make test-gemini-runtime-smoke")
	mustContain(t, string(repoTestsReadme), "Deterministic repo-local Gemini smoke")
	mustContain(t, string(repoTestsReadme), "Advisory `Notification` и `PreCompress`")

	mustContain(t, string(geminiStarterReadme), "managed Gemini hook wiring for `SessionStart`, `SessionEnd`, `Notification`, `PreCompress`, `BeforeModel`, `AfterModel`, `BeforeToolSelection`, `BeforeAgent`, `AfterAgent`, `BeforeTool`, and `AfterTool`")
	mustContain(t, string(geminiStarterReadme), "make test-gemini-runtime-smoke")
	mustContain(t, string(geminiStarterReadme), "make test-gemini-runtime-live")
	mustContain(t, string(geminiStarterReadme), "`plugin-kit-ai capabilities --mode runtime --platform gemini`")
}
