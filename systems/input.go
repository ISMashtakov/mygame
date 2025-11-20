package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/quasilyte/gmath"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	Speed = 20
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (m *Input) Update(world donburi.World) error {
	keys := inpututil.AppendPressedKeys(nil)
	var shift gmath.Vec
	if lo.Contains(keys, ebiten.KeyD) {
		shift.X += Speed
	}
	if lo.Contains(keys, ebiten.KeyA) {
		shift.X -= Speed
	}
	if lo.Contains(keys, ebiten.KeyW) {
		shift.Y -= Speed
	}
	if lo.Contains(keys, ebiten.KeyS) {
		shift.Y += Speed
	}

	for en := range donburi.NewQuery(filter.Contains(components.Character)).Iter(world) {
		components.MovementRequest.Set(en, &components.MovementRequestData{Vec: shift})
	}

	return nil
}
