package pixelmunk

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/vova616/chipmunk"
)

// DrawableType indicates the subtype of the drawable
type DrawableType int

const (
	DrawableBody = iota
	DrawableJoint
)

// Drawable interface for any drawable Object
type Drawable interface {
	Draw(imd *imdraw.IMDraw)
	GetOptions() DrawableOptions
	GetType() DrawableType
	GetBody() *chipmunk.Body
	GetJoint() *chipmunk.PivotJoint
}
