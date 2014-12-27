package photolog

import (
	"github.com/Sirupsen/logrus"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"os"
	"strings"
	"time"
	"path/filepath"
)

type PhotoReader struct {
	Logger *logrus.Logger
}

func (r PhotoReader) Read(fname string) (*PhotoInfo, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	exif.RegisterParsers(mknote.All...)
	readExif, err := exif.Decode(f)
	if err != nil {
		return nil, err
	}

	pi, err := DefaultPhotoInfo()
	if err != nil {
		return nil, err
	}

	pi.FileExt = strings.ToLower(filepath.Ext(f.Name()))

	make, err := readExif.Get("Make")
	if err == nil && make != nil {
		pi.Make = string(make.Val)
	} else {
		r.Logger.Info("missing infomation of a maker")
	}

	model, err := readExif.Get("Model")
	if err == nil && model != nil {
		pi.Model = string(model.Val)
	} else {
		r.Logger.Info("missing infomation of a model")
	}

	software, err := readExif.Get("Software")
	if err == nil && software != nil {
		pi.Software = string(software.Val)
	} else {
		r.Logger.Info("missing infomation of a software")
	}

	dateTime, err := readExif.DateTime()
	if err == nil {
		pi.DateTime = dateTime
	} else {
		r.Logger.Info("missing information of a date time")

		gpsDateStamp, err := readExif.Get("GPSDateStamp")
		if err == nil && gpsDateStamp != nil {
			dateTime, err = time.Parse("2006:01:02", strings.Trim(string(gpsDateStamp.Val[:10]), "utf-8"))
			if err != nil {
				return nil, err
			}
			pi.DateTime = dateTime
		} else {
			r.Logger.Warn("missing infomation of a GPSDateSamp")
		}
	}

	lat, long, err := readExif.LatLong()
	if err == nil {
		pi.GPS = GPS{
			Lat:  lat,
			Long: long,
		}
	} else {
		r.Logger.Info("missing information of GPS Lat and Long")
	}

	return pi, nil
}