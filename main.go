package main

import (
	"github.com/CyrusRoshan/pickture/ui"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(ui.Render)
}
