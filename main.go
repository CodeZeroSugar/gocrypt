package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/CodeZeroSugar/gocrypt/internal/app"
	"github.com/CodeZeroSugar/gocrypt/internal/termfx"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		app.CommandUsage()
		return
	}
	args := os.Args[1:]
	if os.Args[1] == "help" || os.Args[1] == "usage" {
		app.CommandUsage()
		return
	}
	if len(args) != 3 {
		fmt.Print("\nNot enough arguments provided")
		app.CommandUsage()
		return
	}

	fmt.Print("Enter secret:")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Print("\nError reading input:", err)
		return
	}

	inPath := args[1]
	outPath := args[2]

	data, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Print("\nCould not read provided file.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())

	switch command := args[0]; command {
	case "encrypt":
		fmt.Print("\nConfirm secret:")
		confirm, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Print("\nError reading input:", err)
			return
		}

		if !bytes.Equal(password, confirm) {
			fmt.Print("\nConfirmation did not match password")
			return

		}

		go termfx.Scrambler(ctx)

		encrypted, err := app.CommandEncrypt(data, password)
		if err != nil {
			fmt.Print("\nFailed to encrypt input:", err)
			return
		}

		outPath += ".gcrypt"
		err = os.WriteFile(outPath, encrypted, 0o600)
		if err != nil {
			fmt.Print("\nFailed to write encrypted file:", err)
			return
		}

		cancel()
		fmt.Print("\r\033[K")
		fmt.Println("\nSuccessfully encrypted file:", outPath)

	case "decrypt":

		fmt.Println("")
		go termfx.Scrambler(ctx)

		decrypted, err := app.CommandDecrypt(data, password)
		if err != nil {
			fmt.Print("\nFailed to decrypt input:", err)
			return
		}

		err = os.WriteFile(outPath, decrypted, 0o644)
		if err != nil {
			fmt.Print("\nFailed to write plaintext file:", err)
			return
		}

		cancel()
		fmt.Print("\r\033[K")
		fmt.Println("\nSuccessfully decrypted file:", outPath)

	default:
		fmt.Printf("\n'%s' is not a valid command", command)
		app.CommandUsage()
	}
}
