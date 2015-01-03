package photolog

import (
	"testing"
	log "github.com/Sirupsen/logrus"
)

func TestExplore(t *testing.T) {
	logger := log.New()
	logger.Level = log.ErrorLevel
	srcDir := "../test/resources/photos"
	dstDir := srcDir

	err := DirExec(srcDir, NewPhotoLocator(srcDir, dstDir, logger))
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
