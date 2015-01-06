package command

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	. "github.com/foostan/photolog/photolog"
)

var ThumbnailFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "src-dir",
		Value: ".",
		Usage: "base directory path of source files",
	},
	cli.StringFlag{
		Name:  "dst-dir",
		Value: ".",
		Usage: "base directory path of destination files",
	},
	cli.StringFlag{
		Name:  "log-level",
		Value: "warn",
		Usage: "logger level",
	},
}

func ThumbnailCommand(c *cli.Context) {
	// setup logger
	logLvStr := c.String("log-level")
	logLevel, err := log.ParseLevel(logLvStr)
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	logger := log.New()
	logger.Level = logLevel

	// run command
	srcDir := c.String("src-dir")
	dstDir := c.String("dst-dir")
	err = DirExec(srcDir, NewPhotoThumbnailer(srcDir, dstDir, make([]ThumSize, 0), logger))
	if err != nil {
		logger.Fatal(err)
	}
}
