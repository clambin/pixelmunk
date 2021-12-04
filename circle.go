package pixelmunk

import (
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

func NewCircle(options ObjectOptions) Object {
	options.BodyOptions.Type = chipmunk.ShapeType_Circle

	shape := chipmunk.NewCircle(
		vect.Vector_Zero,
		options.BodyOptions.CircleOptions.Radius,
	)
	shape.SetElasticity(options.BodyOptions.Elasticity)
	shape.SetFriction(options.BodyOptions.Friction)

	var body *chipmunk.Body
	if options.BodyOptions.StaticBody {
		body = chipmunk.NewBodyStatic()
	} else {
		body = chipmunk.NewBody(options.BodyOptions.Mass, shape.Moment(float32(options.BodyOptions.Mass)))
	}
	body.AddShape(shape)
	body.SetPosition(options.BodyOptions.Position)
	body.SetAngle(options.BodyOptions.Angle)

	return Object{
		body:    body,
		options: options,
	}
}
