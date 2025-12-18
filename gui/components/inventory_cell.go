package components

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/samber/lo"
)

type InventoryCell struct {
	root *widget.Button
}

func NewInventoryCell(opts ...widget.ButtonOpt) *InventoryCell {
	res := &InventoryCell{}
	res.root = widget.NewButton(opts...)
	res.Disable()

	return res
}

func (InventoryCell) getImage(enabled bool) *widget.ButtonImage {
	background := color.RGBA{R: 246, G: 186, B: 114, A: 255}
	border := lo.Ternary(enabled, color.RGBA{R: 170, G: 29, B: 19, A: 255}, color.RGBA{R: 193, G: 135, B: 72, A: 255})
	im := image.NewAdvancedNineSliceColor(background, image.NewBorder(1, 1, 1, 1, border))

	return &widget.ButtonImage{
		Idle:            im,
		Hover:           im,
		Pressed:         im,
		PressedHover:    im,
		Disabled:        im,
		PressedDisabled: im,
	}
}

func (c InventoryCell) Root() *widget.Button {
	return c.root
}

func (c InventoryCell) Enable() {
	c.root.SetImage(c.getImage(true))
}

func (c InventoryCell) Disable() {
	c.root.SetImage(c.getImage(false))
}
