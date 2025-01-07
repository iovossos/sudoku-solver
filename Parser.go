/*
This program parses and validates a Sudoku puzzle provided via command-line arguments and ensures it is solvable.
The main function, `ParseInput`, reads a 9x9 grid of Sudoku input from the command line, verifies its correctness,
and ensures that there are no conflicts in rows, columns, or 3x3 subgrids. It also checks that the grid contains
at least 17 clues (non-empty cells) and is not completely empty. The supporting functions help ensure that the
grid is valid according to Sudoku rules:

- `validateInitialGrid`: Ensures no duplicates exist in the initial grid's rows, columns, or subgrids.
- `isEmptyGrid`: Checks whether the grid is completely empty.
- `isValid`: Determines whether placing a specific number in a given position is valid according to Sudoku rules.

These functions work together to ensure that the input is valid before attempting to solve the Sudoku puzzle.
*/

package sudokux

import (
	"fmt"
	"os"
)

// ParseInput parses command-line arguments into a Sudoku grid and validates it.
func ParseInput() (map[string]rune, error) {
	// Ensure we have exactly 9 rows from the input (os.Args[1] to os.Args[9])
	if len(os.Args) != 10 { // os.Args[0] is the program name, so we expect 10 arguments in total
		return nil, fmt.Errorf("expected 9 rows of input, got %d", len(os.Args)-1) // Return an error if the count is incorrect
	}

	// Create a map to store the Sudoku grid, where the key is the position (e.g., "A1") and the value is the rune at that position.
	grid := make(map[string]rune)

	// Define the mapping of row indices (0-8) to row names (A-I), corresponding to Sudoku grid rows.
	rowNames := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}

	// Counter for the number of clues (non-empty cells in the grid).
	clueCount := 0

	// Iterate through each of the 9 rows in the command-line input (os.Args[1] to os.Args[9]).
	for i := 1; i <= 9; i++ {
		row := os.Args[i] // Get the row string from the arguments

		// Ensure the row is exactly 9 characters long (standard for a Sudoku row).
		if len(row) != 9 {
			return nil, fmt.Errorf("row %d is not 9 characters long", i) // Return an error if the row length is not 9
		}

		// Iterate through each character in the row to check if it's valid and to populate the grid.
		for j, char := range row {
			// Check if the character is either a number ('1'-'9') or a dot ('.'). If not, return an error.
			if (char < '1' || char > '9') && char != '.' {
				return nil, fmt.Errorf("you can only input numbers and dots (invalid character found in row %d)", i)
			}

			// Generate the grid key in the format "A1", "A2", etc., where 'A' is the row name and '1' is the column number.
			pos := string(rowNames[i-1]) + string(rune('1'+j))

			// Add the character to the grid at the corresponding position. If it's a dot ('.'), it means the cell is empty.
			grid[pos] = char

			// If the character is not a dot ('.'), it's considered a clue, so we increment the clueCount.
			if char != '.' {
				clueCount++
			}
		}
	}

	// After processing all rows, check if there are fewer than 17 clues.
	if clueCount < 17 {
		return nil, fmt.Errorf("invalid grid: less than 17 clues (only %d clues)", clueCount) // Return an error if there are fewer than 17 clues
	}

	// Additional validations
	// 1. Validate the initial grid for conflicts (no duplicates in rows, columns, or subgrids).
	if !validateInitialGrid(grid) {
		return nil, fmt.Errorf("invalid grid: contains conflicts in rows, columns, or subgrids") // Return an error if there are conflicts
	}

	// 2. Check if the grid is completely empty.
	if isEmptyGrid(grid) {
		return nil, fmt.Errorf("invalid grid: the entire grid is empty") // Return an error if the grid is completely empty
	}

	// If all checks pass, return the populated grid and no error.
	return grid, nil
}

// validateInitialGrid checks if the starting grid is valid (no duplicates in rows, columns, or subgrids).
func validateInitialGrid(grid map[string]rune) bool {
	for pos, val := range grid { // Iterate over the grid to check for conflicts
		if val != '.' { // If the cell is not empty
			original := grid[pos]              // Temporarily store the original value
			grid[pos] = '.'                    // Set the cell to empty to test the validity of placing the original value
			if !isValid(grid, pos, original) { // If the original value is not valid in its position
				return false // Return false if there's a conflict
			}
			grid[pos] = original // Restore the original value in the grid
		}
	}
	return true // Return true if no conflicts are found
}

// isEmptyGrid checks if the entire grid is empty (i.e., all cells contain '.').
func isEmptyGrid(grid map[string]rune) bool {
	for _, val := range grid { // Loop through each cell in the grid
		if val != '.' { // If any cell is not empty
			return false // Return false (the grid is not empty)
		}
	}
	return true // Return true if all cells are empty
}

// isValid checks if placing a number (num) at a given position (pos) is valid according to Sudoku rules.
func isValid(grid map[string]rune, pos string, num rune) bool {
	row := rune(pos[0]) // Extract the row from the position
	col := rune(pos[1]) // Extract the column from the position

	// Check the row for duplicates
	for i := '1'; i <= '9'; i++ { // Loop through all columns in the current row
		if grid[string(row)+string(i)] == num { // If the number already exists in the row
			return false // Return false if a duplicate is found
		}
	}

	// Check the column for duplicates
	for i := 'A'; i <= 'I'; i++ { // Loop through all rows in the current column
		if grid[string(i)+string(col)] == num { // If the number already exists in the column
			return false // Return false if a duplicate is found
		}
	}

	// Check the 3x3 subgrid for duplicates
	startRow := rune('A') + (row-'A')/3*3 // Calculate the starting row of the 3x3 subgrid
	startCol := rune('1') + (col-'1')/3*3 // Calculate the starting column of the 3x3 subgrid
	for i := 0; i < 3; i++ {              // Loop through 3 rows of the subgrid
		for j := 0; j < 3; j++ { // Loop through 3 columns of the subgrid
			posCheck := string(startRow+rune(i)) + string(startCol+rune(j)) // Generate the position for subgrid cells
			if grid[posCheck] == num {                                      // If the number already exists in the subgrid
				return false // Return false if a duplicate is found
			}
		}
	}

	return true // Return true if the number can be placed in the position without conflict
}
