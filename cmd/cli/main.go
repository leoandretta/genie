// cmd/hashcli/main.go
package main

import (
	"os"

	"github.com/leoandretta/genie/internal/cli"
	"github.com/leoandretta/genie/internal/cli/commands/generate"
	"github.com/leoandretta/genie/internal/core"
	"github.com/leoandretta/genie/internal/services"
)

func main() {
	actionRegistry := services.SetupRegistry()
	cliService := core.SetupCLIService(actionRegistry)

	cmdRegistry := cli.NewCommandRegistry(
		generate.NewGenerateCommand(cliService),
	)

	runner := cli.NewRunner("genie", cmdRegistry)
	os.Exit(runner.Run())
}
