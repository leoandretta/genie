package cli

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
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

	bundledFlags := preprocessBundledFlags(args)

	cmd, ok := r.registry[bundledFlags[0]]
	if !ok {
		fmt.Fprintf(os.Stderr, "error: unknown command %q\n\n", bundledFlags[0])
		r.printHelp()
		return 1
	}

	if err := cmd.Run(bundledFlags[1:]); err != nil {
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

func preprocessBundledFlags(args []string) []string {
	var result []string
	for _, arg := range args {
		if (strings.HasPrefix(arg, "-") && !strings.Contains(arg, "=")) &&
			(len(arg) > 2 && (strings.HasPrefix(arg, "--") || !strings.Contains(arg[1:], "--"))) {

			if strings.HasPrefix(arg, "--") {
				arg = arg[2:]
			} else {
				arg = arg[1:]
			}
			for _, ch := range arg {
				result = append(result, "-"+string(ch))
			}
		} else {
			result = append(result, arg)
		}
	}
	return result
}
