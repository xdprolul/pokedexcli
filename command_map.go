package main

import (
	"fmt"
	"errors"
)

func commandMapf(cfg *config,args ...string) error {
	req,err:=cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err!=nil {
		return err
	}

	cfg.nextLocationsURL=req.Next
	cfg.prevLocationsURL=req.Previous

	for _,locA:=range req.Results {
		fmt.Println(locA.Name)
	}
	return nil
}

func commandMapb(cfg *config,args ...string) error {
	if cfg.prevLocationsURL==nil {
		return errors.New("you're on the first page")
	}

	req,err:=cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err!=nil {
		return err
	}

	cfg.nextLocationsURL=req.Next
	cfg.prevLocationsURL=req.Previous

	for _,locA:=range req.Results {
		fmt.Println(locA.Name)
	}
	return nil
}
