package workflows

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/developomp/pompup/internal/helper"
	"github.com/google/go-github/github"
	"github.com/pterm/pterm"
)

func init() {
	register(&Workflow{
		Name: "osu!",
		Desc: "osu!lazer and tablet driver",
		Tags: []Tag{Gaming, Gui},
		Setup: func() {
			gearleverWorkflow.Setup()
			downloadOsuAppImage()
		},
		Reminders: append(gearleverWorkflow.Reminders,
			[]string{
				"Install osu.appimage file in Downloads directory",
				"Install osu! skin from https://github.com/developomp/osu-pomp-skin",
			}...,
		),
	})
}

// downloadOsuAppImage downloads osu!lazer AppImage to user's Downloads directory
// without installing it. The user then has to manually install it using gearlever.
func downloadOsuAppImage() {
	// GET https://api.github.com/repos/ppy/osu/releases/latest
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "ppy", "osu")
	if err != nil {
		pterm.Fatal.Println("Failed to get latest osu version:", err)
	}
	var downloadURL string
	for _, asset := range release.Assets {
		if *asset.Name == "osu.AppImage" {
			downloadURL = *asset.BrowserDownloadURL
			break
		}
	}

	osuPath := helper.InHome("Downloads/osu.appimage")

	// create parent directories
	err = os.MkdirAll(filepath.Dir(osuPath), helper.DefaultPerm)
	if err != nil {
		pterm.Fatal.Printfln("Failed to create directory \"%s\": %s", osuPath, err)
	}

	// download latest osu AppImage
	pterm.Debug.Println("Downloading", osuPath, "from", downloadURL)
	cmd := exec.Command("wget", downloadURL, "-q", "--show-progress", "-O", osuPath)
	cmd.Stderr = os.Stderr // show stderr
	cmd.Stdout = os.Stdout // show stdout
	if err := cmd.Run(); err != nil {
		pterm.Fatal.Println("Failed to download osu lazer AppImage:", err)
	}
}
