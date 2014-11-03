package photolog

import (
	"testing"
)

func TestRead(t *testing.T) {
	_, err := Read("../test/resources/photos/2014/07/19/sample01.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = Read("../test/resources/photos/2013/11/06/sample02.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = Read("../test/resources/photos/2013/11/23/sample03.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
