package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
)

func createKey(length int) ([]byte, error) {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func encryptAesGcm(plaintext, key, aadInfo []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block) // Galois/Counter Mode - authenticated encryption
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize()) // 12-byte nonce for GCM
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, aadInfo)
	return ciphertext, nil
}

func decryptAesGcm(ciphertext, key, aadInfo []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce := ciphertext[:nonceSize]
	data := ciphertext[nonceSize:]

	return gcm.Open(nil, nonce, data, aadInfo)
}

func symmetricEncryptExample() {
	plaintext := []byte(FORCE)
	aadInfo := []byte("Star Wars") // optional: Additional Authenticated Data - not encrypted but authenticated
	fmt.Printf("- Plaintext: %s\n", plaintext)

	key, err := createKey(32) // AES-256 requires 32-byte key
	if err != nil {
		log.Fatalf("Error creating key: %v", err)
	}

	ciphertext, err := encryptAesGcm(plaintext, key, aadInfo)
	if err != nil {
		log.Fatalf("Error encrypting data: %v", err)
	}
	fmt.Printf("- Ciphertext (hex): %x\n", ciphertext)

	decryptedText, err := decryptAesGcm(ciphertext, key, aadInfo)
	if err != nil {
		log.Fatalf("Error decrypting data: %v", err)
	}
	fmt.Printf("- Decrypted Text: %s\n", decryptedText)

	ciphertextWithWrongAAD, err := encryptAesGcm(plaintext, key, []byte("Bed Wars"))
	if err != nil {
		log.Fatalf("Error encrypting data with wrong AAD: %v", err)
	}
	_, err = decryptAesGcm(ciphertextWithWrongAAD, key, aadInfo)
	fmt.Printf("- Decryption with wrong AAD error: %v\n", err)
}
