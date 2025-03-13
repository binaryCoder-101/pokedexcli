package main

import (
	"fmt"

	"github.com/binaryCoder-101/pokedexcli/internal/pokeapi"
)

func commandMapForward(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("too many arguments")
	}

	if cfg.next == nil && cfg.prev == nil {
		url := pokeapi.BaseURL + "/location-area"
		err := DisplayLocationAreasUpdatePagination(&url, cfg)
		if err != nil {
			return err
		}
		return nil
	} else if cfg.next == nil && cfg.prev != nil {
		return fmt.Errorf("last location area. Next location area not available")
	} else {
		err := DisplayLocationAreasUpdatePagination(cfg.next, cfg)
		if err != nil {
			return err
		}
		return nil
	}

}

func commandMapBackward(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("too many arguments")
	}

	if cfg.next == nil && cfg.prev == nil {
		return fmt.Errorf("no data")
	} else if cfg.next != nil && cfg.prev == nil {
		return fmt.Errorf("first location area. Previous location area not available")
	} else {
		err := DisplayLocationAreasUpdatePagination(cfg.prev, cfg)
		if err != nil {
			return err
		}
		return nil
	}
}
