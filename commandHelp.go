package main

import (
	"fmt"
)

// Callback for help command
func commandHelp(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("too many arguments")
	}
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range returnCommandMap() {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}
