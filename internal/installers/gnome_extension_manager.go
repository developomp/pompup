package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "GNOME Extension Manager",
		Desc: "GNOME extensions website in an app",
		Setup: func() {
			wrapper.FlatpakOnce("com.mattjakeman.ExtensionManager")
		}})
}
