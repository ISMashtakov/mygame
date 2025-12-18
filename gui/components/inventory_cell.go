package components

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type InventoryCell struct {
	root *widget.Button
}

func NewInventoryCell(opts ...widget.ButtonOpt) *InventoryCell {
	c := color.RGBA{R: 246, G: 186, B: 114, A: 255}
	b := color.RGBA{R: 193, G: 135, B: 72, A: 255}
	return &InventoryCell{
		root: widget.NewButton(
			append(opts,
				widget.ButtonOpts.Image(&widget.ButtonImage{
					Idle:            image.NewAdvancedNineSliceColor(c, image.NewBorder(1, 1, 1, 1, b)),
					Hover:           image.NewAdvancedNineSliceColor(c, image.NewBorder(1, 1, 1, 1, b)),
					Pressed:         image.NewAdvancedNineSliceColor(c, image.NewBorder(1, 1, 1, 1, b)),
					PressedHover:    image.NewAdvancedNineSliceColor(c, image.NewBorder(1, 1, 1, 1, b)),
					Disabled:        image.NewAdvancedNineSliceColor(c, image.NewBorder(1, 1, 1, 1, b)),
					PressedDisabled: image.NewAdvancedNineSliceColor(c, image.NewBorder(1, 1, 1, 1, b)),
				}),
			)...,
		),
	}
}

func (c InventoryCell) Root() *widget.Button {
	return c.root
}
