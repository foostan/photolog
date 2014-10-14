package command

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var RelocateFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "basepath",
		Value: ".",
		Usage: "base directory path of target files",
	},
}

func RelocateCommand(c *cli.Context) {
	basepath := c.String("basepath")

	fmt.Println(basepath)
}
