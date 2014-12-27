package photolog

import (
	"testing"
	log "github.com/Sirupsen/logrus"
)

func TestExplore(t *testing.T) {
	logger := log.New()
	logger.Level = log.ErrorLevel

	err := DirExec("../test/resources/photos", &PhotoLocator{
		BasePath: "../test/resources/photos",
		Logger: logger,
	})
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
