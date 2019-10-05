package main

import (
	"fmt"
	"os"
	"os/exec"
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

func toWrite(n int) string {
	return "Test"
}

func main() {
	checkSpeak()

	var n int
	fmt.Println("Number to Speech. Only accept from 0 to 999999999999")
	fmt.Print("Insert the Desired number:")

	fmt.Scanf("%d", &n)

	numRange(n)
}
