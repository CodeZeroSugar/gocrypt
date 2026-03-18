// Package crypto contains all logic regarding the encryption and decryption of provided data
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/argon2"
)

func Encrypt(plaintext, password []byte) (*EncryptedData, error) {
	salt, err := generateSalt(16)
	if err != nil {
		return nil, err
	}
	key := argon2.Key([]byte(password), salt, Time, Memory, Threads, 32)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, []byte{Version})
	cryptData := EncryptedData{
		Version:    Version,
		Time:       Time,
		Memory:     Memory,
		Threads:    Threads,
		Salt:       salt,
		Nonce:      nonce,
		CipherText: ciphertext,
	}

	return &cryptData, nil
}
