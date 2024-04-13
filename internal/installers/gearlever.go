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
			wrapper.Run(
				"sed",
				"-i",
				fmt.Sprintf("/^appimages-default-folder=/c\\appimages-default-folder='%s/Programs/AppImages'", wrapper.GetHomeDir()),
				"/home/pomp/.var/app/it.mijorus.gearlever/config/glib-2.0/settings/keyfile",
			)
		},
	})
}
