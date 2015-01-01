package photolog

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

type Statistics struct {
	All      int64
	Make     map[string]int64
	Model    map[string]int64
	Software map[string]int64
	DateTime map[string]int64
	FileSize int64
}

type PhotoStatistics struct {
	BasePath   string
	Logger     *log.Logger
	Statistics Statistics
}

func NewPhotoStatistics(basePath string, logger *log.Logger) *PhotoStatistics {
	photoStatistics := &PhotoStatistics{
		BasePath: basePath,
		Logger:   logger,
	}

	photoStatistics.Reset()

	return photoStatistics
}

func (e *PhotoStatistics) Reset() {
	e.Statistics = Statistics{
		All:      0,
		Make:     make(map[string]int64, 0),
		Model:    make(map[string]int64, 0),
		Software: make(map[string]int64, 0),
		DateTime: make(map[string]int64, 0),
		FileSize: 0,
	}
}

func (e *PhotoStatistics) Run(path string) error {
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
	e.sumMake(pi)
	e.sumModel(pi)
	e.sumSoftware(pi)
	e.sumDateTime(pi)
	e.sumFileSize(pi)

	return nil
}

func (e *PhotoStatistics) sumAll() {
	e.Statistics.All++
}

func (e *PhotoStatistics) sumMake(pi *PhotoInfo) {
	if _, ok := e.Statistics.Make[pi.Make]; ok {
		e.Statistics.Make[pi.Make]++
	} else {
		e.Statistics.Make[pi.Make] = 1
	}
}

func (e *PhotoStatistics) sumModel(pi *PhotoInfo) {
	if _, ok := e.Statistics.Model[pi.Model]; ok {
		e.Statistics.Model[pi.Model]++
	} else {
		e.Statistics.Model[pi.Model] = 1
	}
}

func (e *PhotoStatistics) sumSoftware(pi *PhotoInfo) {
	if _, ok := e.Statistics.Software[pi.Software]; ok {
		e.Statistics.Software[pi.Software]++
	} else {
		e.Statistics.Software[pi.Software] = 1
	}
}

func (e *PhotoStatistics) sumDateTime(pi *PhotoInfo) {
	year := string(fmt.Sprintf("%04d", pi.DateTime.Year()))
	month := string(fmt.Sprintf("%02d", pi.DateTime.Month()))
	day := string(fmt.Sprintf("%02d", pi.DateTime.Day()))

	if _, ok := e.Statistics.DateTime[year]; ok {
		e.Statistics.DateTime[year]++
	} else {
		e.Statistics.DateTime[year] = 1
	}

	if _, ok := e.Statistics.DateTime[year+month]; ok {
		e.Statistics.DateTime[year+month]++
	} else {
		e.Statistics.DateTime[year+month] = 1
	}

	if _, ok := e.Statistics.DateTime[year+month+day]; ok {
		e.Statistics.DateTime[year+month+day]++
	} else {
		e.Statistics.DateTime[year+month+day] = 1
	}
}

func (e *PhotoStatistics) sumFileSize(pi *PhotoInfo) {
	e.Statistics.FileSize += pi.FileSize
}
