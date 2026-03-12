package generate

import (
	"flag"

	"github.com/leoandretta/mkdocs-cli/internal/core"
	"github.com/leoandretta/mkdocs-cli/internal/services"
)

func NewCNPJ(s *core.GenerateService) *Subcommand {
	fs := flag.NewFlagSet("generate cnpj", flag.ContinueOnError)
	formatted := fs.Bool("f", false, "Formatar o CNPJ")
	options := services.GenerateOptions{
		Formatted: formatted,
	}
	return &Subcommand{
		Name: "cnpj",
		FS:   fs,
		Run: func(args []string) (string, error) {
			if err := fs.Parse(args); err != nil {
				return "", err
			}
			return s.Generate("cnpj", options)
		},
	}
}
