package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}

	return bs, nil
}

func compareHashPassword(password string, hashPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid Password : %w /n", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error in signMessage while hashing message: %w \n", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSignature(msg, signature []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in checkSignature getting signature of message %w \n", err)
	}

	same := hmac.Equal(newSig, signature)

	return same, nil
}

var key []byte

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))

	pass := 123456789
	hashedPass, err := hashPassword(string(pass))
	if err != nil {
		panic(err)
	}

	err = compareHashPassword(string(pass), hashedPass)
	if err != nil {
		log.Fatalln("not logged in")
	}

	log.Println("Logged in!")
}
