package systems

import (
	"fmt"

	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/gui/components"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

const (
	DownPanelHandlerCodename = "down_panel_handler"
)

type DownPanelHandler struct {
	core.BaseSystem
	downPanel *components.DownPanel
}

func NewDownPanelHandler(downPanel *components.DownPanel) *DownPanelHandler {
	return &DownPanelHandler{
		BaseSystem: core.BaseSystem{
			Codename:        DownPanelHandlerCodename,
			PreviousSystems: []string{InputCodename},
		},
		downPanel: downPanel,
	}
}

func (c *DownPanelHandler) Update(world donburi.World) {
	c.updateSelecting(world)
	c.updateItems(world)
}

func (c *DownPanelHandler) updateItems(world donburi.World) {
	panel := don.GetComponent(world, gui.DownPanel)

	for request := range don.IterByRequests(world, gui.SetItemToDownPanelRequest) {
		if request.Index >= len(c.downPanel.Cells()) || request.Index < 0 {
			panic(fmt.Errorf("invalid value of set item cell %d", request.Index))
		}

		panel.Items[request.Index] = request.Item
		c.downPanel.Cells()[request.Index].SetImage(request.Item.GetImage())
	}
}

func (c *DownPanelHandler) updateSelecting(world donburi.World) {
	cell := don.GetComponent(world, gui.SelectedCell)

	for request := range don.IterByRequests(world, gui.SelectCellRequest) {
		if cell.CellNumber >= len(c.downPanel.Cells()) || cell.CellNumber < 0 {
			panic(fmt.Errorf("invalid value of selected cell %d", cell.CellNumber))
		}

		if request.CellNumber >= len(c.downPanel.Cells()) || request.CellNumber < 0 {
			panic(fmt.Errorf("invalid value of requested selected cell %d", request.CellNumber))
		}

		c.downPanel.Cells()[cell.CellNumber].Disable()
		c.downPanel.Cells()[request.CellNumber].Enable()

		cell.CellNumber = request.CellNumber
	}
}
