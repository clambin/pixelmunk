package pixelmunk

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"image/color"
)

type World struct {
	Name      string
	Bounds    pixel.Rect
	FrameRate int
	Space     *chipmunk.Space
	RunFunc   func(*pixelgl.Window)
	objects   []Drawable
}

type Drawable interface {
	Draw(imd *imdraw.IMDraw)
	GetBody() *chipmunk.Body
}

type Object struct {
	body    *chipmunk.Body
	options ObjectOptions
}

type CustomDrawFunc func(object *Object, draw *imdraw.IMDraw)

type ObjectOptions struct {
	Color          color.Color
	Thickness      float64
	CustomDrawFunc []CustomDrawFunc
	BodyOptions    ObjectBodyOptions
}

type ObjectBodyOptions struct {
	StaticBody    bool
	Position      vect.Vect
	Angle         vect.Float
	Mass          vect.Float
	Velocity      vect.Vect
	Elasticity    vect.Float
	Friction      vect.Float
	Type          chipmunk.ShapeType
	CircleOptions ObjectCircleOptions
	BoxOptions    ObjectBoxOptions
}

type ObjectCircleOptions struct {
	Radius float32
}

type ObjectBoxOptions struct {
	Width  vect.Float
	Height vect.Float
}
