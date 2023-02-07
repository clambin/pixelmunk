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
	world = pixelmunk.NewWorld("pendulum", 0, 0, x, y)
	world.Space.Gravity = vect.Vect{Y: -981}

	midX := vect.Float(x / 3)
	midY := vect.Float(y / 2)
	startY := vect.Float(y * 2 / 3)
	anchor := pixelmunk.NewCircle(pixelmunk.DrawableOptions{
		Color: colornames.Red,
		BodyOptions: pixelmunk.BodyOptions{
			StaticBody: true,
			Position:   vect.Vect{X: midX, Y: startY},
			Mass:       1e3,
			CircleOptions: pixelmunk.CircleOptions{
				Radius: 20,
			},
		},
	})

	ball := pixelmunk.NewCircle(pixelmunk.DrawableOptions{
		Color: colornames.Orange,
		BodyOptions: pixelmunk.BodyOptions{
			Position:   vect.Vect{X: midX + 400, Y: startY + 200},
			Mass:       1e11,
			Elasticity: 0.99,
			Friction:   5.0,
			CircleOptions: pixelmunk.CircleOptions{
				Radius: 50,
			},
		},
	})

	anchorOffset := vect.Vect{X: 0, Y: -vect.Float(anchor.GetOptions().BodyOptions.CircleOptions.Radius)}

	ballOffset := anchor.GetBody().Position()
	ballOffset.Sub(anchorOffset)
	ballOffset.Sub(ball.GetBody().Position())

	joint := pixelmunk.NewJointWithAnchor(anchor, ball,
		anchorOffset,
		ballOffset,
		pixelmunk.DrawableOptions{
			Color:     colornames.Darkgray,
			Thickness: 1,
			JointOptions: pixelmunk.JointOptions{
				Draw: true,
			},
		})

	wall := pixelmunk.NewBox(pixelmunk.DrawableOptions{
		Color:     colornames.Blue,
		Thickness: 0,
		BodyOptions: pixelmunk.BodyOptions{
			StaticBody: true,
			Position:   vect.Vect{X: 20, Y: midY},
			Mass:       1e3,
			Elasticity: 0.99,
			BoxOptions: pixelmunk.BoxOptions{
				Width:  40,
				Height: vect.Float(y),
			},
		},
	})

	world.Add(anchor, ball, joint, wall)

	return
}
