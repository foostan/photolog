package photolog

import (
	"fmt"
	"os"
	"path/filepath"
)

type Executor interface {
	Run(filepath string) error
}

func DirExec(path string, executor Executor) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error reading '%s': %s", path, err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("Error reading '%s': %s", path, err)
	}

	if fi.IsDir() {
		contents, err := f.Readdir(-1)
		if err != nil {
			return fmt.Errorf("Error reading '%s': %s", path, err)
		}

		for _, fi := range contents {
			subpath := filepath.Join(path, fi.Name())
			err = DirExec(subpath, executor)
			if err != nil {
				return err
			}
		}
	} else {
		executor.Run(path)
	}

	return nil
}
