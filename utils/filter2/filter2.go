package filter2

import (
	"github.com/samber/lo"
	"github.com/yohamta/donburi/component"
	"github.com/yohamta/donburi/filter"
)

type containsAny struct {
	components []component.IComponentType
}

// ContainsAny matches layouts that contains all the components specified.
func ContainsAny(components ...component.IComponentType) filter.LayoutFilter {
	return &containsAny{components: components}
}

func (f *containsAny) MatchesLayout(components []component.IComponentType) bool {
	for _, componentType := range f.components {
		if lo.Contains(components, componentType) {
			return true
		}
	}

	return false
}
