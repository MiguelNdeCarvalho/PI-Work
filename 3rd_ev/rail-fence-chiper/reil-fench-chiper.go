package main

import (
	"fmt"
)

func fenchFrase(frase string,rails int) string{


	return "ola"
}


func fench(frase string,rails int)[][]string{

	N_Linhas:=rails
	N_Colunas:=len(frase)
	fmt.Println(N_Linhas)
	fmt.Println(N_Colunas)
	
	matriz:=make([][]string,N_Linhas)
	for i := range matriz {
        matriz[i] = make([]string,N_Colunas)
	}

	for z:=0;z<N_Linhas;z++{
		
		for i:=0;i<N_Colunas;i++{
			matriz[z][i]="."
		}
	}

	z:=0
	for z<N_Colunas{
		
		for i:=0;i<N_Linhas;i++{

			matriz[i][z]=string(frase[z])
			z++
		}

		

		for i:=N_Linhas-1;i>=0;i--{
			
			matriz[i][z]=string(frase[z])
			z++
		}

	}

	for z:=0;z<N_Linhas;z++{
		
		for i:=0;i<N_Colunas;i++{
			fmt.Print(matriz[z][i])
		}
		fmt.Println()
	}

	return matriz
}

func main()  {
	
	fench("BomDia",3)


}