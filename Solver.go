/*
This program solves a Sudoku puzzle using a backtracking algorithm with constraint propagation.
The grid is represented as a map where each key is a string (representing a position like "A1")
and the value is a rune (the number or an empty cell).

The solution process involves:
1. **Tracking used numbers** in each row, column, and 3x3 subgrid using maps.
2. **Minimum Remaining Values (MRV) heuristic**: This is used to select the next cell to fill,
   prioritizing cells with the fewest valid options.
3. **Backtracking algorithm**: The program tries placing numbers in the empty cells while ensuring
   that no duplicates are in any row, column, or subgrid. If a number placement leads to a dead end,
   it backtracks and tries a different number.

Functions:
- **`SolveSudoku`**: Main function that orchestrates solving the Sudoku puzzle. It keeps track of
  the numbers used in rows, columns, and subgrids, and calls helper functions to find the next cell
  to fill and backtrack when necessary. It returns the solved grid if there is exactly one solution,
  or `false` if there are no solutions or multiple solutions.
- **`findNextCell`**: Helper function to find the next empty cell to fill, using the MRV heuristic
  (i.e., the cell with the fewest valid options).
- **`solve`**: Recursive function that attempts to solve the puzzle by placing numbers in empty cells,
  backtracking when necessary.
- **`copyGrid`**: Helper function to copy the current state of the grid when a solution is found.

The goal is to solve the Sudoku puzzle, ensuring there is exactly one solution. If multiple solutions
or no solution exists, the program will return false.
*/

package sudokux

func SolveSudoku(grid map[string]rune) (map[string]rune, bool) {
	// Maps to track used numbers in each row, column, and 3x3 subgrid
	usedRows := make(map[rune]map[rune]bool)       // Map to track numbers used in each row
	usedCols := make(map[rune]map[rune]bool)       // Map to track numbers used in each column
	usedSubgrids := make(map[string]map[rune]bool) // Map to track numbers used in each 3x3 subgrid

	// Initialize the usedRows and usedCols maps for all rows (A-I) and columns (1-9)
	for i := 'A'; i <= 'I'; i++ { // Loop through all rows (A-I)
		usedRows[i] = make(map[rune]bool) // Initialize map to track numbers for each row
	}
	for j := '1'; j <= '9'; j++ { // Loop through all columns (1-9)
		usedCols[j] = make(map[rune]bool) // Initialize map to track numbers for each column
	}

	// Populate the tracking maps based on the initial grid values
	for i := 'A'; i <= 'I'; i++ { // Loop through all rows (A-I)
		for j := '1'; j <= '9'; j++ { // Loop through all columns (1-9)
			pos := string(i) + string(j) // Create a position string, e.g., "A1", "B2"
			val := grid[pos]             // Get the value (rune) at this position
			if val != '.' {              // If the cell is not empty ('.')
				usedRows[i][val] = true                                      // Mark the number as used in the current row
				usedCols[j][val] = true                                      // Mark the number as used in the current column
				subgrid := string('A'+(i-'A')/3*3) + string('1'+(j-'1')/3*3) // Calculate subgrid identifier
				if usedSubgrids[subgrid] == nil {                            // If the subgrid map is not initialized, do so
					usedSubgrids[subgrid] = make(map[rune]bool)
				}
				usedSubgrids[subgrid][val] = true // Mark the number as used in the current subgrid
			}
		}
	}

	// Variable to keep track of how many solutions have been found
	solutionCount := 0
	// Map to store the solved grid (when exactly one solution is found)
	solvedGrid := make(map[string]rune)

	// Function to find the position with the Minimum Remaining Values (MRV)
	findNextCell := func() (string, int) {
		minOptions := 10              // Start with a value greater than the max possible options (9)
		bestCell := ""                // Store the position of the best cell (least options)
		for i := 'A'; i <= 'I'; i++ { // Loop through all rows (A-I)
			for j := '1'; j <= '9'; j++ { // Loop through all columns (1-9)
				pos := string(i) + string(j) // Create the position string
				if grid[pos] == '.' {        // If the cell is empty
					options := 0                                                 // Count valid options for this cell
					subgrid := string('A'+(i-'A')/3*3) + string('1'+(j-'1')/3*3) // Calculate subgrid identifier
					for num := '1'; num <= '9'; num++ {                          // Try numbers 1-9
						if !usedRows[i][num] && !usedCols[j][num] && !usedSubgrids[subgrid][num] { // Check if the number can be placed
							options++ // Increment options count if valid
						}
					}
					if options < minOptions { // If this cell has fewer options than the current minimum
						minOptions = options // Update the minimum options
						bestCell = pos       // Set this cell as the best option
					}
				}
			}
		}
		return bestCell, minOptions // Return the best cell and its number of options
	}

	// Recursive helper function to solve the grid using backtracking with MRV heuristic
	var solve func() bool
	solve = func() bool {
		pos, minOptions := findNextCell() // Find the next cell with the MRV heuristic
		if pos == "" {                    // If no empty cell is found, the grid is fully filled
			return true // Return true to indicate that the grid is solved
		}
		if minOptions == 0 { // If there are no valid options for this cell
			return false // Backtrack by returning false
		}

		row := rune(pos[0])                                              // Extract the row (letter) from the position
		col := rune(pos[1])                                              // Extract the column (number) from the position
		subgrid := string('A'+(row-'A')/3*3) + string('1'+(col-'1')/3*3) // Calculate the subgrid identifier

		for num := '1'; num <= '9'; num++ { // Try numbers 1-9 in the empty cell
			if !usedRows[row][num] && !usedCols[col][num] && !usedSubgrids[subgrid][num] { // Check if number is valid in row, column, and subgrid
				grid[pos] = num                                                                       // Place the number in the grid
				usedRows[row][num], usedCols[col][num], usedSubgrids[subgrid][num] = true, true, true // Mark the number as used

				if solve() { // Recursively attempt to solve the rest of the grid
					solutionCount++         // Increment solution count if a solution is found
					if solutionCount == 1 { // If this is the first solution found
						copyGrid(grid, solvedGrid) // Copy the grid as the solved grid
					}
					if solutionCount > 1 { // If more than one solution is found
						return false // Return false to indicate multiple solutions (not uniquely solvable)
					}
				}

				grid[pos] = '.'                                                                          // Undo the current move (backtrack)
				usedRows[row][num], usedCols[col][num], usedSubgrids[subgrid][num] = false, false, false // Mark the number as unused
			}
		}
		return false // Return false to continue backtracking
	}

	solve() // Start solving the grid

	if solutionCount != 1 { // If there isn't exactly one solution
		return grid, false // Return the grid and false (no solution or multiple solutions)
	}
	return solvedGrid, true // Return the solved grid and true (exactly one solution found)
}

// copyGrid copies the grid state from src to dest.
// This is used to store the first valid solution found during backtracking.
func copyGrid(src, dest map[string]rune) {
	for k, v := range src { // Loop through all key-value pairs in the source grid
		dest[k] = v // Copy each key-value pair to the destination grid
	}
}
