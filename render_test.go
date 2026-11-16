/*
Assignment: Language Study Project - Part 2
Author: Aryaman Mehra & Mayur Valecha
Course: CSc 372
Instructor: L. McCann
TAs: Muhammad Bilal, Hansen Pranata, Claire Srodawa
Due Date: 11/16/2026

Description: Tests the text representation of mazes, including
             cell centers, open passages, dimensions, and
             exterior boundary walls.

Language: Go 1.27.4 (developed and tested; minimum 1.20)
Ex. Packages: testing, strings, math/rand (standard library).
Execution: go test -v

Deficiencies: Tests cover representative cases.
*/

package main

import (
	"math/rand"
	"strings"
	"testing"
)

/*
Function: TestRenderSingleCell
Purpose: Checks the representation of a single-cell maze.
Precondition: NewMaze and Render are implemented.
Postcondition: Test failures are reported if found.
Parameters: t (in/out) - Go test context.
Returns: None.
*/
func TestRenderSingleCell(t *testing.T) {
	maze := NewMaze(1, 1)       // Single-cell test maze.
	expected := "###\n# #\n###" // Expected representation.
	actual := maze.Render()     // Actual representation.

	if actual != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

/*
Function: TestRenderPassages
Purpose: Checks the positions of open horizontal and

	vertical passages.

Precondition: Maze wall operations and Render are implemented.
Postcondition: Test failures are reported if found.
Parameters: t (in/out) - Go test context.
Returns: None.
*/
func TestRenderPassages(t *testing.T) {
	maze := NewMaze(2, 2) // Maze with four cells.

	// Manually construct a known maze.
	maze.RemoveWall(NewCell(0, 0), NewCell(0, 1))
	maze.RemoveWall(NewCell(0, 1), NewCell(1, 1))
	maze.RemoveWall(NewCell(1, 1), NewCell(1, 0))

	expected := "#####\n#   #\n### #\n#   #\n#####" // Expected maze.
	actual := maze.Render()                         // Actual text representation.

	if actual != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}

/*
Function: TestRenderDimensions
Purpose: Checks dimensions and outer walls of a generated maze.
Precondition: GenerateWilson and Render are implemented.
Postcondition: Test failures are reported if found.
Parameters: t (in/out) - Go test context.
Returns: None.
*/
func TestRenderDimensions(t *testing.T) {
	maze := NewMaze(3, 5)                  // Rectangular test maze.
	random := rand.New(rand.NewSource(42)) // Fixed random source.

	GenerateWilson(maze, random)

	lines := strings.Split(maze.Render(), "\n") // Rendered rows.
	height := 2*maze.Rows() + 1                 // Expected text height.
	width := 2*maze.Cols() + 1                  // Expected text width.

	if len(lines) != height {
		t.Fatalf("Expected %d lines, got %d", height, len(lines))
	}

	for r, line := range lines { // Row index and its contents.
		if len(line) != width {
			t.Fatalf("Incorrect width on row %d", r)
		}

		if line[0] != '#' || line[width-1] != '#' {
			t.Error("Side boundary is open")
		}

		if r == 0 || r == height-1 {
			if line != strings.Repeat("#", width) {
				t.Error("Top or bottom boundary is open")
			}
		}
	}
}
