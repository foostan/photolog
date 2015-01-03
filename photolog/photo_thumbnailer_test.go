package photolog

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestPhotoThumbnailer(t *testing.T) {
	logger := log.New()
	logger.Level = log.ErrorLevel

	err := DirExec("../test/resources/photos", NewPhotoThumbnailer(
		"../test/resources/photos",
		"../test/resources/photo_thumbnails",
		make([]ThumSize, 0),
		logger))
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
