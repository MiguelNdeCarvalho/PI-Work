package main

import(
	"fmt"
)

var price = 800

/*/////////////////////////////////
2 LIVROS 5%
3 LIVROS 10%
4 LIVROS 20%
5 LIVROS 25%
////////////////////////////////////
*/
var teste = []int{1,1,2,2,3,3,4,5}



//FUNÇÃO QUE CALCULA O MELHOR PREÇO
func Custo(livros []int) int{

	var preco = 0
	var num_livros = len(livros)

	//prototipo, aqui apenas soma todos os livros, so para testar se funciona
	for i:=0;i<num_livros;i++{
		preco+=price
	}

	
	//fazer todas as combinaçoes possiveis
	//quando a combinaçao de um preço for menor que a outra, o preço vai passar a ter esse valor
	//no fim de estarem feitas todas as combinações, o preço será o menor
	//por fim retorna-se o preço


	return preco
}

func main(){

	fmt.Printf("%d\n",Custo(teste))
}