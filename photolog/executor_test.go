package photolog

import (
	"testing"
	log "github.com/Sirupsen/logrus"
)

type TestExecutor struct {
	Logger *log.Logger
}

func (e *TestExecutor) Run(file_path string) error {
	e.Logger.Level = log.ErrorLevel
	e.Logger.Info(file_path)

	return nil
}

func TestDirExec(t *testing.T) {
	err := DirExec("../test/resources/photos", &TestExecutor{
		Logger: log.New(),
	})
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
