package userdata_test

import (
	"encoding/hex"
	"raverte/userdata"
	"testing"
)

func TestValidDecryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(key)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	}
	pt, err := userdata.DecryptApiKeys(nonce, ciphertext, k)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	}
	if hex.EncodeToString(pt) != plaintext {
		t.Errorf("Plaintext did not match!\nExpected: %s\nGot:%s", plaintext, hex.EncodeToString(pt))
	}
}

func TestInvalidKeyDecryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(invalidkey)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	}

	_, err = userdata.DecryptApiKeys(nonce, ciphertext, k)
	if err != nil {
		if err.Error() != errmsgAuth {
			t.Errorf("Did not get expected error.\nExpected: %s\nGot: %s\n", errmsgAuth, err.Error())
		}
	} else {
		t.Errorf("Did not return expected error: %s\n", errmsgAuth)
	}
}

func TestInvalidKeyLenDecryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(invalidkeylen)
	if err != nil {
		t.Errorf("Unxepected error: %s\n", err.Error())
	}

	_, err = userdata.DecryptApiKeys(nonce, ciphertext, k)
	if err != nil {
		if err.Error() != errmsgKeylen {
			t.Errorf("Did not get expected error.\nExpected: %s\nGot: %s\n", errmsgKeylen, err.Error())
		}
	} else {
		t.Errorf("Did not return expected error:%s\n", errmsgKeylen)
	}
}

func TestInvalidNonceDecryptApiKeys(t *testing.T) {
	k, err := hex.DecodeString(key)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	}
	_, err = userdata.DecryptApiKeys(invalidnonce, ciphertext, k)
	if err != nil {
		if err.Error() != errmsgAuth {
			t.Errorf("Did not get expected error.\nExpected: %s\nGot: %s\n", errmsgAuth, err.Error())
		}
	} else {
		t.Errorf("Did not return expected error: %s\n", errmsgAuth)
	}
}

func TestInvalidNonceLenDecryptKeys(t *testing.T) {
	k, err := hex.DecodeString(key)
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
	}
	_, err = userdata.DecryptApiKeys(invalidnoncelen, ciphertext, k)
	if err != nil {
		if err.Error() != errmsgNonlen {
			t.Errorf("Did not get expected error.\nExpected: %s\nGot: %s\n", errmsgNonlen, err.Error())
		}
	} else {
		t.Errorf("Did not return expected error: %s\n", errmsgNonlen)
	}
}
