package main

import (
	"github.com/codegangsta/cli"

	"github.com/foostan/photolog/command"
)

var Commands = []cli.Command{
	cli.Command{
		Name:        "web",
		Usage:       "Web Interface",
		Description: "Serve web API and user interface",
		Flags:       command.WebFlags,
		Action:      command.WebCommand,
	},
	cli.Command{
		Name:        "relocate",
		Usage:       "Relocate photos",
		Description: "Relocate photos by datetime and rename to the hash of photo information",
		Flags:       command.RelocateFlags,
		Action:      command.RelocateCommand,
	},
}
