# Wilson's Maze Generator

## CSc 372 - Comparative Programming Languages

**Assignment:** Language Study Project - Part 2  
**Authors:** Aryaman Mehra and Mayur Valecha  
**Instructor:** L. McCann  
**TAs:** Muhammad Bilal, Hansen Pranata, Claire Srodawa  
**Due Date:** November 16, 2026  
**Language:** Go (Golang)

## Language Study (Part A)

Our study of the Go programming language is available at the following link:

[Go Language Study](https://docs.google.com/document/d/1bHafbjZN74MTonOZfBBGEpD2qB4PKzyG3n8W0Jk5jNY/edit?usp=sharing)

## Program Description

This program generates a random rectangular maze using Wilson's algorithm. The algorithm constructs a uniform spanning tree through loop-erased random walks.

1. Randomly select an initial cell and add it to the maze.
2. Randomly select a cell that is not yet part of the maze.
3. Perform a random walk, selecting uniformly among legal neighboring cells.
4. Erase loops whenever a cell in the current walk is revisited.
5. Stop the walk when it reaches a cell already in the maze.
6. Remove the walls between consecutive cells in the surviving path and add those cells to the maze.
7. Repeat until all cells are connected.

The resulting maze is displayed using `#` characters for walls and spaces for cells and open passages. Every cell is reachable, and there are no cycles.

## Requirements

- Go 1.20 or later (developed and tested locally using Go 1.27.4).
- A terminal capable of running Go programs.
- No external libraries or dependencies.

## Compilation and Execution

Navigate to the project directory and run:

```bash
go run .
```

Alternatively, compile and execute the program:

```bash
go build -o wilson-maze .
./wilson-maze
```

Enter the desired number of rows and columns when prompted. Both dimensions must be integers between 1 and 100. Invalid inputs are rejected and the user is prompted again.

## Example Output

```text
Wilson's Maze Generator
-----------------------
This program generates a random maze using
Wilson's loop-erased random walk algorithm.

Enter number of rows (1-100): 3
Enter number of columns (1-100): 8

Generated 3 x 8 maze:

#################
#   #   # #     #
# # ### # # #####
# #       #     #
### ##### ##### #
#       #       #
#################
```

The maze will generally differ between executions because the algorithm uses random walks.

## Project Structure

| File | Description |
|---|---|
| `main.go` | Program entry point, user input, and validation. |
| `maze.go` | Cell and maze structures, neighbor selection, and wall operations. |
| `wilson.go` | Wilson's algorithm and loop erasure. |
| `render.go` | ASCII maze rendering. |
| `maze_test.go` | Tests for maze initialization, neighbors, and wall operations. |
| `wilson_test.go` | Tests for connectivity, passage counts, and loop erasure. |
| `render_test.go` | Tests for rendering, dimensions, and exterior boundaries. |
| `go.mod` | Go module configuration. |

## Testing

```bash
go test -v
go vet ./...
go test -cover
```

The automated tests verify maze initialization, shared wall removal, loop-free random walks, connectivity, the required number of passages, and ASCII rendering. Multiple maze sizes and random seeds are tested. Structural tests do not independently establish statistical uniformity.

## Known Limitations

The program supports dimensions from 1 to 100 and displays mazes as text. It generates mazes rather than providing an interactive maze game. There are no known deficiencies in the required functionality.
