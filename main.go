/*
Assignment: Language Study Project - Part 2
Author: Aryaman Mehra & Mayur Valecha
Course: CSc 372
Instructor: L. McCann
TAs: Muhammad Bilal, Hansen Pranata, Claire Srodawa
Due Date: 11/16/2026

Description: Generates a random rectangular maze using Wilson's
             algorithm. The user provides the maze dimensions,
             and the program creates a spanning tree using
             loop-erased random walks. The resulting maze is
             displayed using '#' characters and spaces.

Language: Go 1.27.4 (developed and tested; minimum 1.20)
Ex. Packages: bufio, fmt, math/rand, os, strconv,
              strings, time (standard library).
Execution: go run .
Compilation: go build .
Input: Maze dimensions entered through standard input.

Deficiencies: No known deficiencies in the required features.
              Optional entrances, exits, and gameplay are
              not implemented.
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxDimension = 100 // Maximum permitted rows or columns.

/*
Function: main
Purpose: Reads maze dimensions, generates a random maze,

	and displays its text representation.

Precondition: None.
Postcondition: A maze is displayed if valid input is received.
Parameters: None.
Returns: None.
*/
func main() {
	fmt.Println("Wilson's Maze Generator")
	fmt.Println("-----------------------")
	fmt.Println("This program generates a random maze using")
	fmt.Println("Wilson's loop-erased random walk algorithm.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin) // Reads user input.

	rows, valid := readDimension(scanner, "rows") // Row count and status.

	if !valid {
		return
	}

	cols, valid := readDimension(scanner, "columns") // Column count and status.

	if !valid {
		return
	}

	maze := NewMaze(rows, cols)                               // Maze to be generated.
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Random source.

	GenerateWilson(maze, random)

	fmt.Println()
	fmt.Printf("Generated %d x %d maze:\n\n", rows, cols)
	fmt.Println(maze.Render())
}

/*
Function: readDimension
Purpose: Prompts the user for a valid maze dimension.
Precondition: scanner is initialized for standard input.
Postcondition: A valid dimension is returned, or false is

	returned if input ends or cannot be read.

Parameters: scanner (in/out) - scanner used to read input.

	dimension (in) - name of the requested dimension.

Returns: Valid dimension and true on success; zero and

	false if input is unavailable.
*/
func readDimension(scanner *bufio.Scanner, dimension string) (int, bool) {
	for {
		fmt.Printf("Enter number of %s (1-%d): ", dimension, maxDimension)

		if !scanner.Scan() {
			fmt.Println("\nUnable to read input.")
			return 0, false
		}

		input := strings.TrimSpace(scanner.Text()) // User's input.
		value, err := strconv.Atoi(input)          // Converted integer and error.

		// Accept only integers within the supported range.
		if err == nil && value >= 1 && value <= maxDimension {
			return value, true
		}

		fmt.Printf("Please enter an integer between 1 and %d.\n",
			maxDimension)
	}
}
