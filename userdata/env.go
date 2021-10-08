package userdata

import (
	"errors"
	"fmt"
	"os"
	"raverte/appdata"
	"runtime"
	"strings"
)

var errHomeDirNotFound error = errors.New("couldn't find home folder")
var errAppDataNotFound error = errors.New("couldn't find appdata folder")
var errAppSuppNotFound error = errors.New("couldn't find application support folder")
var errInvalidAsset error = errors.New("invalid raverte asset")

var windowsRaverteFolderPath string = "\\Raverte\\"
var macosRaverteFolderPath string = "/Raverte/"
var linuxRaverteFolderPath string = "/.raverte/"

// Returns path to directory that should be holding the data for Raverte to function effectively.
func GetRaverteAsset(asset string) (string, error) {
	var filePath string

	if asset != appdata.KEYSTORE && asset != appdata.PROFILE {
		return "", errInvalidAsset
	}

	if runtime.GOOS == "windows" {
		appData, err := os.UserCacheDir()
		if err != nil {
			return "", errAppDataNotFound
		}
		filePath = appData + windowsRaverteFolderPath
	} else if runtime.GOOS == "darwin" {
		appData, err := os.UserConfigDir()
		if err != nil {
			return "", errAppSuppNotFound
		}
		filePath = appData + macosRaverteFolderPath
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", errHomeDirNotFound
		}
		filePath = homeDir + linuxRaverteFolderPath

	}

	return filePath + asset, nil
}

// Checks if Raverte asset exists and has permissions defined on creation.
//
// Returns true if folder/file exists and permissions of afore mentioned are 0600.
func checkRaverteAsset(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	if info, err := os.Stat(path); err != nil {
		return false
	} else if info.Mode().Perm() != 0600 {
		return false
	}
	return true
}

// Creates and configures Raverte assets, permissions allow only user to read and write.
func configureRaverteAsset(filePath, asset string) error {
	var raverteFolder string = strings.ReplaceAll(filePath, asset, "")
	if _, err := os.Stat(raverteFolder); os.IsNotExist(err) {
		err = os.Mkdir(filePath, 0600)
		if err != nil {
			return errors.New("couldn't create raverte directory")
		}
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if _, err := os.Create(filePath); err != nil {
			return fmt.Errorf("couldn't create %s", filePath)
		}
	}
	if err := os.Chmod(filePath, 0600); err != nil {
		return fmt.Errorf("couldn't set permissions for %s", filePath)
	}
	return nil
}

// Deletes Raverte asset.
func destroyRaverteAsset(asset string) {
	os.Remove(asset)
}
