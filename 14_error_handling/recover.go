package main

import "math/rand"

func randomlyPanic(funcName string) {
	num := rand.Intn(3)
	if num <= 1 {
		panic(funcName + " panicked")
	}
}