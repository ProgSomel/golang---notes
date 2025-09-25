package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

//? Function to hash password
func HashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func main() {
	password := "password123"
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt: ", err)
		return
	}

	//? Hash the password with salt
	signUpHash := HashPassword(password, salt)

	//? Store the salt and password to the database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt: ", saltStr)
	fmt.Println("Signup Hash: ", signUpHash)

	//! verify
	//? retrieve the saltStr and decode
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt: ", err)
		return
	}
	loginHash := HashPassword("password234", decodedSalt)

	//? compare the stored signUp hash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Password is correct, You are logged in.")
	}else{
		fmt.Println("Login Failed, Please check user credentials.")
	}
}