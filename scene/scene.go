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
	board       *ui.Board
	cells       game.CellGrid
	startButton *ui.Button
	storeButton *ui.Button
	resetButton *ui.Button
	clearButton *ui.Button

	stopped bool
}

func New() *Scene {
	scene := &Scene{}

	scene.board = ui.NewBoard()
	scene.cells = game.NewCellGrid(game.GridMaxX, game.GridMaxY)

	scene.stopped = true
	scene.startButton = ui.NewButton(pixel.V(1000, 600), "START", func(b *ui.Button) {
		if scene.stopped {
			b.SetText("STOP")
		} else {
			b.SetText("START")
		}
		scene.stopped = !scene.stopped
	})

	scene.storeButton = ui.NewButton(pixel.V(1000, 500), "STORE", nil)
	scene.resetButton = ui.NewButton(pixel.V(1000, 400), "RESET", nil)
	scene.clearButton = ui.NewButton(pixel.V(1000, 300), "CLEAR", nil)

	return scene
}

func (s *Scene) Update(win *pixelgl.Window) {
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
	s.board.Draw(win, s.cells)
}
