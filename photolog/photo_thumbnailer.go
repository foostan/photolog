package photolog

import (
	log "github.com/Sirupsen/logrus"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"os"
	"github.com/nfnt/resize"
	"path/filepath"
	"fmt"
)

type ThumSize struct {
	Width  uint
	Height uint
}

type PhotoThumbnailer struct {
	BasePath     string
	BaseThumPath string
	ThumSizes    []ThumSize
	Logger       *log.Logger
}

func init() {
	image.RegisterFormat("jpeg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
}

func (e *PhotoThumbnailer) Run(path string) error {
	e.Logger.Info("read " + path)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	for _, thumSize := range e.ThumSizes {
		thumImg := resize.Resize(thumSize.Width, thumSize.Height, img, resize.Lanczos3)

		name := fmt.Sprintf("%d_%d_%s", thumSize.Width, thumSize.Height, fi.Name())
		out, err := os.Create(filepath.Join(e.BaseThumPath, name))
		if err != nil {
			return err
		}
		defer out.Close()

		err = jpeg.Encode(out, thumImg, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

