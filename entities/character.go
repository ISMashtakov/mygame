package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/yohamta/donburi"
)

type CharacterCreator struct {
}

func NewCharacterCreator() *CharacterCreator {
	return &CharacterCreator{}
}

func (c CharacterCreator) Create(world donburi.World) (donburi.Entity, error) {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.MovementRequest,
		components.Character,
		components.WalkingAnimation,
		components.Collider,
	)

	return entity, nil
}
