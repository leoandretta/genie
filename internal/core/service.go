package core

import (
	"fmt"

	"github.com/leoandretta/mkdocs-cli/internal/services"
)

var ErrUnknownAlgorithm = fmt.Errorf("algoritmo não encontrado")

type GenerateService struct {
	registry services.Registry
}

func SetupGenerateService(registry services.Registry) *GenerateService {
	return &GenerateService{registry: registry}
}

// Generate gera um valor válido para a opção escolhida.
func (s *GenerateService) Generate(algorithm string) (string, error) {
	gen, err := s.registry.Get(algorithm)
	if err != nil {
		return "", err
	}

	return gen.Generate()
}

func (s *GenerateService) AvailableAlgorithms() []string {
	return s.registry.Names()
}
