package render

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

func GetImageScale(originSize image.Rectangle, targetSize gmath.Vec) gmath.Vec {
	return gmath.Vec{X: targetSize.X / float64(originSize.Dx()), Y: targetSize.Y / float64(originSize.Dy())}
}

func AtImage(image *ebiten.Image, vec gmath.Vec) color.Color {
	vec = vec.Add(gmath.VecFromStd(image.Bounds().Min))
	return image.At(int(vec.X), int(vec.Y))
}
