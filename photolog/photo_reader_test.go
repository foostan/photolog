package photolog

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestRead(t *testing.T) {
	logger := log.New()
	logger.Level = log.ErrorLevel

	r := PhotoReader{
		Logger: logger,
	}

	_, err := r.Read("../test/resources/photos/2014/07/19/14057280001212432156.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/06/13837499061012862994.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/23/13851841023797588213.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
