package command

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"encoding/json"
	. "github.com/foostan/photolog/photolog"
)

var StatsFlags = []cli.Flag{
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

func StatsCommand(c *cli.Context) {
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
	ps := NewPhotoStats(basePath, logger)
	err = DirExec(basePath, ps)
	if err != nil {
		logger.Fatal(err)
	}

	json, err := json.Marshal(ps.Stats)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(string(json))
}
