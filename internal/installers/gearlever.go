package installers

import (
	"fmt"

	"github.com/developomp/pompup/internal/wrapper"
)

var gearleverInstaller = Installer{
	Name: "GearLever",
	Desc: "appimage manager",
	Tags: []Tag{System, Gui},
	Setup: func() {
		if wrapper.IsBinInstalled("it.mijorus.gearlever") {
			return
		}

		// set default appimage location to "~/Programs/AppImages"
		wrapper.Run(
			"sed",
			"-i",
			fmt.Sprintf("/^appimages-default-folder=/c\\appimages-default-folder='%s/Programs/AppImages'", wrapper.GetHomeDir()),
			"/home/pomp/.var/app/it.mijorus.gearlever/config/glib-2.0/settings/keyfile",
		)

		wrapper.FlatpakOnce("it.mijorus.gearlever")
	},
}

func init() {
	register(&gearleverInstaller)
}
