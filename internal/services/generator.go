package services

import (
	"fmt"
)

type GenerateOptions struct {
	Formatted *bool
}
type Generator interface {
	Name() string
	Generate(options GenerateOptions) (string, error)
}

type Registry map[string]Generator

func SetupRegistry() Registry {
	r := Registry{}
	r.register(
		&CPFGenerator{},
		&CNPJGenerator{},
	)
	return r
}

func (r Registry) register(gens ...Generator) {
	for _, g := range gens {
		r[g.Name()] = g
	}
}

func (r Registry) Get(name string) (Generator, error) {
	g, ok := r[name]
	if !ok {
		return nil, fmt.Errorf("algoritmo desconhecido: %q", name)
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
