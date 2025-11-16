package resources

import "github.com/hajimehoshi/ebiten/v2"

type IResourceLoader interface {
	LoadImage(imageID ImageID) (*ebiten.Image, error)
}
