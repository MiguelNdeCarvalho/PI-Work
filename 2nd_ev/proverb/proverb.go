// Package proverb generates the relevant proverb
package main
import "fmt"

// Proverb given a list of inputs, generate the relevant proverb
func Proverb(rhyme []string) (result []string) {
	for i := 0; i < len(rhyme); i++ {
		if len(rhyme)-1 != i {
			result = append(result, fmt.Sprint("For want of a ", rhyme[i], " the ", rhyme[i+1], " was lost.\n"))
		} else {
			result = append(result, fmt.Sprint("And all for the want of a ", rhyme[0],".\n"))
		}
	}
	return 
}

var zero = [] string {}
var one = [] string {"nail"}
var two = [] string {"nail", "shoe"}
var three = [] string {"nail", "shoe", "horse"}
var full = [] string {"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}
var four = []string{"pin", "gun", "soldier", "battle"}

func main() {
	

	fmt.Print(Proverb(full))
}