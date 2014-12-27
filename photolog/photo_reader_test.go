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

	_, err := r.Read("../test/resources/photos/2014/07/19/1405728000b5d4582798f000c70a0e97454dd63f5d.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/06/1383749906675e1fa8d8915b256f79601adca1e74e.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/23/13851841026d550cf064aea2c866fc64314257d39b.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
