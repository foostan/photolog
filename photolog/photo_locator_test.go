package photolog

import (
	"testing"
	log "github.com/Sirupsen/logrus"
)

func TestExplore(t *testing.T) {
	err := DirExec("../test/resources/photos", &PhotoLocator{
		basePath: "../test/resources/photos",
		logger: log.New(),
	})
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
