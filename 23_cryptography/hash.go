package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
)

func hashExample() {
	data := []byte(HELLO_WORLD)
	fmt.Printf("- Original: %s\n\n", data)

	sha2_224 := sha256.Sum224(data)
	fmt.Printf("- SHA2-224 Hash: %x\n\n", sha2_224)

	sha2_256 := sha256.Sum256(data)
	fmt.Printf("- SHA2-256 Hash: %x\n\n", sha2_256)

	sha2_384 := sha512.Sum384(data)
	fmt.Printf("- SHA2-384 Hash: %x\n\n", sha2_384)

	sha2_512 := sha512.Sum512(data)
	fmt.Printf("- SHA2-512 Hash: %x\n\n", sha2_512)

	sha3_224 := sha3.Sum224(data)
	fmt.Printf("- SHA3-224 Hash: %x\n", sha3_224)
	fmt.Printf("- Is SHA3-224 equal to SHA2-224? %v\n\n", bytes.Equal(sha3_224[:], sha2_224[:]))

	sha3_256 := sha3.Sum256(data)
	fmt.Printf("- SHA3-256 Hash: %x\n", sha3_256)
	fmt.Printf("- Is SHA3-256 equal to SHA2-256? %v\n\n", bytes.Equal(sha3_256[:], sha2_256[:]))

	sha3_384 := sha3.Sum384(data)
	fmt.Printf("- SHA3-384 Hash: %x\n", sha3_384)
	fmt.Printf("- Is SHA3-384 equal to SHA2-384? %v\n\n", bytes.Equal(sha3_384[:], sha2_384[:]))

	sha3_512 := sha3.Sum512(data)
	fmt.Printf("- SHA3-512 Hash: %x\n", sha3_512)
	fmt.Printf("- Is SHA3-512 equal to SHA2-512? %v\n\n", bytes.Equal(sha3_512[:], sha2_512[:]))

	md5Hash := md5.Sum(data)
	fmt.Printf("- MD5 Hash (vulnerable): %x\n", md5Hash)

	sha1Hash := sha1.Sum(data)
	fmt.Printf("- SHA1 Hash (vulnerable): %x\n", sha1Hash)

}

func md5CollisionExample() {
	firstPart, secondPart := "TEXTCOLLBYfGiJUETHQ4h", "cKSMd5zYpgqf1YRDhkmxHkhPWptrkoyz28wnI9V0aHeAuaKnak"
	collision1 := md5.Sum([]byte(firstPart + "A" + secondPart))
	collision2 := md5.Sum([]byte(firstPart + "E" + secondPart))

	fmt.Printf("- Hashed MD5 (A): %x\n", collision1)
	fmt.Printf("- Hashed MD5 (E): %x\n", collision2)
	fmt.Printf("- Do the hashes collide? %t\n\n", bytes.Equal(collision1[:], collision2[:]))
}

func sha1File(fileName string) []byte {
	file, err := os.Open("./files/" + fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatalf("Error hashing file: %v", err)
	}
	return hash.Sum(nil)
}

func sha1CollisionExample() {
	sha1a := sha1File("shattered-1.pdf")
	sha1b := sha1File("shattered-2.pdf")

	fmt.Printf("- SHA1 of shattered-1.pdf: %x\n", sha1a)
	fmt.Printf("- SHA1 of shattered-2.pdf: %x\n", sha1b)
	fmt.Printf("- Do the SHA1 hashes collide? %t\n", bytes.Equal(sha1a, sha1b))
}
