package commands

import "flag"

type SubCommand struct {
	Name string
	FS   *flag.FlagSet
	Run  func(args []string) (string, error)
}
