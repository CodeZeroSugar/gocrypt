// Package app contains the orchestration of cryptographic functions and encoding features.
package app

import (
	"github.com/CodeZeroSugar/gocrypt/internal/crypto"
	"github.com/CodeZeroSugar/gocrypt/internal/format"
)

func CommandEncrypt(plaintext, password []byte) ([]byte, error) {
	cryptData, err := crypto.Encrypt(plaintext, password)
	if err != nil {
		return nil, err
	}

	serialized := format.Serialize(cryptData)

	return serialized, nil
}
