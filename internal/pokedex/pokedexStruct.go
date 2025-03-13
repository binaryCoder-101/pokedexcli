package pokedex

type PokedexStruct struct {
	CaughtPokemon map[string]Pokemon
}

type Pokemon struct {
	Name           string
	BaseExperience int
}

func NewPokedex() *PokedexStruct {
	return &PokedexStruct{
		CaughtPokemon: make(map[string]Pokemon),
	}
}
