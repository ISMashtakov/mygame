package render

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

func GetImageScale(originSize image.Rectangle, targetSize image.Rectangle) gmath.Vec {
	return gmath.Vec{X: float64(targetSize.Dx()) / float64(originSize.Dx()), Y: float64(targetSize.Dy()) / float64(originSize.Dy())}
}

func AtImage(image *ebiten.Image, vec gmath.Vec) color.Color {
	vec = vec.Add(gmath.VecFromStd(image.Bounds().Min))
	return image.At(int(vec.X), int(vec.Y))
}
