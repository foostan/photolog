package photolog

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestPhotoStatistics(t *testing.T) {
	logger := log.New()
	logger.Level = log.ErrorLevel

	err := DirExec("../test/resources/photos", NewPhotoStatistics("../test/resources/photos", logger))
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
