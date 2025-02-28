package main

import "fmt"

func commandMapForward(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if cfg.next == nil && cfg.prev == nil {
		err := displayLocationAreasUpdatePagination(&url, cfg)
		if err != nil {
			return err
		}
		return nil
	} else if cfg.next == nil && cfg.prev != nil {
		return fmt.Errorf("last location area. Next location area not available")
	} else {
		err := displayLocationAreasUpdatePagination(cfg.next, cfg)
		if err != nil {
			return err
		}
		return nil
	}

}

func commandMapBackward(cfg *config) error {

	if cfg.next == nil && cfg.prev == nil {
		return fmt.Errorf("no data")
	} else if cfg.next != nil && cfg.prev == nil {
		return fmt.Errorf("first location area. Previous location area not available")
	} else {
		err := displayLocationAreasUpdatePagination(cfg.prev, cfg)
		if err != nil {
			return err
		}
		return nil
	}
}
