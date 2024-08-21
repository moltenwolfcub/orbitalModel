package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	WindowWidth  = 1920
	WindowHeight = 1080
	TPS          = 60
)

type Game struct {
	view         Viewport
	lastMousePos image.Point

	bodies []*Body
}

func NewGame() *Game {
	g := &Game{
		view: NewViewport(),
	}

	g.bodies = []*Body{
		NewBody(g, 0),
	}

	return g
}

func (g *Game) Update() (err error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1) {
		g.lastMousePos = image.Pt(ebiten.CursorPosition())
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton1) {
		thisPos := image.Pt(ebiten.CursorPosition())
		change := g.lastMousePos.Sub(thisPos)

		g.view.rect = g.view.rect.Add(change)

		g.lastMousePos = thisPos
	}
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	for _, b := range g.bodies {
		g.view.DrawToScreen(screen, b)
	}
}

func (g Game) Layout(actualWidth, actualHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}

func (g *Game) Run() error {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Orbital Mechanics Simulation")
	// ebiten.SetWindowIcon([]image.Image{})
	ebiten.SetTPS(TPS)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	return ebiten.RunGame(g)
}
