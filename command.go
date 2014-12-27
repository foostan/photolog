package main

import (
	"github.com/codegangsta/cli"

	"github.com/foostan/photolog/command"
)

var Commands = []cli.Command{
	cli.Command{
		Name:        "relocate",
		Usage:       "Relocate photos",
		Description: "Relocate photos by datetime and rename to the hash of photo information",
		Flags:       command.RelocateFlags,
		Action:      command.RelocateCommand,
	},
}
