package photolog

import (
	"github.com/Sirupsen/logrus"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"os"
	"strings"
	"time"
)

type PhotoReader struct {
	logger *logrus.Logger
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

	photoInfo, err := DefaultPhotoInfo()
	if err != nil {
		return nil, err
	}

	make, err := readExif.Get("Make")
	if err == nil && make != nil {
		photoInfo.Make = string(make.Val)
	} else {
		r.logger.Info("missing infomation of a maker")
	}

	model, err := readExif.Get("Model")
	if err == nil && model != nil {
		photoInfo.Model = string(model.Val)
	} else {
		r.logger.Info("missing infomation of a model")
	}

	software, err := readExif.Get("Software")
	if err == nil && software != nil {
		photoInfo.Software = string(software.Val)
	} else {
		r.logger.Info("missing infomation of a software")
	}

	dateTime, err := readExif.DateTime()
	if err == nil {
		photoInfo.DateTime = dateTime
	} else {
		r.logger.Info("missing information of a date time")

		gpsDateStamp, err := readExif.Get("GPSDateStamp")
		if err == nil && gpsDateStamp != nil {
			dateTime, err = time.Parse("2006:01:02", strings.Trim(string(gpsDateStamp.Val[:10]), "utf-8"))
			if err != nil {
				return nil, err
			}
			photoInfo.DateTime = dateTime
		} else {
			r.logger.Warn("missing infomation of a GPSDateSamp")
		}
	}

	lat, long, err := readExif.LatLong()
	if err == nil {
		photoInfo.GPS = GPS{
			Lat:  lat,
			Long: long,
		}
	} else {
		r.logger.Info("missing information of GPS Lat and Long")
	}

	return photoInfo, nil
}
