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

type ResourceLoader struct {
	resources map[ImageID]*ebiten.Image
}

func NewResourceLoader() *ResourceLoader {
	return &ResourceLoader{}
}

func (l *ResourceLoader) Preload() error {
	l.resources = make(map[ImageID]*ebiten.Image, len(imageResources))

	for imageID, path := range imageResources {
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("can't read image file %q: %w", path, err)
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return fmt.Errorf("can't decode image file %q: %w", path, err)
		}

		l.resources[imageID] = ebiten.NewImageFromImage(img)
	}

	return nil
}

func (l *ResourceLoader) LoadImage(imageID ImageID) (*ebiten.Image, error) {
	image, ok := l.resources[imageID]
	if !ok {
		return nil, fmt.Errorf("%q: %w", imageID, errs.ErrUnknowsResourceID)
	}

	return image, nil
}
