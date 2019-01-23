// fifteen - 15-puzzle solver.
// Copyright (C) 2019  Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package fifteen

import (
	"errors"
)

// Puzzle is the base container for 15-puzzles.
type Puzzle struct {
	data  []int
	n     int
	blank Position
}

// NewPuzzle creates a new blank puzzle.
func NewPuzzle(n int) *Puzzle {
	return &Puzzle{
		data:  make([]int, n*n),
		n:     n,
	}
}

// NewSolvedPuzzle creates a new puzzle in the finished form.
func NewSolvedPuzzle(n int) *Puzzle {
	puzzle := NewPuzzle(n)
	for i := 1; i < len(puzzle.data); i++ {
		puzzle.data[i-1] = i
	}
	puzzle.blank.X = 4
	puzzle.blank.Y = 4
	return puzzle
}

// Copy creates a copy of this puzzle.
func (puzzle *Puzzle) Copy() *Puzzle {
	newPuzzle := NewPuzzle(puzzle.n)
	copy(newPuzzle.data, puzzle.data)
	newPuzzle.blank.X = puzzle.blank.X
	newPuzzle.blank.Y = puzzle.blank.Y
	return newPuzzle
}

// Get gets the value of a specific slot.
func (puzzle *Puzzle) Get(x, y int) int {
	if x <= 0 || y <= 0 || x > puzzle.n || y > puzzle.n {
		return -1
	}
	return puzzle.data[puzzle.Index(x, y)]
}

// Set sets the value of a specific slot.
func (puzzle *Puzzle) Set(x, y, val int) {
	if x <= 0 || y <= 0 || x > puzzle.n || y > puzzle.n {
		return
	}
	if val == 0 {
		puzzle.blank = Position{x, y}
	}
	puzzle.data[puzzle.Index(x, y)] = val
}

// Index returns the index of the given coordinates in the puzzle.
func (puzzle *Puzzle) Index(x, y int) int {
	return (y-1)*puzzle.n + (x - 1)
}

// Coordinates returns the X and Y coordinates of the given index in the puzzle.
func (puzzle *Puzzle) Coordinates(index int) (x, y int) {
	y = int(index/puzzle.n) + 1
	x = index % puzzle.n
	return
}

// Move moves the piece at the given coordinates to the empty slot next to it.
// Return false if there is no empty slot next to the given coordinates or if the coordinates are invalid.
// Returns true if the move was successful.
func (puzzle *Puzzle) Move(x, y int) bool {
	val := puzzle.Get(x, y)
	switch {
	case x <= 0 || y <= 0 || x > puzzle.n || y > puzzle.n:
		return false
	case puzzle.Get(x-1, y) == 0:
		puzzle.Set(x-1, y, val)
	case puzzle.Get(x+1, y) == 0:
		puzzle.Set(x+1, y, val)
	case puzzle.Get(x, y-1) == 0:
		puzzle.Set(x, y-1, val)
	case puzzle.Get(x, y+1) == 0:
		puzzle.Set(x, y+1, val)
	default:
		return false
	}
	puzzle.Set(x, y, 0)
	puzzle.blank.X = x
	puzzle.blank.Y = y
	return true
}

func (puzzle *Puzzle) SetData(data [][]int) error {
	if len(data) != puzzle.n {
		return errors.New("invalid input height")
	}
	newData := make([]int, puzzle.n*puzzle.n)
	for i := 0; i < puzzle.n; i++ {
		if len(data[i]) != puzzle.n {
			return errors.New("invalid input width")
		}
		copy(newData[i*puzzle.n : (i+1)*puzzle.n], data[i])
	}
	puzzle.data = newData
	return nil
}

func (puzzle *Puzzle) Data() [][]int {
	data := make([][]int, puzzle.n)
	for i := 0; i < puzzle.n; i++ {
		data[i] = puzzle.data[i*puzzle.n : (i+1)*puzzle.n]
	}
	return data
}

func (puzzle *Puzzle) Size() int {
	return puzzle.n
}

func (puzzle *Puzzle) GetValidMoves() []Position {
	return puzzle.blank.ValidMoves(puzzle.n)
}
