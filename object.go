package pixelmunk

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/vova616/chipmunk"
	"golang.org/x/image/colornames"
	"math"
)

func NewObject(body *chipmunk.Body, options ObjectOptions) Object {
	return Object{
		body:    body,
		options: options,
	}
}

func (o Object) GetBody() *chipmunk.Body {
	return o.body
}

func (o Object) Draw(imd *imdraw.IMDraw) {
	for _, shape := range o.GetBody().Shapes {
		switch shape.ShapeType() {
		case chipmunk.ShapeType_Circle:
			o.drawCircle(imd, shape)
		case chipmunk.ShapeType_Box:
			o.drawBox(imd, shape)
		default:
			panic(fmt.Sprintf("unsupported shape type: %d", shape.ShapeType()))
		}
		for _, customDrawFunc := range o.options.CustomDrawFunc {
			customDrawFunc(&o, imd)
		}
	}
}

func (o Object) drawCircle(imd *imdraw.IMDraw, shape *chipmunk.Shape) {
	lower := shape.BB.Lower
	upper := shape.BB.Upper
	radius := shape.GetAsCircle().Radius

	position := pixel.V(
		float64(lower.X+upper.X)/2,
		float64(lower.Y+upper.Y)/2,
	)

	imd.Color = o.options.Color
	imd.Push(position)
	imd.Circle(float64(radius), o.options.Thickness)
}

func (o Object) drawBox(imd *imdraw.IMDraw, shape *chipmunk.Shape) {
	lower := shape.BB.Lower
	upper := shape.BB.Upper
	angle := float64(shape.Body.Angle())

	box := shape.GetAsBox()
	width := float64(box.Width)
	height := float64(box.Height)

	for angle > math.Pi {
		angle -= math.Pi
	}
	l := width
	if angle > math.Pi/2 {
		angle -= math.Pi / 2
		l = height
	}
	sin, cos := math.Sincos(angle)

	corners := []pixel.Vec{
		pixel.V(float64(upper.X)-l*cos, float64(lower.Y)),
		pixel.V(float64(upper.X), float64(lower.Y)+l*sin),
		pixel.V(float64(lower.X)+l*cos, float64(upper.Y)),
		pixel.V(float64(lower.X), float64(upper.Y)-l*sin),
	}

	const debug = false
	if debug {
		imd.Color = colornames.White
		imd.Push(pixel.V(float64(lower.X), float64(lower.Y)))
		imd.Push(pixel.V(float64(upper.X), float64(upper.Y)))
		imd.Rectangle(0)
	}

	imd.Color = o.options.Color
	imd.Push(corners...)
	imd.Polygon(o.options.Thickness)
}
