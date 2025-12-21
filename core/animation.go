package core

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type IAnimation interface {
	Next(frame int)
}

type AnimationPlayer struct {
	currentFrame int
	countFrames  int
	animations   []IAnimation
	onFinish     func()
}

func NewAnimationPlayer(duration time.Duration, onFinish func(), animations ...IAnimation) *AnimationPlayer {
	return &AnimationPlayer{
		countFrames: int(duration.Seconds() * float64(ebiten.TPS())),
		animations:  animations,
		onFinish:    onFinish,
	}
}

func (a *AnimationPlayer) Next() bool {
	for _, animation := range a.animations {
		animation.Next(a.currentFrame)
	}

	a.currentFrame += 1

	isEnd := a.currentFrame >= a.countFrames

	if isEnd && a.onFinish != nil {
		a.onFinish()
	}

	return isEnd
}
