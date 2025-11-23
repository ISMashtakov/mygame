package resources

import (
	"time"

	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/quasilyte/gmath"
)

type AnimationID int

type animationData struct {
	imageID     ImageID
	startOffset gmath.Vec
	offset      gmath.Vec
	cellSize    gmath.Vec
	frames      int
	directions  []direction.DirectionEnum
	duration    time.Duration
}

var (
	animationResources = map[AnimationID]animationData{
		AnimationCharacterWalking: {
			imageID:     ImageCharacterMoving,
			startOffset: gmath.Vec{X: 80, Y: 60},
			offset:      gmath.Vec{X: 320, Y: 320},
			cellSize:    gmath.Vec{X: 170, Y: 250},
			frames:      4,
			directions:  []direction.DirectionEnum{direction.Down, direction.Left, direction.Right, direction.Up},
			duration:    time.Millisecond * 800,
		},
		AnimationCharacterHoeHitting: {
			imageID:    ImageCharacterHoeHitting,
			offset:     gmath.Vec{X: 250, Y: 170},
			cellSize:   gmath.Vec{X: 250, Y: 170},
			frames:     3,
			directions: []direction.DirectionEnum{direction.Left, direction.Right, direction.Down, direction.Up},
			duration:   time.Millisecond * 1600,
		},
	}
)

const (
	AnimationNone AnimationID = iota
	AnimationCharacterWalking
	AnimationCharacterHoeHitting
)
