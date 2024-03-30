// Package wrapper abstracts away OS-native operations like package installation
// and writing to file as root.
package wrapper

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/pterm/pterm"
)

// DefaultPerm is equivalent to -rw-r--r--
const DefaultPerm = 0644

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		pterm.Fatal.Println("Failed to get user's home directory: ", err)
	}

	return homeDir
}

func GetUserName() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user.Username
}

func InHome(path string) string {
	return fmt.Sprintf("%s/%s", GetHomeDir(), path)
}
