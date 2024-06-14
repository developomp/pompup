package wrapper

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/pterm/pterm"
)

// Startup error
var (
	ErrStartupNotLinux      error = errors.New("platform is not Linux")
	ErrStartupRunningAsRoot error = errors.New("program running as root")
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

// IsFlatpakInstalled checks if a flatpak package has been installed already.
func IsFlatpakInstalled(packageName string) bool {
	return BashRun(fmt.Sprintf("flatpak list | grep -E '%v'", packageName)) == nil
}

// IsNixInstalled checks if a nix package has been installed already.
func IsNixInstalled(packageName string) bool {
	lastSlash := strings.LastIndex(packageName, "/")
	lastPound := strings.LastIndex(packageName, "#")

	// holy fuck so ugly
	max := lastPound
	if lastSlash > max {
		max = lastSlash
	}

	packageName = packageName[max+1:]

	return BashRun(fmt.Sprintf("nix profile list --json | jq '.elements | keys[]' | grep -E '^\"%v\"$'", packageName)) == nil
}

func IsAppImageInstalled(packageName string) bool {
	return IsFileInDir(InHome("Programs/AppImages"), "(?i)gearlever_"+packageName+"_.*.appimage")
}

func IsFileInDir(dirPath, filePattern string) bool {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		pterm.Fatal.Printfln("Failed to read %s: %s", dirPath, err)
	}

	for _, e := range entries {
		if RegexMatch(filePattern, e.Name()) {
			return true // match found
		}
	}

	return false // match not found
}

func RegexMatch(pattern string, s string) bool {
	matched, err := regexp.MatchString(fmt.Sprintf("^%s$", pattern), s)
	if err != nil {
		pterm.Fatal.Printfln("Failed to check regex. pattern='%s' s='%s'", pattern, s)
	}

	return matched
}

// PathExists checks whether the given path exists or not
func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
