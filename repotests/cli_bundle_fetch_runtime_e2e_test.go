package pluginkitairepo_test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestPluginKitAIBundleFetchURLPythonRequirementsFlow(t *testing.T) {
	if !pythonRuntimeAvailable() {
		t.Skip("python runtime not available for bundle fetch flow")
	}

	pluginKitAIBin := buildPluginKitAI(t)
	plugRoot := runtimeProjectRoot(t)
	run := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", "codex-runtime", "--runtime", "python", "-o", plugRoot, "--extras")
	if out, err := run.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai init: %v\n%s", err, out)
	}
	writeRuntimeFile(t, plugRoot, "requirements.txt", "requests==2.32.0\n")

	bootstrap := exec.Command(pluginKitAIBin, "bootstrap", plugRoot)
	if out, err := bootstrap.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai bootstrap before export: %v\n%s", err, out)
	}
	export := exec.Command(pluginKitAIBin, "export", plugRoot, "--platform", "codex-runtime")
	if out, err := export.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai export before fetch: %v\n%s", err, out)
	}

	bundlePath := filepath.Join(plugRoot, "genplug_codex-runtime_python_bundle.tar.gz")
	bundleBody, err := os.ReadFile(bundlePath)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(bundleBody)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/bundle.tar.gz":
			_, _ = w.Write(bundleBody)
		case "/bundle.tar.gz.sha256":
			_, _ = w.Write([]byte(hex.EncodeToString(sum[:]) + "\n"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	dest := filepath.Join(t.TempDir(), "installed")
	fetch := exec.Command(pluginKitAIBin, "bundle", "fetch", "--url", server.URL+"/bundle.tar.gz", "--dest", dest)
	fetch.Env = bundleFetchTestCAEnv(t, server)
	out, err := fetch.CombinedOutput()
	if err != nil {
		t.Fatalf("plugin-kit-ai bundle fetch python requirements: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "Checksum source: "+server.URL+"/bundle.tar.gz.sha256") {
		t.Fatalf("fetch output missing checksum source:\n%s", out)
	}

	doctor := exec.Command(pluginKitAIBin, "doctor", dest)
	out, err = doctor.CombinedOutput()
	if err == nil {
		t.Fatalf("expected doctor to require bootstrap after bundle fetch:\n%s", out)
	}
	if !strings.Contains(string(out), "Status: needs_bootstrap") {
		t.Fatalf("doctor output missing needs_bootstrap:\n%s", out)
	}

	bootstrap = exec.Command(pluginKitAIBin, "bootstrap", dest)
	if out, err := bootstrap.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai bootstrap after bundle fetch: %v\n%s", err, out)
	}
	validate := exec.Command(pluginKitAIBin, "validate", dest, "--platform", "codex-runtime", "--strict")
	if out, err := validate.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai validate after bundle fetch: %v\n%s", err, out)
	}
}

func TestPluginKitAIBundleFetchURLClaudeNodeTypeScriptFlow(t *testing.T) {
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("node not in PATH")
	}
	if _, err := exec.LookPath("npm"); err != nil {
		t.Skip("npm not in PATH")
	}

	pluginKitAIBin := buildPluginKitAI(t)
	plugRoot := runtimeProjectRoot(t)
	run := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", "claude", "--runtime", "node", "--typescript", "-o", plugRoot, "--extras")
	if out, err := run.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai init claude node typescript: %v\n%s", err, out)
	}

	bootstrap := exec.Command(pluginKitAIBin, "bootstrap", plugRoot)
	if out, err := bootstrap.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai bootstrap before claude export: %v\n%s", err, out)
	}
	export := exec.Command(pluginKitAIBin, "export", plugRoot, "--platform", "claude")
	if out, err := export.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai export before fetch: %v\n%s", err, out)
	}

	bundlePath := filepath.Join(plugRoot, "genplug_claude_node_bundle.tar.gz")
	bundleBody, err := os.ReadFile(bundlePath)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(bundleBody)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/bundle.tar.gz":
			_, _ = w.Write(bundleBody)
		case "/bundle.tar.gz.sha256":
			_, _ = w.Write([]byte(hex.EncodeToString(sum[:]) + "  bundle.tar.gz\n"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	dest := filepath.Join(t.TempDir(), "installed")
	fetch := exec.Command(pluginKitAIBin, "bundle", "fetch", "--url", server.URL+"/bundle.tar.gz", "--dest", dest)
	fetch.Env = bundleFetchTestCAEnv(t, server)
	out, err := fetch.CombinedOutput()
	if err != nil {
		t.Fatalf("plugin-kit-ai bundle fetch claude node typescript: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "platform=claude runtime=node") {
		t.Fatalf("fetch output missing claude runtime summary:\n%s", out)
	}

	doctor := exec.Command(pluginKitAIBin, "doctor", dest)
	out, err = doctor.CombinedOutput()
	if err == nil {
		t.Fatalf("expected doctor to require bootstrap after bundle fetch:\n%s", out)
	}
	if !strings.Contains(string(out), "Status: needs_bootstrap") {
		t.Fatalf("doctor output missing needs_bootstrap:\n%s", out)
	}

	bootstrap = exec.Command(pluginKitAIBin, "bootstrap", dest)
	if out, err := bootstrap.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai bootstrap after bundle fetch: %v\n%s", err, out)
	}
	validate := exec.Command(pluginKitAIBin, "validate", dest, "--platform", "claude", "--strict")
	if out, err := validate.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai validate after bundle fetch: %v\n%s", err, out)
	}
}

