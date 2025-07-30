package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var pokemons = []string{"Pikachu", "Bulbasaur", "Charmander", "Squirtle", "Jigglypuff", "Meowth", "Eevee", "Snorlax", "Mewtwo", "Gengar"}

func loadPokemon(pokemon *string) {
	fmt.Println("Loading a random Pokemon...")
	time.Sleep(100 * time.Millisecond)
	*pokemon = pokemons[rand.Intn(len(pokemons))]
	fmt.Println("Pokemon loaded!")
}

func onceExample() {
	// The `sync.Once` type ensures that a function is only executed once,
	// even if called from multiple goroutines
	var once sync.Once
	var wg sync.WaitGroup
	var pokemon string

	for range 10 {
		wg.Go(func() {
			// This function will only run once, regardless of how many goroutines call it
			once.Do(func() { loadPokemon(&pokemon) })
		})
	}
	wg.Wait()
	fmt.Printf("Final Pokemon: %s\n", pokemon)
}
