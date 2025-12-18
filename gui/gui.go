package gui

import (
	"bytes"

	"github.com/ISMashtakov/mygame/gui/components"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type GUI struct {
	root *ebitenui.UI
}

func DefaultFont() text.Face {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		panic(err)
	}
	return &text.GoTextFace{
		Source: s,
		Size:   20,
	}
}

func NewGUI() *GUI {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	downPanel := components.NewDownPanel()

	root.AddChild(downPanel.Root())

	return &GUI{
		root: &ebitenui.UI{Container: root},
	}
}

func (g GUI) Draw(screen *ebiten.Image) {
	g.root.Draw(screen)
}

func (g GUI) Update() {
	g.root.Update()
}
