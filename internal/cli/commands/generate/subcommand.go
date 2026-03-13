package generate

import "flag"

type Subcommand struct {
	Name string
	FS   *flag.FlagSet
	Run  func(args []string) (string, error)
}
