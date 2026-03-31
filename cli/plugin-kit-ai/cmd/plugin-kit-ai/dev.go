package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/777genius/plugin-kit-ai/cli/internal/app"
	"github.com/777genius/plugin-kit-ai/cli/internal/exitx"
	"github.com/spf13/cobra"
)

type devRunner interface {
	Dev(context.Context, app.PluginDevOptions, func(app.PluginDevUpdate)) (app.PluginDevSummary, error)
}

var (
	devPlatform  string
	devEvent     string
	devFixture   string
	devGoldenDir string
	devAll       bool
	devOnce      bool
	devInterval  time.Duration
)

var devCmd = newDevCmd(pluginService)

func newDevCmd(runner devRunner) *cobra.Command {
	devPlatform = ""
	devEvent = ""
	devFixture = ""
	devGoldenDir = ""
	devAll = false
	devOnce = false
	devInterval = 750 * time.Millisecond

	cmd := &cobra.Command{
		Use:           "dev [path]",
		Short:         "Watch the project, re-render, re-validate, rebuild when needed, and rerun fixtures",
		SilenceUsage:  true,
		SilenceErrors: true,
		Long: `Watch launcher-based runtime targets in a fast inner loop.

Each cycle re-renders the selected target, performs runtime-aware rebuilds when needed,
runs strict validation, and reruns the configured stable Claude or Codex fixture smoke tests.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			root := "."
			if len(args) == 1 {
				root = args[0]
			}
			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer stop()

			var lastPassed = true
			summary, err := runner.Dev(ctx, app.PluginDevOptions{
				Root:      root,
				Platform:  devPlatform,
				Event:     devEvent,
				Fixture:   devFixture,
				GoldenDir: devGoldenDir,
				All:       devAll,
				Once:      devOnce,
				Interval:  devInterval,
			}, func(update app.PluginDevUpdate) {
				lastPassed = update.Passed
				for _, line := range update.Lines {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), line)
				}
			})
			if err != nil {
				return err
			}
			if devOnce && !summary.LastPassed {
				return exitx.Wrap(errors.New("dev cycle failed"), 1)
			}
			if devOnce && !lastPassed {
				return exitx.Wrap(errors.New("dev cycle failed"), 1)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&devPlatform, "platform", "", `target override ("claude" or "codex-runtime")`)
	cmd.Flags().StringVar(&devEvent, "event", "", "stable event to execute (for example Stop, PreToolUse, UserPromptSubmit, or Notify)")
	cmd.Flags().BoolVar(&devAll, "all", false, "run every stable event for the selected platform on each cycle")
	cmd.Flags().StringVar(&devFixture, "fixture", "", "fixture JSON path for single-event runs (default: fixtures/<platform>/<event>.json)")
	cmd.Flags().StringVar(&devGoldenDir, "golden-dir", "", "golden output directory (default: goldens/<platform>)")
	cmd.Flags().BoolVar(&devOnce, "once", false, "run a single render/validate/test cycle and exit")
	cmd.Flags().DurationVar(&devInterval, "interval", 750*time.Millisecond, "poll interval for watch mode")
	return cmd
}
