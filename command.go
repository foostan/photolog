package main

import (
	"github.com/codegangsta/cli"

	"github.com/foostan/photolog/command"
)

var Commands = []cli.Command{
	cli.Command{
		Name:        "relocate",
		Usage:       "Relocate photos by taked date",
		Description: "Relocate photos by taked date",
		Flags:       command.RelocateFlags,
		Action:      command.RelocateCommand,
	},
}
