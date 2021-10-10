package userdata_test

import (
	"bytes"
	"raverte/userdata"
	"testing"
)

const (
	password           = "CardiBGreaterThanNicki"
	invalidpassword    = "dojaispartofthisconversation"
	invalidpasswordlen = "CardiB"
	salt               = "5e821054a00b9135cd71dc689e992a3c0ac3e08909ebdc77c034763d7b0076364dc40213fd8f827c5b44e8b3975d674df8b42d9ef9fe85357791a121b9ab52bb27915c2ae35981dc07cb02892fabab3cdbf102eade4fa4576a67be823df3f9c4d6b45c2ddee9dce937769835187f60d0e671eec649bce0e259c428d0532f13c2"
	errPasswordLen     = "invalid password length"
	errSaltLen         = "compromised salt"
)

var testKey []byte = []byte{124, 228, 9, 216, 168, 13, 96, 149, 155, 22, 3, 6, 132, 232, 56, 249, 31, 183, 65, 89, 231, 201, 115, 16, 79, 27, 53, 79, 69, 113, 76, 217}

func TestInvalidPassword(t *testing.T) {
	_, _, err := userdata.GenerateKey(invalidpasswordlen, salt, false)
	if err != nil {
		IncorrectErrReturned(err, errPasswordLen, t)
	} else {
		DidNotReturnErrError(errPasswordLen, t)
	}
}

func TestCompromisedSalt(t *testing.T) {
	var badSalt string = salt + "2a"
	_, _, err := userdata.GenerateKey(password, badSalt, false)
	if err != nil {
		IncorrectErrReturned(err, errSaltLen, t)
	} else {
		DidNotReturnErrError(errSaltLen, t)
	}
}

func TestNewGenerateKey(t *testing.T) {
	s, k, err := userdata.GenerateKey(password, "", true)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	if len(s) != 256 {
		t.Errorf("Salt is not long enough, got: %d, expected: 256\n", len(salt))
	}
	if len(k) != 32 {
		t.Errorf("Key is not long enough, got: %d, expected: 32\n", len(k))
	}
}

func TestExistingGenerateKey(t *testing.T) {
	s, key, err := userdata.GenerateKey(password, salt, false)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	if s != salt {
		t.Errorf("Bad salt:\nExpected: %s\nGot: %s\n", salt, s)
	}
	if !bytes.Equal(testKey, key) {
		t.Errorf("Bad key:\nExpected: % x\nGot: % x\n", testKey, key)
	}
}

func IncorrectErrReturned(err error, errmsg string, t *testing.T) {
	if err.Error() != errmsg {
		t.Errorf("Did not get expected error.\nExpected: %s\nGot: %s\n", errmsg, err.Error())
	}
}

func DidNotReturnErrError(errmsg string, t *testing.T) {
	t.Errorf("Did not return expected error: %s\n", errmsg)
}

func UnexpectedErrError(errmsg string, t *testing.T) {
	t.Errorf("Got unexpected error: %s", errmsg)
}
