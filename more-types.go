package main

import (
	"fmt"
	"math/rand"
)

func ReadMemoryAddress() {
	var x int = rand.Intn(100)
	fmt.Printf("1 - Memory address of x: %p and its value: %d\n", &x, x)
	//? the *T type is a pointer to a value of type T.
	//? The &x operator gives the memory address of a x variable.
	var p *int = &x
	//? The * operator dereferences a pointer to access the value at that address.
	fmt.Printf("2 - Memory address of x: %p and its value: %d\n", p, *p)
}

type Player struct {
	Name    string
	Score   int
	IsAdmin bool
}

func CreatePlayer(name string) Player {
	//? Named fields
	player := Player{Name: name, Score: 100, IsAdmin: false}
	return player
}

func AddScore(player *Player) {
	//? Pointer to a struct, so we can modify the original struct
	player.Score += 100
}

func CreateAdmin(name string) Player {
	//? Positional fields
	return Player{name, 0, true}
}

func GetPoint(x, y int) struct {
	X int
	Y int
} {
	//? Anonymous struct
	point := struct {
		X int
		Y int
	}{X: x, Y: y}

	return point
}

func GetHelloWorldArray() ([2]string, [5]int) {
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	return arr, [5]int{1, 2, 3, 4, 5}
}

func GetHelloWorldSlice() (strSlice []string, intSlice []int) {
	strSlice = append(strSlice, "Hello", "Go")
	//? There isn't such thing as pop or shift in Go
	strSlice = strSlice[:len(strSlice)-1] //? this is like 0:len(slice)-1
	strSlice = append(strSlice, "World")

	numArray := [5]int{1, 2, 3, 4, 5}
	intSlice = numArray[1:4] //? index 1 to 3 (4 is excluded)
	return
}

func PrintCorruptedNames() {
	names := [4]string{"John", "Paul", "George", "Ringo"}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:4]
	fmt.Println(a, b)

	b[0] = "XXX" //? Will change the original array, slices are references to the original array
	fmt.Println(a, b)
	fmt.Println(names)
}
