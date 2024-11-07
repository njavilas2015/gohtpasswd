package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func CreateHtpasswdEntry(username, password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	entry := fmt.Sprintf("%s:%s\n", username, hash)

	return entry, nil
}

func ParseFlags() (*string, *string) {

	username := flag.String("username", "", "username")

	password := flag.String("password", "", "password")

	flag.Parse()

	if *username == "" {
		fmt.Println("Please provide the username --username")
		os.Exit(1)
	}

	if *password == "" {
		fmt.Println("Please provide the password --password")
		os.Exit(1)
	}

	return username, password
}

func main() {

	username, password := ParseFlags()

	entry, err := CreateHtpasswdEntry(*username, *password)

	if err != nil {
		fmt.Println("Error creating htpasswd entry:", err)
		return
	}

	file, err := os.OpenFile(".htpasswd", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		fmt.Println("Error opening .htpasswd file:", err)
		return
	}

	defer file.Close()

	if _, err := file.WriteString(entry); err != nil {
		fmt.Println("Error writing to .htpasswd file:", err)
	} else {
		fmt.Println(".htpasswd entry added successfully!")
	}
}
