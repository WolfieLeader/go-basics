package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func FactorialBasicForLoop(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	result := 1
	//? The init and post statements are optional in Go's for loop.
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

func FactorialWhileForLoop(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	i := n
	result := 1
	//? if you omit the condition, the loop will run forever
	for i > 0 {
		result *= i
		i--
	}

	return result
}
func Power(x, n, lim float64) float64 {
	//? if statements can start with a statement to execute before the condition
	if v := math.Pow(x, n); v < lim {
		return v
	}
	//? here v is not defined, so we cannot use it
	return lim
}

func SqrtWithIterations(x float64) float64 {
	const ITERATIONS = 10
	z := 1.0
	for i := 1; i <= ITERATIONS; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Printf("Iteration %d: z = %f \n", i, z)
	}
	return z
}

func Absolute(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func SqrtWithThreshold(x float64) (float64, int) {
	const THRESHOLD = 1e-10
	counter := 0
	z := 1.0
	for {
		counter++
		prev := z
		z = z - (z*z-x)/(2*z)
		if Absolute(z-prev) < THRESHOLD {
			break
		}
	}
	return z, counter
}

func CurrentOS() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default:
		return fmt.Sprintf("Unknown: %s", os)
	}
}

func IsWeekendClose() string {
	today := int(time.Now().Weekday())
	weekend := int(time.Saturday)

	switch (weekend - today + 7) % 7 {
	case 0:
		return "Today"
	case 1:
		return "Tomorrow"
	case 2:
		return "In two days"
	default:
		return "Too far"
	}
}

func LevelByScore(x int) string {
	//? like switch true
	switch {
	case x <= 0:
		return "New"
	case x <= 100:
		return "Bronze"
	case x <= 500:
		return "Silver"
	case x <= 1000:
		return "Gold"
	case x <= 5000:
		return "Platinum"
	case x <= 10000:
		return "Diamond"
	default:
		return "Master"
	}
}

func DeferredPrint() {
	defer fmt.Println("deferred - end")
	fmt.Println("normal - start")
	for i := 1; i <= 3; i++ {
		fmt.Printf("normal - middle %d \n", i)
		defer fmt.Printf("deferred - middle %d \n", i) //? deferred calls are executed after the surrounding function returns
	}
	fmt.Println("normal - end")
	defer fmt.Println("\ndeferred - start")
}
