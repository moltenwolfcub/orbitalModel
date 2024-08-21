package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type HasHitbox interface {
	// Most objects run a direct call to DefaultHitboxOverlaps.
	// This is basically here incase of non-rectangle hitboxes
	// like circles
	Overlaps([]image.Rectangle) bool
	Origin() image.Point
	Size() image.Point
	GetHitbox() []image.Rectangle
}

func DefaultHitboxOverlaps(mine HasHitbox, other []image.Rectangle) bool {
	for _, otherSub := range other {
		myHitbox := mine.GetHitbox()

		for _, mySub := range myHitbox {
			if otherSub.Overlaps(mySub) {
				return true
			}
		}
	}
	return false
}

type Drawable interface {
	HasHitbox

	DrawAt(*ebiten.Image, image.Point)
}

type Viewport struct {
	rect image.Rectangle
}

func NewViewport() Viewport {
	return Viewport{rect: image.Rectangle{
		Min: image.Point{
			X: -WindowWidth / 2,
			Y: -WindowHeight / 2,
		}, Max: image.Point{
			X: WindowWidth / 2,
			Y: WindowHeight / 2,
		},
	}}
}

func (v Viewport) Overlaps(other []image.Rectangle) bool {
	return DefaultHitboxOverlaps(v, other)
}
func (v Viewport) Origin() image.Point {
	return v.rect.Min
}
func (v Viewport) Size() image.Point {
	return v.rect.Size()
}
func (v Viewport) GetHitbox() []image.Rectangle {
	return []image.Rectangle{v.rect}
}

func (v Viewport) objectInViewport(object HasHitbox) bool {
	hitbox := v.GetHitbox()

	return object.Overlaps(hitbox)
}

func (v Viewport) DrawToScreen(screen *ebiten.Image, drawable Drawable) {
	inViewport := v.objectInViewport(drawable)
	origin := drawable.Origin()

	if !inViewport {
		return
	}
	offsetPos := origin.Sub(v.rect.Min)
	drawable.DrawAt(screen, offsetPos)
}
