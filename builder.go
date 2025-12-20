package main

import (
	"fmt"

	systemssorter "github.com/ISMashtakov/mygame/core/systems_sorter"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/game"
	"github.com/ISMashtakov/mygame/gui"
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
		garden    *background.GardenCreator
	}
	gui *gui.GUI
}

func (b *Builder) Debug() {
	b.renderer.DrawColliders = true
}

func (b *Builder) Resources() {
	b.resourses = resources.NewResourceLoader()
	if err := b.resourses.Preload(); err != nil {
		panic(fmt.Errorf("can't preload resourses: %w", err))
	}
}

func (b *Builder) Renderer() {
	b.renderer = game.NewRenderer()
}

func (b *Builder) GUI() {
	b.gui = gui.NewGUI()
}

func (b *Builder) Entities() {
	// Entity creatores
	b.creators.character = entities.NewCharacterCreator()
	b.creators.garden = background.NewGardenCreator(b.resourses)
	grassCreator := background.NewGrassCreator(b.resourses)
	stoneCreator := entities.NewStoneCreator(b.resourses)
	interfaceCreator := entities.NewInterfaceCreator()

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

	_, err = interfaceCreator.Create(b.world)
	if err != nil {
		panic(fmt.Errorf("can't create interface: %w", err))
	}
}

func (b *Builder) Systems() {
	walkingAnimationSystem, err := systems.NewSwapSpriteByAnimation(b.resourses, b.creators.character)
	if err != nil {
		panic(fmt.Errorf("can't create walking system: %w", err))
	}

	b.systems = []ISystem{
		systems.NewInput(),
		walkingAnimationSystem,
		systems.NewCollisionDetector(),
		systems.NewMovement(),
		systems.NewHoeHitChecker(*b.creators.garden),
		systems.NewCellSelecting(b.gui.DownPanel()),
	}

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
		b.gui,
	)

	panic(ebiten.RunGame(game))
}
