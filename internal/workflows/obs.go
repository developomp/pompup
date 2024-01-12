package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "OBS",
		Desc: "Screen recording and streaming utility",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("com.obsproject.Studio")
			install.Paru("v4l2loopback-dkms") // for virtual camera
		},
	})
}
