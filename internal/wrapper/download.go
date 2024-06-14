package wrapper

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/pterm/pterm"
)

// DownloadFromGitHub downloads a file from GitHub releases to ~/Downloads directory without installing it.
// The user has to manually install it using gearlever.
func DownloadFromGitHub(user, repo, filePattern string) {
	// GET https://api.github.com/repos/<user or org>/<repo>/releases/latest
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), user, repo)
	if err != nil {
		pterm.Fatal.Printfln("Failed to get latest %s/%s version: %s", user, repo, err)
	}
	var downloadURL, fileName string
	for _, asset := range release.Assets {
		matched := RegexMatch(filePattern, *asset.Name)

		if matched {
			fileName = *asset.Name
			downloadURL = *asset.BrowserDownloadURL
			break
		}
	}
	if len(downloadURL) == 0 {
		pterm.Fatal.Printfln("Failed to get latest %s/%s version.", user, repo)
	}

	filePath := InHome("Downloads/" + fileName)

	// create parent directories
	err = os.MkdirAll(filepath.Dir(filePath), DefaultDirPerm)
	if err != nil {
		pterm.Fatal.Printfln("Failed to create directory \"%s\": %s", filePath, err)
	}

	// download file
	pterm.Debug.Println("Downloading", filePath, "from", downloadURL)
	cmd := exec.Command("wget", downloadURL, "-q", "--show-progress", "-O", filePath)
	cmd.Stderr = os.Stderr // show stderr
	cmd.Stdout = os.Stdout // show stdout
	if err := cmd.Run(); err != nil {
		pterm.Fatal.Printfln("Failed to download %s/%s: %s", user, repo, err)
	}
}
