package photolog

import (
	"testing"
	"fmt"
)

type TestExecutor struct {

}

func (e *TestExecutor) Run(file_path string) error {
	fmt.Println(file_path)

	return nil
}

func TestDirExec(t *testing.T) {
	err := DirExec("../test/resources/photos", &TestExecutor{})
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
