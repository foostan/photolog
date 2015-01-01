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
	cli.StringFlag{
		Name:  "log-level",
		Value: "warn",
		Usage: "logger level",
	},
}

func RelocateCommand(c *cli.Context) {
	// setup logger
	logLvStr := c.String("log-level")
	logLevel, err := log.ParseLevel(logLvStr)
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	logger := log.New()
	logger.Level = logLevel

	// run command
	basePath := c.String("basepath")
	ps := NewPhotoStatistics(basePath, logger)
	err = DirExec(basePath, ps)
	if err != nil {
		logger.Fatal(err)
	}
	
	err = DirExec(basePath, &PhotoLocator{
		BasePath: basePath,
		Logger: logger,
	})
	if err != nil {
		logger.Fatal(err)
	}
}
