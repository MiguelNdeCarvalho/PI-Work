package main

import(
	"fmt"
)

////////////////////////////////TESTES////////////////////////////////////////////////////////

var peras = [][]int{{9  ,8 , 7},{5  ,3 , 2 },{6  ,6  ,7}}

//var marco =[][]int { {1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

//var rui =[][]int {{10,20,30,40,50}, {11,21,32,43,51}, {43,53,12,45,78}, {12,43,6,8,63}, {76,46,5,3,54}}

//var namorado = [][]int{{2,4},{3,8}}

///////////////////////////////////////////////////////////////////////////////////////////////////////
func pos_maior(linha []int ) int{
	var maior = 0;

	for i:=0; i<len(linha);i++{
		if linha[i]>linha[maior]{
			maior = i
		}
	}
	return maior
}


func main(){

	var matriz = peras //MATRIZ UTILIZADA

	for i:=0; i<len(matriz);i++{

		var l_do_maior = pos_maior(matriz[i])

		var valor = matriz[i][l_do_maior]

		var j int = 0
		

		for j<len(matriz){

			if matriz[j][l_do_maior] < valor {

				//sai do ciclo

				break

				}

				j++

		}

		if j==len(matriz[i]){

			fmt.Printf("sadle point: (%d,%d)\n", i, l_do_maior)

		} 

	} 
}