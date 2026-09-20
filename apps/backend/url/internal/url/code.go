package url

import (
	"crypto/rand"
	"math/big"
)

func generateCode() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLength = 6

	buffer := make([]byte, codeLength)
	for i := range buffer {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		buffer[i] = alphabet[index.Int64()]
	}
	return string(buffer)
}
