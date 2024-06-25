package pixelmunk

import (
	"fmt"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"time"
)

// World represents the world that chipmunk will simulate
type World struct {
	Name        string
	Bounds      pixel.Rect
	FrameRate   int
	Space       *chipmunk.Space
	RunFunc     func(*opengl.Window)
	RunCallback func(*opengl.Window)
	Objects     []Drawable
}

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
func (w *World) defaultRun(win *opengl.Window) {
	frameTicker := time.NewTicker(time.Second / time.Duration(w.FrameRate))
	timer := time.Now()

	for !win.Closed() {
		w.Space.Step(1.0 / vect.Float(w.FrameRate))

		win.Clear(colornames.Black)
		w.Draw(win)
		win.Update()

		if w.RunCallback != nil {
			w.RunCallback(win)
		}

		win.SetTitle(fmt.Sprintf("%s (%.1f fps)", w.Name, 1/time.Since(timer).Seconds()))
		timer = time.Now()

		<-frameTicker.C
	}
}

// Add adds a new Object to the World
func (w *World) Add(objects ...Drawable) {
	for _, object := range objects {
		w.Objects = append(w.Objects, object)
		switch object.GetType() {
		case DrawableBody:
			w.Space.AddBody(object.GetBody())
		case DrawableJoint:
			w.Space.AddConstraint(object.GetJoint())
		}
	}
}

// Remove removes an Object from the World
func (w *World) Remove(objects ...Drawable) {
	for _, object := range objects {
		if object.GetType() == DrawableBody {
			w.Space.RemoveBody(object.GetBody())
		} else {
			w.Space.RemoveConstraint(object.GetJoint())
		}
		for index, o := range w.Objects {
			if o == object {
				w.Objects = append(w.Objects[:index], w.Objects[index+1:]...)
			}
		}
	}
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
//			w := NewWorld("test", 0, 0, 1024, 1080)
//	     // add some objects
//			opengl.Run(w.Run)
func (w *World) Run() {
	cfg := opengl.WindowConfig{
		Title:  w.Name,
		Bounds: w.Bounds,
	}

	win, err := opengl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	if w.RunFunc != nil {
		w.RunFunc(win)
	} else {
		w.defaultRun(win)
	}
}
