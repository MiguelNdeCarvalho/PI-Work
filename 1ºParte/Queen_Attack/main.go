package main

import (
	"fmt"
	"math"
)

func print(matrix [8][8]string) {
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix); c++ {
			fmt.Printf("%s", matrix[r][c])
		}
		fmt.Println("")
	}
}

func check(lW, cW, lB, cB float64) {
	c := cW - cB
	l := lW - lB
	if lW == lB || cW == cB || c == l || math.Abs(c) == l || c == math.Abs(l) {
		fmt.Println("Queens can attack each other!")
	} else {
		fmt.Println("Queens can't attack each other!")
	}
}

func main() {

	var lW, cW, lB, cB float64
	var matrix = [8][8]string{} //Create the matrix

	//Insert the _ on the matrix
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix); c++ {
			matrix[r][c] = "_"
		}
	}

	fmt.Println("Insert the Coordinates of the White Queen")
	fmt.Scanf("%f,%f", &lW, &cW)
	matrix[int(lW)-1][int(cW)-1] = "W" //Set the White Queen on the Matrix

	fmt.Println("Insert the Coordinates of the Black Queen")
	fmt.Scanf("%f,%f", &lB, &cB)
	matrix[int(lB)-1][int(cB)-1] = "B" //Set the Black Queen on the Matrix

	print(matrix)
	check(lW, cW, lB, cB)
}
