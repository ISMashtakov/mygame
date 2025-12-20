package resources

import (
	"bytes"
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/errs"
	"github.com/hajimehoshi/ebiten/v2"
)

type ResourceLoader struct {
	resources  map[ImageID]*ebiten.Image
	animations map[AnimationID]*images.Animation
}

func NewResourceLoader() *ResourceLoader {
	return &ResourceLoader{}
}

func (l *ResourceLoader) Preload() error {
	// Images
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

	// Animations
	l.animations = make(map[AnimationID]*images.Animation, len(animationResources))

	for animationID, animationData := range animationResources {
		image := l.LoadImage(animationData.imageID)

		spriteSheet := images.NewSpritesSheet(image, animationData.cellSize)
		animationMap := images.NewAnimationsMap(*spriteSheet, animationData.frames, animationData.directions)
		l.animations[animationID] = images.NewAnimation(*animationMap, animationData.duration)
	}

	return nil
}

func (l *ResourceLoader) LoadImage(imageID ImageID) *ebiten.Image {
	image, ok := l.resources[imageID]
	if !ok {
		panic(fmt.Errorf("%d: %w", imageID, errs.ErrUnknowsResourceID))
	}

	return image
}

func (l *ResourceLoader) LoadAnimation(animationID AnimationID) *images.Animation {
	animation, ok := l.animations[animationID]
	if !ok {
		panic(fmt.Errorf("%d: %w", animationID, errs.ErrUnknowsResourceID))
	}

	return animation
}
