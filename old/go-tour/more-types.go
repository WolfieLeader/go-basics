package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
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

func PrintSlice(s []int) {
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))
}

func CreateSlice() []int {
	s := []int{1, 2, 3, 4, 5, 6}
	PrintSlice(s)

	//? The length of the slice is 6, after overflowing it the capacity is doubled to 12
	s = append(s, 7, 8, 9)
	PrintSlice(s)

	s = append(s, 10, 11, 12, 13, 14, 15)
	PrintSlice(s)

	s = s[:3]
	PrintSlice(s) //? Len 3, Cap 24, because there are 4 underlying arrays of 6 elements each

	return s
}

func MakeSlice() {
	//? Make a slice with length 5 and capacity 5
	a := make([]int, 5)
	PrintSlice(a)

	//? Make a slice with length 0 and capacity 5
	b := make([]int, 0, 5)
	PrintSlice(b)

}

func LoopSlice(s []int) {
	//? Looping through a slice
	for i, v := range s {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2) //? Example: simple gradient
		}
	}
	return pic
}

func ShowPic() {
	pic.Show(Pic)
}

type Point struct {
	Lat float64
	Lng float64
}

func MakeMap() map[string]Point {
	//? Map with string keys and Point values
	m := make(map[string]Point)
	m["Bell Labs"] = Point{40.68433, -74.39967}
	m["Google"] = Point{37.42202, -122.08408}
	return m
}

func CreateMap() map[string]Point {
	//? Map literal
	m := map[string]Point{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	return m
}

func DeleteFromMap(m map[string]Point, key string) {
	//? Delete a key from a map
	delete(m, key)
}

func GetElementFromMap(m map[string]Point, key string) (Point, bool) {
	v, ok := m[key]
	return v, ok
}

func WordCount(s string) map[string]int {
	//? Fields separates a string with spaces
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func TestWordCount() {
	wc.Test(WordCount)
}

func Compute(a, b float64, fn func(float64, float64) float64) int {
	return int(math.Floor(fn(a, b)))
}

func CreatePowFn() func(float64, float64) float64 {
	return func(a, b float64) float64 {
		return math.Pow(a, b)
	}
}

func RecursiveFibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
}

func Fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func PanicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("This is a panic example")
}

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Likes int    `json:"likes,omitempty"`
	Admin bool   `json:"isAdmin,omitempty"`
}

func JsonExample() {
	u1 := User{"John Doe", 30, 1, true}

	j1, marshalErr := json.Marshal(u1)
	if marshalErr != nil {
		fmt.Println("Error marshalling JSON:", marshalErr)
		return
	}

	fmt.Println("JSON:", string(j1))
	j2 := `{"name":"Jane Doe","age":25}`
	var u2 User
	unmarshalErr := json.Unmarshal([]byte(j2), &u2)
	if unmarshalErr != nil {
		fmt.Println("Error unmarshalling JSON:", unmarshalErr)
		return
	}
	fmt.Printf("Unmarshalled User: %+v\n", u2)
}
