package resources

import (
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/hajimehoshi/ebiten/v2"
)

type IResourceLoader interface {
	LoadImage(imageID ImageID) *ebiten.Image
	LoadAnimation(animationID AnimationID) *images.Animation
}
