package main

import (
	"fmt"
)

func fenchFrase(frase string,rails int) string{

	var result string=""

	matriz:=fench(frase,rails)
	for z:=0;z<len(matriz);z++{
		
		for i:=0;i<len(matriz[0]);i++{
			
			if matriz[z][i]!="." {

				result=result+matriz[z][i]

			}

		}
	}

	return result
}


func fench(frase string,rails int)[][]string{

	N_Linhas:=rails
	N_Colunas:=len(frase)

	
	matriz:=make([][]string,N_Linhas)
	for i := range matriz {
        matriz[i] = make([]string,N_Colunas)
	}

	for z:=0;z<N_Linhas;z++{
		
		for i:=0;i<N_Colunas;i++{
			matriz[z][i]="."
		}
	}

	//fmt.Println("oi")

	z:=0
	var top int=0
	var bot int=N_Linhas-1
	
	for z<N_Colunas{
		
		for i:=0;i<N_Linhas;i++{

			
			if top<=z && z<bot{

				matriz[i][z]=string(frase[z])
				z++
			}
			if z==N_Colunas{
				break
			}
		}
		top=+top+N_Linhas+rails-2

		if z==N_Colunas{
			break
		}

		for i:=N_Linhas-1;i>=0;i--{

			if bot<=z && z<top  {

				matriz[i][z]=string(frase[z])

				z++
			}
			if z==N_Colunas{
				break
			}
		}
		bot=top+N_Linhas-1


	}

/*
	for z:=0;z<N_Linhas;z++{
		
		for i:=0;i<N_Colunas;i++{
			fmt.Print(matriz[z][i])
		}
		fmt.Println()
	}
*/

	return matriz
}
func main()  {
	

	fmt.Println(fenchFrase("WEAREDISCOVEREDFLEEATONCE",3))


}