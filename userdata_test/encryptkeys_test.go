package userdata_test

import (
	"encoding/hex"
	"raverte/userdata"
	"testing"
)

// Test Vectors https://golang.org/src/crypto/cipher/gcm_test.go
const (
	plaintext       = ""
	ciphertext      = "250327c674aaf477aef2675748cf6971"
	key             = "11754cd72aec309bf52f7687212e8957"
	invalidkey      = "01164cd72aec309bf52f7687212e8957"
	invalidkeylen   = "01164cd72aec309bf52f768721"
	nonce           = "3c819d9a9bed087615030b65"
	invalidnonce    = "c0019d9a9bed087615030b65"
	invalidnoncelen = "c0019d9a9bed087615030b"
	errmsgAuth      = "cipher: message authentication failed"
	errmsgKeylen    = "crypto/aes: invalid key size 13"
	errmsgNonlen    = "invalid nonce length"
)

func TestValidEncryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(key)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	}
	n, err := hex.DecodeString(nonce)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	}
	ct, err := userdata.EncryptApiKeys([]byte(plaintext), k, n)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	} else {
		if ct != ciphertext {
			t.Errorf("Ciphertext did not match!\nExpected: % x\nGot: % x", ciphertext, ct)
		}
	}
}

func TestInvalidKeyLenEncryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(invalidkeylen)
	if err != nil {
		t.Errorf("Unxepected error: %s\n", err.Error())
	}
	n, err := hex.DecodeString(nonce)
	if err != nil {
		t.Errorf("Unxepected error: %s\n", err.Error())
	}

	_, err = userdata.EncryptApiKeys([]byte(plaintext), k, n)
	if err != nil {
		if err.Error() != errmsgKeylen {
			t.Errorf("Did not get expected error.\nExpected: %s\nGot: %s\n", errmsgKeylen, err.Error())
		}
	} else {
		t.Errorf("Did not return expected error:%s\n", errmsgKeylen)
	}
}

func TestInvalidNonceLenEncryptKeys(t *testing.T) {
	k, err := hex.DecodeString(key)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	}
	n, err := hex.DecodeString(invalidnoncelen)
	if err != nil {
		t.Errorf("Unxepected error: %s\n", err.Error())
	}
	_, err = userdata.EncryptApiKeys([]byte(plaintext), k, n)
	if err != nil {
		if err.Error() != errmsgNonlen {
			t.Errorf("Did not get expected error.\nExpected: %s\nGot: %s\n", errmsgNonlen, err.Error())
		}
	} else {
		t.Errorf("Did not return expected error: %s\n", errmsgNonlen)
	}
}
