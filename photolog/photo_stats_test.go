package photolog

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestPhotoStats(t *testing.T) {
	logger := log.New()
	logger.Level = log.ErrorLevel

	err := DirExec("../test/resources/photos", NewPhotoStats("../test/resources/photos", logger))
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
