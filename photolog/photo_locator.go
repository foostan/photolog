package photolog

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"strings"
)

type PhotoLocator struct {
	BasePath string
	Logger   *log.Logger
}

func (e *PhotoLocator) Run(filePath string) error {
	e.Logger.Info("read " + filePath)
	reader := PhotoReader{
		Logger: e.Logger,
	}

	pi, err := reader.Read(filePath)
	if err != nil {
		return err
	}

	location, err := e.getLocation(pi)
	if err != nil {
		return err
	}

	if filePath != location {
		e.Logger.Info("rename to " + location)
		err = os.Rename(filePath, location)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *PhotoLocator) getLocation(pi *PhotoInfo) (string, error) {
	year := fmt.Sprintf("%04d", pi.DateTime.Year())
	month := fmt.Sprintf("%02d", pi.DateTime.Month())
	day := fmt.Sprintf("%02d", pi.DateTime.Day())
	name, err := pi.FileName()
	if err != nil {
		return "", err
	}

	return strings.Join([]string{e.BasePath, year, month, day, name}, "/"), nil
}
