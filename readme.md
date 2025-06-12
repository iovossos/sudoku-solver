# Sudoku Solver Program in Go

This repository contains a Sudoku solver program implemented in Go. The solver uses a combination of backtracking and the Minimum Remaining Values (MRV) heuristic to efficiently solve Sudoku puzzles. The program expects a 9x9 grid as input, parses it, solves it, and outputs the solution if one exists. If no solution or multiple solutions are found, the program notifies the user accordingly.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Sudoku Solving Strategy](#sudoku-solving-strategy)
- [Minimum Remaining Values (MRV) Heuristic](#minimum-remaining-values-mrv-heuristic)
- [Backtracking Algorithm](#backtracking-algorithm)
- [Input Parsing](#input-parsing)
- [How to Run the Program](#how-to-run-the-program)
- [Authors](#authors)

## Overview

The Sudoku solver is a command-line program that solves a given 9x9 Sudoku puzzle. The puzzle is provided as input through command-line arguments, with each argument representing one row of the Sudoku grid. Empty cells are represented by dots (.), while filled cells are represented by numbers ('1' through '9'). The program validates the input, solves the puzzle using the backtracking algorithm enhanced by the MRV heuristic, and then displays the result.

## Project Structure

The project consists of two main components:

- `main.go`: This file contains the entry point of the program. It parses user input, calls the Sudoku solving functions, and prints the result.
- `sudokux/`: This package contains the core logic for solving the Sudoku puzzle. It includes functions for solving the puzzle using backtracking and the MRV heuristic, parsing the input, and validating the grid.

## Sudoku Solving Strategy

The program uses a backtracking algorithm to solve the Sudoku puzzle. The backtracking approach systematically tries all possible values for each empty cell until a solution is found or it is determined that the puzzle cannot be solved. Although backtracking is a brute-force method, it is optimized using heuristics such as Minimum Remaining Values (MRV).

### Backtracking Algorithm Flow

- The algorithm picks an empty cell and tries placing each possible number (1-9) in that cell.
- After placing a number, the algorithm moves to the next empty cell.
- If a number cannot be placed (due to a conflict in the row, column, or 3x3 subgrid), the algorithm backtracks by undoing the last placement and tries a different number.
- This process continues until the puzzle is either solved or determined to be unsolvable.

## Minimum Remaining Values (MRV) Heuristic

The MRV heuristic improves the efficiency of the backtracking algorithm by always selecting the most constrained variable (i.e., the most constrained cell) first. A cell is considered "constrained" if it has fewer legal values that can be placed in it. By prioritizing the most constrained cells, the algorithm reduces the chances of making incorrect guesses and minimizes the search space.

### How MRV Works

- For each empty cell in the grid, the solver counts how many valid numbers can be placed in that cell, considering the current state of the row, column, and subgrid.
- The cell with the fewest valid options is selected first.

This approach leads to faster solutions by resolving highly constrained cells early, reducing unnecessary backtracking.

## Backtracking Algorithm

The backtracking algorithm is the core of the Sudoku solver. It attempts to solve the puzzle by recursively trying different numbers in empty cells. If a conflict arises (i.e., no valid number can be placed in a cell), the algorithm undoes its last move and tries a different number.

### Key Points

- It systematically explores all possible configurations of numbers in the grid.
- When it encounters a cell that cannot be filled with a valid number, it "backtracks" by undoing the last move and trying a different number in the previous cell.
- The algorithm terminates when either the grid is completely filled with valid numbers (indicating a solution) or all possibilities have been exhausted (indicating no solution).
- The MRV heuristic helps reduce the time complexity of the backtracking algorithm by focusing on the cells with the fewest available options first.

## Input Parsing

The program expects a 9x9 Sudoku grid as input, with each row provided as a separate argument. Each row must contain exactly 9 characters, where:

- Numbers ('1' to '9') represent filled cells.
- Dots ('.') represent empty cells.

The input is parsed into a data structure that represents the Sudoku grid, stored as a map where the keys are the positions (e.g., `"A1"` for the first cell in the first row) and the values are the corresponding numbers or empty cells.

### Input Validation

The input is validated to ensure:

- The grid contains exactly 9 rows, each with 9 characters.
- There are at least 17 clues (pre-filled cells), as Sudoku puzzles with fewer than 17 clues may not have a unique solution.
- The initial grid does not contain any conflicts in rows, columns, or subgrids.

## How to Run the Program

To run the program, use the following command format:

```bash
go run main.go "row1" "row2" "row3" "row4" "row5" "row6" "row7" "row8" "row9"
```

Each argument represents a row of the Sudoku grid. For example:
```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```
This input represents the following Sudoku Grid: 
```bash
. 9 6 . 4 . . . 1
1 . . . 6 . . . 4
5 . 4 8 1 . 3 9 .
. . 7 9 5 . . 4 3
. 3 . . 8 . . . .
4 . 5 . 2 3 . 1 8
. 1 . 6 3 . . 5 9
. 5 9 . 7 . 8 3 .
. . 3 5 9 . . . 7
```

## Authors

1. [iovossos](https://github.com/iovossos)
