/*
Assignment: Language Study Project - Part 2
Author: [Team Member Names]
Course: CSc 372
Instructor: L. McCann
TAs: Muhammad Bilal, Hansen Pranata, Claire Srodawa
Due Date: TBD

Description: Converts the internal maze representation into
             a text representation. Walls are displayed using
             '#' characters, while cells and open passages
             are displayed using spaces.

Language: Go (1.20 or later)
Ex. Packages: strings (standard library).
Execution: Compiled as part of the wilson-maze module.
Input: An initialized maze.

Deficiencies: No known deficiencies.
*/

package main

import "strings"

/*
Function: Render
Purpose: Converts a maze into a text representation.
Precondition: m is an initialized maze with positive dimensions

	and consistent shared walls.

Postcondition: The maze is unchanged. The returned representation

	contains walls, cell centers, and open passages.

Parameters: m (in) - maze to be displayed.
Returns: String containing the text representation of the maze.
*/
func (m *Maze) Render() string {
	height := 2*m.Rows() + 1 // Height of the text grid.
	width := 2*m.Cols() + 1  // Width of the text grid.

	grid := make([][]byte, height) // Characters representing the maze.

	// Initially fill the entire grid with wall characters.
	for r := 0; r < height; r++ { // Current text row.
		grid[r] = make([]byte, width)

		for c := 0; c < width; c++ { // Current text column.
			grid[r][c] = '#'
		}
	}

	// Open cell centers and passages between neighboring cells.
	for r := 0; r < m.Rows(); r++ { // Current maze row.
		for c := 0; c < m.Cols(); c++ { // Current maze column.
			cell := NewCell(r, c) // Cell being rendered.

			centerRow := 2*r + 1 // Text row of the cell center.
			centerCol := 2*c + 1 // Text column of the cell center.

			grid[centerRow][centerCol] = ' '

			// Each shared wall only needs to be checked once.
			if !m.HasWall(cell, east) {
				grid[centerRow][centerCol+1] = ' '
			}

			if !m.HasWall(cell, south) {
				grid[centerRow+1][centerCol] = ' '
			}
		}
	}

	lines := make([]string, height) // Individual lines of the maze.

	// Convert each character row into a string.
	for r := 0; r < height; r++ { // Current text row.
		lines[r] = string(grid[r])
	}

	return strings.Join(lines, "\n")
}
