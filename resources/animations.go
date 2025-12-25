package resources

import (
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/quasilyte/gmath"
)

var CharacterImageSize = gmath.Vec{X: 80, Y: 64}

type AnimationID int

type animationData struct {
	imageID    ImageID
	cellSize   gmath.Vec
	frames     int
	directions []direction.Enum
}

var (
	animationResources = map[AnimationID]animationData{
		AnimationCharacterWalking: {
			imageID:    ImageCharacterMoving,
			cellSize:   CharacterImageSize,
			frames:     4,
			directions: []direction.Enum{direction.Down, direction.Right, direction.Up},
		},
		AnimationCharacterHoeHitting: {
			imageID:    ImageCharacterHoeHitting,
			cellSize:   CharacterImageSize,
			frames:     6,
			directions: []direction.Enum{direction.Down, direction.Right, direction.Up},
		},
	}
)

const (
	AnimationCharacterWalking AnimationID = iota
	AnimationCharacterHoeHitting
)
