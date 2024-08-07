package wrapper

import (
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

// WriteFile writes data to path. It can safely perform multiple operations on the same path.
func WriteFile(filePath string, data []byte) {
	// create parent directories
	dirPath := filepath.Dir(filePath)
	err := os.MkdirAll(dirPath, DefaultDirPerm)
	if err != nil {
		pterm.Fatal.Printfln("Failed to create directory %s: %s", dirPath, err)
	}

	err = os.WriteFile(filePath, data, DefaultFilePerm)
	if err != nil {
		pterm.Fatal.Printfln("Failed to write to file %s: %s", filePath, err)
	}
}

// SudoWriteFile writes data to path where file is owned by root. It can safely perform multiple operations on the same path.
func SudoWriteFile(path string, data string) {
	if IsFileUpdated(path, data) {
		return
	}

	// create temporary file
	tmpPath := filepath.Join(TmpDir, filepath.Base(path))
	WriteFile(tmpPath, []byte(data))

	err := Run(
		"sudo",
		"install",
		"-D", // auto-generate parent directories
		"--group", "root",
		"--owner", "root",
		"--mode", "0644", // rw-r--r--
		tmpPath,
		path,
	)
	if err != nil {
		pterm.Fatal.Println("Failed to write to", path)
	}
}

// IsFileUpdated checks if file's content is already s.
func IsFileUpdated(filePath string, s string) bool {
	if !PathExists(filePath) {
		return false
	}

	content, fileExists := ReadFile(filePath)

	if !fileExists {
		return false
	}

	return content == s
}

func ReadFile(filePath string) (content string, fileExists bool) {
	fileExists = true

	if !PathExists(filePath) {
		content = ""
		fileExists = false
		return
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		pterm.Fatal.Printfln("Failed to read %s: %s", filePath, err)
	}

	content = string(b)

	return
}
