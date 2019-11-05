package main

import(
	"strings"
	"fmt"
)


func sentence(frase string)string{

	frase=strings.Split("a man a plan a canal panama", " ")
	newFrase:=""

	
	for i:=0;i<len(frase);i++{

		newFrase=newFrase+translate(frase[i])

	}
	
	return newFrase
}

func translate(word string)string{

	var result string
	
	if (4<len(word) && strings.HasPrefix(word, "xr"))
	
		result:=word+"ay"
	
	}

	if (4<len(word) && strings.HasPrefix(word, "yt"))
	
		result :=word+"ay"

	}

	if (4<len(word) && strings.HasPrefix(word, `[aeiou]`))
	
		result :=word+"ay"

	}

	if (4<len(word) && ~strings.HasPrefix(word, `[aeiou]`))
			
		var n int=0
		for n!=-1{
			result :=word+word.chatAt(n)+"ay"
		    n++
			if (strings.HasPrefix(word, `[aeiou]`))
				n=-1
			}
		
		}

	}

	if (4<len(word) && strings.HasPrefix(word, `[aeiou]`))
	
		result :=word+"ay"

	}



	return result
}


func main()  {
	
	frase:=strings.Split("a man a plan a canal panama", " ")
	fmt.Printf("%q\n", frase)
	fmt.Printf("%d\n",len(frase))

}

