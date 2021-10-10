package userdata_test

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"raverte/appdata"
	"raverte/userdata"
	"strings"
	"testing"
)

const (
	keystoreNonce                 string = "3dafc619595406d21b4da8e1"
	keystoreOneExchangeCiphertext string = "937dae316088044a809f41106742bf24be9c6cd82ef0c12d774c36f15cd084c0a0a0bf8eeb048a2a1f575c36cc394783c5aa6e53b13457ee11575b08b88458716a6377567cb3a9f68852eafcf959e273fe92a1ef39eac004aaec90d30b08b3d3598aacd3c14ba6d0b823a37c99d1ac1aded2ef223e738d825b6bbdeff3a21b0c902ee9ca64112c3bd111f08d05ff92a483c083c987f8e0966a313d574fa89585f10713fa95995504177ff36d9570"
	keystoreTwoExchangeCiphertext string = "937dae316088044a809f41106742bf24be9c6cd82ef0c12d774c36f15cd084c0a0a0bf8eeb048a2a1f575c36cc394783c5aa6e53b13457ee11575b08b88458716a6377567cb3a9f68852eafcf959e273fe92a1ef39eac004aaec90d30b08b3d3598aacd3c14ba6d0b823a37c99d1ac1aded2ef223e738d825b6bbdeff3a21b0c902ee9ca64112c3bd111f08d05ff92a483c083c987f8e0966a313d571ef74cb6788b67263a486a97f0eb3617545d1132b9e7d4874ad04d5f9ad0f73b8ae4dcd36a96b53fa58c518eb59bdbf9ff167c83a9744b5ef3a3239dee2f5ad793f57da728d051418ef7c6278c31a0d917d61853df7426f8a1e78f4d9254783f6053a2803ebf6461152b5e1ded7f0a493905dfb3f55280977015e063b461abfb0d5f1be11ea0da2993d8a22c65cfa4693732455478516e3d96d2add5e869d5429c817b41b59f99"
	exchangeOne                   string = "Binance"
	exchangeTwo                   string = "Coinbase Pro"
	exchangeThree                 string = "Coinbase"
	invalidexchange               string = "NotBinance"
	apikey                        string = "uGOVjbgWA7FunAmGO8lsSUXNsu3eow76sz84Q18fWxnyRzBHCd3pd5nE9qa99"
	apisecret                     string = "sSUXNsu3eow76sz84Q18fWxnyRzBHCd3pd5nE9qa99HAZtuZuj6F1huXgow76"
	errmsgInvExc                  string = "invalid exchange"
	errmsgInvKeyStore             string = "invalid keystore, removing keys"
	errmsgInvPassword             string = "invalid password"
	errmsgKeyExists               string = "key already exists"
	errmsgNoKeyToRemove           string = "no key to remove"
	errmsgNoKeyFound              string = "could not find key associated with"
)

// TODO: Abstract out cases for Keystore Conditions.
func TestReturnApiKeyAndSeretApiKeys(t *testing.T) {
	var apikas userdata.ApiKaS = userdata.ApiKaS{Key: apikey, Secret: apisecret}
	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{Keys: make(map[string]userdata.ApiKaS)}
	apiKeyRing.Keys[exchangeOne] = apikas
	apiKeyRing.Keys[exchangeTwo] = apikas
	apiKeyRing.Keys[exchangeThree] = apikas

	if _, exists := apiKeyRing.ReturnApiKeyAndSecret(exchangeThree); !exists {
		t.Error("Could not find key and secret for exchange.")
	}
}

