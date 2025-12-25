package images

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type SpritesSheet struct {
	image    *ebiten.Image
	cellSize gmath.Vec
}

func NewSpritesSheet(image *ebiten.Image, cellSize gmath.Vec) *SpritesSheet {
	return &SpritesSheet{
		image:    image,
		cellSize: cellSize,
	}
}

func (s SpritesSheet) Get(x, y int) Image {
	start := s.cellSize.Mul(gmath.Vec{X: float64(x), Y: float64(y)})

	rect := image.Rect(int(start.X), int(start.Y), int(start.X+s.cellSize.X), int(start.Y+s.cellSize.Y))

	subImage := s.image.SubImage(rect).(*ebiten.Image)

	return Image{Image: subImage, Scale: gmath.Vec{X: 1, Y: 1}}
}
