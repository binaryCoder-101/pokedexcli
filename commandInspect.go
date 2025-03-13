package main

func(cfg *config, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}
	if len(args) < 1 {
		return fmt.Errorf("too few arguments")
	}

	pokemonName := args[0]

	if _, exists := cfg.pokedexDatabase.CaughtPokemon[pokemonName]; exists {

	}else{
		fmt.Println("You haven't caught the pokemon yet!")
	}

	return nil
}