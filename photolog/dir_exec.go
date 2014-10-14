package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirExec(path string) ([]string, error) {
	list := make([]string, 0)

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading '%s': %s", path, err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("Error reading '%s': %s", path, err)
	}

	if fi.IsDir() {
		contents, err := f.Readdir(-1)
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		for _, fi := range contents {
			subpath := filepath.Join(path, fi.Name())
			list, err = DirExec(subpath)
			if err != nil {
				return nil, err
			}
		}
	} else {
		fmt.Println(path)
	}

	return list, nil
}
