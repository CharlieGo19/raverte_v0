package userdata

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

// TODO: look at adding 2fa.

// Encrypts a users API Keys, return encoded ciphertext.
func EncryptApiKeys(plaintext, key, nonce []byte) (string, error) {
	if len(nonce) != 12 {
		return "", errors.New("invalid nonce length")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	var cipherText []byte = aesgsm.Seal(nil, nonce, plaintext, nil)

	var ct string = hex.EncodeToString(cipherText)

	return ct, nil
}
