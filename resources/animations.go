package resources

import (
	"time"

	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/quasilyte/gmath"
)

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
			cellSize:   constants.CharacterImageSize,
			frames:     4,
			directions: []direction.DirectionEnum{direction.Down, direction.Right, direction.Up},
			duration:   time.Millisecond * 600,
		},
		AnimationCharacterHoeHitting: {
			imageID:    ImageCharacterHoeHitting,
			cellSize:   constants.CharacterImageSize,
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
