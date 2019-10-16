package main

import(
	"fmt"
)

func main() {

	var num int
	var mult2 int
	var mult1 int

	fmt.Println("numero?")
	fmt.Scanf("%d",&num)


	fmt.Println("multiplo 1?")
	fmt.Scanf("%d",&mult1)

	fmt.Println("multiplo 2?")
	fmt.Scanf("%d",&mult2)

	var sum =0
	for i:=1; i<num;i++{

		if i%mult1==0 || i%mult2==0{
			sum+=i
		}
	}
	fmt.Printf("A soma dos multiplos é: %d\n", sum)

	
}

