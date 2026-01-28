package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/xdprolul/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient			pokeapi.Client
	nextLocationsURL		*string
	prevLocationsURL		*string
}

func startRepl(cfg *config) {
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
			err := command.callback(cfg)
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
	callback 		func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:					"help",			
			description:	"Displays a help message",
			callback:			commandHelp,
		},
		"map": {
			name:					"map",
			description:	"Get the next page of location areas",
			callback:			commandMapf,
		},
		"mapb": {
			name:					"mapb",
			description:	"Get the previous page of location areas",
			callback:			commandMapb,
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
