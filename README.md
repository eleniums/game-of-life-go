# game-of-life-go

An adaptation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) created with Go for a hackathon at work.

## Installation

```
go get -u github.com/eleniums/game-of-life-go
dep ensure
```

## Run

```
go run main.go
```

## Prerequisites

- Requires Go 1.8 or later
- Uses [dep](https://github.com/golang/dep) for dependencies
- Uses [pixel](https://github.com/faiface/pixel) for graphics and input

## Explanation

Conway's Game of Life is a zero player game, meaning the player sets an initial state and then sits back as the simulation is run. The simulation consists of a grid of cells that exist in one of two states, alive or dead. These cells are governed by 4 rules, based on the 8 neighboring cells surrounding any given cell:

- If there are less than 2 living cells surrounding a living cell, it will die, as if by underpopulation.
- If there are 2 or 3 living cells surrounding a living cell, it will continue to live.
- If there are more than 3 living cells surrounding a living cell, it will die, as if by overpopulation.
- If there are exactly 3 living cells surrounding a dead cell, it will become alive, as if by reproduction.

This leads to many interesting patterns and is fascinating to mess around with. Interestingly enough, it is also Turing complete, meaning it can simulate a computer. The Game of Life itself has been built with the Game of Life!