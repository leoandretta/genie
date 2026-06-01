package generate

import (
	"flag"

	"github.com/leoandretta/genie/internal/cli/commands"
	"github.com/leoandretta/genie/internal/core"
	"github.com/leoandretta/genie/internal/services"
)

func NewCNPJ(s *core.CLIService) *commands.SubCommand {
	fs := flag.NewFlagSet("generate cnpj", flag.ContinueOnError)
	formatted := fs.Bool("f", false, "Returns the CNPJ in standard format")
	copyToClipboard := fs.Bool("c", false, "Copy response to clipboard")

	options := services.CLIOptions{
		Formatted: formatted,
		Copy:      copyToClipboard,
	}
	return &commands.SubCommand{
		Name: "cnpj",
		FS:   fs,
		Run: func(args []string) (string, error) {
			if err := fs.Parse(args); err != nil {
				return "", err
			}
			return s.Run("cnpj", options)
		},
	}
}
