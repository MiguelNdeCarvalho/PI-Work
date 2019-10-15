package main

import(
	"fmt"
)
func perfect_number( n int) int {
	
	var sum = 0;

	for i := 1; i < n; i++ {
		if n%i==0{
			sum+=i
		}
	}
	return sum
}


func main() {
	 var n int

	fmt.Println("number?")
	fmt.Scanf("%d", &n )

	var x=perfect_number(n)

	fmt.Printf("sum= %d, number= %d\n",x,n)

	if x==n{
		fmt.Println("The number is Perfect!")
	
	}else if x>n{
			fmt.Println("The number is Abundant!")
	}else{
			fmt.Println("The number is Deficient!")
	}	
}

