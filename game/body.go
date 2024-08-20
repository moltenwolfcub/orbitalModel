package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Body struct {
	game *Game
	id   int

	pos   image.Point
	delta image.Point

	color color.RGBA

	image      *ebiten.Image
	imageDirty bool
}

func NewBody(game *Game, id int) *Body {
	b := &Body{
		game:  game,
		id:    id,
		pos:   image.Pt(100, 100),
		delta: image.Pt(0, 0),
		color: color.RGBA{0, 255, 255, 255},
	}
	b.genImage()

	return b
}

func (b *Body) genImage() {
	radius := 50
	img := ebiten.NewImage(radius*2, radius*2)

	vector.DrawFilledCircle(img, float32(radius), float32(radius), float32(radius), b.color, false)

	b.image = img

	b.imageDirty = false
}

func (b Body) Draw(screen *ebiten.Image) {
	if b.imageDirty {
		b.genImage()
	}

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.pos.X), float64(b.pos.Y))
	options.GeoM.Translate(-50, -50) //centre circle by subtracting radius

	screen.DrawImage(b.image, &options)
}
