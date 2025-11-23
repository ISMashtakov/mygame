package components

import (
	"github.com/yohamta/donburi"
)

type CharacterData struct{}

var Character = donburi.NewComponentType[CharacterData]()

type WalkingAnimationData struct{}

var WalkingAnimation = donburi.NewComponentType[WalkingAnimationData]()
