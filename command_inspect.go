package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config,args ...string) error {
	if len(args)!=1 {
		return errors.New("you must provide a pokemon name")
	}

	name:=args[0]
	exists,ok:=cfg.caughtPokemon[name];
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	
	fmt.Println("Name: ",exists.Name)
	fmt.Println("Height: ",exists.Height)
	fmt.Println("Weight: ",exists.Weight)
	fmt.Println("Stats:")
	for _,stat:=range exists.Stats {
		fmt.Printf("  -%s: %v\n",stat.Stat.Name,stat.BaseStat)
	}
	fmt.Println("Types:")
	for _,typeInfo := range exists.Types {
		fmt.Printf("  -%s\n",typeInfo.Type.Name)
	}
	return nil
}
