/*
Assignment: Language Study Project - Part 2
Author: Aryaman Mehra & Mayur Valecha
Course: CSc 372
Instructor: L. McCann
TAs: Muhammad Bilal, Hansen Pranata, Claire Srodawa
Due Date: 11/16/2026

Description: Provides the internal representation of a maze.
             Stores the grid dimensions and the four walls of
             each cell. Provides methods to find legal neighbors,
             inspect walls, and connect adjacent cells.

Language: Go 1.27.4 (developed and tested; minimum 1.20)
Ex. Packages: None.
Execution: Compiled as part of the wilson-maze module.
Input: None directly.

Deficiencies: No known deficiencies in this module.
*/

package main

const (
	north = 0 // Wall above a cell.
	east  = 1 // Wall to the right of a cell.
	south = 2 // Wall below a cell.
	west  = 3 // Wall to the left of a cell.
)

/*
Struct: Cell
Package: main
Dependencies: None.
Inheritance: None.

Purpose: Represents a location in the maze using zero-based

	row and column coordinates.

Public Fields: None.
Constructor: NewCell - creates a cell coordinate.
Methods: Row, Col.
*/
type Cell struct {
	row int // Row coordinate.
	col int // Column coordinate.
}

/*
Struct: Maze
Package: main
Dependencies: None.
Inheritance: None.

Purpose: Stores the dimensions and walls of a rectangular maze.

Public Fields: None.
Constructor: NewMaze - creates a maze with all walls closed.
Methods: Rows, Cols, Neighbors, RemoveWall, HasWall.
*/
type Maze struct {
	rows  int         // Number of rows.
	cols  int         // Number of columns.
	walls [][][4]bool // Four walls for each cell.
}

/*
Function: NewCell
Purpose: Creates a cell with the given coordinates.
Precondition: Coordinates are non-negative.
Postcondition: A cell containing the coordinates is created.
Parameters: row (in) - row coordinate.

	col (in) - column coordinate.

Returns: The created Cell.
*/
func NewCell(row, col int) Cell {
	return Cell{row: row, col: col}
}

/*
Methods: Row, Col
Purpose: Return the coordinates of a cell.
Precondition: The receiver is an initialized Cell.
Postcondition: The receiver is unchanged.
Parameters: c (in) - cell containing the coordinates.
Returns: The requested coordinate.
*/
func (c Cell) Row() int {
	return c.row
}

func (c Cell) Col() int {
	return c.col
}

/*
Function: NewMaze
Purpose: Creates a rectangular maze with all walls closed.
Precondition: rows and cols are positive integers.
Postcondition: All cells have four closed walls.
Parameters: rows (in) - number of rows.

	cols (in) - number of columns.

Returns: Pointer to the initialized Maze.
*/
func NewMaze(rows, cols int) *Maze {
	// Allocate the maze and its rows.
	maze := &Maze{
		rows:  rows,
		cols:  cols,
		walls: make([][][4]bool, rows),
	}

	// Allocate each row and close every cell's walls.
	for r := 0; r < rows; r++ { // Current row.
		maze.walls[r] = make([][4]bool, cols)

		for c := 0; c < cols; c++ { // Current column.
			maze.walls[r][c] = [4]bool{true, true, true, true}
		}
	}

	return maze
}

/*
Methods: Rows, Cols
Purpose: Return the dimensions of the maze.
Precondition: The receiver is an initialized Maze.
Postcondition: The maze is unchanged.
Parameters: m (in) - maze containing the dimensions.
Returns: The requested dimension.
*/
func (m *Maze) Rows() int {
	return m.rows
}

func (m *Maze) Cols() int {
	return m.cols
}

/*
Function: Neighbors
Purpose: Finds all orthogonally adjacent cells within the grid.
Precondition: cell is a valid location in the maze.
Postcondition: The maze is unchanged.
Parameters: m (in) - maze containing the cell.

	cell (in) - location whose neighbors are needed.

Returns: Slice of valid neighboring cells.
*/
func (m *Maze) Neighbors(cell Cell) []Cell {
	neighbors := make([]Cell, 0, 4) // Legal adjacent cells.

	// Check each direction against the maze boundaries.
	if cell.row > 0 {
		neighbors = append(neighbors,
			NewCell(cell.row-1, cell.col))
	}

	if cell.col < m.cols-1 {
		neighbors = append(neighbors,
			NewCell(cell.row, cell.col+1))
	}

	if cell.row < m.rows-1 {
		neighbors = append(neighbors,
			NewCell(cell.row+1, cell.col))
	}

	if cell.col > 0 {
		neighbors = append(neighbors,
			NewCell(cell.row, cell.col-1))
	}

	return neighbors
}

/*
Function: RemoveWall
Purpose: Opens the shared wall between two adjacent cells.
Precondition: Both cells are valid and orthogonally adjacent.
Postcondition: Both sides of the shared wall are removed.
Parameters: m (in/out) - maze whose walls are modified.

	first (in) - first adjacent cell.
	second (in) - second adjacent cell.

Returns: None.
*/
func (m *Maze) RemoveWall(first, second Cell) {
	rowChange := second.row - first.row // Vertical difference.
	colChange := second.col - first.col // Horizontal difference.

	// Remove the corresponding wall from both cells.
	switch {
	case rowChange == -1 && colChange == 0:
		m.walls[first.row][first.col][north] = false
		m.walls[second.row][second.col][south] = false

	case rowChange == 0 && colChange == 1:
		m.walls[first.row][first.col][east] = false
		m.walls[second.row][second.col][west] = false

	case rowChange == 1 && colChange == 0:
		m.walls[first.row][first.col][south] = false
		m.walls[second.row][second.col][north] = false

	case rowChange == 0 && colChange == -1:
		m.walls[first.row][first.col][west] = false
		m.walls[second.row][second.col][east] = false
	}
}

/*
Function: HasWall
Purpose: Determines whether a specified cell wall exists.
Precondition: cell is valid; direction is between 0 and 3.
Postcondition: The maze is unchanged.
Parameters: m (in) - maze containing the wall.

	cell (in) - location of the wall.
	direction (in) - direction of the desired wall.

Returns: True if the wall exists, otherwise false.
*/
func (m *Maze) HasWall(cell Cell, direction int) bool {
	return m.walls[cell.row][cell.col][direction]
}
