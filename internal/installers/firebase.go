package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "firebase CLI",
		Desc: "firebase CLI",
		Setup: func() {
			if !wrapper.IsBinInstalled("firebase") {
				wrapper.Run("npm", "install", "-g", "firebase-tools")
			}
		},
		Reminders: []string{
			"Run firebase login",
		},
	})
}
