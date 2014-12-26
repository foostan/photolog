package photolog

import (
	"testing"
	log "github.com/Sirupsen/logrus"
)

func TestExplore(t *testing.T) {
	err := DirExec("../test/resources/photos", &PhotoLocator{
		BasePath: "../test/resources/photos",
		Logger: log.New(),
	})
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
