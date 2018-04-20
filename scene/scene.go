package scene

import (
	"github.com/eleniums/game-of-life-go/game"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/eleniums/game-of-life-go/ui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Scene struct {
	manager     *game.Manager
	board       *ui.Board
	startButton *ui.Button
	storeButton *ui.Button
	resetButton *ui.Button
	clearButton *ui.Button
}

func New() *Scene {
	manager := game.NewManager()
	board := ui.NewBoard()

	storeButton := ui.NewButton(pixel.V(1000, 500), "STORE", nil)
	resetButton := ui.NewButton(pixel.V(1000, 400), "RESET", nil)
	clearButton := ui.NewButton(pixel.V(1000, 300), "CLEAR", nil)

	startButton := ui.NewButton(pixel.V(1000, 600), "START", func(b *ui.Button) {
		if !manager.Running() {
			b.SetText("STOP")
			storeButton.SetActive(false)
			resetButton.SetActive(false)
			clearButton.SetActive(false)
		} else {
			b.SetText("START")
			storeButton.SetActive(true)
			resetButton.SetActive(true)
			clearButton.SetActive(true)
		}
	})

	return &Scene{
		manager:     manager,
		board:       board,
		startButton: startButton,
		storeButton: storeButton,
		resetButton: resetButton,
		clearButton: clearButton,
	}
}

func (s *Scene) Update(win *pixelgl.Window) {
	s.board.Update(win, s.manager.Cells())

	s.startButton.Update(win)
	s.storeButton.Update(win)
	s.resetButton.Update(win)
	s.clearButton.Update(win)
}

func (s *Scene) Draw(win *pixelgl.Window) {
	win.Clear(colornames.Black)

	// menu
	sprites.Title.Draw(win, pixel.IM.Moved(pixel.V(win.Bounds().Max.X-sprites.Title.Frame().W()/2, win.Bounds().Max.Y-sprites.Title.Frame().H()/2)))
	s.startButton.Draw(win)
	s.storeButton.Draw(win)
	s.resetButton.Draw(win)
	s.clearButton.Draw(win)

	// board
	s.board.Draw(win, s.manager.Cells())
}
