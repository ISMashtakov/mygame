package animations

import (
	"time"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpritesheetAnimation struct {
	animationMap *images.AnimationMap
	countFrames  int
	sprite       *components.SpriteData
	direction    direction.DirectionEnum
}

func NewSpritesheetAnimation(animationMap *images.AnimationMap, direction direction.DirectionEnum, duration time.Duration, sprite *components.SpriteData) *SpritesheetAnimation {
	return &SpritesheetAnimation{
		animationMap: animationMap,
		countFrames:  int(duration.Seconds() * float64(ebiten.TPS())),
		sprite:       sprite,
		direction:    direction,
	}
}

func (a *SpritesheetAnimation) Next(frame int) {
	imageNumber := int(float64(frame) / float64(a.countFrames) * float64(a.animationMap.GetCountFrames()))

	image := a.animationMap.GetByDirection(a.direction, imageNumber)

	a.sprite.Image = image
}
