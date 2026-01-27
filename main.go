package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner:=bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Printf("Pokedex > ")
		input:=""
		if scanner.Scan(){
			input=scanner.Text()
		}
		cleanedWords:=cleanInput(input)
		fmt.Printf("Your command was: %v\n",cleanedWords[0])
	}
}
