package photolog

import (
	log "github.com/Sirupsen/logrus"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"os"
	"github.com/nfnt/resize"
	"path/filepath"
	"fmt"
	"strings"
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
	thumPath := strings.Replace(path, e.BasePath, e.BaseThumPath, 1)
	thumPathBase := filepath.Dir(thumPath)

	err := os.MkdirAll(thumPathBase, os.FileMode(0755))
	if err != nil {
		return fmt.Errorf("Create directories of '%s' : %s", thumPath, err)
	}
	
	originImg, originName, err := e.OriginImg(path)
	if err != nil {
		return err
	}

	for _, thumSize := range e.ThumSizes {
		thumImg := resize.Resize(thumSize.Width, thumSize.Height, originImg, resize.Lanczos3)

		thumNameWithSize := fmt.Sprintf("%d_%d_%s", thumSize.Width, thumSize.Height, originName)
		out, err := os.Create(filepath.Join(thumPathBase, thumNameWithSize))
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

func (e *PhotoThumbnailer) OriginImg(path string) (image.Image, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, "", err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}
	
	return img, fi.Name(), nil
}
