package systems

import (
	"fmt"

	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/gui/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	CellSelectingCodename = "cell_selecting"
)

type CellSelecting struct {
	core.BaseSystem
	downPanel *components.DownPanel
}

func NewCellSelecting(downPanel *components.DownPanel) *CellSelecting {
	return &CellSelecting{
		BaseSystem: core.BaseSystem{
			Codename:        CellSelectingCodename,
			PreviousSystems: []string{InputCodename},
		},
		downPanel: downPanel,
	}
}

func (c *CellSelecting) Update(world donburi.World) error {
	en, ok := donburi.NewQuery(filter.Contains(gui.SelectCellRequest, gui.SelectedCell)).First(world)
	if !ok {
		return nil
	}

	request := gui.SelectCellRequest.Get(en)
	cell := gui.SelectedCell.Get(en)

	if cell.CellNumber >= len(c.downPanel.Cells()) || cell.CellNumber < 0 {
		return fmt.Errorf("invalid value of selected cell %d", cell.CellNumber)
	}

	if request.CellNumber >= len(c.downPanel.Cells()) || request.CellNumber < 0 {
		return fmt.Errorf("invalid value of requested selected cell %d", request.CellNumber)
	}

	c.downPanel.Cells()[cell.CellNumber].Disable()
	c.downPanel.Cells()[request.CellNumber].Enable()

	cell.CellNumber = request.CellNumber

	donburi.Remove[any](en, gui.SelectCellRequest)

	return nil
}
