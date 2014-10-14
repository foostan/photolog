package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "photolog"
	app.Version = Version
	app.Usage = "Cli tools for creating/showing photo catalog"
	app.Author = "foostan"
	app.Email = "ks@fstn.jp"
	app.Commands = Commands

	app.Run(os.Args)
}
