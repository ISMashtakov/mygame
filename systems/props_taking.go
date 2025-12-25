package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/subsystems"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	PropsTakingCodename = "PropsTakingCodename"
)

type PropsTaking struct {
	core.BaseSystem
	subsystems.ColliderSearcher
}

func NewPropsTaking() *PropsTaking {
	return &PropsTaking{
		BaseSystem: core.BaseSystem{
			Codename: PropsTakingCodename,
		},
	}
}

func (s PropsTaking) Update(world donburi.World) {
	charEntry, ok := donburi.NewQuery(filter.Contains(components.Character)).First(world)
	if !ok {
		return
	}

	for _, propEntry := range s.ColliderSearcher.SearchByEntry(world, charEntry, filter.Contains(components.Prop, components.RectCollider)) {
		propEntry.Remove()
	}
}
