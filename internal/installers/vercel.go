package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "vercel CLI",
		Desc: "https://vercel.com/docs/cli",
		Setup: func() {
			wrapper.Run("pnpm", "i", "-g", "vercel")
		},
	})
}
