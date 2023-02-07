package pixelmunk

import (
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

// NewCircle creates a new Object that represents a circle / disc / ball
func NewCircle(options DrawableOptions) *Object {
	options.BodyOptions.Type = chipmunk.ShapeType_Circle
	shape := chipmunk.NewCircle(
		vect.Vector_Zero,
		options.BodyOptions.CircleOptions.Radius,
	)

	return NewObjectWithShape(shape, options)
}
