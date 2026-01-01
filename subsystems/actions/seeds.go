package actions

import (
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Seed struct {
}

func NewSeed(resourcesLoader resources.IResourceLoader) *Seed {
	return &Seed{}
}

func (p *Seed) ProcessAction(world donburi.World, characterEntity *donburi.Entry) bool {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		return false
	}

	don.Create(world, actions.SeedRequest, &actions.SeedRequestData{
		Point: getCursorPos(world),
		Item:  getSelectedItem(world),
	})

	return true
}
