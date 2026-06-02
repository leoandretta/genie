package services

import (
	"encoding/base64"
	"fmt"
)

type Base64Decoder struct{}

func (e *Base64Decoder) Name() string { return "decoding_base64" }

func (e *Base64Decoder) Run(options CLIOptions) (string, error) {
	if options.Input == nil || *options.Input == "" {
		return "", fmt.Errorf("input string is required for decoding")
	}

	decoded, err := base64.StdEncoding.DecodeString(*options.Input)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 string: %w", err)
	}
	return string(decoded), nil
}
