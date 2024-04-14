package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "Snoop",
		Desc: "file content searcher",
		Setup: func() {
			wrapper.FlatpakOnce("de.philippun1.Snoop")
		},
	})
}
