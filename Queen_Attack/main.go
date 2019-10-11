package main

import "fmt"

func generate(matrix [8][8]string) {
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			matrix[r][c] = "_"
		}
	}
}

func set(color string, matrix [8][8]string, x int, y int) {
	if color == "white" {
		matrix[x][y] = "W"
	} else if color == "black" {
		matrix[x][y] = "B"
	}
}

func print(matrix [8][8]string) {
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			fmt.Println(matrix[r][c])
		}
	}
}

func main() {

	var lW, cW, lB, cB int
	white := "white"
	black := "black"
	var matrix = [8][8]string{} //Create the matrix

	generate(matrix)

	fmt.Println("Insert the Coordinates of the White Queen")
	fmt.Scanf("%d", &lW)
	fmt.Scanf("%d", &cW)
	set(white, matrix, lW, cW)

	fmt.Println("Insert the Coordinates of the Black Queen")
	fmt.Scanf("%d", &lB)
	fmt.Scanf("%d", &cB)
	set(black, matrix, lB, cB)

	print(matrix)

}
