package scene

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/eleniums/game-of-life-go/game"
	"github.com/eleniums/game-of-life-go/sprites"
	"github.com/eleniums/game-of-life-go/ui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	// PatternFolder is the folder where saved patterns are located.
	PatternFolder = "testdata"
)

var (
	green = color.RGBA{R: 27, G: 61, B: 0}
)

// Scene represents an entire scene in the game.
type Scene struct {
	manager     *game.Manager
	board       *ui.Board
	startButton *ui.Button
	storeButton *ui.Button
	resetButton *ui.Button
	clearButton *ui.Button
	saveButton  *ui.Button
	cell1Select *ui.Selector
	cell2Select *ui.Selector
	cell3Select *ui.Selector
	cell4Select *ui.Selector
	menuBack    *imdraw.IMDraw
	bounds      pixel.Rect
}

// New creates a new scene.
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

	scene.cell1Select = ui.NewSelector(pixel.V(1050, 120), sprites.Cell1, func(s *ui.Selector) {
		scene.cell2Select.Selected(false)
		scene.cell3Select.Selected(false)
		scene.cell4Select.Selected(false)

		ui.SetCellType = game.CellTypeCross
	})
	scene.cell1Select.Selected(true)

	scene.cell2Select = ui.NewSelector(pixel.V(1120, 120), sprites.Cell2, func(s *ui.Selector) {
		scene.cell1Select.Selected(false)
		scene.cell3Select.Selected(false)
		scene.cell4Select.Selected(false)

		ui.SetCellType = game.CellTypePlus
	})

	scene.cell3Select = ui.NewSelector(pixel.V(1050, 50), sprites.Cell3, func(s *ui.Selector) {
		scene.cell1Select.Selected(false)
		scene.cell2Select.Selected(false)
		scene.cell4Select.Selected(false)

		ui.SetCellType = game.CellTypeCircle
	})

	scene.cell4Select = ui.NewSelector(pixel.V(1120, 50), sprites.Cell4, func(s *ui.Selector) {
		scene.cell1Select.Selected(false)
		scene.cell2Select.Selected(false)
		scene.cell3Select.Selected(false)

		ui.SetCellType = game.CellTypeDot
	})

	return scene
}

// Update the scene.
func (s *Scene) Update(win *pixelgl.Window, dt float64) {
	if s.bounds.W() != win.Bounds().W() || s.bounds.H() != win.Bounds().H() {
		s.bounds = win.Bounds()

		s.menuBack = ui.NewRectangle(pixel.V(s.bounds.W()-300, 0), pixel.V(s.bounds.W(), s.bounds.H()), colornames.Black)

		s.startButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.startButton.Size().W()/2, s.bounds.Max.Y-360))
		s.storeButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.storeButton.Size().W()/2, s.bounds.Max.Y-460))
		s.resetButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.resetButton.Size().W()/2, s.bounds.Max.Y-560))
		s.clearButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.clearButton.Size().W()/2, s.bounds.Max.Y-660))
		s.saveButton.SetPosition(pixel.V(s.bounds.Max.X-150-s.saveButton.Size().W()/2, s.bounds.Max.Y-760))
	}

	s.board.Update(win, dt, s.manager.Running(), s.manager.Cells())

	s.startButton.Update(win)
	s.storeButton.Update(win)
	s.resetButton.Update(win)
	s.clearButton.Update(win)
	s.saveButton.Update(win)

	s.cell1Select.Update(win)
	s.cell2Select.Update(win)
	s.cell3Select.Update(win)
	s.cell4Select.Update(win)

	s.manager.Update()
}

// Draw the scene.
func (s *Scene) Draw(win *pixelgl.Window) {
	win.Clear(green)

	// board
	s.board.Draw(win, s.manager.Cells())

	// menu
	s.menuBack.Draw(win)
	sprites.Title.Draw(win, pixel.IM.Moved(pixel.V(win.Bounds().Max.X-sprites.Title.Frame().W()/2, win.Bounds().Max.Y-sprites.Title.Frame().H()/2)))
	s.startButton.Draw(win)
	s.storeButton.Draw(win)
	s.resetButton.Draw(win)
	s.clearButton.Draw(win)
	s.saveButton.Draw(win)

	s.cell1Select.Draw(win)
	s.cell2Select.Draw(win)
	s.cell3Select.Draw(win)
	s.cell4Select.Draw(win)
}

// Save the scene to file.
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

// Load a pattern from file.
func (s *Scene) Load(pattern string) {
	err := s.manager.Load(fmt.Sprintf("%s/%s", PatternFolder, pattern))
	if err != nil {
		log.Printf("error loading pattern: %v", err)
	}
}
