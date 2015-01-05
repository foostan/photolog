package photolog

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestPhotoLocator(t *testing.T) {
	logger := log.New()
	logger.Level = log.ErrorLevel
	srcDir := "../test/resources/photos"
	dstDir := srcDir
	mode := "move"

	err := DirExec(srcDir, NewPhotoLocator(srcDir, dstDir, mode, logger))
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
