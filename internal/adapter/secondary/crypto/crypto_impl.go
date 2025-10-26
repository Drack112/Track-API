package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

const KeyLength = 64

type KeyGenerator struct{}

func New() *KeyGenerator {
	return &KeyGenerator{}
}

func (g *KeyGenerator) Generate() (string, error) {
	key := make([]byte, KeyLength)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(key), nil
}
