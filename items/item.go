package items

import "github.com/hajimehoshi/ebiten/v2"

type Type int

const (
	Hoe Type = iota + 1
	Pickaxe
	Coal

	// Seeds

	CarrotSeed
)

type IItem interface {
	GetType() Type
	GetImage() *ebiten.Image
	GetMaxStackSize() int
	GetCount() int
	AddCount(int)
}

type SimpleItemOpt func(*SimpleItem)

func WithMaxStackSize(size int) SimpleItemOpt {
	return func(si *SimpleItem) {
		si.maxStackSize = size
	}
}

type SimpleItem struct {
	itemType     Type
	image        *ebiten.Image
	maxStackSize int
	count        int
}

func NewSimpleItem(itemType Type, image *ebiten.Image, opts ...SimpleItemOpt) *SimpleItem {
	res := &SimpleItem{
		itemType:     itemType,
		image:        image,
		maxStackSize: 1,
		count:        1,
	}

	for _, opt := range opts {
		opt(res)
	}

	return res
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

func (s *SimpleItem) GetMaxStackSize() int {
	return s.maxStackSize
}

func (s *SimpleItem) GetCount() int {
	return s.count
}

func (s *SimpleItem) AddCount(add int) {
	s.count += add
}
