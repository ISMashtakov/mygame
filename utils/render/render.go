package render

import (
	"image"

	"github.com/quasilyte/gmath"
)

func GetImageScale(originSize image.Rectangle, targetSize image.Rectangle) gmath.Vec {
	return gmath.Vec{X: float64(targetSize.Dx()) / float64(originSize.Dx()), Y: float64(targetSize.Dy()) / float64(originSize.Dy())}
}
