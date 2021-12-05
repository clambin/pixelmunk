package cup

import (
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"image/color"
)

type Cup struct {
	pixelmunk.Object
	Direction float64
}

func NewCup(position vect.Vect, width, height vect.Float, color color.Color) (cup *Cup) {
	return &Cup{
		Object: pixelmunk.NewObject(
			makeCupBody(position, width, height),
			pixelmunk.ObjectOptions{Color: color},
		),
		Direction: 1.0,
	}
}

func (cup *Cup) SetDirection(direction float64) {
	cup.Direction = direction
}

func (cup *Cup) Move() {
	pos := cup.GetBody().Position()
	pos.X += vect.Float(cup.Direction)
	cup.GetBody().SetPosition(pos)
}

func makeCupBody(position vect.Vect, width, height vect.Float) (body *chipmunk.Body) {
	body = chipmunk.NewBody(1e12, 1e12)
	body.IgnoreGravity = true
	//body.SetVelocity(15, 0)
	boxes := getCupBoxes(float64(width), float64(height))
	for _, box := range boxes {
		boxWidth := vect.Float(box.Max.X - box.Min.X)
		boxHeight := vect.Float(box.Max.Y - box.Min.Y)
		x, y := box.Center().XY()
		pos := vect.Vect{X: vect.Float(x), Y: vect.Float(y)}

		b := chipmunk.NewBox(pos, boxWidth, boxHeight)
		b.SetElasticity(0.4)
		b.SetFriction(200)
		body.AddShape(b)
	}
	body.SetPosition(position)
	body.UserData = "cup"

	return
}

func getCupBoxes(width, height float64) []pixel.Rect {
	delta := width * 0.10
	hw := width / 2
	hh := height / 2
	//return []pixel.Rect{
	//	pixel.R(0, 0, delta, height),
	//	pixel.R(delta, 0, width-delta, delta),
	//	pixel.R(width-delta, 0, width, height),
	//}
	return []pixel.Rect{
		pixel.R(-hw, -hh, -hw+delta, hh),
		pixel.R(-hw+delta, -hh, hw-delta, -hh+delta),
		pixel.R(hw-delta, -hh, hw, hh),
	}
}
