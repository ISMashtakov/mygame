package systemssorter

import (
	"fmt"

	"github.com/oko/toposort"
	"github.com/samber/lo"
)

type node struct{ system string }

func newNode(system string) *node { return &node{system: system} }

func (n node) Id() string { return n.system } //nolint:revive,staticcheck // Специально для имплементации интерфейса для либы

func SortSystems[T ISystem](systems []T) ([]T, error) {
	graph := toposort.NewTopology()
	for _, s := range systems {
		if err := graph.AddNode(newNode(s.GetCodename())); err != nil {
			return nil, err
		}
	}

	for _, system := range systems {
		for _, next := range system.GetNextSystems() {
			if err := graph.AddEdge(newNode(system.GetCodename()), newNode(next)); err != nil {
				return nil, err
			}
		}

		for _, prev := range system.GetPreviousSystems() {
			if err := graph.AddEdge(newNode(prev), newNode(system.GetCodename())); err != nil {
				return nil, err
			}
		}
	}

	resultCodenames, err := graph.Sort()
	if err != nil {
		return nil, fmt.Errorf("system is not sortable: %w", err)
	}

	systemsMap := lo.SliceToMap(systems, func(s T) (string, T) { return s.GetCodename(), s })

	return lo.Map(resultCodenames, func(s toposort.Node, _ int) T { return systemsMap[s.Id()] }), nil
}
