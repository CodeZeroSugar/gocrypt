// Package format implements the encoding functions to convert and revert binary data
package format

import (
	"encoding/binary"

	"github.com/CodeZeroSugar/gocrypt/internal/crypto"
)

func Serialize(e *crypto.EncryptedData) []byte {
	buf := make([]byte, 0)
	buf = append(buf, e.Version)
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, e.Time)
	buf = append(buf, tmp...)
	binary.BigEndian.PutUint32(tmp, e.Memory)
	buf = append(buf, tmp...)
	buf = append(buf, e.Threads)

	buf = append(buf, byte(len(e.Salt)))
	buf = append(buf, e.Salt...)

	buf = append(buf, byte(len(e.Nonce)))
	buf = append(buf, e.Nonce...)

	buf = append(buf, e.CipherText...)

	return buf
}
