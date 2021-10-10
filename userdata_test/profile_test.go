package userdata_test

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"raverte/appdata"
	"raverte/userdata"
	"reflect"
	"strings"
	"testing"
)

const (
	defaultProfileJson  string = `{"name":"","keystore":false}`
	errmsgProfileExists string = "profile already exists"
)

var defaultProfile userdata.Profile = userdata.Profile{Keystore: false}

func TestInitialiseProfile(t *testing.T) {
	var testProfile userdata.Profile = userdata.Profile{}
	err := testProfile.InitialiseProfile()
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}
	// may need to revise this check as profile evolves.
	if !reflect.DeepEqual(defaultProfile, testProfile) {
		t.Error("InitialiseProfile did not set expected default values.")
	}
}

func TestInitialiseProfileWriteProfile(t *testing.T) {
	profileConditions := ProfileConditions("", t)
	defer profileConditions()

	profile, err := userdata.GetRaverteAsset(appdata.PROFILE)
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	var testProfile userdata.Profile = userdata.Profile{}
	err = testProfile.InitialiseProfile()
	if err != nil {
		UnexpectedErrError(err.Error(), t)
	}

	if readProfileData, err := os.ReadFile(profile); err != nil {
		t.Error(err.Error())
	} else {
		if string(readProfileData[:]) != defaultProfileJson {
			t.Errorf("Did not read expected data from default profile.json.\nExpected: %s\nGot: %s", defaultProfileJson, readProfileData)
		}
	}
}

func TestProfileExistsErrInitialiseProfile(t *testing.T) {
	profileConditions := ProfileConditions("TestProfileExistsErrInitialiseProfile", t)
	defer profileConditions()

	var testProfile userdata.Profile = userdata.Profile{}
	err := testProfile.InitialiseProfile()
	if err != nil {
		if err.Error() != errmsgProfileExists {
			IncorrectErrReturned(err, errmsgProfileExists, t)
		}
	} else {
		DidNotReturnErrError(errmsgProfileExists, t)
	}
}

func TestUpdateKeystore(t *testing.T) {
	profileConditions := ProfileConditions("TestUpdateKeystore", t)
	defer profileConditions()

	var testProfile userdata.Profile = userdata.Profile{Keystore: false}
	testProfile.UpdateKeystore(true)

	profilePath, err := userdata.GetRaverteAsset(appdata.PROFILE)
	if err != nil {
		t.Error(err.Error())
	}

	readData, err := os.ReadFile(profilePath)
	if err != nil {
		t.Error(err.Error())
	}

	var validateProfile userdata.Profile

	err = json.Unmarshal(readData, &validateProfile)
	if err != nil {
		t.Error(err.Error())
	}

	if !validateProfile.Keystore {
		t.Errorf("Did not read expected Keystore value from profile.\nExpected: true\nGot: %t", validateProfile.Keystore)
	}
}

func ProfileConditions(testType string, t *testing.T) (substitute func()) {
	profile, err := userdata.GetRaverteAsset(appdata.PROFILE)
	if err != nil {
		t.Error(err.Error())
	}

	var profileCopy string
	if _, err := os.Stat(profile); err == nil {
		fnameArr := strings.Split(profile, ".")
		profileCopy = fnameArr[0] + "_COPY." + fnameArr[1]
		os.Rename(profile, profileCopy)
	}

	switch testType {
	case "TestProfileExistsErrInitialiseProfile":
		if _, err := os.Create(profile); err != nil {
			t.Error(err.Error())
		}
		if err := os.Chmod(profile, 0600); err != nil {
			t.Errorf("couldn't set permissions for %s", profile)
		}
	case "TestUpdateKeystore":
		if _, err := os.Create(profile); err != nil {
			t.Error(err.Error())
		}
		if err := os.Chmod(profile, 0600); err != nil {
			t.Errorf("couldn't set permissions for %s", profile)
		}
		if err := os.WriteFile(profile, []byte(defaultProfileJson), 0); err != nil {
			t.Error(err.Error())
		}
	}

	return func() {
		_ = os.Remove(profile)
		if _, err := os.Stat(profile); errors.Is(err, fs.ErrNotExist) {
			os.Rename(profileCopy, profile)
		} else {
			t.Errorf("Failed to restore original profile from %s", profileCopy)
		}
	}
}
