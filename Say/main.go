package main

import (
	"fmt"
	"os"
)

func numRange(n int) {
	if n < 0 || n > 999999999999 {
		fmt.Println("The number that you inserted is not on the range!")
		os.Exit(1)
	}
}

func toWrite(n int) string {
	return "Test"
}

func main() {
	var n int
	fmt.Println("Number to Speech. Only accept from 0 to 999999999999")
	fmt.Print("Insert the Desired number:")

	fmt.Scanf("%d", &n)

	numRange(n)
}
