package piscine

import "fmt"

func QuadA(x, y int) { // This is the function declaration. It receives the width (x) and height (y) of the rectangle
	if x <= 0 || y <= 0 { // If either x or y is zero or negative, the function exits early and prints nothing
		return
	}
	for row := 0; row < y; row++ { // This loop goes from the top to the bottom of the rectangle, one row at a time
		for col := 0; col < x; col++ { // This inner loop prints the characters from left to right in each row
			if row == 0 || row == y-1 { // Top or bottom row
				if col == 0 || col == x-1 { // First or Last Character 
					fmt.Print("o")
				} else { // Middle Characters
					fmt.Print("-")
				}
			} else { // Middle rows
				if col == 0 || col == x-1 {  // First or Last Character 
					fmt.Print("|")
				} else { // Middle Characters
					fmt.Print(" ")
				}

			}
		}
		fmt.Println() // Move to the next line after each row
	}

}

func QuadB(x, y int) {

	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 {
				if col == 0 {
					fmt.Print("/") // Top-left corner
				} else if col == x-1 {
					fmt.Print("\\") // Top-right corner
				} else {
					fmt.Print("*") // Top edge
				}
			} else if row == y-1 {
				if col == 0 {
					fmt.Print("\\") // Bottom-left corner
				} else if col == x-1 {
					fmt.Print("/") // Bottom-right corner
				} else {
					fmt.Print("*") // Bottom edge
				}

			} else {
				if col == 0 || col == x-1 {
					fmt.Print("*") // Left or right edge
				} else {
					fmt.Print(" ") // Inside the box
				}
			}
		}
		fmt.Println() // Move to the next line after each row
	}
}

func QuadC(x, y int) {

	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 {
				if col == 0 { // A cleaner way to do this: if col == 0 || col == x-1 {, saving the else if 
					fmt.Print("A") // Both corners are 'A'
				} else if col == x-1 {
					fmt.Print("A") // Top edge
				} else {
					fmt.Print("B")
				}
			} else if row == y-1 {
				if col == 0 { // A cleaner way to do this: if col == 0 || col == x-1 {, saving the else if 
					fmt.Print("C")
				} else if col == x-1 {
					fmt.Print("C") // Both corners are 'C'
				} else {
					fmt.Print("B") // Bottom edge
				}

			} else { // Middle Rows
				if col == 0 || col == x-1 {
					fmt.Print("B") // Left and right edges
				} else {
					fmt.Print(" ") // Else print nothing
				}
			}
		    }
		fmt.Println() // Move to the next line after each row
	}
        }
    }
}

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 { //First line
				if col == 0 {
					fmt.Print("A") // Top-left corner
				} else if col == x-1 {
					fmt.Print("C")  // Top-right corner
				} else {
					fmt.Print("B") // Top edge
				}
			} else if row == y-1 { // Last Line
				if col == 0 {
					fmt.Print("A") // Bottom-left corner
				} else if col == x-1 {
					fmt.Print("C")  // Bottom-right corner
				} else {
					fmt.Print("B")  // Bottom edge
				}

			} else { // Middle Rows
				if col == 0 || col == x-1 {
					fmt.Print("B") // Left and right edges
				} else {
					fmt.Print(" ") // Else print nothing
				}
			}
		}
		fmt.Println() // Move to the next line after each row
	}

}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 {
				if col == 0 {
					fmt.Print("A")
				} else if col == x-1 {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
			} else if row == y-1 {
				if col == 0 {
					fmt.Print("C")
				} else if col == x-1 {
					fmt.Print("A")
				} else {
					fmt.Print("B")
				}

			} else {
				if col == 0 || col == x-1 {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println() // Move to the next line after each row
	}

}
