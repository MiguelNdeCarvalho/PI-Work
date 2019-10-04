package main

import (
	"fmt"
	"math"
)

func toBinary(n int) int {

	var rest int
	var counter int = 0
	var result int = 0
	
	
	for n>0 {
        
		rest=n%2
		n=n/2


		if rest!=0{


			result = result + int(math.Pow10(counter))

		
		}
		counter++


	}

	return result



}

func reverse(hand []string) []string {
	
	for i := 0; i < len(hand)/2; i++ {
		
		j := len(hand) - i - 1
		
		hand[i], hand[j] = hand[j], hand[i]
	}

	return hand
}

func main(){

	var user,bin int
	var shake []string
	var i int = 0
	
	fmt.Print("Insira o 'HandShake': ")
	fmt.Scanf("%d", &user)
	bin=toBinary(user)
	fmt.Print("Binario:")
	fmt.Println(bin)
	
    if bin%2!=0{

		shake[i]=",wink,"
		i++

	}

	if bin>=10{

		shake[i]=",double wink,"
		i++

	}

	if bin>=100{
		
		shake[i]=",close your eyes,"
		i++
			
	}

	if bin>=1000{
		
		shake[i]=",jump,"
		i++
			
	}

	if bin>=10000{

	    //reverse(shake) //falta esta parte
			
	}

	fmt.Println(shake)


	
	

	








}