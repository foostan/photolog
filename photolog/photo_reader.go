package photolog

import (
	"github.com/Sirupsen/logrus"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type PhotoReader struct {
	Logger *logrus.Logger
}

func (r PhotoReader) Read(path string) (*PhotoInfo, error) {
	defer func() {
		err := recover()
		if err != nil {
			r.Logger.Errorf("can't resolve as a image file: %s, %s", path, err)
		}
	}()

	r.Logger.Info("read " + path)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	readExif, err := exif.Decode(f)
	if err != nil {
		r.Logger.Warnf("can't resolve as a image file: %s", path)
		return nil, nil
	}

	pi, err := DefaultPhotoInfo()
	if err != nil {
		return nil, err
	}

	pi.FileSize = fi.Size()
	pi.FileExt = strings.ToLower(filepath.Ext(f.Name()))

	make, err := readExif.Get("Make")
	if err == nil && make != nil {
		pi.Make = strings.TrimSpace(strings.Trim(string(make.Val), "\u0000"))
	} else {
		r.Logger.Debug("missing infomation of a maker")
	}

	model, err := readExif.Get("Model")
	if err == nil && model != nil {
		pi.Model = strings.TrimSpace(strings.Trim(string(model.Val), "\u0000"))
	} else {
		r.Logger.Debug("missing infomation of a model")
	}

	software, err := readExif.Get("Software")
	if err == nil && software != nil {
		pi.Software = strings.TrimSpace(strings.Trim(string(software.Val), "\u0000"))
	} else {
		r.Logger.Debug("missing infomation of a software")
	}

	dateTime, err := readExif.DateTime()
	if err == nil {
		pi.DateTime = dateTime
	} else {
		r.Logger.Debug("missing information of a date time")

		gpsDateStamp, err := readExif.Get("GPSDateStamp")
		if err == nil && gpsDateStamp != nil {
			dateTime, err = time.Parse("2006:01:02", strings.Trim(string(gpsDateStamp.Val[:10]), "utf-8"))
			if err != nil {
				return nil, err
			}
			pi.DateTime = dateTime
		} else {
			r.Logger.Debug("missing infomation of a GPSDateSamp")
			r.Logger.Warn("missing information all of date time")
		}
	}

	lat, long, err := readExif.LatLong()
	if err == nil {
		pi.GPS = GPS{
			Lat:  lat,
			Long: long,
		}
	} else {
		r.Logger.Debug("missing information of GPS Lat and Long")
	}

	return pi, nil
}
