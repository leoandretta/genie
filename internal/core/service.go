package core

import (
	"github.com/leoandretta/genie/internal/services"
	"golang.design/x/clipboard"
)

type CLIService struct {
	registry services.Registry
}

func SetupCLIService(registry services.Registry) *CLIService {
	return &CLIService{registry: registry}
}

func (s *CLIService) Run(algorithm string, options services.CLIOptions) (string, error) {
	gen, err := s.registry.Get(algorithm)
	if err != nil {
		return "", err
	}

	response, genErr := gen.Run(options)
	if genErr != nil {
		return "", genErr
	}
	if *options.Copy == true {
		clipboard.Write(clipboard.FmtText, []byte(response))
		response += " - 🗐 Copied to clipboard!"
	}
	return response, nil
}

func (s *CLIService) AvailableAlgorithms() []string {
	return s.registry.Names()
}
