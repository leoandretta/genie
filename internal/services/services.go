package services

import (
	"fmt"
)

type CLIOptions struct {
	Formatted *bool
	Copy      *bool
	Input     *string
}
type CLIAction interface {
	Name() string
	Run(options CLIOptions) (string, error)
}

type Registry map[string]CLIAction

func SetupRegistry() Registry {
	r := Registry{}
	r.register(
		&CPFGenerator{},
		&CNPJGenerator{},
		&Base64Encoder{},
		&Base64Decoder{},
	)
	return r
}

func (r Registry) register(gens ...CLIAction) {
	for _, g := range gens {
		r[g.Name()] = g
	}
}

func (r Registry) Get(name string) (CLIAction, error) {
	g, ok := r[name]
	if !ok {
		return nil, fmt.Errorf("unknown option: %q", name)
	}
	return g, nil
}

func (r Registry) Names() []string {
	names := make([]string, 0, len(r))
	for k := range r {
		names = append(names, k)
	}
	return names
}
