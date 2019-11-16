package main

import (
	"fmt"
)

func diamond(letra rune){

	N_letras:=int(letra)-int('A')+1
	length:=2*N_letras-1
	meio:=length/2
	size:=1
	sizeMax:=length
	fmt.Printf("%d\n",N_letras)
	fmt.Printf("%d\n",length)
	fmt.Println(meio)

	z:=1
	for size<sizeMax+1{
		
		if size<=N_letras{

			for x:=0;x<length;x++{
				
				if (z==1||z==sizeMax) && meio==x{
					
					fmt.Printf("%c",'A')

				}else if z!=1 && (x==meio-z+1 || x==meio+z-1){
				
					value:=int('A')+z-1
					character := rune(value)
					fmt.Printf("%c",character)

				}else{
				
					fmt.Printf(".")

				}
			
			}
			fmt.Printf("\n")
			z++
		
		}else{

			if size==N_letras+1{
				z--
			}
			
			z--
			for x:=0;x<length;x++{
				
				if (size==sizeMax) && meio==x{
					
					fmt.Printf("%c",'A')

				}else if x==meio-z+1|| x==meio+z-1{
				
					value:=int('A')+z-1
					character := rune(value)
					fmt.Printf("%c",character)

				}else{
				
					fmt.Printf(".")

				}
			
			}
			fmt.Printf("\n")
	
		}
		size++
	}



}

func main()  {
	
	value1:=int('A'+1)
	fmt.Println(value1)
	character1 := rune(value1)
	fmt.Printf("Character corresponding to Ascii Code %d = %c\n", value1, character1)
	diamond('E')

}