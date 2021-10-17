package main

import (
	"fmt"
	"raverte/userdata"
)

type Raverte struct {
	ChartOnlyMode bool
	Profile       *userdata.Profile
	KeyRing       *userdata.ApiKeyRing
}

func RaverteInit() *Raverte {
	return &Raverte{}
}

func (r *Raverte) GetChartOnlyMode() bool {
	return r.ChartOnlyMode
}

func (r *Raverte) StartUp() error {
	r.Profile = &userdata.Profile{}
	var err error = r.Profile.LoadProfile()
	if err != nil {
		if err.Error() == "file does not exist" {
			err = r.Profile.InitialiseProfile()
			if err != nil {
				// TODO: Add wiki.
				var raverteAssetError string = "https://someGood.Stuff"
				//lint:ignore ST1005 error will not be wrapped.
				return fmt.Errorf("Could not initialise profile, because: %s, please consult: %s.", err.Error(), raverteAssetError)
			}
		}
	}

	r.ChartOnlyMode = !r.Profile.Keystore

	return nil
}

func (r *Raverte) UnlockKeys(password string) error {
	r.KeyRing = &userdata.ApiKeyRing{}

	err := r.KeyRing.UnlockKeys(password, *r.Profile)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
