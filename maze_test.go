/*
Assignment: Language Study Project - Part 2
Author: Aryaman Mehra & Mayur Valecha
Course: CSc 372
Instructor: L. McCann
TAs: Muhammad Bilal, Hansen Pranata, Claire Srodawa
Due Date: 11/16/2026

Description: Tests maze initialization, neighbor selection,
             and removal of shared walls.

Language: Go 1.27.4 (developed and tested; minimum 1.20)
Ex. Packages: testing (standard library).
Execution: go test -v

Deficiencies: Tests cover representative cases.
*/

package main

import "testing"

/*
Function: TestNewMaze
Purpose: Checks dimensions and initial wall values.
Precondition: NewMaze is implemented.
Postcondition: Test failures are reported if found.
Parameters: t (in/out) - Go test context.
Returns: None.
*/
func TestNewMaze(t *testing.T) {
	maze := NewMaze(3, 4) // Rectangular test maze.

	if maze.Rows() != 3 || maze.Cols() != 4 {
		t.Fatal("Incorrect maze dimensions")
	}

	for r := 0; r < maze.Rows(); r++ { // Current row.
		for c := 0; c < maze.Cols(); c++ { // Current column.
			cell := NewCell(r, c) // Cell being inspected.

			for d := 0; d < 4; d++ { // Wall direction.
				if !maze.HasWall(cell, d) {
					t.Fatal("A wall is initially open")
				}
			}
		}
	}
}

/*
Function: TestNeighbors
Purpose: Checks neighbors of corner, edge, and interior cells.
Precondition: Neighbors is implemented.
Postcondition: Test failures are reported if found.
Parameters: t (in/out) - Go test context.
Returns: None.
*/
func TestNeighbors(t *testing.T) {
	maze := NewMaze(3, 3) // Test maze.

	if len(maze.Neighbors(NewCell(0, 0))) != 2 {
		t.Error("Corner should have two neighbors")
	}

	if len(maze.Neighbors(NewCell(0, 1))) != 3 {
		t.Error("Edge should have three neighbors")
	}

	if len(maze.Neighbors(NewCell(1, 1))) != 4 {
		t.Error("Interior should have four neighbors")
	}

	single := NewMaze(1, 1) // Maze with no adjacent cells.

	if len(single.Neighbors(NewCell(0, 0))) != 0 {
		t.Error("Single cell should have no neighbors")
	}
}

/*
Function: TestRemoveWall
Purpose: Checks that shared walls are opened on both sides.
Precondition: RemoveWall and HasWall are implemented.
Postcondition: Test failures are reported if found.
Parameters: t (in/out) - Go test context.
Returns: None.
*/
func TestRemoveWall(t *testing.T) {
	maze := NewMaze(2, 2)  // Test maze.
	first := NewCell(0, 0) // Starting cell.
	right := NewCell(0, 1) // Cell east of first.
	below := NewCell(1, 0) // Cell south of first.

	maze.RemoveWall(first, right)

	if maze.HasWall(first, east) || maze.HasWall(right, west) {
		t.Error("Horizontal passage was not opened correctly")
	}

	maze.RemoveWall(first, below)

	if maze.HasWall(first, south) || maze.HasWall(below, north) {
		t.Error("Vertical passage was not opened correctly")
	}

	if !maze.HasWall(first, north) || !maze.HasWall(first, west) {
		t.Error("An outer boundary wall was removed")
	}
}
