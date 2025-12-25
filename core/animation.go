package core

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationPlayerOpt func(*AnimationPlayer)

func WithOnFinish(f func()) AnimationPlayerOpt {
	return func(ap *AnimationPlayer) {
		ap.onFinish = f
	}
}

func WithAnimations(animations ...IAnimation) AnimationPlayerOpt {
	return func(ap *AnimationPlayer) {
		ap.animations = append(ap.animations, animations...)
	}
}

func WithRepetable(value bool) AnimationPlayerOpt {
	return func(ap *AnimationPlayer) {
		ap.repeatable = value
	}
}

type IAnimation interface {
	Next(frame int)
}

type AnimationPlayer struct {
	currentFrame int
	countFrames  int
	animations   []IAnimation
	onFinish     func()
	repeatable   bool
}

func NewAnimationPlayer(duration time.Duration, opts ...AnimationPlayerOpt) *AnimationPlayer {
	player := &AnimationPlayer{
		countFrames: int(duration.Seconds() * float64(ebiten.TPS())),
	}

	for _, opt := range opts {
		opt(player)
	}

	return player
}

func (a *AnimationPlayer) Next() bool {
	for _, animation := range a.animations {
		animation.Next(a.currentFrame)
	}

	a.currentFrame += 1

	isEnd := a.currentFrame >= a.countFrames

	if isEnd && a.repeatable {
		a.currentFrame = 0
		isEnd = false
	}

	if isEnd && a.onFinish != nil {
		a.onFinish()
	}

	return isEnd
}
