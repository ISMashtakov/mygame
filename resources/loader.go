package resources

import (
	"bytes"
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/ISMashtakov/mygame/errs"
	"github.com/hajimehoshi/ebiten/v2"
)

type ResourceLoader struct{}

func (l *ResourceLoader) LoadImage(imageID ImageID) (*ebiten.Image, error) {
	path, ok := imageResources[imageID]
	if !ok {
		return nil, fmt.Errorf("%q: %w", imageID, errs.ErrUnknowsResourceID)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't read image file %q: %w", path, err)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("can't decode image file %q: %w", path, err)
	}

	return ebiten.NewImageFromImage(img), nil
}

type ImageID = int

var (
	imageResources = map[ImageID]string{
		ImageCharacter: "resources/images/char.png",
	}
)

const (
	ImageNone ImageID = iota
	ImageCharacter
)
