package images

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type SpritesSheet struct {
	image       *ebiten.Image
	startOffset gmath.Vec
	offset      gmath.Vec
	cellSize    gmath.Vec
}

func NewSpritesSheet(image *ebiten.Image, startOffset gmath.Vec, offset gmath.Vec, cellSize gmath.Vec) *SpritesSheet {
	return &SpritesSheet{
		image:       image,
		offset:      offset,
		cellSize:    cellSize,
		startOffset: startOffset,
	}
}

func (s SpritesSheet) Get(x, y int) *ebiten.Image {
	start := s.startOffset.Add(s.offset.Mul(gmath.Vec{X: float64(x), Y: float64(y)}))

	rect := image.Rect(int(start.X), int(start.Y), int(start.X+s.cellSize.X), int(start.Y+s.cellSize.Y))

	subImage := s.image.SubImage(rect)

	return subImage.(*ebiten.Image)
}
