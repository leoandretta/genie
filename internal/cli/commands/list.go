package commands

import (
	"flag"
	"fmt"
	"sort"

	"github.com/leoandretta/mkdocs-cli/internal/cli"
)

type ListCommand struct {
	registry cli.CommandRegistry
	fs       *flag.FlagSet
}

func NewListCommand(r cli.CommandRegistry) *ListCommand {
	return &ListCommand{
		registry: r,
		fs:       flag.NewFlagSet("list", flag.ContinueOnError),
	}
}

func (c *ListCommand) Name() string           { return "list" }
func (c *ListCommand) Description() string    { return "Lista os comandos disponíveis" }
func (c *ListCommand) FlagSet() *flag.FlagSet { return c.fs }

func (c *ListCommand) Run(args []string) error {
	if err := c.fs.Parse(args); err != nil {
		return err
	}
	names := make([]string, 0, len(c.registry))
	for k := range c.registry {
		names = append(names, k)
	}
	sort.Strings(names)
	fmt.Println("Comandos disponíveis:")
	for _, name := range names {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
