package check

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/developomp/pompup/internal/helper"
)

// Startup error
var (
	ErrStartupNotLinux  error = errors.New("platform is not Linux")
	ErrStartupIsRoot    error = errors.New("program running as root")
	ErrStartupNoPing    error = errors.New("ping is not installed")
	ErrStartupNotOnline error = errors.New("no internet connection")
	ErrStartupNotArch   error = errors.New("OS is not Arch Linux")
	ErrStartupNoSudo    error = errors.New("sudo is not installed")
	ErrStartupNoWget    error = errors.New("wget is not installed")
)

// StartupCheck checks various aspect of the system required by the program.
// It only needs to be called once during the initialization process.
func StartupCheck() error {
	if !IsLinux() {
		return ErrStartupNotLinux
	}

	if IsRoot() {
		return ErrStartupIsRoot
	}

	if !IsBinInstalled("ping") {
		return ErrStartupNoPing
	}

	if !IsOnline() {
		return ErrStartupNotOnline
	}

	if !IsBinInstalled("pacman") {
		return ErrStartupNotArch
	}

	if !IsBinInstalled("sudo") {
		return ErrStartupNoSudo
	}

	if !IsBinInstalled("wget") {
		return ErrStartupNoWget
	}

	return nil
}

// IsRoot checks if the program is running with root permission.
func IsRoot() bool {
	return os.Geteuid() == 0
}

// IsLinux checks if the program is running on a Linux-based Operating System.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsBinInstalled checks if a command exists.
func IsBinInstalled(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}

// IsArchPkgInstalled checks if an Arch/AUR package has been installed already.
func IsArchPkgInstalled(installer string, packageName string) bool {
	return helper.BashRun(fmt.Sprintf("%s -Q | grep -E '(^|\\s)%v($|\\s)'", installer, packageName)) == nil
}

// IsOnline checks if the program has working internet connection.
func IsOnline() bool {
	// ping archlinux.org by sending one packet (-c 1)
	return helper.Run("ping", "-c", "1", "archlinux.org") == nil
}

// PathExists checks whether the given path exists or not
func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
