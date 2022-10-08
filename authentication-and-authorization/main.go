package main

import (
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

func main() {
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
