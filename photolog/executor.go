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
		return fmt.Errorf("Read '%s': %s", path, err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return fmt.Errorf("Read '%s': %s", path, err)
	}

	if fi.IsDir() {
		contents, err := f.Readdir(-1)
		if err != nil {
			return fmt.Errorf("Read '%s': %s", path, err)
		}

		errChan := make(chan error)
		for _, fi := range contents {
			subpath := filepath.Join(path, fi.Name())
			go func(subpath string) {
				errChan <- DirExec(subpath, executor)
			}(subpath)
		}
		for i := 0; i < len(contents); i++ {
			if <-errChan != nil {
				return err
			}
		}
	} else {
		err := executor.Run(path)
		if err != nil {
			return fmt.Errorf("Run at '%s': %s", path, err)
		}
	}

	return nil
}
