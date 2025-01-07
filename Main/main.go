package main

import (
	"fmt"
	"os"
	"sudokux" // Import the sudokux package where the Sudoku functions are defined
)

// main is the entry point of the program. It parses the command-line input to create a Sudoku grid, solves it using a backtracking algorithm, and prints the result.
// The program expects 9 rows of input, each with 9 characters (numbers '1'-'9' or dots '.' representing empty cells).
func main() {
	// Parse the command-line input to create the Sudoku grid, using the ParseInput function from the sudokux package.
	grid, err := sudokux.ParseInput()
	if err != nil {
		// If there's an error during parsing (e.g., invalid input format), print the error and exit the program with a non-zero status.
		fmt.Println("Error:", err)
		os.Exit(1) // Exit the program if input is invalid
	}

	// Solve the Sudoku puzzle using the SolveSudoku function from the sudokux package.
	solvedGrid, solved := sudokux.SolveSudoku(grid)
	if solved {
		// If the puzzle is successfully solved, print a success message and display the solved grid.
		fmt.Println("Sudoku solved successfully:")
		// Call a helper function to print the solved Sudoku grid in a readable 9x9 format.
		printSudoku(solvedGrid)
	}
}

// printSudoku is a helper function to print the Sudoku grid in a formatted 9x9 layout.
// It iterates through rows 'A' to 'I' and columns '1' to '9', printing the grid values for each position.
func printSudoku(grid map[string]rune) {
	// Iterate through each row (from 'A' to 'I')
	for i := 'A'; i <= 'I'; i++ {
		// Iterate through each column (from '1' to '9') for the current row
		for j := '1'; j <= '9'; j++ {
			// Print the value at the current grid position (i, j) followed by a space
			fmt.Print(string(grid[string(i)+string(j)]), " ")
		}
		// After printing all columns for the current row, print a newline to move to the next row
		fmt.Println()
	}
}
