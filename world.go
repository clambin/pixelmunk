package pixelmunk

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func (w *World) Add(object Drawable) {
	w.Objects = append(w.Objects, object)
	w.Space.AddBody(object.GetBody())
}

func (w *World) Draw(win pixel.Target) {
	imd := imdraw.New(nil)
	for _, object := range w.Objects {
		object.Draw(imd)
	}
	imd.Draw(win)
}

func (w *World) Run() {
	if w.RunFunc == nil {
		panic("no run function provided")
	}

	cfg := pixelgl.WindowConfig{
		Title:  w.Name,
		Bounds: w.Bounds,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	w.RunFunc(win)
}
