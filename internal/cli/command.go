package cli

import "flag"

type Command interface {
	Name() string
	Description() string
	FlagSet() *flag.FlagSet
	Run(args []string) error
}

type CommandRegistry map[string]Command

func NewCommandRegistry(cmds ...Command) CommandRegistry {
	r := CommandRegistry{}
	for _, c := range cmds {
		r[c.Name()] = c
	}
	return r
}
