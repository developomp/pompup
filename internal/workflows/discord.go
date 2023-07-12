package workflows

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"

	"github.com/developomp/pompup/internal/install"
	"github.com/pterm/pterm"
)

//go:embed assets/betterdiscord/plugins
var _BDPlugins embed.FS

// BDPlugins is a collection of BetterDiscord plugin and configuration files
var BDPlugins, _ = fs.Sub(_BDPlugins, "assets/betterdiscord/plugins")

func init() {
	register(&Workflow{
		Name: "Discord",
		Desc: "discord and betterdiscord plugins",
		Tags: []Tag{Gui},
		Setup: func() {
			install.Flatpak("com.discordapp.Discord")

			// BetterDiscord stuff
			install.Paru("betterdiscordctl-git")                               // BetterDiscord installer
			exec.Command("betterdiscordctl", "-i", "flatpak", "install").Run() // install BetterDiscord
			installBDPlugins()                                                 // install Plugins
		},
		PostSetup: func() {
			pterm.Info.Println("Enable BetterDiscord plugins in discord settings")
		},
	})
}

func installBDPlugins() {
	pterm.Info.Printfln("Installing BetterDiscord plugins")

	// go through each BetterDiscord plugin files
	err := fs.WalkDir(BDPlugins, ".", func(pluginName string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip if not file
		if entry.IsDir() {
			return nil
		}

		// e.g. /home/pomp
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		pluginPath := fmt.Sprintf("%s/.var/app/com.discordapp.Discord/config/BetterDiscord/plugins/%s", homeDir, pluginName)

		file, err := BDPlugins.Open(pluginName)
		if err != nil {
			return err
		}

		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		// write to file with -rw-r--r-- permission (0644)
		err = os.WriteFile(pluginPath, data, 0644)
		if err != nil {
			return err
		}

		pterm.Info.Printfln("created %s", pluginPath)

		return nil
	})

	if err != nil {
		pterm.Fatal.Println("Failed to install BetterDiscord plugin:", err)
	}
}