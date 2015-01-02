package photolog

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
	"io/ioutil"
)

type PhotoLocator struct {
	BasePath string
	Logger   *log.Logger
}

func (e *PhotoLocator) Run(path string) error {
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

	photoLocation, err := e.getLocation(pi)
	if err != nil {
		return err
	}

	if path != photoLocation {
		e.Logger.Info("Rename to " + photoLocation)

		err = os.MkdirAll(filepath.Dir(photoLocation), os.FileMode(0755))
		if err != nil {
			return fmt.Errorf("Create directories of '%s' : %s", photoLocation, err)
		}

		err = os.Rename(path, photoLocation)
		if err != nil {
			return err
		}

		err = removeAllEmpDir(filepath.Dir(path))
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

	return filepath.Join(e.BasePath, year, month, day, name), nil
}

func removeAllEmpDir(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		err := os.Remove(path)
		if err != nil {
			return fmt.Errorf("Error while removing '%s' : %s", path, err)
		}

		return removeAllEmpDir(filepath.Dir(path))
	}

	return nil
}
