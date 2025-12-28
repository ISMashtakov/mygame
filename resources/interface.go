package resources

import (
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type IResourceLoader interface {
	LoadImage(imageID ImageID) *ebiten.Image
	LoadAnimationMap(animationID AnimationID) *images.AnimationMap
	LoadFont(fontID FontID) *text.GoTextFaceSource
}
