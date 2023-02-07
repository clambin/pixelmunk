package main

import (
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
)

func main() {
	w := createWorld(1024, 1080)
	pixelgl.Run(w.Run)
}

func createWorld(x, y float64) (world *pixelmunk.World) {
	world = pixelmunk.NewWorld("ragdoll", 0, 0, x, y)
	world.Space.Gravity = vect.Vect{Y: -981}
	world.FrameRate = 120

	// Floor
	world.Add(pixelmunk.NewBox(pixelmunk.DrawableOptions{
		Color: colornames.Blue,
		BodyOptions: pixelmunk.BodyOptions{
			StaticBody: true,
			Position:   vect.Vect{X: vect.Float(x) / 2, Y: 20},
			Mass:       10e3,
			Elasticity: 1.2,
			Friction:   1,
			BoxOptions: pixelmunk.BoxOptions{
				Width:  vect.Float(x) - 200,
				Height: 40,
			},
		},
	}))

	// doll
	_ = NewRagDoll(world, x, y)
	return
}

const jointThickness = 0

type RagDoll struct {
	Head      *pixelmunk.Object
	Torso     *pixelmunk.Object
	Neck      *pixelmunk.Joint
	Arms      []*pixelmunk.Object
	Shoulders []*pixelmunk.Joint
	Legs      []*pixelmunk.Object
	Hips      []*pixelmunk.Joint
}

func NewRagDoll(w *pixelmunk.World, x, y float64) (ragdoll RagDoll) {
	startX := vect.Float(x / 2)
	startY := vect.Float(y * 0.9)

	ragdoll.Head = pixelmunk.NewCircle(pixelmunk.DrawableOptions{
		Color: colornames.Orange,
		BodyOptions: pixelmunk.BodyOptions{
			Position:      vect.Vect{X: startX, Y: startY},
			Mass:          1e2,
			Elasticity:    0.2,
			Friction:      1,
			CircleOptions: pixelmunk.CircleOptions{Radius: 15},
		},
	})
	w.Add(ragdoll.Head)

	ragdoll.Torso = pixelmunk.NewBox(pixelmunk.DrawableOptions{
		Color: colornames.Blue,
		BodyOptions: pixelmunk.BodyOptions{
			Position:   vect.Vect{X: startX, Y: startY - 15 - 30},
			Mass:       1e4,
			Elasticity: 0.5,
			Friction:   1,
			BoxOptions: pixelmunk.BoxOptions{
				Width:  40,
				Height: 60,
			},
		},
	})
	w.Add(ragdoll.Torso)

	p1 := ragdoll.Head.GetBody().Position()
	p1.Add(vect.Vect{X: 0, Y: -15})
	p2 := ragdoll.Torso.GetBody().Position()
	p2.Add(vect.Vect{X: 0, Y: 30})
	ragdoll.Neck = pixelmunk.NewJointWithAnchor(
		ragdoll.Head, ragdoll.Torso,
		vect.Vect{X: 0, Y: -15}, vect.Vect{X: 0, Y: 30},
		pixelmunk.DrawableOptions{Color: colornames.Yellow, Thickness: jointThickness})
	w.Add(ragdoll.Neck)

	ragdoll.Arms = append(ragdoll.Arms, pixelmunk.NewBox(pixelmunk.DrawableOptions{
		Color: colornames.Orange,
		BodyOptions: pixelmunk.BodyOptions{
			Position:   vect.Vect{X: startX - 20 - 20, Y: startY - 15 - 5},
			Mass:       1e3,
			Elasticity: 0.8,
			Friction:   1,
			BoxOptions: pixelmunk.BoxOptions{
				Width:  40,
				Height: 10,
			},
		},
	}))
	w.Add(ragdoll.Arms[0])

	ragdoll.Shoulders = append(ragdoll.Shoulders, pixelmunk.NewJointWithAnchor(ragdoll.Torso, ragdoll.Arms[0],
		vect.Vect{X: -20, Y: 30 - 5}, vect.Vect{X: 20, Y: 0},
		pixelmunk.DrawableOptions{Color: colornames.Yellow, Thickness: jointThickness}))
	w.Add(ragdoll.Shoulders[0])

	ragdoll.Arms = append(ragdoll.Arms, pixelmunk.NewBox(pixelmunk.DrawableOptions{
		Color: colornames.Orange,
		BodyOptions: pixelmunk.BodyOptions{
			Position:   vect.Vect{X: startX + 20 + 20, Y: startY - 15 - 5},
			Mass:       1e3,
			Elasticity: 0.8,
			Friction:   1,
			BoxOptions: pixelmunk.BoxOptions{
				Width:  40,
				Height: 10,
			},
		},
	}))
	w.Add(ragdoll.Arms[1])

	ragdoll.Shoulders = append(ragdoll.Shoulders, pixelmunk.NewJointWithAnchor(ragdoll.Torso, ragdoll.Arms[1],
		vect.Vect{X: 20, Y: 30 - 5}, vect.Vect{X: -20, Y: 0},
		pixelmunk.DrawableOptions{Color: colornames.Yellow, Thickness: jointThickness}))
	w.Add(ragdoll.Shoulders[1])

	ragdoll.Legs = append(ragdoll.Legs, pixelmunk.NewBox(pixelmunk.DrawableOptions{
		Color: colornames.Blue,
		BodyOptions: pixelmunk.BodyOptions{
			Position:   vect.Vect{X: startX - 15, Y: startY - 15 - 60 - 20},
			Mass:       5e3,
			Elasticity: 0.8,
			Friction:   1,
			BoxOptions: pixelmunk.BoxOptions{
				Width:  10,
				Height: 40,
			},
		},
	}))
	w.Add(ragdoll.Legs[0])

	ragdoll.Hips = append(ragdoll.Hips, pixelmunk.NewJointWithAnchor(ragdoll.Torso, ragdoll.Legs[0],
		vect.Vect{X: -15, Y: -30}, vect.Vect{X: 0, Y: 20},
		pixelmunk.DrawableOptions{Color: colornames.Yellow, Thickness: jointThickness}))
	w.Add(ragdoll.Hips[0])

	ragdoll.Legs = append(ragdoll.Legs, pixelmunk.NewBox(pixelmunk.DrawableOptions{
		Color: colornames.Blue,
		BodyOptions: pixelmunk.BodyOptions{
			Position:   vect.Vect{X: startX + 15, Y: startY - 15 - 60 - 20},
			Mass:       5e3,
			Elasticity: 0.8,
			Friction:   1,
			BoxOptions: pixelmunk.BoxOptions{
				Width:  10,
				Height: 40,
			},
		},
	}))
	w.Add(ragdoll.Legs[1])

	ragdoll.Hips = append(ragdoll.Hips, pixelmunk.NewJointWithAnchor(ragdoll.Torso, ragdoll.Legs[1],
		vect.Vect{X: 15, Y: -30}, vect.Vect{X: 0, Y: 20},
		pixelmunk.DrawableOptions{Color: colornames.Yellow, Thickness: jointThickness}))
	w.Add(ragdoll.Hips[1])

	return
}
