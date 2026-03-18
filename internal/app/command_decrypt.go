package app

import (
	"github.com/CodeZeroSugar/gocrypt/internal/crypto"
	"github.com/CodeZeroSugar/gocrypt/internal/format"
)

func CommandDecrypt(ciphertext, password []byte) ([]byte, error) {
	deserialized, err := format.Deserialize(ciphertext)

	plaintext, err := crypto.Decrypt(deserialized, password)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
