package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "texlive",
		Desc: "LaTeX document for true scholar",
		Setup: func() {
			wrapper.ParuOnce("texlive-meta")

			// https://github.com/Glavin001/atom-beautify/issues/1792#issuecomment-327071117
			wrapper.Run("sudo", "cpan", "Unicode::GCString")
			wrapper.Run("sudo", "cpan", "App::cpanminus")
			wrapper.Run("sudo", "cpan", "YAML::Tiny")
			wrapper.Run("sudo", "perl", "-MCPAN", "-e", "install \"File::HomeDir\"")
		},
	})
}
