package photolog

import (
	log "github.com/Sirupsen/logrus"
)

type PhotoLocator struct {
	base_path string
	logger *log.Logger
}

func (e *PhotoLocator) Run(file_path string) error {
	reader := PhotoReader{
		logger: e.logger,
	}

	pi, err := reader.Read(file_path)
	if err != nil {
		return err
	}

	e.logger.Warn("base_path: ",e.base_path)
	e.logger.Warn("read: ",file_path)
	e.logger.Info(pi)

	return nil
}

