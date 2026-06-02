package services

import (
	"encoding/base64"
	"fmt"
)

type Base64Encoder struct{}

func (e *Base64Encoder) Name() string { return "encoding_base64" }

func (e *Base64Encoder) Run(options CLIOptions) (string, error) {
	if options.Input == nil || *options.Input == "" {
		return "", fmt.Errorf("input string is required for encoding")
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(*options.Input))
	return encoded, nil
}
