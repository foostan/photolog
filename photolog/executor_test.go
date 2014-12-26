package photolog

import (
	"testing"
	log "github.com/Sirupsen/logrus"
)

type TestExecutor struct {
	Logger *log.Logger
}

func (e *TestExecutor) Run(file_path string) error {
	e.logger.Level = log.ErrorLevel
	e.logger.Info(file_path)

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
