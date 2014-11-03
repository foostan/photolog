package photolog

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestRead(t *testing.T) {
	logger := log.New()

	r := PhotoReader{
		logger: logger,
	}

	_, err := r.Read("../test/resources/photos/2014/07/19/sample01.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/06/sample02.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/23/sample03.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
