package pixelmunk

import (
	"fmt"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

// Object represents an object inside the World
type Object struct {
	//Drawable
	body    *chipmunk.Body
	options DrawableOptions
}

// DrawableOptions for a drawable
type DrawableOptions struct {
	Color          color.Color
	Thickness      float64
	CustomDrawFunc []CustomDrawFunc
	BodyOptions    BodyOptions
	JointOptions   JointOptions
}

// CustomDrawFunc is the callback function to provide additional functionality when drawing an Option.
// This is called after the main shape is drawn.
type CustomDrawFunc func(object *Object, draw *imdraw.IMDraw)

// BodyOptions holds the physical attributes for the Object
type BodyOptions struct {
	StaticBody    bool
	Position      vect.Vect
	Angle         vect.Float
	Mass          vect.Float
	Velocity      vect.Vect
	Elasticity    vect.Float
	Friction      vect.Float
	Type          chipmunk.ShapeType
	CircleOptions CircleOptions
	BoxOptions    BoxOptions
}

// CircleOptions holds the attributes for a Circle object
type CircleOptions struct {
	Radius float32
}

// BoxOptions holds the attributes for a Box object
type BoxOptions struct {
	Width  vect.Float
	Height vect.Float
}

var _ Drawable = &Object{}

// NewObject creates a new Object for the provided Body and DrawableOptions
func NewObject(body *chipmunk.Body, options DrawableOptions) *Object {
	return &Object{
		body:    body,
		options: options,
	}
}

// NewObjectWithShape creates a new Object for the provided Shape and DrawableOptions
func NewObjectWithShape(shape *chipmunk.Shape, options DrawableOptions) *Object {
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
	body.SetVelocity(float32(options.BodyOptions.Velocity.X), float32(options.BodyOptions.Velocity.Y))
	body.SetAngle(options.BodyOptions.Angle)

	return NewObject(body, options)
}

// GetType returns the type of drawable
func (o Object) GetType() DrawableType {
	return DrawableBody
}

// GetBody returns the chipmunk.Body that the Object represents
func (o Object) GetBody() *chipmunk.Body {
	return o.body
}

// GetJoint returns the chipmunk.PivotJoint that the Object represents
func (o Object) GetJoint() *chipmunk.PivotJoint {
	return nil
}

// GetOptions returns the DrawableOptions that were used to create the Object
func (o Object) GetOptions() DrawableOptions {
	return o.options
}

// Draw draws the Object on the provided imdraw.IMDraw
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

	for angle < 0 {
		angle += 2 * math.Pi
	}
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
