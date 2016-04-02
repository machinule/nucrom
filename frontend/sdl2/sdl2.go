package sdl2

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle            = "Nuclear Romance"
	winWidth, winHeight = 800, 600
)

type sdl2Frontend struct{}

func New() (*sdl2Frontend, error) {
	return &sdl2Frontend{}, nil
}

func (fe *sdl2Frontend) Run() error {
	var window *sdl.Window
	var context sdl.GLContext
	var event sdl.Event
	var running bool
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		return err
	}
	defer window.Destroy()
	context, err = sdl.GL_CreateContext(window)
	if err != nil {
		return err
	}
	defer sdl.GL_DeleteContext(context)

	if err = gl.Init(); err != nil {
		return err
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.2, 0.2, 0.3, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, int32(winWidth), int32(winHeight))

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			}
		}
		sdl.GL_SwapWindow(window)
	}
	return nil
}
