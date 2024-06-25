package main

import (
	"fmt"
	"github.com/clambin/pixelmunk"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"math/rand"
	"time"
)

var world *pixelmunk.World

func main() {
	world = createWorld(1024, 1080)
	opengl.Run(world.Run)
}

func createWorld(x, y float64) (world *pixelmunk.World) {
	world = pixelmunk.NewWorld("angled boxes", 0, 0, x, y)
	world.RunFunc = run

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

	angle := 0.0
	for i := 0; i < 8; i++ {
		world.Add(pixelmunk.NewBox(pixelmunk.DrawableOptions{
			Color: colors[rand.Intn(len(colors))],
			BodyOptions: pixelmunk.BodyOptions{
				Mass:     1,
				Position: vect.Vect{X: vect.Float(50 + i*150), Y: 500},
				Angle:    vect.Float(angle),
				BoxOptions: pixelmunk.BoxOptions{
					Width:  40,
					Height: 80,
				},
			},
		}))
		angle += math.Pi / 4
	}

	return
}

func run(win *opengl.Window) {
	frameTicker := time.NewTicker(time.Second / time.Duration(world.FrameRate))
	timer := time.Now()

	for !win.Closed() {
		for _, body := range world.Space.Bodies {
			angle := body.Angle()
			body.SetAngle(angle - vect.Float(math.Pi/float64(world.FrameRate*2)))
		}
		world.Space.Step(1.0 / vect.Float(world.FrameRate))

		win.Clear(colornames.Black)
		world.Draw(win)
		win.Update()

		win.SetTitle(fmt.Sprintf("%s (%.1f fps)", world.Name, 1/time.Now().Sub(timer).Seconds()))
		timer = time.Now()

		<-frameTicker.C
	}
}
