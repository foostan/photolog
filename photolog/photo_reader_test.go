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

	_, err := r.Read("../test/resources/photos/2014/07/19/14057280003812191413.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/06/13837499061493942213.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/23/13851841024249406430.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
