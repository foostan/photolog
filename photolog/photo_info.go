package photolog

import (
	"strings"
	"time"
	"crypto/md5"
	"encoding/json"
	"io"
	"fmt"
	"strconv"
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
	FileExt  string
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
		FileExt:   "",
	}

	return photoInfo, nil
}

func (pi *PhotoInfo) FileName() (string, error) {
	json, err := json.Marshal(pi)
	if err != nil {
		return "", err
	}

	timeStr := strconv.FormatInt(pi.DateTime.Unix(), 10)

	h := md5.New()
	io.WriteString(h, string(json))
	nameHashStr := fmt.Sprintf("%x", h.Sum(nil))

	return timeStr+nameHashStr+pi.FileExt, nil
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
