package wrapper

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

// DefaultFilePerm is equivalent to -rw-r--r--
const DefaultFilePerm = 0644

// DefaultDirPerm is equivalent to -rwxr-xr-x
const DefaultDirPerm = 0755

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		pterm.Fatal.Println("Failed to get user's home directory: ", err)
	}

	return homeDir
}

func GetTmpDir() string {
	// ⚠️ ⚠️ ⚠️ ⚠️ ⚠️
	// BE FUCKING CAREFUL WHEN CHANGING THIS
	// IT IS USED WITH "rm -rf"
	// IT MIGHT KDE6 THE SHIT OUT OF YOUR HOME DIR
	// ⚠️ ⚠️ ⚠️ ⚠️ ⚠️
	return "/tmp/pompup"
}

func InHome(path string) string {
	return fmt.Sprintf("%s/%s", GetHomeDir(), path)
}
