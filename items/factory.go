package items

import "github.com/ISMashtakov/mygame/resources"

type ItemsFactory struct {
	loader resources.IResourceLoader
}

func NewItemsFactory(loader resources.IResourceLoader) *ItemsFactory {
	return &ItemsFactory{
		loader: loader,
	}
}

func (f ItemsFactory) Hoe() IItem {
	return NewSimpleItem(Hoe, f.loader.LoadImage(resources.ImageItemHoe))
}

func (f ItemsFactory) Pickaxe() IItem {
	return NewSimpleItem(Pickaxe, f.loader.LoadImage(resources.ImageItemPickaxe))
}
