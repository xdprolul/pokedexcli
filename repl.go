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
	nextLocationsURL	*string
	prevLocationsURL	*string
	caughtPokemon			map[string]pokeapi.Pokemon
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

		args:=[]string{}
		if len(cleanedWords)>1 {
			args=cleanedWords[1:]
		}

		command,ok:=getCommands()[cleanedWords[0]]
		if ok {
			err := command.callback(cfg,args...)
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
	callback 		func(*config,...string) error
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
		"explore": {
			name:					"explore <location_name>",
			description:	"Get the pokemons in the area",
			callback: 		commandExplore,
		},
		"catch": {
			name:					"catch <pokemon_name>",
			description:	"Catch a pokemon and add it to the pokedex",
			callback:			commandCatch,
		},
		"inspect": {
			name:					"inspect <pokemon_name>",
			description:	"Inspect a pokemon that has been caught",
			callback:			commandInspect,
		},
		"pokedex": {
			name:					"pokedex",
			description:	"List all the pokemon caught",
			callback:			commandPokedex,
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
