package core

import (
	"github.com/leoandretta/genie/internal/services"
)

type GenerateService struct {
	registry services.Registry
}

func SetupGenerateService(registry services.Registry) *GenerateService {
	return &GenerateService{registry: registry}
}

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
