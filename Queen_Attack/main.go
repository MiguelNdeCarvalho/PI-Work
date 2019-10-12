package main

import "fmt"

/* func generate(matrix [8][8]string) [8][8]string {
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix); c++ {
			matrix[r][c] = "_"
		}
	}
	return matrix
} */

/* func set(color string, matrix [8][8]string, x int, y int) [8][8]string {
	if color == "white" {
		matrix[x-1][y-1] = "W"
	} else if color == "black" {
		matrix[x-1][y-1] = "B"
	}

	print(matrix)
	return matrix
} */

func print(matrix [8][8]string) {
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix); c++ {
			fmt.Printf("%s", matrix[r][c])
		}
		fmt.Println("")
	}
}

func check(lW, cW, lB, cB int) {
	c := cW - cB
	l := lW - lB
	if lW == lB || cW == cB || c == l || c*-1 == l || c == l*-1 {
		fmt.Println("works")
	} else {
		fmt.Println("dont work")
	}
}

func main() {

	var lW, cW, lB, cB int
	var matrix = [8][8]string{} //Create the matrix

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix); c++ {
			matrix[r][c] = "_"
		}
	}

	//generate(matrix)

	fmt.Println("Insert the Coordinates of the White Queen")
	fmt.Scanf("%d", &lW)
	fmt.Scanf("%d", &cW)

	matrix[lW-1][cW-1] = "W"

	/* set(white, matrix, lW, cW) */

	fmt.Println("Insert the Coordinates of the Black Queen")
	fmt.Scanf("%d", &lB)
	fmt.Scanf("%d", &cB)
	matrix[lB-1][cB-1] = "B"
	/* 	set(black, matrix, lB, cB) */
	print(matrix)
	check(lW, cW, lB, cB)
}
