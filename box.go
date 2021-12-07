package pixelmunk

import (
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

// NewBox creates a new Object for a rectangle
func NewBox(options ObjectOptions) Object {
	options.BodyOptions.Type = chipmunk.ShapeType_Box
	shape := chipmunk.NewBox(
		vect.Vector_Zero,
		options.BodyOptions.BoxOptions.Width,
		options.BodyOptions.BoxOptions.Height,
	)

	return NewObjectWithShape(shape, options)
}
