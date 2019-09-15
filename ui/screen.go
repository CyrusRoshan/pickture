package ui

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func ScreenBounds() (width float64, height float64) {
	return pixelgl.PrimaryMonitor().Size()
}

func LimitFPS(fps int, loopContents func()) {
	d := time.Second / time.Duration(fps)

	select {
	case <-time.After(d):
		loopContents()
	}
}

var (
	frames = 0
	second = time.Tick(time.Second)
)

func ShowFPSInTitle(title string, window *pixelgl.Window) {
	frames++
	select {
	case <-second:
		window.SetTitle(fmt.Sprintf("%s | FPS: %d", title, frames))
		frames = 0
	default:
	}
}

func BuildWindow(title string) (*pixelgl.Window, error) {
	width, height := ScreenBounds()
	cfg := pixelgl.WindowConfig{
		Title:     title,
		Bounds:    pixel.R(0, 0, width/2, height/2),
		VSync:     true,
		Resizable: true,
	}

	var err error
	window, err = pixelgl.NewWindow(cfg)
	return window, err
}
