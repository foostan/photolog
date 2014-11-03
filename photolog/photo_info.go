package photolog

import (
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
	defaultTime, err := time.Parse("2006", "0001")
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

func (pi *PhotoInfo) CameraName() string {
	names := make([]string, 0)

	if pi.Make != "" {
		names = append(names, pi.Make)
	}

	if pi.Model != "" {
		names = append(names, pi.Model)
	}

	if pi.Software != "" {
		names = append(names, pi.Software)
	}

	return strings.Join(names, " ")
}
