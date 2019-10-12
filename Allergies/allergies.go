package main

import(
	"fmt"
	"math"
)

type alergia struct {
    nome string
    pontos  float64
}


func nova_alergia(nome_a string, pontuacao float64) *alergia {
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

	var num float64
    
    fmt.Print("number: ")
	fmt.Scanf("%f", &num)

	fmt.Println("Alergias:")

	var temp = num

	//ciclo que vai descobrir as alergias
	for temp>0{

		var expoente = math.Log2(temp)

		//MACUMBAS PARA FAZER OS EXPOENTES CERTOS/////////////////
		var inteiro = int(expoente)

		var inteiro_f = float64(inteiro)


		var v_alergia = math.Pow(2,inteiro_f)
		///////////////////////////////////////////////////////////


		for j:=0;j<len(lista);j++{

			if lista[j].pontos==v_alergia{
				fmt.Printf("%s\n",lista[j].nome)
			}
		}

		temp = temp - v_alergia

	}

}
