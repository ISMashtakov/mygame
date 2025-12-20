package entities

import (
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/utils/don"
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
		gui.DownPanel,
	)

	// Стартовый выбор первой ячейки в нижней панели
	don.CreateRequest(world, gui.SelectCellRequest, &gui.SelectCellRequestData{})
	return entity, nil
}
