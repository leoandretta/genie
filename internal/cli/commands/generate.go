package commands

import (
	"flag"
	"fmt"
	"sort"

	"github.com/leoandretta/mkdocs-cli/internal/cli/commands/generate"
	"github.com/leoandretta/mkdocs-cli/internal/core"
)

// GenerateCommand é o comando `mkdocs generate <subcomando>`.
type GenerateCommand struct {
	service     *core.GenerateService
	fs          *flag.FlagSet
	subcommands map[string]*generate.Subcommand
}

func NewGenerateCommand(s *core.GenerateService) *GenerateCommand {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)

	subList := []*generate.Subcommand{
		generate.NewCPF(s),
		generate.NewCNPJ(s),
	}

	subs := make(map[string]*generate.Subcommand, len(subList))
	for _, sub := range subList {
		subs[sub.Name] = sub
	}

	return &GenerateCommand{service: s, fs: fs, subcommands: subs}
}

func (c *GenerateCommand) Name() string { return "generate" }
func (c *GenerateCommand) Description() string {
	return "Gera um valor válido para a opção selecionada"
}
func (c *GenerateCommand) FlagSet() *flag.FlagSet { return c.fs }

func (c *GenerateCommand) Run(args []string) error {
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		c.printHelp()
		return nil
	}

	subName := args[0]

	// `mkdocs generate list` lista subcomandos disponíveis
	if subName == "list" {
		c.listSubcommands()
		return nil
	}

	sub, ok := c.subcommands[subName]
	if !ok {
		fmt.Printf("erro: subcomando desconhecido %q\n\n", subName)
		c.printHelp()
		return fmt.Errorf("subcomando desconhecido: %q", subName)
	}

	result, err := sub.Run(args[1:])
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func (c *GenerateCommand) printHelp() {
	fmt.Println("Uso: mkdocs generate <subcomando> [flags]")
	fmt.Println()
	c.listSubcommands()
	fmt.Println()
	fmt.Println("Use mkdocs generate <subcomando> --help para detalhes.")
}

func (c *GenerateCommand) listSubcommands() {
	names := make([]string, 0, len(c.subcommands))
	for k := range c.subcommands {
		names = append(names, k)
	}
	sort.Strings(names)
	fmt.Println("Subcomandos disponíveis:")
	for _, name := range names {
		fmt.Printf("  - %s\n", name)
	}
}
