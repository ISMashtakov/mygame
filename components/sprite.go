package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image *ebiten.Image
	Scale gmath.Vec
}

var Sprite = donburi.NewComponentType[SpriteData]()
