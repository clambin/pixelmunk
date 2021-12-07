package pixelmunk

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"image/color"
)

// World represents the world that chipmunk will simulate
type World struct {
	Name      string
	Bounds    pixel.Rect
	FrameRate int
	Space     *chipmunk.Space
	RunFunc   func(*pixelgl.Window)
	Objects   []Drawable
}

// Drawable interface for any drawable Object
type Drawable interface {
	Draw(imd *imdraw.IMDraw)
	GetBody() *chipmunk.Body
	GetOptions() ObjectOptions
}

// Object represents an object inside the World
type Object struct {
	body    *chipmunk.Body
	options ObjectOptions
}

// CustomDrawFunc is the callback function to provide additional functionality when drawing an Option.
// This is called after the main shape is drawn.
type CustomDrawFunc func(object *Object, draw *imdraw.IMDraw)

// ObjectOptions for creating a new Object
type ObjectOptions struct {
	Color          color.Color
	Thickness      float64
	CustomDrawFunc []CustomDrawFunc
	BodyOptions    ObjectBodyOptions
}

// ObjectBodyOptions holds the physical attributes for the Object
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

// ObjectCircleOptions holds the attributes for a Circle object
type ObjectCircleOptions struct {
	Radius float32
}

// ObjectBoxOptions holds the attributes for a Box object
type ObjectBoxOptions struct {
	Width  vect.Float
	Height vect.Float
}
