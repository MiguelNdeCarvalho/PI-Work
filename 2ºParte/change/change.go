package main

import "fmt"

func change(coins []int, desejado int)[]int{

	N_Coins := len(coins)
	count := make([]int, desejado + 1)
	
	if desejado < 0 {
		fmt.Printf("%d:valor desejado invalido", desejado)
	}

	for z := 0; z < desejado; z++{
		
    }

	
}



func main(){
	
	coins := []int{1,5,10,20}
	var desejado int = 25
	
}