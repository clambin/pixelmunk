package pixelmunk

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"math"
)

// Joint joins two bodies together
type Joint struct {
	//Drawable
	pivotJoint  *chipmunk.PivotJoint
	options     DrawableOptions
	origOffsetA vect.Vect
	origOffsetB vect.Vect
}

// JointOptions contains options for drawing joints
type JointOptions struct {
	Draw bool
}

var _ Drawable = &Joint{}

/*
// NewJoint creates a constraint between two bodies
func NewJoint(object1, object2 *Object, options DrawableOptions) *Joint {
	return &Joint{
		pivotJoint: chipmunk.NewPivotJoint(object1.GetBody(), object2.GetBody()),
		options:    options,
	}
}
*/

// NewJointWithAnchor creates a constraint between two bodies at specified positions
func NewJointWithAnchor(object1, object2 *Object, offset1, offset2 vect.Vect, options DrawableOptions) *Joint {
	j := &Joint{
		pivotJoint:  chipmunk.NewPivotJointAnchor(object1.GetBody(), object2.GetBody(), offset1, offset2),
		options:     options,
		origOffsetA: offset1,
		origOffsetB: offset2,
	}
	return j
}

// GetType returns the type of drawable
func (j Joint) GetType() DrawableType {
	return DrawableJoint
}

// GetBody returns the chipmunk.Body that the Object represents
func (j Joint) GetBody() *chipmunk.Body {
	return nil
}

// GetJoint returns the chipmunk.PivotJoint that the Object represents
func (j Joint) GetJoint() *chipmunk.PivotJoint {
	return j.pivotJoint
}

// GetOptions returns the DrawableOptions that were used to create the Object
func (j Joint) GetOptions() DrawableOptions {
	return j.options
}

// Draw draws the Joint on the provided imdraw.IMDraw
func (j Joint) Draw(imd *imdraw.IMDraw) {
	if j.options.JointOptions.Draw {
		pA := j.pivotJoint.BodyA.Position()
		pB := j.pivotJoint.BodyB.Position()

		imd.Color = j.options.Color

		// line from first body to offset
		imd.Push(pixel.V(float64(pA.X), float64(pA.Y)))
		pA.Add(rotateVector(j.origOffsetA, j.pivotJoint.BodyA.Angle()))
		imd.Push(pixel.V(float64(pA.X), float64(pA.Y)))
		imd.Line(j.options.Thickness)

		// Line from 2nd body to offset
		imd.Push(pixel.V(float64(pB.X), float64(pB.Y)))
		pB.Add(rotateVector(j.origOffsetB, j.pivotJoint.BodyB.Angle()))
		imd.Push(pixel.V(float64(pB.X), float64(pB.Y)))
		imd.Line(j.options.Thickness)

		// line between 2 offsets
		imd.Push(pixel.V(float64(pA.X), float64(pA.Y)), pixel.V(float64(pB.X), float64(pB.Y)))
		imd.Line(j.options.Thickness)
	}
}

func rotateVector(v vect.Vect, angle vect.Float) (a vect.Vect) {
	//𝑥2=cos𝛽𝑥1−sin𝛽𝑦1
	//𝑦2=sin𝛽𝑥1+cos𝛽𝑦1
	a = vect.Vect{
		X: v.X*vect.Float(math.Cos(float64(angle))) - v.Y*vect.Float(math.Sin(float64(angle))),
		Y: v.X*vect.Float(math.Sin(float64(angle))) + v.Y*vect.Float(math.Cos(float64(angle))),
	}
	return
}