func TestKeyAndSecretDoesNotExistReturnApiKeyAndSecret(t *testing.T) {
	var apiKeyring userdata.ApiKeyRing = userdata.ApiKeyRing{}
	if _, exists := apiKeyring.ReturnApiKeyAndSecret(exchangeOne); exists {
		t.Error("False positive on key and secret search.")
	}
}
func TestUnlockKeys(t *testing.T) {
	// Test the integrity of unlockKeys unlock KNOWNKEYSTORE compare exchange and apikey &  secret
	testKeystoreSetup := KeystoreConditions("TestUnlockKeys", t)
	defer testKeystoreSetup()

	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	var profile userdata.Profile = userdata.Profile{Keystore: true}

	err := apiKeyRing.UnlockKeys(password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
		return
	}

	knownKeyAndSecret := apiKeyRing.Keys[exchangeOne]
	if knownKeyAndSecret.Key != apikey && knownKeyAndSecret.Secret != apisecret {
		t.Errorf("Did not return expected key and secret.\nExpected Key: %s\nRecovered Key: %s\nExpected Secret: %s\nRecovered Secret: %s\n", apikey, knownKeyAndSecret.Key, apisecret, knownKeyAndSecret.Secret)
	}
}

func TestRaverteAssetDoesNotExistUnlockKeys(t *testing.T) {
	testKeystoreSetup := KeystoreConditions("", t)
	defer testKeystoreSetup()

	keystore, err := userdata.GetRaverteAsset(appdata.KEYSTORE)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	var profile userdata.Profile = userdata.Profile{Keystore: true}

	var expectedErrMsg string = fmt.Sprintf("file does not exist: %s", keystore)
	err = apiKeyRing.UnlockKeys(password, profile)
	if err != nil {
		IncorrectErrReturned(err, expectedErrMsg, t)
	} else {
		DidNotReturnErrError(expectedErrMsg, t)
	}
}

func TestRaverteAssetBadPermissionsUnlockKeys(t *testing.T) {
	testKeystoreSetup := KeystoreConditions("TestRaverteAssetBadPermissionsUnlockKeys", t)
	defer testKeystoreSetup()

	keystore, err := userdata.GetRaverteAsset(appdata.KEYSTORE)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	var profile userdata.Profile = userdata.Profile{Keystore: true}

	var expectedErrMsg string = fmt.Sprintf("incorrect permissions: %s", keystore)
	err = apiKeyRing.UnlockKeys(password, profile)
	if err != nil {
		IncorrectErrReturned(err, expectedErrMsg, t)
	} else {
		DidNotReturnErrError(expectedErrMsg, t)
	}
}

func TestCompromisedKeystoreUnlockKeys(t *testing.T) {
	testKeystoreSetup := KeystoreConditions("TestCompromisedKeystoreUnlockKeys", t)
	defer testKeystoreSetup()

	keystore, err := userdata.GetRaverteAsset(appdata.KEYSTORE)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	var profile userdata.Profile = userdata.Profile{Keystore: true}

	err = apiKeyRing.UnlockKeys(password, profile)
	if err != nil {
		IncorrectErrReturned(err, errmsgInvKeyStore, t)
		if err.Error() == errmsgInvKeyStore {
			if _, err := os.Stat(keystore); err == nil {
				t.Errorf("Did not remove compromised keyfile.")
			}
		}
	} else {
		DidNotReturnErrError(errmsgInvKeyStore, t)
	}
}

func TestInvalidPasswordUnlockKeys(t *testing.T) {
	testKeystoreSetup := KeystoreConditions("TestUnlockKeys", t) // use expected conditions, as we're testing password capture.
	defer testKeystoreSetup()

	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	var profile userdata.Profile = userdata.Profile{Keystore: true}

	err := apiKeyRing.UnlockKeys(invalidpassword, profile)
	if err != nil {
		IncorrectErrReturned(err, errmsgInvPassword, t)
	} else {
		DidNotReturnErrError(errmsgInvPassword, t)
	}
}

