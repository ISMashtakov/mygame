package systems

import (
	"log"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/quasilyte/gmath"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	Speed         = 2
	InputCodename = "input"
)

type Input struct {
	core.BaseSystem
}

func NewInput() *Input {
	return &Input{
		core.BaseSystem{
			Codename:    InputCodename,
			NextSystems: []string{CollissionDetectorCodename, MovementCodename},
		},
	}
}

func (m *Input) Update(world donburi.World) error {
	en, ok := donburi.NewQuery(filter.Contains(components.Character)).First(world)
	if !ok {
		log.Println("can't found character")
		return nil
	}

	if en.HasComponent(actions.Action) {
		return nil
	}

	keys := inpututil.AppendPressedKeys(nil)

	if lo.Contains(keys, ebiten.KeySpace) {
		donburi.Add(en, actions.Action, &actions.HoeHit)
	}

	m.processMoving(en, keys)

	return nil
}

func (m *Input) processMoving(en *donburi.Entry, keys []ebiten.Key) {
	var shift gmath.Vec
	if lo.Contains(keys, ebiten.KeyD) {
		shift.X += 1
		direction.Direction.SetValue(en, direction.Right)
	}
	if lo.Contains(keys, ebiten.KeyA) {
		shift.X -= 1
		direction.Direction.SetValue(en, direction.Left)
	}
	if lo.Contains(keys, ebiten.KeyW) {
		shift.Y -= 1
		direction.Direction.SetValue(en, direction.Up)
	}
	if lo.Contains(keys, ebiten.KeyS) {
		shift.Y += 1
		direction.Direction.SetValue(en, direction.Down)
	}

	if !shift.IsZero() {
		donburi.Add(en, components.MovementRequest, &components.MovementRequestData{Vec: shift})
	}
}
