package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "OBS",
		Desc: "Screen recording and streaming utility",
		Setup: func() {
			wrapper.FlatpakOnce("com.obsproject.Studio")
			wrapper.FlatpakOnce("com.obsproject.Studio.Plugin.OBSVkCapture")
			wrapper.ParuOnce("v4l2loopback-dkms") // for virtual camera
		},
	})
}
