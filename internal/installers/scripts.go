package installers

import (
	"embed"
	"path/filepath"

	"github.com/developomp/pompup/internal/wrapper"
	"github.com/pterm/pterm"
)

//go:embed assets/home/.local/bin/*
var scripts embed.FS

func init() {
	register(&Installer{
		Name: "Scripts",
		Desc: "~/.local/bin scripts",
		Setup: func() {
			paths, err := getAllFilenames(&scripts, "")
			if err != nil {
				pterm.Fatal.Println("Failed to read files from embed.FS:", err)
			}

			for _, filePath := range paths {
				b, err := scripts.ReadFile(filePath)
				if err != nil {
					pterm.Fatal.Println("Failed to read file from embed.FS:", err)
				}

				newPath := filepath.Join(wrapper.InHome(".local/bin"), filepath.Base(filePath))
				wrapper.WriteFile(newPath, b)
				err = wrapper.Run("chmod", "+x", newPath)
				if err != nil {
					pterm.Fatal.Printfln("Failed to make %s executable: %s", newPath, err)
				}
			}
		},
	})
}

func getAllFilenames(fs *embed.FS, path string) (out []string, err error) {
	// https://gist.github.com/clarkmcc/1fdab4472283bb68464d066d6b4169bc

	if len(path) == 0 {
		path = "."
	}

	entries, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}

	return
}
