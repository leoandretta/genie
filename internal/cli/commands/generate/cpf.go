package generate

import (
	"flag"

	"github.com/leoandretta/mkdocs-cli/internal/core"
)

func NewCPF(s *core.GenerateService) *Subcommand {
	fs := flag.NewFlagSet("generate cpf", flag.ContinueOnError)
	return &Subcommand{
		Name: "cpf",
		FS:   fs,
		Run: func(args []string) (string, error) {
			if err := fs.Parse(args); err != nil {
				return "", err
			}
			return s.Generate("cpf")
		},
	}
}
