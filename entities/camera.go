package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/yohamta/donburi"
)

type CameraCreator struct{}

func NewCameraCreator() *CameraCreator {
	return &CameraCreator{}
}

func (c CameraCreator) Create(world donburi.World) donburi.Entity {
	entity := world.Create(
		components.Position,
		components.Camera,
	)

	return entity
}
