.PHONY: test test-required test-install-compat test-extended test-live test-live-cli test-install-live test-e2e-live generated-check release-gate release-rehearsal build-hookplex vet

GOCACHE ?= /tmp/hookplex-gocache
export GOCACHE

EXTENDED_TEST_ARGS ?=

test:
	$(MAKE) test-required

test-required:
	go test ./...

test-install-compat:
	go test -count=1 -run '^TestHookplexInstall_' ./repotests

test-extended:
	go test -count=1 -run '^TestClaudeCLIHooks$$' ./repotests $(EXTENDED_TEST_ARGS)
	go test -count=1 -run '^TestCodexCLINotify$$' ./repotests $(EXTENDED_TEST_ARGS)

# Live E2E: real GitHub + real claude-notifications-go release (needs network). Optional: GITHUB_TOKEN.
# Package is ./repotests (tests moved out of repo root).
test-live: test-e2e-live

test-live-cli:
	go test -count=1 -run 'TestClaudeHooks_LiveHaikuLow' ./repotests $(EXTENDED_TEST_ARGS)

test-install-live:
	HOOKPLEX_E2E_LIVE=1 go test -count=1 -timeout=15m -run '^TestLiveInstall_' ./repotests

test-e2e-live: test-install-live

# Root module is workspace-only; submodules are vetted explicitly.
vet:
	go vet ./...
	cd cli/hookplex && go vet ./...
	cd install/plugininstall && go vet ./...
	cd sdk/hookplex && go vet ./...

generated-check:
	bash ./scripts/check-generated-sync.sh

release-gate:
	$(MAKE) test-required
	$(MAKE) vet

release-rehearsal: release-gate
	$(MAKE) generated-check
	$(MAKE) test-install-compat
	@echo "Release rehearsal deterministic checks complete. Record install compatibility, extended/live evidence, audit updates, and release notes draft."

build-hookplex:
	go build -o bin/hookplex ./cli/hookplex/cmd/hookplex
