package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/eleniums/game-of-life-go/assets"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	// pixel will run on the main thread
	pixelgl.Run(run)
}

func run() {
	//interval := flag.Int("interval", 1000, "simulation update interval in ms")
	resizable := flag.Bool("resizable", false, "allow resizing of the main window")
	flag.Parse()

	// load all assets
	err := assets.Load()
	if err != nil {
		log.Fatalf("unable to load assets: %v", err)
	}

	sprites.Load()

	// create new window
	cfg := pixelgl.WindowConfig{
		Title:     "Game of Life",
		Bounds:    pixel.R(0, 0, 1260, 960),
		VSync:     true, // update at the refresh rate of the monitor
		Resizable: *resizable,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatalf("unable to create window: %v", err)
	}
	win.SetSmooth(true) // remove pixelation

	frames := 0
	second := time.Tick(time.Second)

	// main update loop
	//last := time.Now()
	for !win.Closed() {
		//dt := time.Since(last).Seconds()
		//last = time.Now()

		win.Clear(colornames.Black)

		sprites.Title.Draw(win, pixel.IM.Moved(pixel.V(1100, 800)))

		win.Update()

		// calculate FPS
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}
