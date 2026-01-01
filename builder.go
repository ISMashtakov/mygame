package main

import (
	"fmt"
	"log/slog"

	systemssorter "github.com/ISMashtakov/mygame/core/systems_sorter"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/game"

	"github.com/ISMashtakov/mygame/gui"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/systems"
	"github.com/ISMashtakov/mygame/systems/actions"
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
		character    *entities.CharacterCreator
		garden       *background.GardenCreator
		simpleSprite *entities.SimpeSpriteCreator
		props        *entities.PropsCreator
		plants       *entities.PlantCreator
	}
	gui *gui.GUI

	itemsFactory *items.Factory
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

func (b *Builder) ItemsFactory() {
	b.itemsFactory = items.NewItemsFactory(b.resourses)
}

func (b *Builder) GUI() {
	b.gui = gui.NewGUI(b.resourses)
}

func (b *Builder) Entities() {
	// Entity creatores
	b.creators.character = entities.NewCharacterCreator()
	b.creators.garden = background.NewGardenCreator(b.resourses)
	b.creators.simpleSprite = entities.NewSimpeSpriteCreator()
	b.creators.props = entities.NewPropsCreator()
	b.creators.plants = entities.NewPlantCreator(b.resourses)
	grassCreator := background.NewGrassCreator(b.resourses)
	coalCreator := entities.NewCoalCreator(b.resourses, *b.itemsFactory)
	interfaceCreator := entities.NewInterfaceCreator(b.itemsFactory, b.gui.Inventory(), b.gui.DownPanel())
	cameraCreator := entities.NewCameraCreator()

	// ----------
	worldBuilder := game.NewWorldBuilder(*grassCreator, *coalCreator)

	err := worldBuilder.Build(b.world)
	if err != nil {
		panic(fmt.Errorf("can't build world: %w", err))
	}

	_ = b.creators.character.Create(b.world)

	_ = cameraCreator.Create(b.world)

	_ = interfaceCreator.Create(b.world)
}

func (b *Builder) Systems() {
	walkingAnimationSystem := systems.NewAnimation(b.creators.character)

	b.systems = []ISystem{
		systems.NewInput(b.resourses),
		walkingAnimationSystem,
		systems.NewCollisionDetector(),
		systems.NewMovement(),
		actions.NewPickaxeHitRequestHandler(*b.creators.simpleSprite, *b.creators.props),
		actions.NewGardenCreatingRequestHandler(*b.creators.garden),
		actions.NewSeedHandler(b.creators.plants),
		systems.NewPlants(b.creators.garden),
		systems.NewCameraMoving(),
		systems.NewPropsTaking(b.gui.Inventory(), b.gui.DownPanel()),
		systems.NewInventory(b.gui.Inventory(), b.gui.DownPanel()),
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
		b.gui,
	)

	if err := ebiten.RunGame(game); err != nil {
		slog.Error(err.Error())
	}
}
