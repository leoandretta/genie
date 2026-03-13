package core

import (
	"github.com/leoandretta/mkdocs-cli/internal/services"
)

type GenerateService struct {
	registry services.Registry
}

func SetupGenerateService(registry services.Registry) *GenerateService {
	return &GenerateService{registry: registry}
}

// Generate gera um valor válido para a opção escolhida.
func (s *GenerateService) Generate(algorithm string, options services.GenerateOptions) (string, error) {
	gen, err := s.registry.Get(algorithm)
	if err != nil {
		return "", err
	}

	return gen.Generate(options)
}

func (s *GenerateService) AvailableAlgorithms() []string {
	return s.registry.Names()
}
