package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
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
	disableVsync := flag.Bool("disable-vsync", false, "disable vertical sync with refresh rate of monitor")
	pattern := flag.String("pattern", "", "name of pattern file in testdata to load as initial state (ex: -pattern diehard)")
	reproduce := flag.String("reproduce", "majority-wins", "how to determine cell type when cell becomes alive (majority-wins|random-percentage)")
	flag.Parse()

	// seed random numbers
	rand.Seed(time.Now().UTC().UnixNano())

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
		VSync:     !*disableVsync, // update at the refresh rate of the monitor
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

	// load pattern if requested
	if *pattern != "" {
		scene.Load(*pattern)
	}

	// set reproduction type
	switch *reproduce {
	case "majority-wins":
		game.ReproduceMethod = game.ReproduceTypeMajorityWins
	case "random-percentage":
		game.ReproduceMethod = game.ReproduceTypeRandomPercentage
	default:
		game.ReproduceMethod = game.ReproduceTypeMajorityWins
	}

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
