package userdata_test

import (
	"encoding/hex"
	"raverte/userdata"
	"testing"
)

func TestValidDecryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(key)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	pt, err := userdata.DecryptApiKeys(nonce, ciphertext, k)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	if hex.EncodeToString(pt) != plaintext {
		t.Errorf("Plaintext did not match!\nExpected: %s\nGot:%s", plaintext, hex.EncodeToString(pt))
	}
}

func TestInvalidKeyDecryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(invalidkey)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	_, err = userdata.DecryptApiKeys(nonce, ciphertext, k)
	if err != nil {
		IncorrectErrReturned(err, errmsgAuth, t)
	} else {
		DidNotReturnErrError(errmsgAuth, t)
	}
}

func TestInvalidKeyLenDecryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(invalidkeylen)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	_, err = userdata.DecryptApiKeys(nonce, ciphertext, k)
	if err != nil {
		IncorrectErrReturned(err, errmsgKeylen, t)
	} else {
		DidNotReturnErrError(errmsgKeylen, t)
	}
}

func TestInvalidNonceDecryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(key)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	_, err = userdata.DecryptApiKeys(invalidnonce, ciphertext, k)
	if err != nil {
		IncorrectErrReturned(err, errmsgAuth, t)
	} else {
		DidNotReturnErrError(errmsgAuth, t)
	}
}

func TestInvalidNonceLenDecryptKeys(t *testing.T) {
	k, err := hex.DecodeString(key)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	_, err = userdata.DecryptApiKeys(invalidnoncelen, ciphertext, k)
	if err != nil {
		IncorrectErrReturned(err, errmsgNonlen, t)
	} else {
		DidNotReturnErrError(errmsgNonlen, t)
	}
}
