package main

import(
	"strings"
	"fmt"
)


func sentence(frase string)string{

	frase1:=strings.Split(frase, " ")
	newFrase:=""

	
	for i:=0;i<len(frase);i++{

		newFrase=newFrase+translate(frase1[i])

	}
	
	return newFrase
}

func translate(word string)string{

	var result string
	
	if 4<len(word) && strings.HasPrefix(word, "xr"){ 
	
		result=word+"ay"
	
	}else if 4<len(word) && strings.HasPrefix(word, "yt"){ 
	
		result=word+"ay"

	}else if 4<len(word) && strings.HasPrefix(word, `[aeiou]`){ 
	
		result=word+"ay"

	}else if 4<len(word) && !strings.HasPrefix(word, `[aeiou]`){ 
			
		if strings.HasPrefix(word, `^[aeiou]qu`){

			result=strings.TrimPrefix(word,substring(word,0,2))+substring(word,0,2)+"ay"
		
		}else{
	
			var n int
			n=0
			for n!=-1{
				result=strings.TrimPrefix(word,string(word[n]))+string(word[n])+"ay"
		    	n++
				if strings.HasPrefix(word, `[aeiou]`){ 
					return result
				}
			}
		
		}

	}else if 4<len(word) && strings.HasPrefix(word, `[aeiou]`){ 
	
		//result :=word+"ay"

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
	
	frase:="a man"
	frase=sentence(frase)
	fmt.Printf("%s\n", frase)


}

