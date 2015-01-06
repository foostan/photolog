package photolog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Executor interface {
	Run(filepath string) error
}

func DirExec(path string, executor Executor) error {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	cnt := 0
	paths := []string{path}
	errChan := make(chan error)
	semaphore := make(chan int, cpus)
	for len(paths) > 0 {
		path, paths = paths[len(paths)-1], paths[:len(paths)-1]

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

			for _, fi := range contents {
				subpath := filepath.Join(path, fi.Name())
				paths = append(paths, subpath)
			}
		} else {
			cnt++
			go func(path string) {
				semaphore <- 1
				errChan <- executor.Run(path)
				<-semaphore
			}(path)
		}
	}
	for i := 0; i < cnt; i++ {
		err := <-errChan
		if err != nil {
			return err
		}
	}

	return nil
}
