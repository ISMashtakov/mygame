package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type Sprite struct {
	Image *ebiten.Image
	Scale *gmath.Vec
}
