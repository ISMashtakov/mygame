package don

import (
	"fmt"
	"iter"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/component"
	"github.com/yohamta/donburi/filter"
)

type iComponent[T any] interface {
	component.IComponentType
	Get(entry *donburi.Entry) *T
}

func CreateRequest[T any](world donburi.World, component component.IComponentType, requestData *T) {
	entity := world.Create(component)
	entry := world.Entry(entity)

	donburi.Add(entry, component, requestData)
}

func DeleteRequest(world donburi.World, entry *donburi.Entry) {
	world.Remove(entry.Entity())
}

func GetComponent[T any](world donburi.World, comp iComponent[T]) *T {
	entry, ok := donburi.NewQuery(filter.Contains(comp)).First(world)
	if !ok {
		panic(fmt.Errorf("can't find component %+v", comp))
	}

	return comp.Get(entry)
}

func IterByRequests[T any](world donburi.World, comp iComponent[T]) iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for reqEn := range donburi.NewQuery(filter.Contains(comp)).Iter(world) {
			res := comp.Get(reqEn)
			DeleteRequest(world, reqEn)
			if !yield(res) {
				return
			}
		}
	}
}
