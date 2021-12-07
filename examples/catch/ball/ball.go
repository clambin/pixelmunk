package ball

import (
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"math/rand"
)

// Ball represents a tennis ball
type Ball struct {
	pixelmunk.Object
}

// NewBall creates a new ball
func NewBall(position vect.Vect, radius float32, color color.Color) (ball *Ball) {
	ball = &Ball{
		Object: pixelmunk.NewCircle(pixelmunk.ObjectOptions{
			Color:          color,
			Thickness:      0,
			CustomDrawFunc: []pixelmunk.CustomDrawFunc{draw},
			BodyOptions: pixelmunk.ObjectBodyOptions{
				Position:      position,
				Angle:         vect.Float(rand.Float32()),
				Mass:          1,
				Elasticity:    0.9,
				Friction:      2e8,
				Type:          chipmunk.ShapeType_Circle,
				CircleOptions: pixelmunk.ObjectCircleOptions{Radius: radius},
			},
		}),
	}
	return
}

func draw(object *pixelmunk.Object, imd *imdraw.IMDraw) {
	body := object.GetBody()
	position := body.Position()
	angle := body.Angle()
	radius := body.Shapes[0].GetAsCircle().Radius

	dx := float64(radius) * math.Cos(float64(angle))
	dy := float64(radius) * math.Sin(float64(angle))

	imd.Color = colornames.Lightgrey
	imd.Push(
		pixel.V(float64(position.X)+dx, float64(position.Y)+dy),
		pixel.V(float64(position.X)-dx, float64(position.Y)-dy),
	)
	imd.Line(1)
}
