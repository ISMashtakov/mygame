package components

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type DownPanel struct {
	root  *widget.Container
	cells []*InventoryCell
}

func NewDownPanel(opts ...widget.ButtonOpt) *DownPanel {
	result := &DownPanel{}

	c := color.RGBA{R: 255, G: 220, B: 152, A: 255}
	b := color.RGBA{R: 216, G: 122, B: 36, A: 255}

	downPanel := widget.NewContainer(
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
		widget.ContainerOpts.BackgroundImage(image.NewBorderedNineSliceColor(c, b, 1)),
	)

	for i := 0; i < 9; i++ {
		cell := NewInventoryCell(
			widget.ButtonOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					MaxWidth:  20,
					MaxHeight: 20,
				}),
			),
		)

		downPanel.AddChild(cell.Root())

		result.cells = append(result.cells, cell)
	}

	result.root = downPanel

	return result
}

func (c DownPanel) Root() *widget.Container {
	return c.root
}

func (c DownPanel) Cells() []*InventoryCell {
	return c.cells
}
