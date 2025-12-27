package components

import (
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/gui/colors"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type DownPanel struct {
	root  *widget.Container
	cells []*InventoryCell
}

func NewDownPanel() *DownPanel {
	result := &DownPanel{}

	result.root = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(2)),
			widget.RowLayoutOpts.Spacing(1),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
				Padding:            widget.NewInsetsSimple(5),
			}),
		),
		widget.ContainerOpts.BackgroundImage(
			image.NewBorderedNineSliceColor(colors.InventoryBackgroud, colors.InventoryBorder, 1),
		),
	)

	for range constants.DownPanelLength {
		cell := NewInventoryCell()

		result.root.AddChild(cell.Root())

		result.cells = append(result.cells, cell)
	}

	return result
}

func (c DownPanel) Root() *widget.Container {
	return c.root
}

func (c DownPanel) Cells() []*InventoryCell {
	return c.cells
}
