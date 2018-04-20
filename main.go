package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/eleniums/game-of-life-go/assets"
	"github.com/eleniums/game-of-life-go/game"
	"github.com/eleniums/game-of-life-go/scene"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	// pixel will run on the main thread
	pixelgl.Run(run)
}

func run() {
	interval := flag.Int("interval", 100, "simulation update interval in ms")
	resizable := flag.Bool("resizable", false, "allow resizing of the main window")
	flag.Parse()

	// load all assets and sprites
	err := assets.Load()
	if err != nil {
		log.Fatalf("unable to load assets: %v", err)
	}
	sprites.Load()

	// create new window
	cfg := pixelgl.WindowConfig{
		Title:     "Game of Life",
		Icon:      []pixel.Picture{assets.Icon16x16},
		Bounds:    pixel.R(0, 0, 1260, 960),
		VSync:     true, // update at the refresh rate of the monitor
		Resizable: *resizable,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatalf("unable to create window: %v", err)
	}
	win.SetSmooth(true) // remove pixelation

	// create new scene containing all elements of the game
	game.Interval = *interval
	scene := scene.New()

	frames := 0
	second := time.Tick(time.Second)

	// main update loop
	for !win.Closed() {
		// update all objects in the scene
		scene.Update(win)

		// draw all objects in the scene
		scene.Draw(win)

		// swap buffers and poll for events
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
