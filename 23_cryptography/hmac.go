package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func hmacExample() {
	secret := []byte("mysecretkey")
	data := []byte(helloWorld)

	hmacSha256 := hmac.New(sha256.New, secret)
	hmacSha256.Write(data)
	tag := hmacSha256.Sum(nil)

	fmt.Printf("- Message: %s\n", data)
	fmt.Printf("- HMAC Tag: %x\n", tag)

	verify := hmac.New(sha256.New, secret)
	verify.Write(data)
	expectedTag := verify.Sum(nil)
	fmt.Printf("- HMAC Verified with correct key: %v\n", hmac.Equal(tag, expectedTag)) // Constant time comparison

	verifyWrong := hmac.New(sha256.New, []byte("wrongkey"))
	verifyWrong.Write(data)
	wrongTag := verifyWrong.Sum(nil)
	fmt.Printf("- HMAC Verified with wrong key: %v\n", hmac.Equal(tag, wrongTag))
}

var now = time.Date(2025, 10, 23, 12, 0, 0, 0, time.UTC)
var past = now.Add(-2 * time.Hour)

func hmacJwtExample() {
	secret := []byte("myjwtsecretkey")
	payload := jwtPayload{Sub: "id-12345", Iat: now.Unix(), Exp: now.Add(time.Hour).Unix(), Name: "John Doe"}
	jwtToken, err := signJwt(payload, secret)
	if err != nil {
		log.Fatalf("Error signing JWT: %v", err)
	}
	fmt.Printf("- JWT Token: %s\n", jwtToken)
	verifiedPayload, err := verifyJwt(jwtToken, secret)
	if err != nil {
		log.Fatalf("Error verifying JWT: %v", err)
	}
	fmt.Printf("- Verified JWT Payload: %+v\n", verifiedPayload)

	_, err = verifyJwt(jwtToken, []byte("wrongsecretkey"))
	fmt.Printf("- JWT Verification with wrong key error: %v\n", err)
	_, err = verifyJwt(jwtToken[2:], secret)
	fmt.Printf("- JWT Verification with tampered token error: %v\n", err)

	expiredPayload := jwtPayload{Sub: "id-12345", Iat: past.Unix(), Exp: past.Add(time.Hour).Unix(), Name: "John Doe"}
	expiredToken, err := signJwt(expiredPayload, secret)
	if err != nil {
		log.Fatalf("Error signing expired JWT: %v", err)
	}
	_, err = verifyJwt(expiredToken, secret)
	fmt.Printf("- JWT Verification with expired token error: %v\n", err)
}

type jwtHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type jwtPayload struct {
	Sub  string `json:"sub"`
	Iat  int64  `json:"iat"`
	Exp  int64  `json:"exp"`
	Name string `json:"name"`
}

func signJwt(payload jwtPayload, secret []byte) (string, error) {
	header := jwtHeader{Alg: "HS256", Typ: "JWT"}
	headerJson, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	encodedHeader := base64.RawURLEncoding.EncodeToString(headerJson)
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadJson)

	hmacSha256 := hmac.New(sha256.New, secret)
	hmacSha256.Write([]byte(encodedHeader + "." + encodedPayload))
	signature := hmacSha256.Sum(nil)
	encodedSignature := base64.RawURLEncoding.EncodeToString(signature)

	jwtToken := encodedHeader + "." + encodedPayload + "." + encodedSignature

	return jwtToken, nil
}

func verifyJwt(jwtToken string, secret []byte) (jwtPayload, error) {
	var payload jwtPayload

	parts := strings.Split(jwtToken, ".")
	if len(parts) != 3 {
		return payload, errors.New("invalid JWT token format")
	}

	decodedHeader, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return payload, err
	}
	var header jwtHeader
	if err = json.Unmarshal(decodedHeader, &header); err != nil {
		return payload, err
	}
	if header.Alg != "HS256" || header.Typ != "JWT" {
		return payload, errors.New("unsupported JWT header")
	}

	decodedPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return payload, err
	}
	if err = json.Unmarshal(decodedPayload, &payload); err != nil {
		return payload, err
	}

	decodedSignature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return payload, err
	}
	hmacSha256 := hmac.New(sha256.New, secret)
	hmacSha256.Write([]byte(parts[0] + "." + parts[1]))
	expectedSignature := hmacSha256.Sum(nil)

	if !hmac.Equal(decodedSignature, expectedSignature) {
		return payload, errors.New("invalid JWT signature")
	}
	if now.Unix() < payload.Iat {
		return payload, errors.New("JWT token used before issued")
	}
	if now.Unix() >= payload.Exp {
		return payload, errors.New("JWT token has expired")
	}
	return payload, nil
}
