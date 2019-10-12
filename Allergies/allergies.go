package main

import(
	"fmt"
)

type alergia struct {
    nome string
    pontos  int
}


func nova_alergia(nome_a string, pontuacao int) *alergia {
	p := alergia{nome: nome_a, pontos: pontuacao}
    //p.pontos = pontuacao
    return &p
}


func main(){


	var lista []*alergia   //array de struct 

	////valores dos testes no array
	lista = append(lista,nova_alergia("eggs",1))
	lista = append(lista,nova_alergia("peanuts",2))
	lista = append(lista,nova_alergia("shellfish",4))
	lista = append(lista,nova_alergia("strawberries",8))
	lista = append(lista,nova_alergia("tomatoes",16))
	lista = append(lista,nova_alergia("chocolate",32))
	lista = append(lista,nova_alergia("pollen",64))
	
	/////////////////TESTE//////////////////////////////
	/*
	for i:=0;i<len(lista);i++{
		fmt.Printf(("(%s,%d)\n"),lista[i].nome, lista[i].pontos)
	}
	*/
	//////////////////////////////////////////////////////////////

	var num int
    
    fmt.Print("number: ")
	fmt.Scanf("%d", &num)


	var temp = num

	



}
