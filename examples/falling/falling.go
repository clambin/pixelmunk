package main

import (
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"math"
)

func main() {
	w := New(1024, 1080)
	pixelgl.Run(w.world.Run)
}

type App struct {
	world *pixelmunk.World
}

func New(x, y float64) (app *App) {
	app = &App{
		world: pixelmunk.NewWorld("falling blocks", 0, 0, x, y),
	}
	app.world.Space.Gravity = vect.Vect{Y: -981}

	// Floor
	app.world.Add(pixelmunk.NewBox(pixelmunk.ObjectOptions{
		Color: colornames.Blue,
		BodyOptions: pixelmunk.ObjectBodyOptions{
			StaticBody: true,
			Position:   vect.Vect{X: vect.Float(x) / 2, Y: 20},
			Mass:       10e3,
			Elasticity: 0.1,
			Friction:   1.0,
			BoxOptions: pixelmunk.ObjectBoxOptions{
				Width:  vect.Float(x) - 200,
				Height: 40,
			},
		},
	}))

	// Flat box
	app.world.Add(pixelmunk.NewBox(pixelmunk.ObjectOptions{
		Color:          colornames.Green,
		CustomDrawFunc: []pixelmunk.CustomDrawFunc{drawVelocity},
		BodyOptions: pixelmunk.ObjectBodyOptions{
			Position:   vect.Vect{X: 500, Y: 200},
			Mass:       10e10,
			Elasticity: 0.2,
			Friction:   0.1,
			BoxOptions: pixelmunk.ObjectBoxOptions{
				Width:  500,
				Height: 25,
			},
		},
	}))

	// First falling object
	app.world.Add(pixelmunk.NewBox(pixelmunk.ObjectOptions{
		Color:          colornames.Red,
		CustomDrawFunc: []pixelmunk.CustomDrawFunc{drawVelocity},
		BodyOptions: pixelmunk.ObjectBodyOptions{
			Position:   vect.Vect{X: 600, Y: 1000},
			Velocity:   vect.Vect{X: -10, Y: 0},
			Angle:      math.Pi * 5.9 / 3,
			Mass:       10e3,
			Elasticity: 0.2,
			Friction:   5.0,
			BoxOptions: pixelmunk.ObjectBoxOptions{
				Width:  50,
				Height: 100,
			},
		},
	}))

	// Second falling object
	app.world.Add(pixelmunk.NewCircle(pixelmunk.ObjectOptions{
		Color:          colornames.Purple,
		CustomDrawFunc: []pixelmunk.CustomDrawFunc{drawVelocity},
		BodyOptions: pixelmunk.ObjectBodyOptions{
			Position:   vect.Vect{X: 600, Y: 2000},
			Angle:      math.Pi * 4 / 3,
			Mass:       10e3,
			Elasticity: 0.9,
			Friction:   10.0,
			CircleOptions: pixelmunk.ObjectCircleOptions{
				Radius: 25,
			},
		},
	}))

	return
}

func drawVelocity(object *pixelmunk.Object, imd *imdraw.IMDraw) {
	body := object.GetBody()
	pos := body.Position()
	velocity := body.Velocity()

	imd.Color = colornames.Red
	imd.Push(
		pixel.V(float64(pos.X), float64(pos.Y)),
		pixel.V(float64(pos.X+velocity.X), float64(pos.Y+velocity.Y)),
	)
	imd.Line(1)
}
