package photolog

import (
	log "github.com/Sirupsen/logrus"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"os"
	"strings"
	"time"
)

type GPS struct {
	Lat  float64
	Long float64
}

type PhotoInfo struct {
	Make     string
	Model    string
	Software string
	DateTime time.Time
	GPS      GPS
}

func DefaultPhotoInfo() (*PhotoInfo, error) {
	defaultTime, err := time.Parse("2006", "2014")
	if err != nil {
		return nil, err
	}

	photoInfo := &PhotoInfo{
		Make:     "",
		Model:    "",
		Software: "",
		DateTime: defaultTime,
		GPS: GPS{
			Lat:  0,
			Long: 0,
		},
	}

	return photoInfo, nil
}

func Read(fname string) (*PhotoInfo, error) {
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
		log.Info("missing infomation of a maker")
	}

	model, err := readExif.Get("Model")
	if err == nil && model != nil {
		photoInfo.Model = string(model.Val)
	} else {
		log.Info("missing infomation of a model")
	}

	software, err := readExif.Get("Software")
	if err == nil && software != nil {
		photoInfo.Software = string(software.Val)
	} else {
		log.Info("missing infomation of a software")
	}

	dateTime, err := readExif.DateTime()
	if err == nil {
		photoInfo.DateTime = dateTime
	} else {
		log.Info("missing information of a date time")

		gpsDateStamp, err := readExif.Get("GPSDateStamp")
		if err == nil && gpsDateStamp != nil {
			dateTime, err = time.Parse("2006:01:02", strings.Trim(string(gpsDateStamp.Val[:10]), "utf-8"))
			if err != nil {
				return nil, err
			}
			photoInfo.DateTime = dateTime
		} else {
			log.Warn("missing infomation of a GPSDateSamp")
		}
	}

	lat, long, err := readExif.LatLong()
	if err == nil {
		photoInfo.GPS = GPS{
			Lat:  lat,
			Long: long,
		}
	} else {
		log.Info("missing information of GPS Lat and Long")
	}

	return photoInfo, nil
}

func (pi *PhotoInfo) CameraName() string {
	names := make([]string, 0)

	if(pi.Make != "") {
		names = append(names, pi.Make)
	}

	if(pi.Model != "") {
		names = append(names, pi.Model)
	}

	if(pi.Software != "") {
		names = append(names, pi.Software)
	}

	return strings.Join(names, " ")
}
