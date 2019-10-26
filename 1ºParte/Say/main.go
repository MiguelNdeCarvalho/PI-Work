package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func checkSpeak() { //Check if espeak-ng is present
	cmd := exec.Command("espeak-ng")
	if err := cmd.Run(); err != nil {
		fmt.Println("You dont have 'eSpeakNG' package installed so the program can`t be executed correctly.")
		os.Exit(1)
	}
}

func numRange(n int) { //Check if number is on the desired range
	if n < 0 || n > 999999999999 {
		speech := exec.Command("espeak-ng", "-a 200", "'The number you inserted is out of range!'")
		speech.Run()
		os.Exit(1)
	}
}

var englishMegas = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion"}
var englishUnits = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var englishTens = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var englishTeens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

func integerToTriplets(number int) []int {
	triplets := []int{}

	for number > 0 {
		triplets = append(triplets, number%int(1000))
		number = number / int(1000)
	}

	return triplets
}

func toWrite(n int) string {
	words := []string{}

	triplets := integerToTriplets(n)

	if len(triplets) == 0 {
		return "zero"
	}

	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]

		if triplet == 0 {
			continue
		}

		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		if hundreds > 0 {
			words = append(words, englishUnits[hundreds], "hundred")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, englishUnits[units])
		case 1:
			words = append(words, englishTeens[units])
			break
		default:
			if units > 0 {
				word := fmt.Sprintf("%s-%s", englishTens[tens], englishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, englishTens[tens])
			}
			break
		}

	tripletEnd:
		if mega := englishMegas[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	return strings.Join(words, " ")
}

func readIt(input string) {
	speech := exec.Command("espeak-ng", input)
	speech.Run()
}

func main() {
	checkSpeak() //Check if exists the package to talk

	var n int
	fmt.Println("Number to Speech. Only accept from 0 to 999999999999")
	fmt.Print("Insert the Desired number:")

	fmt.Scanf("%d", &n)

	numRange(n)

	fmt.Println(toWrite(n)) //Need to print instead of speaking the number
	//readIt(toWrite(n)) //Tranform the numbers into string and say it
}
