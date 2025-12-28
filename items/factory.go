package items

import "github.com/ISMashtakov/mygame/resources"

type Factory struct {
	loader resources.IResourceLoader
}

func NewItemsFactory(loader resources.IResourceLoader) *Factory {
	return &Factory{
		loader: loader,
	}
}

func (f Factory) Hoe() IItem {
	return NewSimpleItem(Hoe, f.loader.LoadImage(resources.ImageItemHoe))
}

func (f Factory) Pickaxe() IItem {
	return NewSimpleItem(Pickaxe, f.loader.LoadImage(resources.ImageItemPickaxe))
}

func (f Factory) Coal() IItem {
	return NewSimpleItem(Coal, f.loader.LoadImage(resources.ImageItemCoal), WithMaxStackSize(3))
}
