package photolog

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"strings"
	"os"
	"path/filepath"
)

type PhotoLocator struct {
	BasePath string
	Logger   *log.Logger
}

func (e *PhotoLocator) Run(path string) error {
	e.Logger.Info("read " + path)
	reader := PhotoReader{
		Logger: e.Logger,
	}

	pi, err := reader.Read(path)
	if err != nil {
		return err
	}
	if pi == nil {
		return nil
	}

	location, err := e.getLocation(pi)
	if err != nil {
		return err
	}

	if path != location {
		e.Logger.Info("Rename to " + location)

		err = os.MkdirAll(filepath.Dir(location), os.FileMode(0755))
		if err != nil {
			return fmt.Errorf("Create directories of '%s' : %s", location, err)
		}

		err = os.Rename(path, location)
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
