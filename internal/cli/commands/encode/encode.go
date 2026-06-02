package encode

import (
	"flag"
	"fmt"

	"github.com/leoandretta/genie/internal/core"
	"github.com/leoandretta/genie/internal/services"
)

type EncodeCommand struct {
	service *core.CLIService
	fs      *flag.FlagSet
}

func NewEncodeCommand(s *core.CLIService) *EncodeCommand {
	fs := flag.NewFlagSet("encode", flag.ContinueOnError)
	return &EncodeCommand{service: s, fs: fs}
}

func (c *EncodeCommand) Name() string { return "encode" }
func (c *EncodeCommand) Description() string {
	return "Encodes the input string using the specified encoding"
}
func (c *EncodeCommand) FlagSet() *flag.FlagSet { return c.fs }

func (c *EncodeCommand) Run(args []string) error {
	copyToClipboard := c.fs.Bool("c", false, "Copy response to clipboard")
	_ = c.fs.Bool("base64", true, "Use Base64 encoding (default)")

	if err := c.fs.Parse(args); err != nil {
		return err
	}

	remaining := c.fs.Args()
	if len(remaining) == 0 {
		c.printHelp()
		return fmt.Errorf("input string is required")
	}

	input := remaining[0]

	options := services.CLIOptions{
		Copy:  copyToClipboard,
		Input: &input,
	}

	result, err := c.service.Run("encoding_base64", options)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func (c *EncodeCommand) printHelp() {
	fmt.Println("Usage: genie encode [flags] <input>")
	fmt.Println()
	fmt.Println("Flags:")
	c.fs.PrintDefaults()
	fmt.Println()
	fmt.Println("Example: genie encode \"hello world\"")
}
