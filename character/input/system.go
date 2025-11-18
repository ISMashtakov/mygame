package input

import (
	"github.com/ISMashtakov/mygame/character"
	"github.com/ISMashtakov/mygame/physics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/ark/ecs"
	"github.com/quasilyte/gmath"
	"github.com/samber/lo"
)

const (
	Speed = 10
)

type Input struct {
	characterFilter *ecs.Filter2[character.Character, physics.Speed]
}

func NewInput(world *ecs.World) *Input {
	return &Input{
		characterFilter: ecs.NewFilter2[character.Character, physics.Speed](world),
	}
}

func (m *Input) Update() error {
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

	query := m.characterFilter.Query()
	for query.Next() {
		_, speed := query.Get()
		speed.Vec = speed.Vec.Add(shift)
	}

	return nil
}
