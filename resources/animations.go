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
	}
)

const (
	AnimationNone AnimationID = iota
	AnimationCharacterWalking
)
