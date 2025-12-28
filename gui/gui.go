package gui

import (
	"github.com/ISMashtakov/mygame/gui/guicomponents"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type GUI struct {
	root      *ebitenui.UI
	downPanel *guicomponents.DownPanel
	inventory *guicomponents.Inventory
}

func NewGUI(resourceLoaded resources.IResourceLoader) *GUI {
	g := &GUI{}

	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	g.downPanel = guicomponents.NewDownPanel(resourceLoaded)
	g.inventory = guicomponents.NewInventory(resourceLoaded)

	root.AddChild(g.downPanel.Root())
	root.AddChild(g.inventory.Root())

	g.root = &ebitenui.UI{Container: root}

	return g
}

func (g GUI) Draw(screen *ebiten.Image) {
	g.root.Draw(screen)
}

func (g GUI) Update() {
	g.root.Update()
}

func (g GUI) DownPanel() *guicomponents.DownPanel {
	return g.downPanel
}

func (g GUI) Inventory() *guicomponents.Inventory {
	return g.inventory
}
