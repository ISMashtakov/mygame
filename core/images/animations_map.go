package images

import (
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/hajimehoshi/ebiten/v2"
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

func (m AnimationMap) GetByDirection(dir direction.DirectionEnum, i int) *ebiten.Image {
	return m.spritesSheet.Get(i, lo.IndexOf(m.directions, dir))
}

func (m AnimationMap) GetCountFrames() int {
	return m.framesInAnimation
}
