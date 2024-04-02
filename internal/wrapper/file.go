package wrapper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
)

func WriteFile(path string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(path), DefaultFilePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, DefaultFilePerm)
}

func SudoWriteFile(path string, data string) error {
	return Run(
		"sudo",
		"bash",
		"-c",
		fmt.Sprintf("cat >%s <<EOL\n%s\nEOL", path, data),
	)
}

// IsFileUpdated checks if file's content is already s.
func IsFileUpdated(filePath string, s string) bool {
	b, err := os.ReadFile(filePath)
	if err != nil {
		pterm.Fatal.Printfln("Failed to read %s: %s", filePath, err)
	}

	return strings.TrimSpace(string(b)) == s
}
