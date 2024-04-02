package wrapper

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Startup error
var (
	ErrStartupNotLinux      error = errors.New("platform is not Linux")
	ErrStartupRunningAsRoot error = errors.New("program running as root")
	ErrStartupNotOnline     error = errors.New("no internet connection")
	ErrStartupNoPacman      error = errors.New("pacman is not installed")
	ErrStartupNoSudo        error = errors.New("sudo is not installed")
)

// StartupCheck checks whether there are any irrecoverable issues in the system
// that prevents the program to begin installing.
func StartupCheck() error {
	if !IsLinux() {
		return ErrStartupNotLinux
	}

	if IsRoot() {
		return ErrStartupRunningAsRoot
	}

	if !IsOnline() {
		return ErrStartupNotOnline
	}

	if !IsBinInstalled("pacman") {
		return ErrStartupNoPacman
	}

	if !IsBinInstalled("sudo") {
		return ErrStartupNoSudo
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

// IsBinInstalled checks if a file can be found in PATH.
func IsBinInstalled(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}

// IsArchPkgInstalled checks if an arch package has been installed already.
func IsArchPkgInstalled(packageName string) bool {
	return BashRun(fmt.Sprintf("pacman -Q | grep -E '(^|\\s)%v($|\\s)'", packageName)) == nil
}

// IsFlatpakInstalled checks if an flatpak package has been installed already.
func IsFlatpakInstalled(packageName string) bool {
	return BashRun(fmt.Sprintf("flatpak list | grep -E '%v'", packageName)) == nil
}

// IsOnline checks if the program has working internet connection.
func IsOnline() bool {
	// ping archlinux.org by sending one packet (-c 1)
	return Run("ping", "-c", "1", "archlinux.org") == nil
}

// PathExists checks whether the given path exists or not
func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
