package command

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	. "github.com/foostan/photolog/photolog"
)

var RelocateFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "basepath",
		Value: ".",
		Usage: "base directory path of target files",
	},
}

func RelocateCommand(c *cli.Context) {
	basePath := c.String("basepath")

	err := DirExec(basePath, &PhotoLocator{
		BasePath: basePath,
		Logger: log.New(),
	})
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
}
