package photolog

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

type Stats struct {
	All        int64
	CameraName map[string]int64
	DateTime   map[string]int64
	FileSize   int64
}

type PhotoStats struct {
	BasePath string
	Logger   *log.Logger
	Stats    Stats
}

func NewPhotoStats(basePath string, logger *log.Logger) *PhotoStats {
	photoStats := &PhotoStats{
		BasePath: basePath,
		Logger:   logger,
	}

	photoStats.Reset()

	return photoStats
}

func (e *PhotoStats) Reset() {
	e.Stats = Stats{
		All:        0,
		CameraName: make(map[string]int64, 0),
		DateTime:   make(map[string]int64, 0),
		FileSize:   0,
	}
}

func (e *PhotoStats) Run(path string) error {
	reader := PhotoReader{
		Logger: e.Logger,
	}

	pi, err := reader.Read(path)
	if err != nil {
		return err
	}
	if pi == nil {
		return nil
	}

	e.sumAll()
	e.sumCameraName(pi)
	e.sumDateTime(pi)
	e.sumFileSize(pi)

	return nil
}

func (e *PhotoStats) sumAll() {
	e.Stats.All++
}

func (e *PhotoStats) sumCameraName(pi *PhotoInfo) {
	name := pi.CameraName()

	if _, ok := e.Stats.CameraName[name]; ok {
		e.Stats.CameraName[name]++
	} else {
		e.Stats.CameraName[name] = 1
	}
}

func (e *PhotoStats) sumDateTime(pi *PhotoInfo) {
	year := string(fmt.Sprintf("%04d", pi.DateTime.Year()))
	month := string(fmt.Sprintf("%02d", pi.DateTime.Month()))
	day := string(fmt.Sprintf("%02d", pi.DateTime.Day()))

	if _, ok := e.Stats.DateTime[year]; ok {
		e.Stats.DateTime[year]++
	} else {
		e.Stats.DateTime[year] = 1
	}

	if _, ok := e.Stats.DateTime[year+month]; ok {
		e.Stats.DateTime[year+month]++
	} else {
		e.Stats.DateTime[year+month] = 1
	}

	if _, ok := e.Stats.DateTime[year+month+day]; ok {
		e.Stats.DateTime[year+month+day]++
	} else {
		e.Stats.DateTime[year+month+day] = 1
	}
}

func (e *PhotoStats) sumFileSize(pi *PhotoInfo) {
	e.Stats.FileSize += pi.FileSize
}
