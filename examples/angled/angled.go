package main

import (
	"fmt"
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"math/rand"
	"time"
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
		world: pixelmunk.NewWorld("angled boxes", 0, 0, x, y),
	}
	app.world.RunFunc = app.Run

	colors := []color.Color{
		colornames.White,
		colornames.Red,
		colornames.Orange,
		colornames.Yellow,
		colornames.Green,
		colornames.Blue,
		colornames.Indigo,
		colornames.Violet,
	}

	rand.Seed(time.Now().Unix())
	angle := 0.0
	for i := 0; i < 8; i++ {
		app.world.Add(pixelmunk.NewBox(pixelmunk.ObjectOptions{
			Color: colors[rand.Intn(len(colors))],
			BodyOptions: pixelmunk.ObjectBodyOptions{
				Mass:     1,
				Position: vect.Vect{X: vect.Float(50 + i*150), Y: 500},
				Angle:    vect.Float(angle),
				BoxOptions: pixelmunk.ObjectBoxOptions{
					Width:  40,
					Height: 80,
				},
			},
		}))
		angle += math.Pi / 4
	}

	return
}

func (app *App) Run(win *pixelgl.Window) {
	frameTicker := time.NewTicker(time.Second / time.Duration(app.world.FrameRate))
	timer := time.Now()

	for !win.Closed() {
		for _, body := range app.world.Space.Bodies {
			angle := body.Angle()
			body.SetAngle(angle + math.Pi/150)
		}
		app.world.Space.Step(1.0 / vect.Float(app.world.FrameRate))

		win.Clear(colornames.Black)
		app.world.Draw(win)
		win.Update()

		win.SetTitle(fmt.Sprintf("%s (%.1f fps)", app.world.Name, 1/time.Now().Sub(timer).Seconds()))
		timer = time.Now()

		<-frameTicker.C
	}
}
