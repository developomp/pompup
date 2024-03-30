package installers

import "github.com/developomp/pompup/internal/wrapper"

var gearleverInstaller = Installer{
	Name: "GearLever",
	Desc: "appimage manager",
	Tags: []Tag{System, Gui},
	Setup: func() {
		wrapper.Flatpak("it.mijorus.gearlever")
	},
}

func init() {
	register(&gearleverInstaller)
}
