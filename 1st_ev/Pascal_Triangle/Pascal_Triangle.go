package main

import (
	"fmt"
)

func factorial(n int) int{

	if n==0 {
		
		return 1

	}else if n>0 {

		return n*factorial(n-1)

	}else{

		fmt.Println("Invalid number")
		return 0

	}

}

func combination(n int,k int) int{

	return factorial(n)/(factorial(k)*factorial(n-k))

}

func main()  {
	
	var user int
	var countLine int = 0
	
	fmt.Print("How many Pascal Triangle`s lines you want?:")
	fmt.Scanf("%d", &user)
	fmt.Println()

	for countLine<user {

		for z:= user;z>countLine;z--{
			fmt.Print(" ")
		}

		for countRow:=0;countRow<=countLine;countRow++{

			fmt.Print(combination(countLine,countRow))
			fmt.Print(" ")


		}

		countLine++
		fmt.Println()

		
	} 



}