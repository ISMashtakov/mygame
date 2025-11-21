package systemssorter_test

import (
	"testing"

	"github.com/ISMashtakov/mygame/core"
	systemssorter "github.com/ISMashtakov/mygame/core/systems_sorter"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSortSystems(t *testing.T) {
	t.Parallel()

	A := core.BaseSystem{Codename: "A"}
	A2B := core.BaseSystem{Codename: "A", NextSystems: []string{"B"}}
	B2A := core.BaseSystem{Codename: "B", PreviousSystems: []string{"A"}}
	B := core.BaseSystem{Codename: "B"}
	C := core.BaseSystem{Codename: "C"}
	D2A2B := core.BaseSystem{Codename: "D", PreviousSystems: []string{"A"}, NextSystems: []string{"B"}}

	tests := []struct {
		name    string
		systems []core.BaseSystem
		want    []core.BaseSystem
		wantErr bool
	}{
		{
			name:    "A -> B",
			systems: []core.BaseSystem{B, A2B},
			want:    []core.BaseSystem{A2B, B},
		},
		{
			name:    "A <-> B",
			systems: []core.BaseSystem{B2A, A2B},
			want:    []core.BaseSystem{A2B, B2A},
		},
		{
			name:    "A <-> B правильно",
			systems: []core.BaseSystem{A2B, B2A},
			want:    []core.BaseSystem{A2B, B2A},
		},
		{
			name:    "A <-> B C",
			systems: []core.BaseSystem{C, B2A, A2B},
			want:    []core.BaseSystem{A2B, B2A, C},
		},
		{
			name:    "A <- D -> B",
			systems: []core.BaseSystem{B, A, D2A2B},
			want:    []core.BaseSystem{A, D2A2B, B},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := systemssorter.SortSystems(tt.systems)
			assert.Equal(t, tt.want, got)
			lo.Ternary(tt.wantErr, assert.Error, assert.NoError)(t, err)
		})
	}
}
