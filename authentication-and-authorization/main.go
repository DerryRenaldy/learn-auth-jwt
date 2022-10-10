package main

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var key []byte

type UserClaims struct {
	jwt.RegisteredClaims
	SessionId int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now(), true) {
		return fmt.Errorf("Token has expired")
	}

	if u.SessionId == 0 {
		return fmt.Errorf("Invalid Session ID")
	}

	return nil
}

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Effor in createToken when signing token: %w \n", err)
	}
	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	claims := &UserClaims{}
	t, err := jwt.ParseWithClaims(signedToken, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() == jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}
		return key, nil
	})

	// check error
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorSignatureInvalid:
			// token invalid
			return nil, fmt.Errorf("Error signature invalid: %w \n", jwt.ValidationErrorSignatureInvalid)
		case jwt.ValidationErrorExpired:
			return nil, fmt.Errorf("Error signature expired: %w \n", jwt.ValidationErrorExpired)
		default:
			return nil, fmt.Errorf("some error occured: %w \n", err)
		}

	}

	if !t.Valid {
		return nil, fmt.Errorf("Error token is invalid")
	}

	claims = t.Claims.(*UserClaims)
	return claims, nil
}

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
