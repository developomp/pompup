package workflows

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os/exec"

	"github.com/developomp/pompup/internal/helper"
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
		Reminders: []string{
			"Enable BetterDiscord plugins in discord settings",
		},
	})
}

func installBDPlugins() {
	pterm.Debug.Printfln("Installing BetterDiscord plugins")

	// go through each BetterDiscord plugin files
	err := fs.WalkDir(BDPlugins, ".", func(pluginName string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip if not file
		if entry.IsDir() {
			return nil
		}

		pluginPath := fmt.Sprintf(
			"%s/.var/app/com.discordapp.Discord/config/BetterDiscord/plugins/%s",
			helper.GetHomeDir(),
			pluginName,
		)

		file, err := BDPlugins.Open(pluginName)
		if err != nil {
			return err
		}

		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		err = helper.WriteFile(pluginPath, data)
		if err != nil {
			return err
		}

		pterm.Debug.Printfln("created %s", pluginPath)

		return nil
	})

	if err != nil {
		pterm.Fatal.Println("Failed to install BetterDiscord plugin:", err)
	}
}
