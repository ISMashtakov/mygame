package animations

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/utils/funcs"
	"github.com/quasilyte/gmath"
)

type ScaleAnimation struct {
	xFunc      funcs.Func
	yFunc      funcs.Func
	spriteData *components.SpriteData
	baseSprite components.SpriteData
}

func NewScaleAnimation(
	xFunc funcs.Func,
	yFunc funcs.Func,
	spriteData *components.SpriteData,
) *ScaleAnimation {
	return &ScaleAnimation{
		xFunc:      xFunc,
		yFunc:      yFunc,
		spriteData: spriteData,
		baseSprite: *spriteData,
	}
}

func (a *ScaleAnimation) Next(frame int) {
	a.spriteData.Image.Scale = a.baseSprite.Image.Scale.Add(gmath.Vec{
		X: a.xFunc(float64(frame)),
		Y: a.yFunc(float64(frame)),
	})
}
