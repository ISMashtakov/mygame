package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/yohamta/donburi"
)

type CharacterCreator struct {
	world donburi.World
}

func NewCharacterCreator(world donburi.World) *CharacterCreator {
	return &CharacterCreator{
		world: world,
	}
}

func (c CharacterCreator) Create() (donburi.Entity, error) {
	entity := c.world.Create(components.Position, components.Sprite, components.MovementRequest, components.Character, components.WalkingAnimation)

	return entity, nil
}
