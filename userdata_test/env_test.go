package userdata_test

import (
	"os"
	"os/user"
	"raverte/appdata"
	"raverte/userdata"
	"runtime"
	"testing"
)

// Note: Unexported functions tested via apikeys_test.

const (
	windowsProfilePath string = "\\Raverte\\profile.json"
	macosProfilePath   string = "/Library/Application Support/Raverte/profile.json"
	linuxProfilePath   string = "/.raverte/profile.json"

	windowsKeystorePath string = "\\Raverte\\keystore.ks"
	macosKeystorePath   string = "/Library/Application Support/Raverte/keystore.ks"
	linuxKeystorePath   string = "/.raverte/keystore.ks"

	errmsgHomeDirNotFound string = "couldn't find home folder"
	errmsgAppDataNotFound string = "couldn't find appdata folder"
	errmsgAppSuppNotFound string = "couldn't find application support folder"
	errInvalidAssetMsg    string = "invalid raverte asset"
)

func TestInvalidAssetGetRaverteAsset(t *testing.T) {
	_, err := userdata.GetRaverteAsset("beans on toast")
	if err != nil {
		IncorrectErrReturned(err, errInvalidAssetMsg, t)
	} else {
		DidNotReturnErrError(errInvalidAssetMsg, t)
	}

}

func TestNoEnvErrGetRaverteAsset(t *testing.T) {
	reset := ResetEnvVars(t)
	defer reset()

	if runtime.GOOS == "windows" {
		_, err := userdata.GetRaverteAsset(appdata.PROFILE)
		if err != nil {
			IncorrectErrReturned(err, errmsgAppDataNotFound, t)
		} else {
			DidNotReturnErrError(errmsgAppDataNotFound, t)
		}
	} else if runtime.GOOS == "darwin" {
		_, err := userdata.GetRaverteAsset(appdata.PROFILE)
		if err != nil {
			IncorrectErrReturned(err, errmsgAppSuppNotFound, t)
		} else {
			DidNotReturnErrError(errmsgAppSuppNotFound, t)
		}
	} else {
		_, err := userdata.GetRaverteAsset(appdata.PROFILE)
		if err != nil {
			IncorrectErrReturned(err, errmsgHomeDirNotFound, t)
		} else {
			DidNotReturnErrError(errmsgHomeDirNotFound, t)
		}
	}
}

func TestGetRaverteAssetProfile(t *testing.T) {
	var validProfilePath string
	var err error

	currUser, err := user.Current()
	if err != nil {
		t.Errorf(err.Error())
	}
	// Refactor to use env?
	if runtime.GOOS == "windows" {
		validProfilePath = "C:\\Users\\" + currUser.Username + "\\AppData\\Local" + windowsProfilePath
	} else if runtime.GOOS == "darwin" {
		validProfilePath = "/Users/" + currUser.Username + macosProfilePath
	} else {
		validProfilePath = "/home/" + currUser.Username + linuxProfilePath
	}

	if raverteAsset, err := userdata.GetRaverteAsset(appdata.PROFILE); err != nil {
		t.Errorf(err.Error())
	} else {
		if raverteAsset != validProfilePath {
			t.Errorf("Did not return valid profile path!\nExpected: %s\nGot: %s", validProfilePath, raverteAsset)
		}
	}
}

// Refactor to use env?
func TestGetRaverteAssetKeystore(t *testing.T) {
	var validKeystorePath string
	var err error

	currUser, err := user.Current()
	if err != nil {
		t.Errorf(err.Error())
	}

	if runtime.GOOS == "windows" {
		validKeystorePath = "C:\\Users\\" + currUser.Username + "\\AppData\\Local" + windowsKeystorePath
	} else if runtime.GOOS == "darwin" {
		validKeystorePath = "/Users/" + currUser.Username + macosKeystorePath
	} else {
		validKeystorePath = "/home/" + currUser.Username + linuxKeystorePath
	}

	if raverteAsset, err := userdata.GetRaverteAsset(appdata.KEYSTORE); err != nil {
		t.Errorf(err.Error())
	} else {
		if raverteAsset != validKeystorePath {
			t.Errorf("Did not return valid profile path!\nExpected: %s\nGot: %s", validKeystorePath, raverteAsset)
		}
	}
}

func ResetEnvVars(t *testing.T) (reset func()) {
	var originalHome string
	var err error

	if runtime.GOOS == "windows" {
		originalHome, err = os.UserCacheDir()
		if err != nil {
			t.Errorf(err.Error())
		}
		err = os.Unsetenv("LocalAppData")
		if err != nil {
			t.Errorf(err.Error())
		}
	} else {
		// this is sufficient for macOS, UserConfigDir() appends to $HOME.
		originalHome, err = os.UserHomeDir()
		if err != nil {
			t.Errorf(err.Error())
		}
		err = os.Unsetenv("HOME")
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	return func() {
		if runtime.GOOS == "windows" {
			err = os.Setenv("LocalAppData", originalHome)
			if err != nil {
				t.Errorf(err.Error())
			}
		} else {
			err = os.Setenv("HOME", originalHome)
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	}
}
