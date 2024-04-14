package installers

import (
	"fmt"

	"github.com/developomp/pompup/internal/wrapper"
)

func init() {
	register(&Installer{
		Name: "GearLever",
		Desc: "appimage manager",
		Setup: func() {
			wrapper.FlatpakOnce("it.mijorus.gearlever")

			// set default appimage location to "~/Programs/AppImages"
			// Will fail if the app has never been launched in a graphical environment.
			// Requires re-run of pompup
			wrapper.Run(
				"sed",
				"-i",
				fmt.Sprintf("/^appimages-default-folder=/c\\appimages-default-folder='%s/Programs/AppImages'", wrapper.GetHomeDir()),
				wrapper.InHome(".var/app/it.mijorus.gearlever/config/glib-2.0/settings/keyfile"),
			)
		},
	})
}
