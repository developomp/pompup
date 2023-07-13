package workflows

import (
	"github.com/developomp/pompup/internal/install"
)

var gearleverWorkflow = Workflow{
	Name: "GearLever",
	Desc: "appimage manager",
	Tags: []Tag{System, Gui},
	Setup: func() {
		install.Flatpak("it.mijorus.gearlever")
	},
}

func init() {
	register(&gearleverWorkflow)
}
