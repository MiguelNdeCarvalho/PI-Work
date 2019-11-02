package main

import "fmt"

func change(coins []int, desejado int)[]int{

	N_Coins := len(coins)
	count := make([]int, desejado + 1)
	from := make([]int, desejado + 1)

	count[0] = 1
	
	if desejado < 0 {
		fmt.Printf("%d:valor desejado invalido", desejado)
	}

	for z := 0; z < desejado; z++{
		
		if count[z] > 0{
			
			for i := 0; i < N_Coins; i++{
				
				p := z + coins[i]
				if p <= desejado{
					
					if count[p] == 0 || count[p] > count[z] + 1{
						
						count[p] = count[z] + 1
						from[p] = i

					}
				}
			}
		}
	}

	result := make([]int, count [desejado] - 1)
	for l:= desejado; l > 0; {
		
		result[count[l] - 2 ] = coins[from[l]]
		l = l - coins[from[l]]
	
	}

	return result
}

func main(){
	
	coins := []int{1,5,10,20}
	var desejado int = 26
	var resultado []int = change(coins,desejado)

	fmt.Print(resultado)
}