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

	// Cria o registry com os comandos principais primeiro
	cmdRegistry := cli.NewCommandRegistry(
		commands.NewGenerateCommand(service),
	)

	runner := cli.NewRunner("mkdocs", cmdRegistry)
	os.Exit(runner.Run())
}
