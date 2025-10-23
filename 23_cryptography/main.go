package main

import (
	"fmt"
)

const (
	HELLO_WORLD = "ĤèlĬϴ ₩órł⫒🌍"
	FORCE       = "May the Force be with you⭐"
)

func main() {
	fmt.Println("Hashing Example")
	hashExample()
	fmt.Println()

	fmt.Println("MD5 Collision Example")
	md5CollisionExample()
	fmt.Println()

	fmt.Println("SHA1 Collision Example")
	sha1CollisionExample()
	fmt.Println()

	fmt.Println("HMAC Example")
	hmacExample()
	fmt.Println()

	fmt.Println("HMAC JWT Example")
	hmacJwtExample()
	fmt.Println()

	fmt.Println("Symmetric Encryption Example")
	symmetricEncryptExample()
	fmt.Println()
}
