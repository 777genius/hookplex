package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/777genius/plugin-kit-ai/cli/internal/app"
	"github.com/777genius/plugin-kit-ai/cli/internal/exitx"
	"github.com/spf13/cobra"
)

type testRunner interface {
	Test(context.Context, app.PluginTestOptions) (app.PluginTestResult, error)
}

var (
	testPlatform     string
	testEvent        string
	testFixture      string
	testGoldenDir    string
	testFormat       string
	testUpdateGolden bool
	testAll          bool
)

var testCmd = newTestCmd(pluginService)

func newTestCmd(runner testRunner) *cobra.Command {
	testPlatform = ""
	testEvent = ""
	testFixture = ""
	testGoldenDir = ""
	testFormat = "text"
	testUpdateGolden = false
	testAll = false

	cmd := &cobra.Command{
		Use:           "test [path]",
		Short:         "Run stable fixture-driven smoke tests against the launcher entrypoint",
		SilenceUsage:  true,
		SilenceErrors: true,
		Long: `Run stable Claude or Codex runtime smoke tests from JSON fixtures.

The command loads a fixture, invokes the configured launcher entrypoint with the correct carrier
(stdin JSON for Claude stable hooks, argv JSON for Codex notify), and optionally compares or updates
golden stdout/stderr/exitcode files for CI-grade regression checks.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := "."
			if len(args) == 1 {
				root = args[0]
			}
			if testFormat != "text" && testFormat != "json" {
				return fmt.Errorf("unsupported format %q (use text or json)", testFormat)
			}
			result, err := runner.Test(cmd.Context(), app.PluginTestOptions{
				Root:         root,
				Platform:     testPlatform,
				Event:        testEvent,
				Fixture:      testFixture,
				GoldenDir:    testGoldenDir,
				UpdateGolden: testUpdateGolden,
				All:          testAll,
			})
			if err != nil {
				return err
			}
			if testFormat == "json" {
				body, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					return err
				}
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(body))
			} else {
				for _, line := range result.Lines {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), line)
				}
			}
			if result.Passed {
				return nil
			}
			return exitx.Wrap(errors.New("test failures"), 1)
		},
	}
	cmd.Flags().StringVar(&testPlatform, "platform", "", `target override ("claude" or "codex-runtime")`)
	cmd.Flags().StringVar(&testEvent, "event", "", "stable event to execute (for example Stop, PreToolUse, UserPromptSubmit, or Notify)")
	cmd.Flags().BoolVar(&testAll, "all", false, "run every stable event for the selected platform")
	cmd.Flags().StringVar(&testFixture, "fixture", "", "fixture JSON path for single-event runs (default: fixtures/<platform>/<event>.json)")
	cmd.Flags().StringVar(&testGoldenDir, "golden-dir", "", "golden output directory (default: goldens/<platform>)")
	cmd.Flags().BoolVar(&testUpdateGolden, "update-golden", false, "write current stdout/stderr/exitcode outputs into the golden files")
	cmd.Flags().StringVar(&testFormat, "format", "text", "output format: text or json")
	return cmd
}
