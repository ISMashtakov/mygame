package guicomponents

import (
	"image/color"
	"strconv"

	"github.com/ISMashtakov/mygame/gui/colors"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/samber/lo"
)

type InventoryCell struct {
	root     *widget.Container
	cell     *widget.Container
	text     *widget.Text
	image    *ebiten.Image
	count    int
	selected bool

	onDrag func()
	onDrop func()
}

func ToFont(f *text.GoTextFaceSource) text.Face {
	return &text.GoTextFace{
		Source: f,
		Size:   12,
	}
}

func NewInventoryCell(resourceLoaded resources.IResourceLoader) *InventoryCell {
	res := &InventoryCell{}
	res.root = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout()))
	res.cell = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(20, 20),
			// Drag and drop настройки
			widget.WidgetOpts.EnableDragAndDrop(NewDragAndDropItem(dragData{cell: res})),
			widget.WidgetOpts.CanDrop(func(_ *widget.DragAndDropDroppedEventArgs) bool {
				return true
			}),
			widget.WidgetOpts.Dropped(func(_ *widget.DragAndDropDroppedEventArgs) {
				if res.onDrop != nil {
					res.onDrop()
				}
			}),
			/////////
		))

	res.text = widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
		),
		widget.TextOpts.Text(
			"",
			lo.ToPtr(ToFont(resourceLoaded.LoadFont(resources.FontSimple))),
			colors.InventoryCellText,
		),
	)

	res.root.AddChild(res.cell)
	res.cell.AddChild(res.text)
	res.updateImage()

	return res
}

func (c *InventoryCell) updateImage() {
	border := lo.Ternary(
		c.selected,
		colors.InventoryCellBorderSelected,
		colors.InventoryCellBorder,
	)

	backgroundImage := image.NewAdvancedNineSliceColor(
		colors.InventoryCellBackground,
		image.NewBorder(1, 1, 1, 1, border),
	)

	c.root.SetBackgroundImage(backgroundImage)

	if c.image != nil {
		c.cell.SetBackgroundImage(image.NewNineSliceBorder(c.image, 1))
	} else {
		c.cell.SetBackgroundImage(image.NewNineSliceColor(color.Transparent))
	}

	c.text.Label = lo.Ternary(c.count > 0, strconv.Itoa(c.count), "")
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

func (c *InventoryCell) SetCount(count int) {
	c.count = count
	c.updateImage()
	c.root.RequestRelayout()
}

func (c *InventoryCell) SetOnDrag(fun func()) {
	c.onDrag = fun
}

func (c *InventoryCell) SetOnDrop(fun func()) {
	c.onDrop = fun
}
