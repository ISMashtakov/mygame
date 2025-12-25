package images

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type Image struct {
	*ebiten.Image

	Flip  bool
	Scale gmath.Vec
}
