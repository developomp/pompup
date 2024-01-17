package check

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
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

	if !IsInstalled("ping") {
		return ErrStartupNoPing
	}

	if !IsOnline() {
		return ErrStartupNotOnline
	}

	if !IsInstalled("pacman") {
		return ErrStartupNotArch
	}

	if !IsInstalled("sudo") {
		return ErrStartupNoSudo
	}

	if !IsInstalled("wget") {
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

// IsInstalled checks if a command exists.
func IsInstalled(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}

// IsOnline checks if the program has working internet connection.
func IsOnline() bool {
	// ping archlinux.org by sending one packet (-c 1)
	cmd := exec.Command("ping", "-c", "1", "archlinux.org")
	cmd.Stderr = os.Stderr
	return cmd.Run() == nil
}

// PathExists checks whether the given path exists or not
func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
