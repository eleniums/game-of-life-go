package main

import (
	"log"

	"github.com/eleniums/game-of-life-go/assets"
	_ "github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "github.com/faiface/pixel/text"
)

func main() {
	// pixel will run on the main thread
	pixelgl.Run(run)
}

func run() {
	err := assets.Load()
	if err != nil {
		log.Fatalf("unable to load assets: %v", err)
	}
}
