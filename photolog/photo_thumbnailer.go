package photolog

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type ThumSize struct {
	Name  string
	Scale uint
}

type PhotoThumbnailer struct {
	SrcDir    string
	DstDir    string
	ThumSizes []ThumSize
	Logger    *log.Logger
}

func init() {
	image.RegisterFormat("jpeg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
}

func NewPhotoThumbnailer(srcDir string, dstDir string, thumSizes []ThumSize, logger *log.Logger) *PhotoThumbnailer {
	if len(thumSizes) == 0 {
		thumSizes = []ThumSize{
			ThumSize{
				Name:  "s",
				Scale: 240,
			},
			ThumSize{
				Name:  "l",
				Scale: 960,
			},
		}
	}

	pt := &PhotoThumbnailer{
		SrcDir:    srcDir,
		DstDir:    dstDir,
		ThumSizes: thumSizes,
		Logger:    logger,
	}

	return pt
}

func (e *PhotoThumbnailer) Run(originPath string) error {
	e.Logger.Info("make thumbnail(s) of " + originPath)

	p := strings.Replace(originPath, e.SrcDir, e.DstDir, 1)
	dstDirBase := filepath.Dir(p)

	err := os.MkdirAll(dstDirBase, os.FileMode(0755))
	if err != nil {
		return fmt.Errorf("Create directories of '%s' : %s", dstDirBase, err)
	}

	originImg, originName, err := e.OriginImg(originPath)
	if err != nil {
		return err
	}

	for _, thumSize := range e.ThumSizes {
		w := thumSize.Scale
		h := uint(0)
		if originImg.Bounds().Max.X < originImg.Bounds().Max.Y {
			w = uint(0)
			h = thumSize.Scale
		}
		thumImg := resize.Resize(w, h, originImg, resize.Lanczos3)
		fileName := fmt.Sprintf("%s_%s", thumSize.Name, originName)

		out, err := os.Create(filepath.Join(dstDirBase, fileName))
		if err != nil {
			return err
		}
		defer out.Close()

		err = jpeg.Encode(out, thumImg, nil)
		if err != nil {
			return err
		}

		e.Logger.Info("made new file: " + filepath.Join(dstDirBase, fileName))
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
