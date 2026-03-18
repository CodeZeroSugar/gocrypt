package format

import (
	"encoding/binary"

	"github.com/CodeZeroSugar/gocrypt/internal/crypto"
)

func Deserialize(data []byte) (*crypto.EncryptedData, error) {
	cryptData := crypto.EncryptedData{}
	i := 0

	cryptData.Version = data[i]
	i++

	cryptData.Time = binary.BigEndian.Uint32(data[i : i+4])
	i += 4

	cryptData.Memory = binary.BigEndian.Uint32(data[i : i+4])
	i += 4

	cryptData.Threads = data[i]
	i++

	saltLen := int(data[i])
	i++

	cryptData.Salt = data[i : i+saltLen]
	i += saltLen

	nonceLen := int(data[i])
	i++

	cryptData.Nonce = data[i : i+nonceLen]
	i += nonceLen

	cryptData.CipherText = data[i:]

	return &cryptData, nil
}
