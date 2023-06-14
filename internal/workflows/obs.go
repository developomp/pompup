package workflows

import "github.com/developomp/pompup/internal/install"

func init() {
	register(&Workflow{
		Name: "OBS",
		Desc: "Screen recording and streaming utility",
		Setup: func() {
			install.Flatpak("com.obsproject.Studio")
			install.Flatpak("com.obsproject.Studio.Plugin.InputOverlay")
		},
	})
}
