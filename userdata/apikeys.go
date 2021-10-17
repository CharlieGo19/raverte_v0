package userdata

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"raverte/appdata"
	"strings"
)

// API Key & Secret
type ApiKaS struct {
	Key    string `json:"k"`
	Secret string `json:"s"`
}

type ApiKeyRing struct {
	Keys map[string]ApiKaS `json:"keys"`
}

// TODO: Add WAILS logger here.

// Searches for and returns API Key and Secret associated with the provided exchange.
//
// Returns ApiKaS and true if the key (exchange) exists on the keyring, else returns nil and false.
func (a *ApiKeyRing) ReturnApiKeyAndSecret(exchange string) (ApiKaS, bool) {
	key, exists := a.Keys[exchange]
	return key, exists
}

// Adds API Key and Secret associated with the provided exchange and updates the keystore and profile (if required).
//
// Returns error if: unsupported exchange, password is invalid or problem submitting new data to profile & keystore.
func (a *ApiKeyRing) AddApiKeyAndSecret(exchange, newKey, newSecret, password string, profile Profile) error {
	// to determine if this is a new keystore & to be able to relflect such in profile after initialising keystore.
	var exchangeExists bool = false
	for _, v := range appdata.EXCHANGES {
		if strings.EqualFold(exchange, v) {
			exchangeExists = true
			break
		}
	}

	if !exchangeExists {
		return errors.New("invalid exchange")
	}

	var updateProfile bool = false
	// TODO: Do some checks on keys, platform specific.
	if len(a.Keys) == 0 {
		// If first key, keystore should not exist, create asset.
		a.Keys = make(map[string]ApiKaS)
		updateProfile = true
		if filePath, err := GetRaverteAsset(appdata.KEYSTORE); err != nil {
			return err
		} else {
			if err = configureRaverteAsset(filePath, appdata.KEYSTORE); err != nil {
				return err
			}
		}
	} else {
		if _, e := a.ReturnApiKeyAndSecret(exchange); e {
			return errors.New("key already exists")
		}
	}

	a.Keys[exchange] = ApiKaS{Key: newKey, Secret: newSecret}
	if err := a.writeKeysAndSecrets(password, profile); err != nil {
		return err
	} else {
		if updateProfile {
			if err := profile.UpdateKeystore(updateProfile); err != nil {
				return err
			}
		}
	}

	return nil
}

// Removes API Key and Secret associated with the provided exchange and updates the keystore and profile (if required).
//
// Returns error if: there is no key to remove, there is no key associated with provided exchange, there is a problem destroying keystore and updating profile.
func (a *ApiKeyRing) RemoveApiKeyAndSecret(exchange, password string, profile Profile) error {
	if len(a.Keys) == 0 {
		return errors.New("no key to remove")
		// set keystoreFalse: false.
	}

	keyPath, err := GetRaverteAsset(appdata.KEYSTORE)
	if err != nil {
		return err
	}

	if _, e := a.ReturnApiKeyAndSecret(exchange); !e {
		return fmt.Errorf("could not find key associated with %s", exchange)
	}

	if len(a.Keys) == 1 {
		delete(a.Keys, exchange)
		if err := profile.UpdateKeystore(false); err != nil {
			return err
		}
		destroyRaverteAsset(keyPath)
	} else {
		if err := a.writeKeysAndSecrets(password, profile); err != nil {
			return err
		}
	}

	return nil
}

// Unlocks keystore and loads users API Keys and Secrets into memory.
//
// Returns error if: invalid keystore (does not exist/permissions have been changed), will also remove keystore if the keystore has been altered or the password is invalid.
func (a *ApiKeyRing) UnlockKeys(password string, profile Profile) error {
	keyPath, err := GetRaverteAsset(appdata.KEYSTORE)
	if err != nil {
		return err
	}

	if err := checkRaverteAsset(keyPath); err != nil {
		// check keystore exists and permissions are appropriate.
		return fmt.Errorf("%s: %s", err.Error(), keyPath)
	}

	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return err
	}
	keystore := strings.Split(string(keyData), " ")
	var nonce, salt, ciphertext string
	if len(keystore) != 3 {
		// Someone has potentially tampered with keystore.
		profile.UpdateKeystore(false)
		destroyRaverteAsset(keyPath)
		return errors.New("invalid keystore, removing keys")
	} else {
		nonce = keystore[0]
		salt = keystore[1]
		ciphertext = keystore[2]
	}

	_, key, err := GenerateKey(password, salt, false)
	if err != nil {
		if err.Error() == "invalid password length" {
			return errors.New("invalid password") // this will not be a new password, therefore it must be incorrect.
		}
		return err
	}
	if data, err := DecryptApiKeys(nonce, ciphertext, key); err != nil {
		if err.Error() == "cipher: message authentication failed" {
			// TODO: Add a counter, after x attempts offer to delete keys.
			return errors.New("invalid password") // or nonce or salt
		} else {
			return err
		}
	} else {
		if err = json.Unmarshal(data, a); err != nil {
			return err
		}
	}

	return nil
}

// Commits API Keys and Secrets that are currently stored in memory to the keystore.
//
// Returns error if: invalid keystore (does not exist/permissions have been changed) or invalid password
func (a *ApiKeyRing) writeKeysAndSecrets(password string, profile Profile) error {
	keyPath, err := GetRaverteAsset(appdata.KEYSTORE)
	if err != nil {
		return err
	}

	var nonce, salt, ciphertext string
	var key []byte

	if !profile.Keystore {
		// When we have a new keystore.

		salt, key, err = GenerateKey(password, "", true)
		if err != nil {
			return err
		}

	} else {
		keyData, err := os.ReadFile(keyPath)
		if err != nil {
			// Shouldn't really hit this, if we do, env.go is not good enough!
			return err
		}

		keystore := strings.Split(string(keyData), " ")

		if len(keystore) != 3 {
			// Someone has potentially tampered with keystore.
			profile.UpdateKeystore(false)
			destroyRaverteAsset(keyPath)
			return errors.New("invalid keystore, removing keys")
		} else {
			nonce = keystore[0]
			salt = keystore[1]
			ciphertext = keystore[2]
		}

		// Unlock keys to ensure password continuity.
		_, key, err = GenerateKey(password, salt, false)
		if err != nil {
			return err
		} else {
			if _, err = DecryptApiKeys(nonce, ciphertext, key); err != nil {
				if err.Error() == "cipher: message authentication failed" {
					// TODO: Add a counter, after x attempts offer to delete keys.
					return errors.New("invalid password") // or nonce or salt
				} else {
					return err
				}
			}
		}
	}

	var newNonce []byte = make([]byte, 12)
	_, err = io.ReadFull(rand.Reader, newNonce)
	if err != nil {
		return err
	}
	plaintext, err := json.Marshal(a)
	if err != nil {
		return err
	}
	newCiphertext, err := EncryptApiKeys(plaintext, key, newNonce)
	if err != nil {
		return err
	}

	var newKeyFileData string = hex.EncodeToString(newNonce) + " " + salt + " " + newCiphertext
	if err = os.WriteFile(keyPath, []byte(newKeyFileData), 0); err != nil {
		return err
	}

	return nil
}
