package scene

import (
	"fmt"
	"log"
	"os"

	"github.com/eleniums/game-of-life-go/game"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/eleniums/game-of-life-go/ui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	PatternFolder = "testdata"
)

type Scene struct {
	manager     *game.Manager
	board       *ui.Board
	startButton *ui.Button
	storeButton *ui.Button
	resetButton *ui.Button
	clearButton *ui.Button
	saveButton  *ui.Button
	bounds      pixel.Rect
}

func New() *Scene {
	scene := &Scene{}

	scene.manager = game.NewManager()
	scene.board = ui.NewBoard()

	scene.storeButton = ui.NewButton(pixel.V(1000, 500), "STORE", func(b *ui.Button) {
		scene.manager.Store()
	})

	scene.resetButton = ui.NewButton(pixel.V(1000, 400), "RESET", func(b *ui.Button) {
		scene.manager.Reset()
	})

	scene.clearButton = ui.NewButton(pixel.V(1000, 300), "CLEAR", func(b *ui.Button) {
		scene.manager.Clear()
	})

	scene.saveButton = ui.NewButton(pixel.V(1000, 200), "SAVE", func(b *ui.Button) {
		scene.Save("saved")
	})

	scene.startButton = ui.NewButton(pixel.V(1000, 600), "START", func(b *ui.Button) {
		if !scene.manager.Running() {
			b.SetText("STOP")
			scene.storeButton.SetActive(false)
			scene.resetButton.SetActive(false)
			scene.clearButton.SetActive(false)
			scene.saveButton.SetActive(false)
			scene.manager.Start()
		} else {
			b.SetText("START")
			scene.storeButton.SetActive(true)
			scene.resetButton.SetActive(true)
			scene.clearButton.SetActive(true)
			scene.saveButton.SetActive(true)
			scene.manager.Stop()
		}
	})

	return scene
}

func (s *Scene) Update(win *pixelgl.Window) {
	if s.bounds.W() != win.Bounds().W() || s.bounds.H() != win.Bounds().H() {
		s.bounds = win.Bounds()

		s.startButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.startButton.Size().W()/2, s.bounds.Max.Y-360))
		s.storeButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.storeButton.Size().W()/2, s.bounds.Max.Y-460))
		s.resetButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.resetButton.Size().W()/2, s.bounds.Max.Y-560))
		s.clearButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.clearButton.Size().W()/2, s.bounds.Max.Y-660))
		s.saveButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.saveButton.Size().W()/2, s.bounds.Max.Y-760))
	}

	s.board.Update(win, s.manager.Cells())

	s.startButton.Update(win)
	s.storeButton.Update(win)
	s.resetButton.Update(win)
	s.clearButton.Update(win)
	s.saveButton.Update(win)

	s.manager.Update()
}

func (s *Scene) Draw(win *pixelgl.Window) {
	win.Clear(colornames.Black)

	// menu
	sprites.Title.Draw(win, pixel.IM.Moved(pixel.V(win.Bounds().Max.X-sprites.Title.Frame().W()/2, win.Bounds().Max.Y-sprites.Title.Frame().H()/2)))
	s.startButton.Draw(win)
	s.storeButton.Draw(win)
	s.resetButton.Draw(win)
	s.clearButton.Draw(win)
	s.saveButton.Draw(win)

	// board
	s.board.Draw(win, s.manager.Cells())
}

func (s *Scene) Save(pattern string) {
	if _, err := os.Stat(PatternFolder); os.IsNotExist(err) {
		err := os.Mkdir(PatternFolder, os.ModePerm)
		if err != nil {
			log.Printf("error creating pattern directory: %v", err)
			return
		}
	}

	err := s.manager.Save(fmt.Sprintf("%s/%s", PatternFolder, pattern))
	if err != nil {
		log.Printf("error saving pattern: %v", err)
	}
}

func (s *Scene) Load(pattern string) {
	err := s.manager.Load(fmt.Sprintf("%s/%s", PatternFolder, pattern))
	if err != nil {
		log.Printf("error loading pattern: %v", err)
	}
}
