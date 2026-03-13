// cmd/hashcli/main.go
package main

import (
	"os"

	"github.com/leoandretta/genie/internal/cli"
	"github.com/leoandretta/genie/internal/cli/commands"
	"github.com/leoandretta/genie/internal/core"
	"github.com/leoandretta/genie/internal/services"
)

func main() {
	generatorRegistry := services.SetupRegistry()
	service := core.SetupGenerateService(generatorRegistry)

	// Cria o registry com os comandos principais primeiro
	cmdRegistry := cli.NewCommandRegistry(
		commands.NewGenerateCommand(service),
	)

	runner := cli.NewRunner("genie", cmdRegistry)
	os.Exit(runner.Run())
}
