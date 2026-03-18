package crypto

import (
	"crypto/aes"
	"crypto/cipher"

	"golang.org/x/crypto/argon2"
)

func Decrypt(ciphertext *EncryptedData, password []byte) ([]byte, error) {
	key := argon2.Key([]byte(password), ciphertext.Salt, ciphertext.Time, ciphertext.Memory, ciphertext.Threads, 32)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	daesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := daesgcm.Open(nil, ciphertext.Nonce, ciphertext.CipherText, []byte{ciphertext.Version})
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
