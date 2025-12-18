package entities

import (
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/yohamta/donburi"
)

type InterfaceCreator struct {
}

func NewInterfaceCreator() *InterfaceCreator {
	return &InterfaceCreator{}
}

func (c InterfaceCreator) Create(world donburi.World) (donburi.Entity, error) {
	entity := world.Create(
		gui.SelectedCell,
		gui.SelectCellRequest,
	)

	return entity, nil
}
