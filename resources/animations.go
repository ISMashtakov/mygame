package resources

import (
	"time"

	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/quasilyte/gmath"
)

var CharacterImageSize = gmath.Vec{X: 80, Y: 64}

type AnimationID int

type animationData struct {
	imageID    ImageID
	cellSize   gmath.Vec
	frames     int
	directions []direction.DirectionEnum
	duration   time.Duration
}

var (
	animationResources = map[AnimationID]animationData{
		AnimationCharacterWalking: {
			imageID:    ImageCharacterMoving,
			cellSize:   CharacterImageSize,
			frames:     4,
			directions: []direction.DirectionEnum{direction.Down, direction.Right, direction.Up},
			duration:   time.Millisecond * 600,
		},
		AnimationCharacterHoeHitting: {
			imageID:    ImageCharacterHoeHitting,
			cellSize:   CharacterImageSize,
			frames:     6,
			directions: []direction.DirectionEnum{direction.Down, direction.Right, direction.Up},
			duration:   time.Millisecond * 600,
		},
	}
)

const (
	AnimationNone AnimationID = iota
	AnimationCharacterWalking
	AnimationCharacterHoeHitting
)
