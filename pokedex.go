package main

import (
	"bufio"
	// "encoding/json"
	"fmt"
	// "net/http"
	"os"
	"strings"

	"github.com/binaryCoder-101/pokedexcli/internal/pokeapi"
)

func startRepl(cfg *config) {

	//REPL
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		wordSlice := cleanInput(scanner.Text())
		if len(wordSlice) == 0 {
			continue
		}

		command := wordSlice[0]
		found := 0
		for key, value := range returnCommandMap() {
			if command == key {
				found = 1
				err := value.callback(cfg)
				if err != nil {
					fmt.Println("Error:", err)
				}
				break
			}
		}
		if found == 0 {
			fmt.Println("Unknown command")
		}
	}
}

// Description of each command
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// For holding stateful information about pagination
type config struct {
	httpClient pokeapi.Client
	prev       *string
	next       *string
}

// Maps the command names to their name, description and callback
func returnCommandMap() map[string]cliCommand {

	commandMap := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays names of next 20 locations",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of previous 20 locations",
			callback:    commandMapBackward,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commandMap
}

// Split users input into words based on whitespace, trim leading and trailing whitespaces, and lowercase the input.
func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	output := strings.Fields(lowerCase)
	return output
}

// Displays the locations for a particular URL and updates the pagination state
func DisplayLocationAreasUpdatePagination(urlInput *string, cfg *config) error {

	respData, err := cfg.httpClient.ResponseData(urlInput)
	if err != nil {
		return err
	}

	locationResults := respData.Results

	for _, locationResult := range locationResults {
		fmt.Println(locationResult.Name)
	}

	cfg.prev = respData.Previous
	cfg.next = respData.Next

	return nil
}
