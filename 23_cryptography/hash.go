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
	helloWorld := []byte("ĤèlĬϴ ₩órł⫒🌍")
	fmt.Printf("- Original: %s\n\n", helloWorld)

	sha2_224 := sha256.Sum224(helloWorld)
	fmt.Printf("- SHA2-224 Hash: %x\n\n", sha2_224)

	sha2_256 := sha256.Sum256(helloWorld)
	fmt.Printf("- SHA2-256 Hash: %x\n\n", sha2_256)

	sha2_384 := sha512.Sum384(helloWorld)
	fmt.Printf("- SHA2-384 Hash: %x\n\n", sha2_384)

	sha2_512 := sha512.Sum512(helloWorld)
	fmt.Printf("- SHA2-512 Hash: %x\n\n", sha2_512)

	sha3_224 := sha3.Sum224(helloWorld)
	fmt.Printf("- SHA3-224 Hash: %x\n", sha3_224)
	fmt.Printf("- Is SHA3-224 equal to SHA2-224? %v\n\n", bytes.Equal(sha3_224[:], sha2_224[:]))

	sha3_256 := sha3.Sum256(helloWorld)
	fmt.Printf("- SHA3-256 Hash: %x\n", sha3_256)
	fmt.Printf("- Is SHA3-256 equal to SHA2-256? %v\n\n", bytes.Equal(sha3_256[:], sha2_256[:]))

	sha3_384 := sha3.Sum384(helloWorld)
	fmt.Printf("- SHA3-384 Hash: %x\n", sha3_384)
	fmt.Printf("- Is SHA3-384 equal to SHA2-384? %v\n\n", bytes.Equal(sha3_384[:], sha2_384[:]))

	sha3_512 := sha3.Sum512(helloWorld)
	fmt.Printf("- SHA3-512 Hash: %x\n", sha3_512)
	fmt.Printf("- Is SHA3-512 equal to SHA2-512? %v\n\n", bytes.Equal(sha3_512[:], sha2_512[:]))

	md5Hash := md5.Sum(helloWorld)
	fmt.Printf("- MD5 Hash (vulnerable): %x\n", md5Hash)

	firstPart, secondPart := "TEXTCOLLBYfGiJUETHQ4h", "cKSMd5zYpgqf1YRDhkmxHkhPWptrkoyz28wnI9V0aHeAuaKnak"
	collision1 := md5.Sum([]byte(firstPart + "A" + secondPart))
	collision2 := md5.Sum([]byte(firstPart + "E" + secondPart))
	fmt.Printf("- Hashed MD5 (A): %x\n", collision1)
	fmt.Printf("- Hashed MD5 (E): %x\n", collision2)
	fmt.Printf("- Do the hashes collide? %t\n\n", bytes.Equal(collision1[:], collision2[:]))

	sha1Hash := sha1.Sum(helloWorld)
	fmt.Printf("- SHA1 Hash (vulnerable): %x\n", sha1Hash)

	file1, err := os.Open("./files/shattered-1.pdf")
	if err != nil {
		log.Fatalf("Error opening file 1: %v", err)
	}
	defer file1.Close()

	file2, err := os.Open("./files/shattered-2.pdf")
	if err != nil {
		log.Fatalf("Error opening file 2: %v", err)
	}
	defer file2.Close()

	hash1 := sha1.New()
	hash2 := sha1.New()

	if _, err := io.Copy(hash1, file1); err != nil {
		log.Fatalf("Error hashing file 1: %v", err)
	}
	if _, err := file1.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("Error resetting file 1 pointer: %v", err)
	}

	if _, err := io.Copy(hash2, file2); err != nil {
		log.Fatalf("Error hashing file 2: %v", err)
	}
	if _, err := file2.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("Error resetting file 2 pointer: %v", err)
	}

	sha1a := hash1.Sum(nil)
	sha1b := hash2.Sum(nil)

	fmt.Printf("- SHA1 of shattered-1.pdf: %x\n", sha1a)
	fmt.Printf("- SHA1 of shattered-2.pdf: %x\n", sha1b)
	fmt.Printf("- Do the SHA1 hashes collide? %t\n", bytes.Equal(sha1a, sha1b))
}
