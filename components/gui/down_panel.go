package gui

import (
	"fmt"

	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

type cellStorage struct {
	size     int
	items    []items.IItem
	location LocationEnum
}

type DownPanelData struct {
	cellStorage
}

func NewDownPanelData() *DownPanelData {
	return &DownPanelData{
		cellStorage{
			size:     9,
			items:    make([]items.IItem, 9),
			location: DownPanelLocation,
		},
	}
}

func (s *cellStorage) SetItem(world donburi.World, index int, item items.IItem) {
	s.validateIndex(index)

	s.items[index] = item

	s.updateGUI(world, index)
}

func (s *cellStorage) updateGUI(world donburi.World, index int) {
	don.Create(world, SetItemToGUIInventoryRequest, &SetItemToGUIInventoryRequestData{
		Location: CellLocation{
			CellNumber: index,
			Location:   s.location,
		},
		Item: s.items[index],
	})
}

func (s *cellStorage) AddItem(world donburi.World, index int, count int) {
	s.validateIndex(index)

	s.items[index].AddCount(count)
	if s.items[index].GetCount() <= 0 {
		s.items[index] = nil
	}

	s.updateGUI(world, index)
}

func (s *cellStorage) GetItem(index int) items.IItem {
	s.validateIndex(index)

	return s.items[index]
}

func (s *cellStorage) GetItems() []items.IItem {
	return s.items[:]
}

func (s *cellStorage) validateIndex(index int) {
	if index >= s.size || index < 0 {
		panic(fmt.Errorf("can't set items to index %d", index))
	}
}

var DownPanel = donburi.NewComponentType[DownPanelData]()
