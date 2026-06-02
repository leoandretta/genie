package decode

import (
	"flag"
	"fmt"

	"github.com/leoandretta/genie/internal/core"
	"github.com/leoandretta/genie/internal/services"
)

type DecodeCommand struct {
	service *core.CLIService
	fs      *flag.FlagSet
}

func NewDecodeCommand(s *core.CLIService) *DecodeCommand {
	fs := flag.NewFlagSet("decode", flag.ContinueOnError)
	return &DecodeCommand{service: s, fs: fs}
}

func (c *DecodeCommand) Name() string { return "decode" }
func (c *DecodeCommand) Description() string {
	return "Decodes the input string using the specified decoding"
}
func (c *DecodeCommand) FlagSet() *flag.FlagSet { return c.fs }

func (c *DecodeCommand) Run(args []string) error {
	copyToClipboard := c.fs.Bool("c", false, "Copy response to clipboard")
	_ = c.fs.Bool("base64", true, "Use Base64 decoding (default)")

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

	result, err := c.service.Run("decoding_base64", options)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func (c *DecodeCommand) printHelp() {
	fmt.Println("Usage: genie decode [flags] <input>")
	fmt.Println()
	fmt.Println("Flags:")
	c.fs.PrintDefaults()
	fmt.Println()
	fmt.Println("Example: genie decode \"hello world\"")
}
