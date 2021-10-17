package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	// TODO: Create a checklist that must be satisfied when adding new Exchanges
	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Raverte",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
		//DisableInspector: true, // disable for release.
	})
	raverte := RaverteInit()

	app.Bind(raverte)
	app.Bind(raverte.Profile.ReturnSelf())
	app.Run()
}
