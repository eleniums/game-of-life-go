package main

import (
	_ "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "github.com/faiface/pixel/text"
)

func main() {
	// pixel will run on the main thread
	pixelgl.Run(run)
}

func run() {
	// TODO
}
