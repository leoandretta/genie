package services

import (
	"fmt"
	"math/rand/v2"
)

type CPFGenerator struct{}

func (g *CPFGenerator) Name() string { return "cpf" }

func (g *CPFGenerator) Generate(options GenerateOptions) (string, error) {
	d := make([]int, 9)
	for i := range d {
		d[i] = rand.IntN(10)
	}

	// Calcula o 1º dígito verificador
	sum := 0
	for i, v := range d {
		sum += v * (10 - i)
	}
	resto := sum % 11
	v1 := 0
	if resto >= 2 {
		v1 = 11 - resto
	}

	// Calcula o 2º dígito verificador
	sum = 0
	for i, v := range d {
		sum += v * (11 - i)
	}
	sum += v1 * 2
	resto = sum % 11
	v2 := 0
	if resto >= 2 {
		v2 = 11 - resto
	}

	if *options.Formatted {
		return fmt.Sprintf("%d%d%d.%d%d%d.%d%d%d-%d%d",
			d[0], d[1], d[2],
			d[3], d[4], d[5],
			d[6], d[7], d[8],
			v1, v2), nil
	}
	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d",
		d[0], d[1], d[2],
		d[3], d[4], d[5],
		d[6], d[7], d[8],
		v1, v2), nil
}
