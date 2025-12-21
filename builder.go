package main

import (
	"fmt"

	guicomponents "github.com/ISMashtakov/mygame/components/gui"
	systemssorter "github.com/ISMashtakov/mygame/core/systems_sorter"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/game"
	"github.com/ISMashtakov/mygame/utils/don"

	"github.com/ISMashtakov/mygame/gui"
	"github.com/ISMashtakov/mygame/items"
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
		character    *entities.CharacterCreator
		garden       *background.GardenCreator
		simpleSprite *entities.SimpeSpriteCreator
	}
	gui *gui.GUI

	itemsFactory *items.ItemsFactory
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
	b.gui = gui.NewGUI()
}

func (b *Builder) Entities() {
	// Entity creatores
	b.creators.character = entities.NewCharacterCreator()
	b.creators.garden = background.NewGardenCreator(b.resourses)
	b.creators.simpleSprite = entities.NewSimpeSpriteCreator()
	grassCreator := background.NewGrassCreator(b.resourses)
	coilCreator := entities.NewCoilCreator(b.resourses)
	interfaceCreator := entities.NewInterfaceCreator()
	cameraCreator := entities.NewCameraCreator()

	// ----------
	worldBuilder := game.NewWorldBuilder(*grassCreator, *coilCreator)

	err := worldBuilder.Build(b.world)

	_ = b.creators.character.Create(b.world)

	_ = cameraCreator.Create(b.world)

	_, err = interfaceCreator.Create(b.world)
	if err != nil {
		panic(fmt.Errorf("can't create interface: %w", err))
	}

	don.CreateRequest(b.world, guicomponents.SetItemToDownPanelRequest, &guicomponents.SetItemToDownPanelRequestData{
		Index: 0,
		Item:  b.itemsFactory.Hoe(),
	})

	don.CreateRequest(b.world, guicomponents.SetItemToDownPanelRequest, &guicomponents.SetItemToDownPanelRequestData{
		Index: 1,
		Item:  b.itemsFactory.Pickaxe(),
	})
}

func (b *Builder) Systems() {
	walkingAnimationSystem := systems.NewSwapSpriteByAnimation(b.resourses, b.creators.character)

	b.systems = []ISystem{
		systems.NewInput(),
		walkingAnimationSystem,
		systems.NewCollisionDetector(),
		systems.NewMovement(),
		systems.NewHoeHitChecker(*b.creators.garden, *b.creators.simpleSprite),
		systems.NewDownPanelHandler(b.gui.DownPanel()),
		systems.NewCameraMoving(),
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

	fmt.Println(ebiten.RunGame(game))
}
