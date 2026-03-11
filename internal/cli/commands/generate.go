package commands

import (
	"flag"
	"fmt"

	"github.com/leoandretta/mkdocs-cli/internal/core"
)

type GenerateCommand struct {
	service *core.GenerateService
	fs      *flag.FlagSet
	tool    *string
}

func NewGenerateCommand(s *core.GenerateService) *GenerateCommand {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)
	return &GenerateCommand{
		service: s,
		fs:      fs,
		tool:    fs.String("tool", "", "Valor a ser gerado (para consultar as opções disponíveis, usar 'mkdocs list')"),
	}
}

func (c *GenerateCommand) Name() string { return "generate" }
func (c *GenerateCommand) Description() string {
	return "Gera um valor válido para a opção selecionada"
}
func (c *GenerateCommand) FlagSet() *flag.FlagSet { return c.fs }

func (c *GenerateCommand) Run(args []string) error {
	if err := c.fs.Parse(args); err != nil {
		return err
	}

	if *c.tool == "" {
		c.fs.Usage()
		return fmt.Errorf("flag obrigatória ausente: --tool")
	}

	result, err := c.service.Generate(*c.tool)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
