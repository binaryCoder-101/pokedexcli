package main

import (
	"fmt"
	"os"
)

// Callback for exit command
func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
