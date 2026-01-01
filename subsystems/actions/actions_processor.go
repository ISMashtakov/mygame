package actions

import (
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/yohamta/donburi"
)

type IProccesor interface {
	ProcessAction(world donburi.World, characterEntity *donburi.Entry) bool
}

type ActionsProcessor struct {
	processors map[items.Type]IProccesor
}

func NewActionProcessor(resourcesLoader resources.IResourceLoader) *ActionsProcessor {
	return &ActionsProcessor{
		processors: map[items.Type]IProccesor{
			items.Hoe:        NewHoe(resourcesLoader),
			items.Pickaxe:    NewPickaxe(resourcesLoader),
			items.CarrotSeed: NewSeed(resourcesLoader),
		},
	}
}

func (p *ActionsProcessor) Process(characterEntity *donburi.Entry) bool {
	item := getSelectedItem(characterEntity.World)
	processor, ok := p.processors[item.GetType()]
	if !ok {
		return false
	}

	return processor.ProcessAction(characterEntity.World, characterEntity)
}
