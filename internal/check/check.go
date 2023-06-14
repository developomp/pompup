package check

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

// Startup error
var (
	ErrStartupIsRoot    error = errors.New("program running as root")
	ErrStartupNotLinux  error = errors.New("platform is not Linux")
	ErrStartupNotArch   error = errors.New("OS is not Arch Linux")
	ErrStartupNoSudo    error = errors.New("sudo is not installed")
	ErrStartupNotOnline error = errors.New("no internet connection")
)

// StartupCheck checks various aspect of the system required by the program.
// It only needs to be called once during the initialization process.
func StartupCheck() error {
	if IsRoot() {
		return ErrStartupIsRoot
	}

	if !IsLinux() {
		return ErrStartupNotLinux
	}

	if !IsInstalled("pacman") {
		return ErrStartupNotArch
	}

	if !IsInstalled("sudo") {
		return ErrStartupNoSudo
	}

	if !IsOnline() {
		return ErrStartupNotOnline
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

// IsInstalled checks if a command exists.
func IsInstalled(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}

// IsOnline checks if the program has working internet connection.
func IsOnline() bool {
	// ping archlinux.org by sending one packet (-c 1)
	return exec.Command("ping", "-c", "1", "archlinux.org").Run() == nil
}
