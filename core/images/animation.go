package images

import (
	"time"

	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	animationMap AnimationMap
	currentFrame int
	countFrames  int
}

func NewAnimation(animationMap AnimationMap, duration time.Duration) *Animation {
	return &Animation{
		animationMap: animationMap,
		currentFrame: 0,
		countFrames:  int(duration.Seconds() * float64(ebiten.TPS())),
	}
}

func (a *Animation) Next(dir direction.DirectionEnum) Image {
	if a.currentFrame >= a.countFrames {
		a.currentFrame = 0
	}

	frame := int(float64(a.currentFrame) / float64(a.countFrames) * float64(a.animationMap.GetCountFrames()))

	a.currentFrame += 1

	return a.animationMap.GetByDirection(dir, frame)
}

func (a *Animation) IsFinish() bool {
	return a.currentFrame == a.countFrames
}

func (a *Animation) Reset() {
	a.currentFrame = 0
}
