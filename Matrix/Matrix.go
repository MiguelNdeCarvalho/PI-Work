package main

import (
	"fmt"
)

func colunM(matriz [][]int ,n_column int)[]int{
	
	a := make([]int,len(matriz)) 
	
	for i:=0;i<len(matriz);i++{
		
		a[i]=matriz[i][n_column]

	}

	return a

}


func main()  {

	var matrix = [][]int{ {9,8,7}, {5,3,2}, {6,6,7},{6,6,8}}	

	fmt.Println("Linhas:")
	
	for i:=0;i<len(matrix);i++{

		fmt.Println(matrix[i])

	}

	fmt.Println("Colunas:")

	for z:=0;z<len(matrix[z]);z++{

		fmt.Println(colunM(matrix,z))	

	}

	

}

