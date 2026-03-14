package core

import (
	"github.com/leoandretta/genie/internal/services"
	"golang.design/x/clipboard"
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

	response, genErr := gen.Generate(options)
	if genErr != nil {
		return "", genErr
	}
	if *options.Copy == true {
		clipboard.Write(clipboard.FmtText, []byte(response))
		response += " - 🗐 Copied to clipboard!"
	}
	return response, nil
}

func (s *GenerateService) AvailableAlgorithms() []string {
	return s.registry.Names()
}
