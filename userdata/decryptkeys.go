package userdata

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

// Decrypts users API Keys, returns plaintext and an error.
func DecryptApiKeys(nonce, ciphertext string, key []byte) ([]byte, error) {
	n, err := hex.DecodeString(nonce)
	if err != nil {
		return make([]byte, 0), err
	}
	if len(n) != 12 {
		return make([]byte, 0), errors.New("invalid nonce length")
	}

	ct, err := hex.DecodeString(ciphertext)
	if err != nil {
		return make([]byte, 0), err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return make([]byte, 0), err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return make([]byte, 0), err
	}

	plaintext, err := aesgcm.Open(nil, n, ct, nil)
	if err != nil {
		return make([]byte, 0), err
	}
	return plaintext, nil
}