func TestPluginKitAIBundleFetchGitHubClaudeNodeTypeScriptFlow(t *testing.T) {
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("node not in PATH")
	}
	if _, err := exec.LookPath("npm"); err != nil {
		t.Skip("npm not in PATH")
	}

	pluginKitAIBin := buildPluginKitAI(t)
	plugRoot := runtimeProjectRoot(t)
	run := exec.Command(pluginKitAIBin, "init", "genplug", "--platform", "claude", "--runtime", "node", "--typescript", "-o", plugRoot, "--extras")
	if out, err := run.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai init claude node typescript: %v\n%s", err, out)
	}

	bootstrap := exec.Command(pluginKitAIBin, "bootstrap", plugRoot)
	if out, err := bootstrap.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai bootstrap before github fetch export: %v\n%s", err, out)
	}
	export := exec.Command(pluginKitAIBin, "export", plugRoot, "--platform", "claude")
	if out, err := export.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai export before github fetch: %v\n%s", err, out)
	}

	bundleName := "genplug_claude_node_bundle.tar.gz"
	bundlePath := filepath.Join(plugRoot, bundleName)
	bundleBody, err := os.ReadFile(bundlePath)
	if err != nil {
		t.Fatal(err)
	}
	server := newMockBundleFetchGitHubServer(t, bundleName, bundleBody)
	defer server.Close()

	dest := filepath.Join(t.TempDir(), "installed")
	fetch := exec.Command(
		pluginKitAIBin,
		"bundle", "fetch", "o/r",
		"--tag", "v1",
		"--dest", dest,
		"--platform", "claude",
		"--runtime", "node",
		"--github-api-base", server.URL,
	)
	out, err := fetch.CombinedOutput()
	if err != nil {
		t.Fatalf("plugin-kit-ai bundle fetch github claude node typescript: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "Bundle source: github release o/r@v1 (tag) asset="+bundleName) {
		t.Fatalf("fetch output missing github bundle source:\n%s", out)
	}
	if !strings.Contains(string(out), "Checksum source: release asset checksums.txt") {
		t.Fatalf("fetch output missing checksums source:\n%s", out)
	}

	doctor := exec.Command(pluginKitAIBin, "doctor", dest)
	out, err = doctor.CombinedOutput()
	if err == nil {
		t.Fatalf("expected doctor to require bootstrap after github bundle fetch:\n%s", out)
	}
	if !strings.Contains(string(out), "Status: needs_bootstrap") {
		t.Fatalf("doctor output missing needs_bootstrap:\n%s", out)
	}

	bootstrap = exec.Command(pluginKitAIBin, "bootstrap", dest)
	if out, err := bootstrap.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai bootstrap after github bundle fetch: %v\n%s", err, out)
	}
	validate := exec.Command(pluginKitAIBin, "validate", dest, "--platform", "claude", "--strict")
	if out, err := validate.CombinedOutput(); err != nil {
		t.Fatalf("plugin-kit-ai validate after github bundle fetch: %v\n%s", err, out)
	}
}

func bundleFetchTestCAEnv(t *testing.T, server *httptest.Server) []string {
	t.Helper()
	certPath := filepath.Join(t.TempDir(), "bundle-fetch-test-ca.pem")
	certDER := server.TLS.Certificates[0].Certificate[0]
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	if err := os.WriteFile(certPath, certPEM, 0o644); err != nil {
		t.Fatal(err)
	}
	return append(os.Environ(), "PLUGIN_KIT_AI_TEST_CA_FILE="+certPath)
}

func newMockBundleFetchGitHubServer(t *testing.T, bundleName string, bundleBody []byte) *httptest.Server {
	t.Helper()
	sum := sha256.Sum256(bundleBody)
	checksums := hex.EncodeToString(sum[:]) + "  " + bundleName + "\n"
	type ghAsset struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
		Size               int64  `json:"size"`
	}
	release := struct {
		TagName string    `json:"tag_name"`
		Assets  []ghAsset `json:"assets"`
	}{}

	var srv *httptest.Server
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		base := srv.URL
		switch r.URL.Path {
		case "/repos/o/r/releases/tags/v1":
			release.TagName = "v1"
			release.Assets = []ghAsset{
				{Name: "checksums.txt", BrowserDownloadURL: base + "/checksums.txt", Size: int64(len(checksums))},
				{Name: bundleName, BrowserDownloadURL: base + "/bundle.tar.gz", Size: int64(len(bundleBody))},
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(release)
		case "/checksums.txt":
			_, _ = w.Write([]byte(checksums))
		case "/bundle.tar.gz":
			_, _ = w.Write(bundleBody)
		default:
			http.NotFound(w, r)
		}
	}))
	return srv
}
