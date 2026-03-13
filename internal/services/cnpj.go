package services

import (
	"fmt"
	"math/rand/v2"
)

type CNPJGenerator struct{}

func (g *CNPJGenerator) Name() string { return "cnpj" }

func (g *CNPJGenerator) Generate(options GenerateOptions) (string, error) {
	d := make([]int, 12)
	for i := range d {
		d[i] = rand.IntN(10)
	}
	// Default 4 digits (d8-d11) for main office are usually 0001
	d[8], d[9], d[10], d[11] = 0, 0, 0, 1

	pesos1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	pesos2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// First verifier digit (d12)
	sum := 0
	for i, v := range d {
		sum += v * pesos1[i]
	}
	resto := sum % 11
	v1 := 0
	if resto >= 2 {
		v1 = 11 - resto
	}

	// Second verifier digit (d13)
	sum = 0
	for i, v := range d {
		sum += v * pesos2[i]
	}
	sum += v1 * pesos2[12]
	resto = sum % 11
	v2 := 0
	if resto >= 2 {
		v2 = 11 - resto
	}

	if *options.Formatted {
		return fmt.Sprintf("%d%d.%d%d%d.%d%d%d/%d%d%d%d-%d%d",
			d[0], d[1],
			d[2], d[3], d[4],
			d[5], d[6], d[7],
			d[8], d[9], d[10], d[11],
			v1, v2), nil
	}
	return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d",
		d[0], d[1],
		d[2], d[3], d[4],
		d[5], d[6], d[7],
		d[8], d[9], d[10], d[11],
		v1, v2), nil
}
