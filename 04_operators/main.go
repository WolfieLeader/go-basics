package main

import (
	"fmt"
)

func main() {
	// Arithmetic Operators: +, -, *, /, %, ++, --
	fmt.Println("Arithmetic Operators:")
	a, b := 6, 4
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	a++
	fmt.Printf("After increment, a = %d\n", a)
	b--
	fmt.Printf("After decrement, b = %d\n", b)
	fmt.Printf("%.2f / %.2f = %.2f\n", float64(a), float64(b), float64(a)/float64(b))
	fmt.Println()

	// Comparison Operators: ==, !=, <, <=, >, >=
	fmt.Println("Comparison Operators:")
	fmt.Printf("%d == %d = %t\n", a, b, a == b)
	fmt.Printf("%d != %d = %t\n", a, b, a != b)
	fmt.Printf("%d < %d = %t\n", a, b, a < b)
	fmt.Printf("%d <= %d = %t\n", a, b, a <= b)
	fmt.Printf("%d > %d = %t\n", a, b, a > b)
	fmt.Printf("%d >= %d = %t\n", a, b, a >= b)
	fmt.Println()

	// Logical Operators: &&, ||, !
	fmt.Println("Logical Operators:")
	x, y := true, false
	fmt.Printf("%t && %t = %t\n", x, y, x && y)
	fmt.Printf("%t || %t = %t\n", x, y, x || y)
	fmt.Printf("!%t = %t\n", x, !x)
	fmt.Println()

	// Bitwise Operators: &(AND), |(OR), ^(XOR), &^(AND NOT), ^X(NOT) <<(LEFT SHIFT), >>(RIGHT SHIFT)
	fmt.Println("Bitwise Operators:")

	var c, d uint8 = 92, 45 // 01011100 and 00101101 in binary

	// & (AND) - both bits must be 1
	fmt.Printf("   %08b (%d)\n", c, c)
	fmt.Printf("&  %08b (%d)\n", d, d)
	fmt.Printf("=  %08b (%d)\n\n", c&d, c&d) // 00001100 = 12

	// | (OR) - at least one bit must be 1
	fmt.Printf("   %08b (%d)\n", c, c)
	fmt.Printf("|  %08b (%d)\n", d, d)
	fmt.Printf("=  %08b (%d)\n\n", c|d, c|d) // 01111101 = 125

	// ^ (XOR - Exclusive OR) - bits must be different
	fmt.Printf("   %08b (%d)\n", c, c)
	fmt.Printf("^  %08b (%d)\n", d, d)
	fmt.Printf("=  %08b (%d)\n\n", c^d, c^d) // 01110001 = 113

	// &^ (AND NOT) - bits in d are cleared from c
	fmt.Printf("   %08b (%d)\n", c, c)
	fmt.Printf("&^ %08b (%d)\n", d, d)
	fmt.Printf("=  %08b (%d)\n\n", c&^d, c&^d) // 01010000 = 80

	// ^ (NOT) - flips all bits
	fmt.Printf("^  %08b (%d)\n", c, c)
	fmt.Printf("=  %08b (%d)\n\n", ^c, ^c) // 1010011 = 163 (for uint8)

	// << (LEFT SHIFT) - shifts bits to the left, filling with 0 like multiplying by 2 powers
	fmt.Printf("   %08b (%d) << 1\n", c, c)
	fmt.Printf("=  %08b (%d)\n\n", c<<1, c<<1) // 10111000 = 184

	fmt.Printf("   %08b (%d) << 2\n", d, d)
	fmt.Printf("=  %08b (%d)\n\n", d<<2, d<<2) // 10111100 = 180

	// >> (RIGHT SHIFT) - shifts bits to the right, filling with 0 like dividing by 2 powers
	fmt.Printf("   %08b (%d) >> 1\n", c, c)
	fmt.Printf("=  %08b (%d)\n\n", c>>1, c>>1) // 00101110 = 46

	fmt.Printf("   %08b (%d) >> 2\n", d, d)
	fmt.Printf("=  %08b (%d)\n", d>>2, d>>2) // 00001011 = 11
}
