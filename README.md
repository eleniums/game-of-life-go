# game-of-life-go

[![Build Status](https://travis-ci.org/eleniums/game-of-life-go.svg?branch=master)](https://travis-ci.org/eleniums/game-of-life-go) [![Go Report Card](https://goreportcard.com/badge/github.com/eleniums/game-of-life-go)](https://goreportcard.com/report/github.com/eleniums/game-of-life-go) [![GoDoc](https://godoc.org/github.com/eleniums/game-of-life-go?status.svg)](https://godoc.org/github.com/eleniums/game-of-life-go) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/eleniums/game-of-life-go/blob/master/LICENSE)

An adaptation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) created with Go for a hackathon at work.

*To see the code as it was when the hackathon finished, see [Release 1.0.0](https://github.com/eleniums/game-of-life-go/releases/tag/v1.0.0).*

## Installation

```
go get -u github.com/eleniums/game-of-life-go
```

## Run

```
go run ./cmd/game/main.go
```

## Prerequisites

- Requires Go 1.8 or later
- Requires GCC (follow instructions [here](https://github.com/faiface/pixel/wiki/Building-Pixel-on-Windows) to install [MSYS2](http://www.msys2.org) and [MinGW](http://www.mingw.org))
- Uses [dep](https://github.com/golang/dep) for dependencies
- Uses [pixel](https://github.com/faiface/pixel) for graphics and input

## Explanation

Conway's Game of Life is a zero player game, meaning the player sets an initial state and then sits back as the simulation is run. The simulation consists of a grid of cells that exist in one of two states, alive or dead. These cells are governed by 4 rules, based on the 8 neighboring cells surrounding any given cell:

- If there are less than 2 living cells surrounding a living cell, it will die, as if by underpopulation.
- If there are 2 or 3 living cells surrounding a living cell, it will continue to live.
- If there are more than 3 living cells surrounding a living cell, it will die, as if by overpopulation.
- If there are exactly 3 living cells surrounding a dead cell, it will become alive, as if by reproduction.

This leads to many interesting patterns and is fascinating to mess around with. Interestingly enough, it is also Turing complete, meaning it can simulate a computer. The Game of Life itself has been built with the Game of Life!

## Instructions

- Input
    - Right mouse button: Add a new cell to the grid.
    - Left mouse button: Remove an existing cell from the grid.
    - Arrow keys: Scroll around the grid.
    - Spacebar: Reset the scroll position to the default.
- Menu
    - Start/Stop: Start or stop the simulation.
    - Store: Store current grid state in memory.
    - Reset: Reset the grid to the state saved in memory.
    - Clear: Clear the grid completely.
    - Save: Save the current grid state to a file named "saved".
    - Cells: Select the cell type to place on the grid.