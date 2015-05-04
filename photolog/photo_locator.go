package photolog

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

type PhotoLocator struct {
	SrcDir string
	DstDir string
	Mode   string
	Logger *log.Logger
}

func NewPhotoLocator(srcDir string, dstDir string, mode string, logger *log.Logger) *PhotoLocator {
	photoLocator := &PhotoLocator{
		SrcDir: srcDir,
		DstDir: dstDir,
		Mode:   mode,
		Logger: logger,
	}

	return photoLocator
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

		switch e.Mode {
		case "move":
			err = os.Rename(path, photoLocation)
			if err != nil {
				return err
			}
		case "link":
			err = os.Link(path, photoLocation)
			if err != nil {
				return err
			}
		case "symlink":
			abspath, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			err = os.Symlink(abspath, photoLocation)
			if err != nil {
				return err
			}
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

	return filepath.Join(e.DstDir, year, year+month, year+month+day, name), nil
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
