package main

import (
	"fmt"

	systemssorter "github.com/ISMashtakov/mygame/core/systems_sorter"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/game"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
)

type ISystem interface {
	game.ISystem
	systemssorter.ISystem
}

type Builder struct {
	resourses *resources.ResourceLoader
	renderer  *game.Renderer
	systems   []ISystem
	world     donburi.World
	creators  struct {
		character *entities.CharacterCreator
	}
}

func (b *Builder) Resources() {
	b.resourses = resources.NewResourceLoader()
	b.resourses.Preload()
}

func (b *Builder) Renderer() {
	b.renderer = game.NewRenderer()
	// renderer.DrawColliders = true
}

func (b *Builder) Entities() {
	// Entity creatores
	b.creators.character = entities.NewCharacterCreator()
	grassCreator := background.NewGrassCreator(b.resourses)
	stoneCreator := entities.NewStoneCreator(b.resourses)

	// ----------
	worldBuilder := game.NewWorldBuilder(*grassCreator, *stoneCreator)

	err := worldBuilder.Build(b.world)
	if err != nil {
		panic(fmt.Errorf("can't build world: %w", err))
	}

	_, err = b.creators.character.Create(b.world)
	if err != nil {
		panic(fmt.Errorf("can't create character: %w", err))
	}
}

func (b *Builder) Systems() {
	b.systems = []ISystem{
		systems.NewInput(),
		systems.NewSwapSpriteByWalkingAnimation(60, b.resourses, b.creators.character),
		systems.NewCollisionDetector(),
		systems.NewMovement(),
	}

	var err error
	b.systems, err = systemssorter.SortSystems(b.systems)
	if err != nil {
		panic(fmt.Errorf("can't sort systems: %w", err))
	}
}

func (b *Builder) World() {
	b.world = donburi.NewWorld()
}

func (b *Builder) RunGame() {
	game := game.NewGame(
		*b.renderer,
		b.world,
		lo.Map(b.systems, func(s ISystem, _ int) game.ISystem { return s }),
	)

	panic(ebiten.RunGame(game))
}
