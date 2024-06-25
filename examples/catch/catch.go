package main

import (
	"fmt"
	"github.com/clambin/pixelmunk"
	"github.com/clambin/pixelmunk/examples/catch/ball"
	"github.com/clambin/pixelmunk/examples/catch/cup"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

const (
	width       = 1024
	height      = 1080
	floorHeight = 40
	cupWidth    = 300
	cupHeight   = 200
)

func main() {
	app := createApp()
	opengl.Run(app.world.Run)
}

type catch struct {
	world *pixelmunk.World
	cup   *cup.Cup
}

func createApp() (app *catch) {
	app = &catch{
		world: pixelmunk.NewWorld("catch!", 0, 0, width, height),
	}
	app.world.RunFunc = app.run
	app.world.Space.Gravity = vect.Vect{X: 0, Y: -900}

	// Floor
	app.world.Add(pixelmunk.NewBox(pixelmunk.DrawableOptions{
		Color: colornames.Blue,
		BodyOptions: pixelmunk.BodyOptions{
			StaticBody: true,
			Position:   vect.Vect{X: vect.Float(width) / 2, Y: floorHeight / 2},
			Mass:       10e3,
			Elasticity: 0.6,
			Friction:   1.0,
			BoxOptions: pixelmunk.BoxOptions{
				Width:  vect.Float(width),
				Height: floorHeight,
			},
		},
	}))

	// Cup
	app.cup = cup.NewCup(vect.Vect{X: 400, Y: floorHeight + cupHeight/2}, cupWidth, cupHeight, colornames.Brown)
	app.world.Add(app.cup)

	return
}

func (c *catch) run(win *opengl.Window) {
	timer := time.Now()
	frameTicker := time.NewTicker(time.Second / time.Duration(c.world.FrameRate))
	ballTicker := time.NewTicker(1 * time.Second)

	for !win.Closed() {
		select {
		case <-ballTicker.C:
			c.addBall()
			c.cleanup()
		case <-frameTicker.C:
			c.world.Space.Step(1.0 / vect.Float(c.world.FrameRate))
			win.Clear(colornames.Black)
			c.world.Draw(win)
			win.Update()
			c.processEvents(win)

			win.SetTitle(fmt.Sprintf("gravity (%.1f fps)", 1/time.Now().Sub(timer).Seconds()))
			timer = time.Now()
		}
	}
}

func (c *catch) addBall() {
	pos := vect.Vect{
		X: vect.Float(c.world.Bounds.Min.X + float64(rand.Intn(int(c.world.Bounds.Max.X-c.world.Bounds.Min.X)))),
		Y: vect.Float(c.world.Bounds.Max.Y),
	}
	c.world.Add(ball.NewBall(pos, 20.0, colornames.Yellow))
}

func (c *catch) cleanup() {
	for index := 0; index < len(c.world.Objects); index++ {
		if c.world.Objects[index].GetBody().Position().Y < 0 {
			c.world.Space.RemoveBody(c.world.Objects[index].GetBody())
			c.world.Objects = append(c.world.Objects[:index], c.world.Objects[index+1:]...)
			index--
		}
	}
}

func (c *catch) processEvents(win *opengl.Window) {
	if win.JustReleased(pixel.KeyRight) {
		c.cup.SetDirection(1.0)
	}
	if win.JustReleased(pixel.KeyLeft) {
		c.cup.SetDirection(-1.0)
	}

	c.cup.Move()
}
