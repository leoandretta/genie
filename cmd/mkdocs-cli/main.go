// cmd/hashcli/main.go
package main

import (
	"os"

	"github.com/leoandretta/mkdocs-cli/internal/cli"
	"github.com/leoandretta/mkdocs-cli/internal/cli/commands"
	"github.com/leoandretta/mkdocs-cli/internal/core"
	"github.com/leoandretta/mkdocs-cli/internal/services"
)

func main() {
	generatorRegistry := services.SetupRegistry()
	service := core.SetupGenerateService(generatorRegistry)

	cmdRegistry := cli.NewCommandRegistry(
		commands.NewGenerateCommand(service),
		commands.NewListCommand(generatorRegistry),
	)

	runner := cli.NewRunner("mkdocs-cli", cmdRegistry)
	os.Exit(runner.Run())
}
