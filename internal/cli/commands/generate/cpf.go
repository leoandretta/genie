package generate

import (
	"flag"

	"github.com/leoandretta/mkdocs-cli/internal/core"
	"github.com/leoandretta/mkdocs-cli/internal/services"
)

func NewCPF(s *core.GenerateService) *Subcommand {
	fs := flag.NewFlagSet("generate cpf", flag.ContinueOnError)
	formatted := fs.Bool("f", false, "Formatar o CPF")
	options := services.GenerateOptions{
		Formatted: formatted,
	}
	return &Subcommand{
		Name: "cpf",
		FS:   fs,
		Run: func(args []string) (string, error) {
			if err := fs.Parse(args); err != nil {
				return "", err
			}
			return s.Generate("cpf", options)
		},
	}
}
