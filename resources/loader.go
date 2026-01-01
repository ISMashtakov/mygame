package resources

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/errs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ResourceLoader struct {
	resources    map[ImageID]*ebiten.Image
	animations   map[AnimationID]*images.AnimationMap
	fonts        map[FontID]*text.GoTextFaceSource
	spriteSheets map[SpriteSheetID]*images.SpritesSheet
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
	l.animations = make(map[AnimationID]*images.AnimationMap, len(animationResources))

	for animationID, animationData := range animationResources {
		image := l.LoadImage(animationData.imageID)

		spriteSheet := images.NewSpritesSheet(image, animationData.cellSize)
		l.animations[animationID] = images.NewAnimationsMap(
			*spriteSheet,
			animationData.frames,
			animationData.directions,
		)
	}

	l.fonts = make(map[FontID]*text.GoTextFaceSource, len(fontResources))

	for fontID, path := range fontResources {
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("can't read font file %q: %w", path, err)
		}

		goFont, err := text.NewGoTextFaceSource(bytes.NewReader(data))
		if err != nil {
			return fmt.Errorf("can't decode font file %q: %w", path, err)
		}

		l.fonts[fontID] = goFont
	}

	l.spriteSheets = make(map[SpriteSheetID]*images.SpritesSheet, len(spriteSheetResources))

	for spriteID, info := range spriteSheetResources {
		data, err := os.ReadFile(info.path)
		if err != nil {
			return fmt.Errorf("can't read image file %q: %w", info.path, err)
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return fmt.Errorf("can't decode image file %q: %w", info.path, err)
		}

		spriteSheet := images.NewSpritesSheet(ebiten.NewImageFromImage(img), info.cellSize)

		l.spriteSheets[spriteID] = spriteSheet
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

func (l *ResourceLoader) LoadAnimationMap(animationID AnimationID) *images.AnimationMap {
	animation, ok := l.animations[animationID]
	if !ok {
		panic(fmt.Errorf("%d: %w", animationID, errs.ErrUnknowsResourceID))
	}

	return animation
}

func (l *ResourceLoader) LoadFont(fontID FontID) *text.GoTextFaceSource {
	font, ok := l.fonts[fontID]
	if !ok {
		panic(fmt.Errorf("%d: %w", fontID, errs.ErrUnknowsResourceID))
	}

	return font
}

func (l *ResourceLoader) LoadSpriteSheet(spriteSheetID SpriteSheetID) *images.SpritesSheet {
	sheet, ok := l.spriteSheets[spriteSheetID]
	if !ok {
		panic(fmt.Errorf("%d: %w", spriteSheetID, errs.ErrUnknowsResourceID))
	}

	return sheet
}
