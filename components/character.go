package components

import (
	"github.com/ISMashtakov/mygame/core/direction"
	"github.com/yohamta/donburi"
)

type CharacterData struct{}

var Character = donburi.NewComponentType[CharacterData]()

type WalkingAnimationData struct {
	Frame     int
	Direction direction.Direction
}

var WalkingAnimation = donburi.NewComponentType[WalkingAnimationData]()
