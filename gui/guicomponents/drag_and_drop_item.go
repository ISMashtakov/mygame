package guicomponents

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type dragData struct {
	cell *InventoryCell
}

func (d *dragData) Create(_ widget.HasWidget) (*widget.Container, any) {
	if d.cell.image == nil {
		return nil, nil
	}

	root := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(20, 20),
		),
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceBorder(d.cell.image, 0)),
	)

	if d.cell.onDrag != nil {
		d.cell.onDrag()
	}

	return root, nil
}

func NewDragAndDropItem(data dragData) *widget.DragAndDrop {
	return widget.NewDragAndDrop(
		widget.DragAndDropOpts.ContentsCreater(&data),
	)
}
