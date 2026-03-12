package generate

import "flag"

// Subcommand representa um sub-comando do comando generate (cpf, cnpj, etc.).
type Subcommand struct {
	Name string
	FS   *flag.FlagSet
	Run  func(args []string) (string, error)
}
