package images

import (
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/samber/lo"
)

type AnimationMap struct {
	spritesSheet      SpritesSheet
	framesInAnimation int
	directions        []direction.DirectionEnum
}

func NewAnimationsMap(spritesSheet SpritesSheet, framesInAnimation int, directions []direction.DirectionEnum) *AnimationMap {
	return &AnimationMap{
		spritesSheet:      spritesSheet,
		framesInAnimation: framesInAnimation,
		directions:        directions,
	}
}

func (m AnimationMap) GetByDirection(dir direction.DirectionEnum, i int) Image {
	resDir := dir
	if dir == direction.Left {
		resDir = direction.Right
	}

	im := m.spritesSheet.Get(i, lo.IndexOf(m.directions, resDir))
	im.Flip = dir == direction.Left

	return im
}

func (m AnimationMap) GetCountFrames() int {
	return m.framesInAnimation
}
