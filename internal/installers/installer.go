package installers

import (
	"fmt"
)

type Installer struct {
	Name      string   // Name is the display name of the installer.
	Desc      string   // Desc briefly explains what the installer does.
	Setup     func()   // Setup contains logic regarding the setup process.
	Reminders []string // Reminders for manual tasks user has to perform. Shows after all Setup functions are executed.
}

// list of installers
var Installers []*Installer

func register(installer *Installer) {
	for i, reminder := range installer.Reminders {
		installer.Reminders[i] = fmt.Sprintf("[%s] %s", installer.Name, reminder)
	}

	Installers = append(Installers, installer)
}
