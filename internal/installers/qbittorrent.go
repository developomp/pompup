package installers

import "github.com/developomp/pompup/internal/wrapper"

func init() {
	register(&Installer{
		Name: "qBittorrent",
		Desc: "torrent done right",
		Setup: func() {
			wrapper.FlatpakOnce("org.qbittorrent.qBittorrent")
		},
	})
}
