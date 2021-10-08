package userdata

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (
	n int = 1048576 // 2^20
	r int = 8
	p int = 1
	k int = 32
)

// Generates key, requires password length of at least 8 characters, if newDoc is true, you can pass an empty string for userSalt.
//
// Returns encoded salt, key and an error.
//
// Returns error if salt or password length do not meet requirements.
func GenerateKey(password, userSalt string, newDoc bool) (string, []byte, error) {
	var salt []byte

	if len(password) < 8 {
		return "", make([]byte, 0), errors.New("invalid password length")
	}

	if !newDoc && len(userSalt) != 256 {
		return "", make([]byte, 0), errors.New("compromised salt")
	}

	if newDoc {
		salt = make([]byte, 128)
		_, err := io.ReadFull(rand.Reader, salt)
		if err != nil {
			return "", make([]byte, 0), err
		}
	} else {
		var err error // required otherwise salt scope will change?
		salt, err = hex.DecodeString(userSalt)
		if err != nil {
			return "", make([]byte, 0), err
		}
	}

	key, err := scrypt.Key([]byte(password), salt, n, r, p, k)
	if err != nil {
		return "", make([]byte, 0), err
	}

	var s string = hex.EncodeToString(salt)

	return s, key, nil
}
