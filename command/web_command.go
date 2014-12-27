package command

import (
	"github.com/codegangsta/cli"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"path/filepath"
	"os"
)

var WebFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "port",
		Value: "3000",
	},
	cli.StringFlag{
		Name:  "ui-dir",
		Value: "./ui",
		Usage: "ui directory path",
	},
	cli.StringFlag{
		Name:  "photos-dir",
		Value: "./photos",
		Usage: "base directory path of photo files",
	},
}

func WebCommand(c *cli.Context) {
	m := martini.Classic()

	err := os.Setenv("PORT", c.String("port"))
	if err != nil {
		fmt.Errorf("%v", err)
	}

	uiDir, err := filepath.Abs(c.String("ui-dir"))
	if err != nil {
		fmt.Errorf("%v", err)
	}

	m.Use(martini.Static(uiDir, martini.StaticOptions{
		Prefix: "ui",
	}))

	photosDir, err := filepath.Abs(c.String("photos-dir"))
	if err != nil {
		fmt.Errorf("%v", err)
	}

	m.Use(martini.Static(photosDir, martini.StaticOptions{
		Prefix: "photos",
	}))

	m.Use(render.Renderer())

	m.Get("/v1", func(r render.Render) {
			r.JSON(200, map[string]interface{}{"ui_dir": uiDir, "photos_dir": photosDir})
		})

	m.Run()
}
