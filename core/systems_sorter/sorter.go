package systemssorter

import (
	"fmt"

	"github.com/oko/toposort"
	"github.com/samber/lo"
)

type node struct{ system string }

func newNode(system string) *node { return &node{system: system} }

func (n node) Id() string { return n.system }

func SortSystems[T ISystem](systems []T) ([]T, error) {
	graph := toposort.NewTopology()
	lo.ForEach(systems, func(s T, _ int) { graph.AddNode(newNode(s.GetCodename())) })

	for _, system := range systems {
		for _, next := range system.GetNextSystems() {
			graph.AddEdge(newNode(system.GetCodename()), newNode(next))
		}

		for _, prev := range system.GetPreviousSystems() {
			graph.AddEdge(newNode(prev), newNode(system.GetCodename()))
		}
	}

	resultCodenames, err := graph.Sort()
	if err != nil {
		return nil, fmt.Errorf("system is not sortable: %w", err)
	}

	systemsMap := lo.SliceToMap(systems, func(s T) (string, T) { return s.GetCodename(), s })

	return lo.Map(resultCodenames, func(s toposort.Node, _ int) T { return systemsMap[s.Id()] }), nil
}
