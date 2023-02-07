package main

import (
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

func main() {
	app := createWorld(1024, 1080)
	pixelgl.Run(app.world.Run)
}

type App struct {
	world      *pixelmunk.World
	anchor     *pixelmunk.Object
	fire       bool
	fireTicker *time.Ticker
}

func createWorld(x, y float64) (app App) {
	app = App{
		world: pixelmunk.NewWorld("ball and chain", 0, 0, x, y),
	}

	app.world.Space.Gravity = vect.Vect{Y: -981}
	app.world.FrameRate = 120
	app.world.RunCallback = app.Process
	//app.fireTicker = time.NewTicker(time.Second)

	midX := vect.Float(x / 2)
	midY := vect.Float(y * 3 / 4)

	app.anchor = createAnchor(midX, midY)
	ball := createBall(midX+300, midY+200)
	links, joints := createChain(app.anchor, ball)

	app.world.Add(app.anchor, ball)
	for _, link := range links {
		app.world.Add(link)
	}
	for _, joint := range joints {
		app.world.Add(joint)
	}

	return
}

func createAnchor(x, y vect.Float) (anchor *pixelmunk.Object) {
	anchor = pixelmunk.NewCircle(pixelmunk.DrawableOptions{
		Color: colornames.Red,
		BodyOptions: pixelmunk.BodyOptions{
			StaticBody: true,
			Position:   vect.Vect{X: x, Y: y},
			Mass:       1e3,
			Elasticity: 1,
			CircleOptions: pixelmunk.CircleOptions{
				Radius: 20,
			},
		},
	})
	return
}

func createBall(x, y vect.Float) (ball *pixelmunk.Object) {
	ball = pixelmunk.NewCircle(pixelmunk.DrawableOptions{
		Color: colornames.Orange,
		BodyOptions: pixelmunk.BodyOptions{
			Position:   vect.Vect{X: x, Y: y},
			Mass:       1e3,
			Elasticity: 0.99,
			CircleOptions: pixelmunk.CircleOptions{
				Radius: 50,
			},
		},
	})
	return
}

func createChain(anchor, ball *pixelmunk.Object) (links []*pixelmunk.Object, joints []*pixelmunk.Joint) {

	const linkWidth = 10
	const linkCount = 20

	deltaX := ball.GetBody().Position().X - anchor.GetBody().Position().X - vect.Float(ball.GetOptions().BodyOptions.CircleOptions.Radius) - vect.Float(anchor.GetOptions().BodyOptions.CircleOptions.Radius)
	deltaX /= linkCount
	deltaY := ball.GetBody().Position().Y - anchor.GetBody().Position().Y
	deltaY /= linkCount

	x := anchor.GetBody().Position().X + vect.Float(anchor.GetOptions().BodyOptions.CircleOptions.Radius)
	y := anchor.GetBody().Position().Y

	for count := 0; count < linkCount; count++ {
		links = append(links, pixelmunk.NewBox(pixelmunk.DrawableOptions{
			Color: colornames.Darkgray,
			BodyOptions: pixelmunk.BodyOptions{
				Position: vect.Vect{X: x, Y: anchor.GetBody().Position().Y + deltaY*vect.Float(count)},
				Mass:     5e2,
				BoxOptions: pixelmunk.BoxOptions{
					Width:  linkWidth,
					Height: 6,
				},
			},
		}))
		x += deltaX
		y += deltaY
	}

	jointOptions := pixelmunk.DrawableOptions{
		Color:     colornames.Darkgoldenrod,
		Thickness: 1,
		JointOptions: pixelmunk.JointOptions{
			Draw: true,
		},
	}

	joints = append(joints, pixelmunk.NewJointWithAnchor(
		anchor, links[0],
		vect.Vect{Y: -vect.Float(anchor.GetOptions().BodyOptions.CircleOptions.Radius)},
		vect.Vect{X: -linkWidth / 2},
		jointOptions,
	),
	)

	for count := 1; count < linkCount; count++ {
		joints = append(joints, pixelmunk.NewJointWithAnchor(
			links[count-1], links[count],
			vect.Vect{X: linkWidth / 2},
			vect.Vect{X: -linkWidth / 2},
			jointOptions,
		),
		)
	}

	joints = append(joints, pixelmunk.NewJointWithAnchor(
		links[linkCount-1], ball,
		vect.Vect{X: linkWidth / 2},
		vect.Vect{X: -vect.Float(ball.GetOptions().BodyOptions.CircleOptions.Radius)},
		jointOptions,
	),
	)

	return
}

func (app *App) Process(win *pixelgl.Window) {
	if win.JustReleased(pixelgl.KeySpace) {
		if app.fire {
			app.fire = false
			app.fireTicker.Stop()
		} else {
			app.fire = true
			app.fireBullet()
			app.fireTicker = time.NewTicker(time.Second)
		}
	}
	if app.fire {
		select {
		case <-app.fireTicker.C:
			app.fireBullet()
			app.cleanup()
		default:
		}
	}
}

func (app *App) fireBullet() {
	position := vect.Vect{X: 20, Y: 700}
	velocity := vect.Vect{X: 2e3, Y: 0}

	if rand.Intn(2) == 0 {
		position.X = vect.Float(app.world.Bounds.Max.X) - 20
		velocity.X = -velocity.X
	}
	bullet := pixelmunk.NewCircle(pixelmunk.DrawableOptions{
		Color:     colornames.Silver,
		Thickness: 0,
		BodyOptions: pixelmunk.BodyOptions{
			Position:   position,
			Velocity:   velocity,
			Elasticity: 0,
			Mass:       1e4,
			CircleOptions: pixelmunk.CircleOptions{
				Radius: 5,
			},
		},
	})
	app.world.Add(bullet)
}

func (app *App) cleanup() {
	for _, object := range app.world.Objects {
		if object.GetType() == pixelmunk.DrawableBody {
			if object.GetBody().Position().Y < 0 {
				app.world.Remove(object)
			}
		}
	}
}
