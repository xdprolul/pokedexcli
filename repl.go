package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	scanner:=bufio.NewScanner(os.Stdin)

  for {
		fmt.Printf("Pokedex > ")
    input:=""
    if scanner.Scan(){
	    input=scanner.Text()
    }
    cleanedWords:=cleanInput(input)
    //fmt.Printf("Your command was: %v\n",cleanedWords[0])
		command,ok:=getCommands()[cleanedWords[0]]
		if ok {
			err := command.callback()
			if err!=nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}	
  }
}

type cliCommand struct {
	name 				string
	description string
	callback 		func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:					"help",			
			description:	"Displays a help message",
			callback:			commandHelp,
		},
		"exit": {
			name:					"exit",
			description:	"Exit the Pokedex",
			callback: 		commandExit,
		},
	}
}

func cleanInput(text string) []string {
	words:=strings.Fields(strings.ToLower(text))	
	return words
}
