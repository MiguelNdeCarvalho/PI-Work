package main

import (
	"fmt"
)

func rowM(matriz [][]int , size_column int ,n_row int)[]int{
	
	a := make([]int,size_column) 
	
	for i:=0;i<size_column;i++{
		
		a[i]=matriz[n_row][i]

	}

	return a

}

func colunM(matriz [][]int ,size_row int ,n_column int)[]int{
	
	a := make([]int,size_row) 
	
	for i:=0;i<size_row;i++{
		
		a[i]=matriz[i][n_column]

	}

	return a

}


func main()  {

	var sizeRow,sizeComumn int
	var matrix = [][]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	

	fmt.Print("Insert number of rows:")
	fmt.Scanf("%d",&sizeRow)

	fmt.Print("Insert number of columns:")
	fmt.Scanf("%d",&sizeComumn)

	fmt.Println("Linhas:")
	
	for i:=0;i<sizeRow;i++{

		fmt.Println(rowM(matrix,sizeComumn,i))

	}

	fmt.Println("Colunas:")

	for z:=0;z<sizeComumn;z++{

		fmt.Println(colunM(matrix,sizeRow,z))	

	}

	//falta uma func para preencher o array, mas penso que é o suficiente 

}

