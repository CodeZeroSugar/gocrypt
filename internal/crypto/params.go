package crypto

const (
	Version = 1
	Time    = 3
	Memory  = 64 * 1024
	Threads = 4
)

type EncryptedData struct {
	Version    byte
	Time       uint32
	Memory     uint32
	Threads    uint8
	Salt       []byte
	Nonce      []byte
	CipherText []byte
}
