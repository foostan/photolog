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
	acCN := pi.CameraName()
	exCN := "Make Model Software"

	if acCN != exCN {
		t.Fatalf("expected result is %s, but %s", acCN, exCN)
	}
}
