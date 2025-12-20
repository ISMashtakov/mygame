package components

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

type InventoryCell struct {
	root     *widget.Container
	cell     *widget.Container
	image    *ebiten.Image
	selected bool
}

func NewInventoryCell() *InventoryCell {
	res := &InventoryCell{}
	res.root = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout()))
	res.cell = widget.NewContainer(widget.ContainerOpts.WidgetOpts(
		widget.WidgetOpts.MinSize(20, 20),
	))
	res.root.AddChild(res.cell)

	res.updateImage()

	return res
}

func (c *InventoryCell) updateImage() {
	background := color.RGBA{R: 246, G: 186, B: 114, A: 255}
	border := lo.Ternary(c.selected, color.RGBA{R: 170, G: 29, B: 19, A: 255}, color.RGBA{R: 193, G: 135, B: 72, A: 255})

	backgroundImage := image.NewAdvancedNineSliceColor(background, image.NewBorder(1, 1, 1, 1, border))

	c.root.SetBackgroundImage(backgroundImage)

	if c.image != nil {
		c.cell.SetBackgroundImage(image.NewNineSliceBorder(c.image, 1))
	} else {
		c.cell.SetBackgroundImage(image.NewNineSliceColor(color.Transparent))
	}
}

func (c *InventoryCell) Root() *widget.Container {
	return c.root
}

func (c *InventoryCell) Enable() {
	c.selected = true
	c.updateImage()
}

func (c *InventoryCell) Disable() {
	c.selected = false
	c.updateImage()
}

func (c *InventoryCell) SetImage(image *ebiten.Image) {
	c.image = image
	c.updateImage()
}
