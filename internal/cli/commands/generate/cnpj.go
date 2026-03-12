package generate

import (
	"flag"

	"github.com/leoandretta/mkdocs-cli/internal/core"
)

func NewCNPJ(s *core.GenerateService) *Subcommand {
	fs := flag.NewFlagSet("generate cnpj", flag.ContinueOnError)
	return &Subcommand{
		Name: "cnpj",
		FS:   fs,
		Run: func(args []string) (string, error) {
			if err := fs.Parse(args); err != nil {
				return "", err
			}
			return s.Generate("cnpj")
		},
	}
}
