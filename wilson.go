/*
Assignment: Language Study Project - Part 2
Author: Aryaman Mehra
Course: CSc 372
Instructor: L. McCann
TAs: Muhammad Bilal, Hansen Pranata, Claire Srodawa
Due Date: TBD

Description: Generates a random maze using Wilson's algorithm.
             Selects an initial maze cell and repeatedly performs
             loop-erased random walks from unvisited cells.
             Each completed walk joins the existing maze until
             every cell belongs to a single spanning tree.

Language: Go (1.20 or later)
Ex. Packages: math/rand (standard library).
Execution: Compiled as part of the wilson-maze module.
Input: An initialized maze and random number generator.

Deficiencies: No known deficiencies in this module.
*/

package main

import "math/rand"

/*
Function: GenerateWilson
Purpose: Generates a maze using loop-erased random walks.
Precondition: m is an initialized maze with positive dimensions

	and all walls closed; random is not nil.

Postcondition: All cells are connected without cycles.
Parameters: m (in/out) - maze whose walls are modified.

	random (in/out) - random number generator.

Returns: None.
*/
func GenerateWilson(m *Maze, random *rand.Rand) {
	totalCells := m.Rows() * m.Cols() // Number of maze cells.
	visited := make(map[Cell]bool)    // Permanent maze membership.

	// Randomly select the initial cell of the maze.
	root := NewCell(
		random.Intn(m.Rows()),
		random.Intn(m.Cols()),
	) // Initial maze location.

	visited[root] = true

	// Add loop-free paths until every cell joins the maze.
	for len(visited) < totalCells {
		start := chooseStart(m, visited, random)      // Walk origin.
		path := randomWalk(m, start, visited, random) // Surviving walk.

		connectPath(m, path, visited)
	}
}

/*
Function: chooseStart
Purpose: Randomly selects a cell outside the existing maze.
Precondition: At least one cell is not visited.
Postcondition: The maze and visited map are unchanged.
Parameters: m (in) - maze containing all possible cells.

	visited (in) - cells already in the maze.
	random (in/out) - random number generator.

Returns: A randomly selected unvisited cell.
*/
func chooseStart(m *Maze, visited map[Cell]bool,
	random *rand.Rand) Cell {

	unvisited := make([]Cell, 0) // Available starting cells.

	// Collect all cells outside the permanent maze.
	for r := 0; r < m.Rows(); r++ { // Current row.
		for c := 0; c < m.Cols(); c++ { // Current column.
			cell := NewCell(r, c) // Candidate starting cell.

			if !visited[cell] {
				unvisited = append(unvisited, cell)
			}
		}
	}

	// Give every unvisited cell an equal chance of selection.
	return unvisited[random.Intn(len(unvisited))]
}

/*
Function: randomWalk
Purpose: Builds a loop-free random path to the existing maze.
Precondition: start is a valid unvisited cell; visited contains

	at least one cell; random is not nil.

Postcondition: The returned path has no repeated cells and

	ends at a cell already in the maze.
	The permanent maze and visited map are unchanged.

Parameters: m (in) - maze containing legal neighbors.

	start (in) - starting cell of the random walk.
	visited (in) - cells already in the maze.
	random (in/out) - random number generator.

Returns: Loop-erased path including the existing maze endpoint.
*/
func randomWalk(m *Maze, start Cell, visited map[Cell]bool,
	random *rand.Rand) []Cell {

	path := []Cell{start}           // Temporary walk.
	positions := make(map[Cell]int) // Cell indices in the walk.
	current := start                // Current walking location.

	positions[start] = 0

	// Continue until the walk reaches the permanent maze.
	for !visited[current] {
		neighbors := m.Neighbors(current)              // Legal next cells.
		next := neighbors[random.Intn(len(neighbors))] // Next step.

		// If the next cell is already on the current path,
		// erase all cells following its earlier occurrence.
		index, exists := positions[next] // Previous index and presence.

		if exists {
			for i := index + 1; i < len(path); i++ { // Removed index.
				delete(positions, path[i])
			}

			path = path[:index+1]
		} else {
			// Extend the walk when no loop has been formed.
			positions[next] = len(path)
			path = append(path, next)
		}

		current = next
	}

	return path
}

/*
Function: connectPath
Purpose: Permanently adds a completed random walk to the maze.
Precondition: path is loop-free, contains adjacent cells, and

	ends at a visited cell. All earlier cells are
	unvisited.

Postcondition: Walls between consecutive cells are removed,

	and all path cells belong to the maze.

Parameters: m (in/out) - maze whose walls are modified.

	path (in) - completed loop-erased random walk.
	visited (in/out) - permanent maze membership.

Returns: None.
*/
func connectPath(m *Maze, path []Cell, visited map[Cell]bool) {
	// Open the shared wall between each consecutive pair.
	for i := 0; i < len(path)-1; i++ { // Current path index.
		m.RemoveWall(path[i], path[i+1])
	}

	// Mark the completed path as part of the permanent maze.
	for _, cell := range path { // Cell being added.
		visited[cell] = true
	}
}
