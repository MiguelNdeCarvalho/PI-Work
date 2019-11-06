package main

import(
	"strings"
	"fmt"
)


func sentence(frase string)string{

	frase1:=strings.Split(frase, " ")
	newFrase:=""

	
	for i:=0;i<len(frase1);i++{

		if i==0{
			newFrase=translate(frase1[i])
		}else{

			newFrase=newFrase+" "+translate(frase1[i])
	    }

	}
	
	return newFrase
}

func translate(word string)string{

	var result string
	
	if (len(word)>3==true) && strings.HasPrefix(word, "xr"){ 
	
		fmt.Println("oi1")

		result=word+"ay"
	
	}else if (len(word)>3==true) && strings.HasPrefix(word, "yt"){ 
	
		fmt.Println("oi2")

		result=word+"ay"

	}else if (len(word)>3==true) && strings.HasPrefix(word, `[aeiou]`){ 
	
		fmt.Println("oi3")

		result=word+"ay"

	}else if (len(word)>4==true) && !strings.HasPrefix(word, `[aeiou]`){ 
			
		fmt.Println(word)
		fmt.Println(strings.Contains(word, "[^aeiou]qu"))

		if strings.Contains(word, "[^aeiou]qu"){

			fmt.Println("oi4")

			result=strings.TrimPrefix(word,substring(word,0,3))+substring(word,0,3)+"ay"
		
		}else if (strings.HasPrefix(word, "ch")){
	
			fmt.Println("oi5")
			result=strings.TrimPrefix(word,substring(word,0,2))+substring(word,0,2)+"ay"

		}else{

			fmt.Println("oi6")

			result=strings.TrimPrefix(word,string(word[0]))+string(word[0])+"ay"


		}


	}else{

		result=word

	}

	return result
}

func substring(s string, start int, end int) string {
	
	start_str_idx := 0
    i := 0
    for j := range s {
        if i == start {
            start_str_idx = j
        }
        if i == end {
            return s[start_str_idx:j]
        }
        i++
    }
    return s[start_str_idx:]
}


func main()  {
	
	frase:="the xray yttria chair square caaaaaaa"
	frase=sentence(frase)
	fmt.Println("/"+frase)
	//falta a regra 3 e 4 


}

