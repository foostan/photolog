package photolog

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestPhotoThumbnailer(t *testing.T) {
	logger := log.New()
	logger.Level = log.ErrorLevel

	err := DirExec("../test/resources/photos", &PhotoThumbnailer{
		BasePath:     "../test/resources/photos",
		BaseThumPath: "../test/resources/photo_thumbnails",
		ThumSizes: []ThumSize{
			ThumSize{
				Width:  240,
				Height: 0,
			},
			ThumSize{
				Width:  32,
				Height: 0,
			},
		},
		Logger: logger,
	})
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
