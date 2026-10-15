// Package cli builds the tally command tree: sum, count, version.
package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/paiml/go-capstone-cli/config"
	"github.com/paiml/go-capstone-cli/logx"
	"github.com/paiml/go-capstone-cli/version"
	"github.com/spf13/cobra"
)

// Env is what the commands read from the outside world.
type Env struct {
	In       io.Reader
	Out, Err io.Writer
	Getenv   func(string) string
}

// OS returns an Env wired to the real process.
func OS() Env { return Env{In: os.Stdin, Out: os.Stdout, Err: os.Stderr, Getenv: os.Getenv} }

// New builds the root command.
func New(e Env) *cobra.Command {
	var verbose bool
	var unit string
	root := &cobra.Command{Use: "tally", Short: "Add numbers and count lines", SilenceUsage: true, SilenceErrors: true}
	root.SetOut(e.Out)
	root.SetErr(e.Err)
	root.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "log debug lines to stderr")
	root.PersistentFlags().StringVar(&unit, "unit", "", "suffix for results (flag > TALLY_UNIT > none)")
	suffix := func() string {
		u, _ := config.Resolve("unit", unit, e.Getenv("TALLY_UNIT"), nil, "")
		if u == "" {
			return ""
		}
		return " " + u
	}
	root.AddCommand(
		&cobra.Command{
			Use: "sum N...", Short: "Add whole numbers", Args: cobra.MinimumNArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				log := logx.New(e.Err, verbose)
				total := 0
				for _, a := range args {
					n, err := strconv.Atoi(a)
					if err != nil {
						return fmt.Errorf("sum: %q is not a whole number", a)
					}
					log.Debug("add", "n", n)
					total += n
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%d%s\n", total, suffix())
				return nil
			},
		},
		&cobra.Command{
			Use: "count", Short: "Count lines and words on stdin", Args: cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				sc := bufio.NewScanner(e.In)
				lines, words := 0, 0
				for sc.Scan() {
					lines++
					words += len(strings.Fields(sc.Text()))
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%d lines %d words\n", lines, words)
			},
		},
		&cobra.Command{
			Use: "version", Short: "Print the build version", Args: cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) { fmt.Fprintln(cmd.OutOrStdout(), version.String()) },
		},
	)
	return root
}

// Run executes args and returns the exit code: 0 ok, 1 error, 2 usage.
func Run(e Env, args []string) int {
	root := New(e)
	root.SetArgs(args)
	err := root.Execute()
	if err == nil {
		return 0
	}
	fmt.Fprintln(e.Err, "tally:", err)
	if strings.HasPrefix(err.Error(), "sum:") {
		return 1
	}
	return 2
}
