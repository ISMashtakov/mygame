package guicomponents

import (
	"github.com/ISMashtakov/mygame/gui/colors"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/samber/lo"
)

type Inventory struct {
	root     *widget.Container
	cells    []*InventoryCell
	size     int
	disabled bool
}

func NewInventory(resourceLoaded resources.IResourceLoader) *Inventory {
	result := &Inventory{
		size:     4,
		root:     widget.NewContainer(),
		disabled: true,
	}

	result.root = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(result.size),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(3)),
			widget.GridLayoutOpts.Spacing(2, 2),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				Padding:            widget.NewInsetsSimple(5),
			}),
		),
		widget.ContainerOpts.BackgroundImage(
			image.NewBorderedNineSliceColor(colors.InventoryBackgroud, colors.InventoryBorder, 1),
		),
	)

	for range result.size * result.size {
		cell := NewInventoryCell(resourceLoaded)

		result.root.AddChild(cell.Root())

		result.cells = append(result.cells, cell)
	}

	result.update()

	return result
}

func (c *Inventory) update() {
	c.root.GetWidget().Visibility = lo.Ternary(c.disabled, widget.Visibility_Hide_Blocking, widget.Visibility_Show)
}

func (c *Inventory) Root() *widget.Container {
	return c.root
}

func (c *Inventory) Cells() []*InventoryCell {
	return c.cells
}

func (c *Inventory) Disable() {
	c.disabled = true
	c.update()
}

func (c *Inventory) Enable() {
	c.disabled = false
	c.update()
}

func (c *Inventory) Switch() {
	lo.Ternary(c.disabled, c.Enable, c.Disable)()
}
