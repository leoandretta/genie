package generate

import (
	"flag"

	"github.com/leoandretta/genie/internal/core"
	"github.com/leoandretta/genie/internal/services"
)

func NewCNPJ(s *core.GenerateService) *Subcommand {
	fs := flag.NewFlagSet("generate cnpj", flag.ContinueOnError)
	formatted := fs.Bool("f", false, "Returns the CNPJ in standard format")
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
