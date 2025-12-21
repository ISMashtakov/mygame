package animations

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/utils/funcs"
	"github.com/quasilyte/gmath"
)

type MoveAnimation struct {
	xFunc        funcs.Func
	yFunc        funcs.Func
	positionData *components.PositionData
	basePosition components.PositionData
}

func NewMoveAnimation(
	xFunc funcs.Func,
	yFunc funcs.Func,
	positionData *components.PositionData,
) *MoveAnimation {
	return &MoveAnimation{
		xFunc:        xFunc,
		yFunc:        yFunc,
		positionData: positionData,
		basePosition: *positionData,
	}
}

func (a *MoveAnimation) Next(frame int) {
	a.positionData.Vec = a.basePosition.Vec.Add(gmath.Vec{
		X: a.xFunc(float64(frame)),
		Y: a.yFunc(float64(frame)),
	})
}