func TestAddApiKeyAndSecret(t *testing.T) {
	testKeystoreSetup := KeystoreConditions("TestUnlockKeys", t) // use expected conditions.
	defer testKeystoreSetup()

	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	var profile userdata.Profile = userdata.Profile{Keystore: true}

	err := apiKeyRing.UnlockKeys(password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	err = apiKeyRing.AddApiKeyAndSecret(exchangeTwo, apikey, apisecret, password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	if apiKeyRing.Keys[exchangeTwo].Key != apikey && apiKeyRing.Keys[exchangeTwo].Secret != apisecret {
		t.Error("AddApiKeyAndSecret did not commit new key and secret to ApiKeyRing.")
	}

	var verifyKeyringCommitedToKeystore userdata.ApiKeyRing = userdata.ApiKeyRing{}
	err = verifyKeyringCommitedToKeystore.UnlockKeys(password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	if verifyKeyringCommitedToKeystore.Keys[exchangeTwo].Key != apikey && verifyKeyringCommitedToKeystore.Keys[exchangeTwo].Secret != apisecret {
		t.Error("AddApiKeyAndSecret did not commit new key and secret to the keystore.ks.")
	}

}

func TestNewKeystoreAddApiKeyAndSecret(t *testing.T) {
	testKeystoreSetup := KeystoreConditions("", t) // We do not want a keystore.
	defer testKeystoreSetup()

	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	var profile userdata.Profile = userdata.Profile{Keystore: false}

	err := apiKeyRing.AddApiKeyAndSecret(exchangeOne, apikey, apisecret, password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	if apiKeyRing.Keys[exchangeOne].Key != apikey && apiKeyRing.Keys[exchangeOne].Secret != apisecret {
		t.Error("AddApiKeyAndSecret did not commit new key and secrey to ApiKeyRing.")
	}

	var verifyKeyringCommitedToKeystore userdata.ApiKeyRing = userdata.ApiKeyRing{}
	err = verifyKeyringCommitedToKeystore.UnlockKeys(password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	if verifyKeyringCommitedToKeystore.Keys[exchangeOne].Key != apikey && verifyKeyringCommitedToKeystore.Keys[exchangeOne].Secret != apisecret {
		t.Error("AddApiKeyAndSecret did not commit new key and secrey to ApiKeyRing.")
	}
}

func TestKeyAlreadyExistsAddApiKeyAndSecret(t *testing.T) {
	var apikas userdata.ApiKaS = userdata.ApiKaS{Key: apikey, Secret: apisecret}
	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{Keys: make(map[string]userdata.ApiKaS)}
	apiKeyRing.Keys[exchangeOne] = apikas

	err := apiKeyRing.AddApiKeyAndSecret(exchangeOne, "", "", "", userdata.Profile{})
	if err != nil {
		IncorrectErrReturned(err, errmsgKeyExists, t)
	} else {
		DidNotReturnErrError(errmsgKeyExists, t)
	}
}

func TestInvalidExchangeAddApiKeyAndSecret(t *testing.T) {
	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	err := apiKeyRing.AddApiKeyAndSecret(invalidexchange, "", "", "", userdata.Profile{})
	if err != nil {
		IncorrectErrReturned(err, errmsgInvExc, t)
	} else {
		DidNotReturnErrError(errmsgInvExc, t)
	}
}

func TestRemoveApiKeyAndSecret(t *testing.T) {
	testKeystoreSetup := KeystoreConditions("TestRemoveApiKeyAndSecret", t)
	defer testKeystoreSetup()

	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{}
	var profile userdata.Profile = userdata.Profile{Keystore: true}

	err := apiKeyRing.UnlockKeys(password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
		return
	}
	err = apiKeyRing.RemoveApiKeyAndSecret(exchangeThree, password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
		return
	}
	var verifyKeyringCommitedToKeystore userdata.ApiKeyRing = userdata.ApiKeyRing{}
	err = verifyKeyringCommitedToKeystore.UnlockKeys(password, profile)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	_, verifyExchangeOneOnChain := verifyKeyringCommitedToKeystore.ReturnApiKeyAndSecret(exchangeOne)
	if len(verifyKeyringCommitedToKeystore.Keys) != 1 && !verifyExchangeOneOnChain {
		t.Errorf("Did not return expected key and secret, keychain should only contain %s", exchangeOne)
	}
}

func TestNoKeyRemoveApiKeyAndSecret(t *testing.T) {
	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{Keys: make(map[string]userdata.ApiKaS)}
	err := apiKeyRing.RemoveApiKeyAndSecret(exchangeOne, "", userdata.Profile{})
	if err != nil {
		IncorrectErrReturned(err, errmsgNoKeyToRemove, t)
	} else {
		DidNotReturnErrError(errmsgNoKeyToRemove, t)
	}

}

func TestNoKeyFoundRemoveApiKeyAndSecret(t *testing.T) {
	var expectedErr string = fmt.Sprintf("%s %s", errmsgNoKeyFound, exchangeThree)
	var apiKeyRing userdata.ApiKeyRing = userdata.ApiKeyRing{Keys: make(map[string]userdata.ApiKaS)}
	var apikas userdata.ApiKaS = userdata.ApiKaS{Key: apikey, Secret: apisecret}
	apiKeyRing.Keys[exchangeOne] = apikas
	apiKeyRing.Keys[exchangeTwo] = apikas

	err := apiKeyRing.RemoveApiKeyAndSecret(exchangeThree, "", userdata.Profile{})

	if err != nil {
		IncorrectErrReturned(err, expectedErr, t)
	} else {
		DidNotReturnErrError(expectedErr, t)
	}
}

// Write keystore conditions as per test requirements.
func KeystoreConditions(testType string, t *testing.T) (substitute func()) {
	// Revisit this backup required lark.
	var backupRequired bool = false
	keystore, err := userdata.GetRaverteAsset(appdata.KEYSTORE)
	if err != nil {
		t.Error(err.Error())
	}
	var keystoreCopy string
	if _, err := os.Stat(keystore); err == nil {
		backupRequired = true
		fnameArr := strings.Split(keystore, ".")
		keystoreCopy = fnameArr[0] + "_COPY." + fnameArr[1]
		os.Rename(keystore, keystoreCopy)
	}
	switch testType {
	case "TestUnlockKeys":
		if _, err := os.Create(keystore); err != nil {
			t.Error(err.Error())
		} else {
			var knownKeystoreData string = keystoreNonce + " " + salt + " " + keystoreOneExchangeCiphertext
			if err = os.WriteFile(keystore, []byte(knownKeystoreData), 0); err != nil {
				t.Error(err.Error())
			}
		}
		if err := os.Chmod(keystore, 0600); err != nil {
			t.Errorf("couldn't set permissions for %s", keystore)
		}
	case "TestRaverteAssetBadPermissionsUnlockKeys":
		if _, err := os.Create(keystore); err != nil {
			t.Error(err.Error())
		} else {
			var knownKeystoreData string = keystoreNonce + " " + salt + " " + keystoreOneExchangeCiphertext
			if err = os.WriteFile(keystore, []byte(knownKeystoreData), 0); err != nil {
				t.Error(err.Error())
			}
		}
		if err := os.Chmod(keystore, 0000); err != nil {
			t.Errorf("couldn't set permissions for %s", keystore)
		}
	case "TestCompromisedKeystoreUnlockKeys":
		if _, err := os.Create(keystore); err != nil {
			t.Error(err.Error())
		} else {
			var knownKeystoreData string = keystoreNonce + " " + salt
			if err = os.WriteFile(keystore, []byte(knownKeystoreData), 0); err != nil {
				t.Error(err.Error())
			}
		}
		if err := os.Chmod(keystore, 0600); err != nil {
			t.Errorf("couldn't set permissions for %s", keystore)
		}
	case "TestRemoveApiKeyAndSecret":
		if _, err := os.Create(keystore); err != nil {
			t.Error(err.Error())
		} else {
			var knownKeystoreData string = keystoreNonce + " " + salt + " " + keystoreTwoExchangeCiphertext
			if err = os.WriteFile(keystore, []byte(knownKeystoreData), 0); err != nil {
				t.Error(err.Error())
			}
			if err := os.Chmod(keystore, 0600); err != nil {
				t.Errorf("couldn't set permissions for %s", keystore)
			}
		}
	}

	return func() {
		if backupRequired {
			_ = os.Remove(keystore)
			if _, err := os.Stat(keystore); errors.Is(err, fs.ErrNotExist) {
				os.Rename(keystoreCopy, keystore)
			} else {
				t.Errorf("Failed to restore original keystore from %s", keystoreCopy)
			}
		}
	}
}
