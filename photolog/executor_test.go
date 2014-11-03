package photolog

import (
	"testing"
	log "github.com/Sirupsen/logrus"
)

type TestExecutor struct {
}

func (e *TestExecutor) Run(file_path string) error {
	log.SetLevel(log.ErrorLevel)
	log.Info(file_path)

	return nil
}

func TestDirExec(t *testing.T) {
	err := DirExec("../test/resources/photos", &TestExecutor{})
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
