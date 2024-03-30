package wrapper

import (
	"fmt"
	"os"
	"path/filepath"
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
