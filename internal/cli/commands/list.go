package commands

import (
	"flag"
	"fmt"
	"sort"

	"github.com/leoandretta/mkdocs-cli/internal/services"
)

type ListCommand struct {
	registry services.Registry
	fs       *flag.FlagSet
}

func NewListCommand(r services.Registry) *ListCommand {
	return &ListCommand{
		registry: r,
		fs:       flag.NewFlagSet("list", flag.ContinueOnError),
	}
}

func (c *ListCommand) Name() string           { return "list" }
func (c *ListCommand) Description() string    { return "Lista as opções de gerador disponíveis" }
func (c *ListCommand) FlagSet() *flag.FlagSet { return c.fs }

func (c *ListCommand) Run(args []string) error {
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	names := c.registry.Names()
	sort.Strings(names)
	fmt.Println("Opções disponíveis:")
	for _, name := range names {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
