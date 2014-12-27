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

	_, err := r.Read("../test/resources/photos/2014/07/19/14057280006d69d2f804acdaf59f2a7d51b9eaba0f.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/06/13837499064ce0c4fe2c204de959707e0f877d38f8.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}

	_, err = r.Read("../test/resources/photos/2013/11/23/1385184102d37d22df6fd6294094f752313b88e362.jpg")
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
