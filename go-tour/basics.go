package main

import (
	"math"
	"math/cmplx"
	"math/rand"
)

func Swap(x, y int) (int, int) {
	return y, x
}

func GetRandomNumber(x int) int {
	return rand.Intn(x)
}

func Add(x, y int) int {
	return x + y
}

func GetAreaOfCircle(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func NakedSplit(sum int) (x, y int) { //? naked return
	x = sum * 4 / 9
	y = sum - x
	return //? naked return
}

var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1 //? This is 2^64 - 1, Shifting 1 left by 64 bits and subtracting 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func GetStats() (bool, uint64, complex128) {
	return ToBe, MaxInt, z
}

var (
	i int
	f float64
	b bool
	s string
)

func GetZeroValues() (int, float64, bool, string) { return i, f, b, s }

func Divide(x, y int) float64 {
	if y == 0 {
		return 0
	}
	return float64(x) / float64(y) //? Convert x and y to float64 before division
}
