package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WindowWidth  = 1920
	WindowHeight = 1080
	TPS          = 60
)

type Game struct {
	bodies []*Body
}

func NewGame() *Game {
	g := &Game{}

	g.bodies = []*Body{
		NewBody(g, 0),
	}

	return g
}

func (g *Game) Update() (err error) {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	for _, b := range g.bodies {
		b.Draw(screen)
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
