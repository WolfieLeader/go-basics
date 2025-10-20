package sync

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func configPokemon(pokemonPtr *string, pokemon string) {
	fmt.Println("Configuring Pokemon...")
	time.Sleep(100 * time.Millisecond)
	*pokemonPtr = pokemon
	fmt.Printf("Pokemon %q configured successfully!\n", pokemon)
}

func OnceExample() {
	pokemons := []string{"Pikachu", "Bulbasaur", "Charmander", "Squirtle", "Turtwig", "Chimchar", "Piplup"}
	randPokemon := pokemons[rand.Intn(len(pokemons))]
	var finalPokemon string

	// The `sync.Once` type ensures that a function is only executed once, even if called from multiple goroutines
	// This is good for singleton patterns or one-time initialization tasks
	var once sync.Once
	var wg sync.WaitGroup

	for range 10 {
		wg.Go(func() {
			if expected := pokemons[rand.Intn(len(pokemons))]; randPokemon == expected {
				// This function will only run once, regardless of how many goroutines call it
				once.Do(func() { configPokemon(&finalPokemon, randPokemon) })
			}
		})
	}
	wg.Wait()

	if finalPokemon == "" {
		fmt.Printf("Pokemon %q was not configured successfully.\n", randPokemon)
		return
	}
	fmt.Printf("Final Pokemon: %q\n", finalPokemon)
}
