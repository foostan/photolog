package photolog

import (
	"testing"
)

func TestCameraName(t *testing.T) {
	pi, err := DefaultPhotoInfo()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	pi.Make = "Make"
	pi.Model = "Model"
	pi.Software = "Software"
}
