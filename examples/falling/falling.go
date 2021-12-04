package main

import (
	"fmt"
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"math"
	"time"
)

func main() {
	w := New(1024, 1080)
	pixelgl.Run(w.World.Run)
}

type World struct {
	World *pixelmunk.World
}

func New(x, y float64) (w *World) {
	w = &World{
		World: &pixelmunk.World{
			Name:   "falling block",
			Bounds: pixel.R(0, 0, x, y),
			Space:  chipmunk.NewSpace(),
		},
	}
	w.World.Space.Gravity = vect.Vect{X: 0, Y: -981}
	w.World.RunFunc = w.Run

	// Floor
	w.World.Add(pixelmunk.NewBox(pixelmunk.ObjectOptions{
		Color:          colornames.Blue,
		BodyOptions:    pixelmunk.ObjectBodyOptions{
			StaticBody: true,
			Position:   vect.Vect{X: vect.Float(x)/2, Y: 20},
			Mass:       10e3,
			Elasticity: 0.1,
			Friction:   1.0,
			BoxOptions: pixelmunk.ObjectBoxOptions{
				Width:  vect.Float(x),
				Height: 40,
			},
		},
	}))

	// Flat box
	w.World.Add(pixelmunk.NewBox(pixelmunk.ObjectOptions{
		Color:          colornames.Green,
		CustomDrawFunc: []pixelmunk.CustomDrawFunc{drawVelocity},
		BodyOptions:    pixelmunk.ObjectBodyOptions{
			Position:   vect.Vect{X: 500, Y: 200},
			Mass:       10e10,
			Elasticity: 0.2,
			Friction:   1.0,
			BoxOptions: pixelmunk.ObjectBoxOptions{
				Width:  500,
				Height: 10,
			},
		},
	}))

	// First falling box
	w.World.Add(pixelmunk.NewBox(pixelmunk.ObjectOptions{
		Color:          colornames.Red,
		CustomDrawFunc: []pixelmunk.CustomDrawFunc{drawVelocity},
		BodyOptions:    pixelmunk.ObjectBodyOptions{
			Position:   vect.Vect{X: 600, Y: 1000},
			Angle:      math.Pi * 4/3,
			Mass:       10e3,
			Elasticity: 0.2,
			Friction:   1.0,
			BoxOptions: pixelmunk.ObjectBoxOptions{
				Width:  50,
				Height: 100,
			},
		},
	}))

	// Second falling object
	w.World.Add(pixelmunk.NewCircle(pixelmunk.ObjectOptions{
		Color:          colornames.Purple,
		CustomDrawFunc: []pixelmunk.CustomDrawFunc{drawVelocity},
		BodyOptions:    pixelmunk.ObjectBodyOptions{
			Position:   vect.Vect{X: 600, Y: 3000},
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

func (w *World) Run(win *pixelgl.Window) {
	const frameRate = 60
	frameTicker := time.NewTicker(time.Second / frameRate)
	timer := time.Now()

	for !win.Closed() {
		select {
		case <-frameTicker.C:
			w.World.Space.Step(1.0 / frameRate)
			win.Clear(colornames.Black)
			w.World.Draw(win)
			win.Update()

			win.SetTitle(fmt.Sprintf("gravity (%.1f fps)", 1/time.Now().Sub(timer).Seconds()))
			timer = time.Now()
		}
	}
}

func drawVelocity(object *pixelmunk.Object, imd *imdraw.IMDraw) {
	body := object.GetBody()
	pos := body.Position()
	velocity := body.Velocity()

	imd.Color = colornames.Red
	imd.Push(
		pixel.V(float64(pos.X), float64(pos.Y)),
		pixel.V(float64(pos.X + velocity.X), float64(pos.Y + velocity.Y)),
	)
	imd.Line(1)
}
