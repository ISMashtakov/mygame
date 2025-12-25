package items

import "github.com/hajimehoshi/ebiten/v2"

type Type int

const (
	Hoe Type = iota + 1
	Pickaxe
	Coal
)

type IItem interface {
	GetType() Type
	GetImage() *ebiten.Image
}

type SimpleItem struct {
	itemType  Type
	image     *ebiten.Image
	propImage *ebiten.Image
}

func NewSimpleItem(itemType Type, image *ebiten.Image) *SimpleItem {
	return &SimpleItem{
		itemType: itemType,
		image:    image,
	}
}

func (s *SimpleItem) GetType() Type {
	return s.itemType
}

func (s *SimpleItem) GetImage() *ebiten.Image {
	if s == nil {
		return nil
	}

	return s.image
}
