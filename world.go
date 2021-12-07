package pixelmunk

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"time"
)

const defaultFrameRate = 60

// NewWorld creates a new World
func NewWorld(name string, minX, minY, maxX, maxY float64) *World {
	return &World{
		Name:      name,
		Bounds:    pixel.R(minX, minY, maxX, maxY),
		FrameRate: defaultFrameRate,
		Space:     chipmunk.NewSpace(),
	}
}

// defaultRun is the default run function for a world. This is used if the run function isn't overridden by World.RunFunc
func (w *World) defaultRun(win *pixelgl.Window) {
	frameTicker := time.NewTicker(time.Second / time.Duration(w.FrameRate))
	timer := time.Now()

	for !win.Closed() {
		w.Space.Step(1.0 / vect.Float(w.FrameRate))

		win.Clear(colornames.Black)
		w.Draw(win)
		win.Update()

		win.SetTitle(fmt.Sprintf("%s (%.1f fps)", w.Name, 1/time.Now().Sub(timer).Seconds()))
		timer = time.Now()

		<-frameTicker.C
	}
}

// Add adds a new Object to the World
func (w *World) Add(object Drawable) {
	w.Objects = append(w.Objects, object)
	w.Space.AddBody(object.GetBody())
}

// Draw draws all Objects in the World
func (w *World) Draw(win pixel.Target) {
	imd := imdraw.New(nil)
	for _, object := range w.Objects {
		object.Draw(imd)
	}
	imd.Draw(win)
}

// Run runs the world simulation. This should be called from main():
//
//		w := NewWorld("test", 0, 0, 1024, 1080)
//      // add some objects
//		pixelgl.Run(w.Run)
//
func (w *World) Run() {
	cfg := pixelgl.WindowConfig{
		Title:  w.Name,
		Bounds: w.Bounds,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	if w.RunFunc != nil {
		w.RunFunc(win)
	} else {
		w.defaultRun(win)
	}
}
