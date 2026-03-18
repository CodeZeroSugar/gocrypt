package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CodeZeroSugar/gocrypt/internal/app"
	"golang.org/x/term"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Println("Not enough arguments provided")
	}

	fmt.Print("Enter secret:")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal("\nError reading input:", err)
	}
	path := args[1]
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not read provided file.")
	}

	switch command := args[0]; command {
	case "encrypt":

		encrypted, err := app.CommandEncrypt(data, password)
		if err != nil {
			log.Fatal("\nFailed to encrypt input:", err)
		}

		err = os.WriteFile(path, encrypted, 0o644)
		if err != nil {
			log.Fatal("\nFailed to write encrypted file:", err)
		}

	case "decrypt":

		decrypted, err := app.CommandDecrypt(data, password)
		if err != nil {
			log.Fatal("\nFailed to decrypt input:", err)
		}

		err = os.WriteFile(path, decrypted, 0o664)
		if err != nil {
			log.Fatal("\nFailed to write plaintext file:", err)
		}

	}
}
