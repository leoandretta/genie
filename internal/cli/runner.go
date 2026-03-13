package cli

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

type Runner struct {
	registry CommandRegistry
	appName  string
}

func NewRunner(appName string, r CommandRegistry) *Runner {
	return &Runner{registry: r, appName: appName}
}

func (r *Runner) Run() int {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		r.printHelp()
		return 0
	}

	cmd, ok := r.registry[args[0]]
	if !ok {
		fmt.Fprintf(os.Stderr, "error: unknown command %q\n\n", args[0])
		r.printHelp()
		return 1
	}

	if err := cmd.Run(args[1:]); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		return 1
	}
	return 0
}

func (r *Runner) printHelp() {
	fmt.Printf("Usage: %s <command> [flags]\n\nAvailable commands:\n", r.appName)
	// ordena para output determinístico
	names := make([]string, 0, len(r.registry))
	for k := range r.registry {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("  %-14s %s\n", name, r.registry[name].Description())
	}
	fmt.Printf("\nTry: %s <command> --help for help with each available command.\n", r.appName)
}
