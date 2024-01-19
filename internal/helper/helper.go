package helper

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

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

func Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	// show error messages
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func BashRun(command string) error {
	return Run("bash", "-c", command)
}

func WriteFile(path string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, DefaultPerm)
}

func SudoWriteFile(path string, data string) error {
	return Run(
		"sudo",
		"bash",
		"-c",
		fmt.Sprintf("cat >%s <<EOL\n%s\nEOL", path, data),
	)
}
