package wrapper

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
)

// WriteFile writes data to path. It can safely perform multiple operations on the same path.
func WriteFile(path string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(path), DefaultFilePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, DefaultFilePerm)
}

// SudoWriteFile writes data to path where file is owned by root. It can safely perform multiple operations on the same path.
func SudoWriteFile(path string, data string) {
	if IsFileUpdated(path, data) {
		return
	}

	// create temporary file
	tmpPath := filepath.Join(TmpDir, filepath.Base(path))
	WriteFile(tmpPath, []byte(data))
	err := WriteFile(tmpPath, []byte(data))
	if err != nil {
		pterm.Fatal.Println("Failed to write to", path)
	}

	err = Run(
		"sudo",
		"install",
		"--group",
		"root",
		"--owner",
		"root",
		"--mode",
		"0644", // rw-r--r--
		tmpPath,
		path,
	)
	if err != nil {
		pterm.Fatal.Println("Failed to write to", path)
	}
}

// IsFileUpdated checks if file's content is already s.
func IsFileUpdated(filePath string, s string) bool {
	return strings.TrimSpace(ReadFile(filePath)) == s
}

func ReadFile(filePath string) string {
	b, err := os.ReadFile(filePath)
	if err != nil {
		pterm.Fatal.Printfln("Failed to read %s: %s", filePath, err)
	}

	return string(b)
}
